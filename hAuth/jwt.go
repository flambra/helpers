package hAuth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID    uint
	Email string
}

// GenerateJWT generates a JSON Web Token (JWT) for the given user.
// It takes a User object as input and returns the generated JWT as a string.
// If an error occurs during token generation, it returns an empty string and the error.
//
// Example usage:
//   userAuth := hAuth.User{
//     ID:    user.ID,
//     Email: user.Email,
//   }
//
//   token, err := hAuth.GenerateJWT(userAuth)
//   if err != nil {
//     return hResp.InternalServerErrorResponse(c, err.Error())
//   }
//
//   return hResp.SuccessResponse(c, fiber.Map{"token": token})
//
// The function requires the following environment variables to be set:
//   - TOKEN_SECRET_KEY: The secret key used for signing the token.
//   - TOKEN_DURATION: The duration (in hours) for which the token is valid, e.g., "1", "0.5", etc.
func GenerateJWT(user User) (string, error) {
	secretKey := os.Getenv("TOKEN_SECRET_KEY")
	durationStr := os.Getenv("TOKEN_DURATION")

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
