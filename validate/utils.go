package validate

import (
	"regexp"
	"strings"
)

func hasEmptySpaces(s string) bool {
	re := regexp.MustCompile(`\s`)
	return re.MatchString(s)
}

func isAllLowerCase(s string) bool {
	re := regexp.MustCompile(`[A-Z]`)
	return re.MatchString(s)
}

func hasNonDigits(s string) bool {
	re := regexp.MustCompile(`[^0-9]`)
	return re.MatchString(s)
}

func removeEmptySpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
