package validate

import (
	"github.com/flambra/helpers/errgen" 
	"regexp"
)

func Email(email string) error {
	if hasEmptySpaces(email) {
		return errgen.New("email has empty characters")
	}

	if isAllLowerCase(email) {
		return errgen.New("email has upper case characters")
	}

	if !checkAt(email) {
		return errgen.New("email missing @ character")
	}

	return nil
}

func checkAt(s string) bool {
	re := regexp.MustCompile(`@`)
	return re.MatchString(s)
}
