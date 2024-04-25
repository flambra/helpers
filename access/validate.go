package token

import (
	"os"
	"strings"

	"github.com/flambra/helpers/errgen"
	"github.com/golang-jwt/jwt"
)

func ValidateToken(authorization string) error {
	parts := strings.Split(authorization, " ")
	if len(parts) != 2 {
		return errgen.New("Token error")
	}

	scheme, token := parts[0], parts[1]
	if !strings.EqualFold(scheme, "Bearer") {
		return errgen.New("Token malformatted")
	}

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

	if err != nil || !parsedToken.Valid {
		return err
	}

	return nil
}
