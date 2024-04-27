package hValidate

import (
	"regexp"

	"github.com/flambra/helpers/hError"
)

func Email(email string) error {
	if hasEmptySpaces(email) {
		return hError.New("email has empty characters")
	}

	if isAllLowerCase(email) {
		return hError.New("email has upper case characters")
	}

	if !checkAt(email) {
		return hError.New("email missing @ character")
	}

	return nil
}

func checkAt(s string) bool {
	re := regexp.MustCompile(`@`)
	return re.MatchString(s)
}
