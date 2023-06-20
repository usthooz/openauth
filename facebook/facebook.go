package facebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	facebookAppTokenURL   = "https://graph.facebook.com/oauth/access_token"
	facebookDebugTokenURL = "https://graph.facebook.com/debug_token"
)

// FacebookVerify
func FacebookVerify(appId, appSecret, accessToken string) (*DebugTokenResp, error) {
	client := &http.Client{}

	// getAccessToken
	appToken, err := getAccessToken(appId, appSecret, client)
	if err != nil {
		return nil, err
	}

	// verify auth token
	return debugToken(appId, client, accessToken, appToken)
}

// getAccessToken
func getAccessToken(appId, appSecret string, client *http.Client) (string, error) {
	req, err := http.NewRequest("GET", facebookAppTokenURL, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("client_id", appId)
	q.Add("client_secret", appSecret)
	q.Add("grant_type", "client_credentials")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	res.Body.Close()

	var (
		at *AccessTokenResp
	)
	err = json.Unmarshal(body, &at)
	if err != nil {
		return "", err
	}

	if at == nil || len(at.AccessToken) <= 0 {
		return "", errors.New("failed to get app access token")
	}

	return at.AccessToken, nil
}

// debugToken
func debugToken(appId string, client *http.Client, accessToken, appToken string) (*DebugTokenResp, error) {
	req, err := http.NewRequest("GET", facebookDebugTokenURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("input_token", accessToken)
	q.Add("access_token", appToken)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var (
		responseData *DebugTokenResp
	)

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	if responseData == nil {
		return nil, fmt.Errorf("wrong appID. verify resp is nil")
	}

	if responseData.Data.AppID != appId {
		return nil, fmt.Errorf("wrong appID. The token is not for our app")
	}

	if responseData.Data.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("expired access token")
	}

	return responseData, nil
}
