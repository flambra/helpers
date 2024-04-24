package password

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func DecryptPassword(encrypted, entered string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(entered)); err != nil {
		return err
	}
	return nil
}
