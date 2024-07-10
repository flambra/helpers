package hToken

import (
	"github.com/flambra/helpers/hError"
	"github.com/golang-jwt/jwt/v5"
)

func Parse(token string) (map[string]interface{}, error) {
	parsedToken, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, hError.New("invalid token claims")
	}

	data, ok := claims["dat"].(map[string]interface{})
	if !ok {
		return nil, hError.New("data claim not found or invalid")
	}

	return data, nil
}
