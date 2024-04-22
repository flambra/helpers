package validate

import (
	"errors"
	"regexp"
)

func Email(email string) error {
	if hasEmptySpaces(email) {
		return errors.New("email has empty characters")
	}

	if isAllLowerCase(email) {
		return errors.New("email has upper case characters")
	}

	if !checkAt(email) {
		return errors.New("email missing @ character")
	}

	return nil
}

func checkAt(s string) bool {
	re := regexp.MustCompile(`@`)
	return re.MatchString(s)
}
