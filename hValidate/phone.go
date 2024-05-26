package hValidate

import (
	"regexp"
	"strings"

	"github.com/flambra/helpers/hError"
)

// Cellphone validates and formats a phone number to the international format (E.164).
func Cellphone(phone string) (string, error) {
	phone = strings.TrimSpace(phone)
	if !isValidInternationalPhone(phone) {
		return "", hError.New("failed to format phone: invalid international format")
	}
	return phone, nil
}

// isValidInternationalPhone validates a phone number in the international format (E.164).
func isValidInternationalPhone(s string) bool {
	re := regexp.MustCompile(`^\+\d{1,3}\d{1,14}$`)
	return re.MatchString(s)
}
