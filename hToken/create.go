package hToken

import (
	"encoding/base64"
	"encoding/json"
	"os"

	"github.com/flambra/helpers/hError"
	"github.com/flambra/helpers/hReq"
)

type Access struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func Create(data map[string]interface{}) (*Access, error) {
	username := os.Getenv("AUTH_USERNAME")
	password := os.Getenv("AUTH_PASSWORD")
	url := os.Getenv("AUTH_URL")
	if username == "" || password == "" || url == "" {
		return nil, hError.New("Missing environment variables: AUTH_USERNAME, AUTH_PASSWORD, or AUTH_URL")
	}

	authorization := BasicAuth(username, password)
	request := hReq.Request{
		Url:           url + "/token",
		Authorization: authorization,
		Body:          data,
	}

	responseBytes, err := request.Post()
	if err != nil {
		return nil, err
	}

	var response Access
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func BasicAuth(username, password string) string {
	credentials := username + ":" + password
	base64Credentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	return "Basic " + base64Credentials
}
