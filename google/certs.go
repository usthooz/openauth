package google

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const (
	googleCertsURL = "https://www.googleapis.com/oauth2/v3/certs"
)

var (
	// used to extract max-age from cache-control HTTP header.
	reMaxAge = regexp.MustCompile(`(?i)max-age=(\d+)`)
)

// Key is a cert key.
type Key struct {
	Use string `json:"use"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

// Certs are google certs.
type Certs struct {
	Keys   map[string]rsa.PublicKey
	Expiry time.Time
}

var (
	// global unique google certs.
	gCerts *Certs
)

// listCerts lists google certs.
func listCerts() (*Certs, error) {
	// use cached
	if gCerts != nil && time.Now().Before(gCerts.Expiry) {
		return gCerts, nil
	}

	// fetch certs
	resp, err := http.Get(googleCertsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// respect cache-control
	cacheAge := 5 * 60
	if cacheControl := resp.Header.Get("Cache-Control"); cacheControl != "" {
		if matches := reMaxAge.FindStringSubmatch(cacheControl); len(matches) == 2 {
			maxAge, err := strconv.ParseInt(matches[1], 10, 64)
			if err == nil {
				cacheAge = int(maxAge)
			}
		}
	}

	// parse all keys
	keysObj := struct {
		Keys []Key `json:"keys"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&keysObj); err != nil {
		return nil, err
	}

	pubKeys := map[string]rsa.PublicKey{}
	for _, key := range keysObj.Keys {
		if key.Use == "sig" && key.Kty == "RSA" {
			n, err := base64.RawURLEncoding.DecodeString(key.N)
			if err != nil {
				return nil, err
			}
			e, err := base64.RawURLEncoding.DecodeString(key.E)
			if err != nil {
				return nil, err
			}
			pubKeys[key.Kid] = rsa.PublicKey{
				N: big.NewInt(0).SetBytes(n),
				E: int(big.NewInt(0).SetBytes(e).Int64()),
			}
		}
	}

	// save to certs
	gCerts = &Certs{
		Keys:   pubKeys,
		Expiry: time.Now().Add(time.Second * time.Duration(cacheAge)),
	}

	return gCerts, nil
}
