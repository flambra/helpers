package hToken

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func Validate(token string) error {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("PUBLIC_KEY")))
	if err != nil {
		return err
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, err
		}
		return publicKey, nil
	})

	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return err
	}

	return nil
}
