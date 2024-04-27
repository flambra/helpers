package passwords

import (
	"fmt"
	"regexp"
)

func Validate(password string) error {
	if !containsMinLength(password) {
		return fmt.Errorf("a senha deve ter pelo menos 8 caracteres")
	}
	if !containsUpperCase(password) {
		return fmt.Errorf("a senha deve conter pelo menos uma letra maiúscula")
	}
	if !containsLowerCase(password) {
		return fmt.Errorf("a senha deve conter pelo menos uma letra minúscula")
	}
	if !containsDigit(password) {
		return fmt.Errorf("a senha deve conter pelo menos um número")
	}
	if !containsSpecialChar(password) {
		return fmt.Errorf("a senha deve conter pelo menos um caractere especial")
	}
	return nil
}

func containsMinLength(s string) bool {
	return len(s) >= 8
}
func containsUpperCase(s string) bool {
	upperCase := regexp.MustCompile(`[A-Z]`)
	return upperCase.MatchString(s)
}

func containsLowerCase(s string) bool {
	lowerCase := regexp.MustCompile(`[a-z]`)
	return lowerCase.MatchString(s)
}

func containsDigit(s string) bool {
	digit := regexp.MustCompile(`[0-9]`)
	return digit.MatchString(s)
}

func containsSpecialChar(s string) bool {
	specialChar := regexp.MustCompile(`[^A-Za-z0-9]`)
	return specialChar.MatchString(s)
}
