package token

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"time"

	"github.com/flambra/helpers/http"
)

var (
	accessToken string
	expireAt    time.Time
)

type GetTokenResponse struct {
	AccessToken string    `json:"access_token"`
	ExpireAt    time.Time `json:"expire_at"`
}

func Get() (string, error) {
	if accessToken == "" || isTokenExpired() {
		authoritazion := BasicAuth(os.Getenv("AUTH_USERNAME"), os.Getenv("AUTH_PASSWORD"))
		request := http.HttpRequest{
			Url:           os.Getenv("AUTH_URL"),
			Authorization: authoritazion,
			StatusCode:    200,
		}

		responseBytes, err := request.Get()
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
