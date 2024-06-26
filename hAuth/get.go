package hAuth

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"time"

	"github.com/flambra/helpers/hError"
	"github.com/flambra/helpers/hReq"
)

var (
	accessToken string
	expireAt    time.Time
)

type GetTokenResponse struct {
	AccessToken string    `json:"access_token"`
	ExpireAt    time.Time `json:"expire_at"`
}

func GetToken() (string, error) {
	username := os.Getenv("AUTH_USERNAME")
	password := os.Getenv("AUTH_PASSWORD")
	url := os.Getenv("AUTH_URL")
	if username == "" || password == "" || url == "" {
		return "", hError.New("Missing environment variables: AUTH_USERNAME, AUTH_PASSWORD, or AUTH_URL")
	}

	if accessToken == "" || isTokenExpired() {
		authoritazion := BasicAuth(username, password)
		request := hReq.Request{
			Url:           url + "/client/auth",
			Authorization: authoritazion,
		}

		responseBytes, err := request.Post()
		if err != nil {
			return "", err
		}

		var response GetTokenResponse
		err = json.Unmarshal(responseBytes, &response)
		if err != nil {
			return "", err
		}

		accessToken = response.AccessToken
		expireAt = response.ExpireAt
	}
	return accessToken, nil
}

func isTokenExpired() bool {
	now := time.Now()
	return now.After(expireAt)
}

func BasicAuth(username, password string) string {
	credentials := username + ":" + password
	base64Credentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	return "Basic " + base64Credentials
}
