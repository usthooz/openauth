package google

import (
	"errors"
)

var (
	ErrBadToken          = errors.New("jwt: invalid token")
	ErrBadHeader         = errors.New("jwt: bad jwt header")
	ErrBadClaimSet       = errors.New("jwt: bad jwt claim set")
	ErrPublicKeyNotFound = errors.New("jwt: no correspond public key found")
	ErrBadClientID       = errors.New("jwt: client ID mismatch")
	ErrBadIssuer         = errors.New("jwt: bad issuer")
	ErrExpired           = errors.New("jwt: token expired")
)

// Header ...
type Header struct {
	Alg string `json:"alg"`
	Kid string `json:"kid"`
	Typ string `json:"typ"`
}

// ClaimSet ...
type ClaimSet struct {
	Iss        string `json:"iss"`
	Aud        string `json:"aud"`
	Exp        int64  `json:"exp"`
	Iat        int64  `json:"iat"`
	Typ        string `json:"typ"`
	Sub        string `json:"sub"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Picture    string `json:"picture"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Locale     string `json:"locale"`
	Domain     string `json:"hd"`
}
