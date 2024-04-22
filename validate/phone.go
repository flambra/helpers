package validate

import (
	"errors"
	"fmt"
	"regexp"
)

func Cellphone(phone string) (string, error) {
	if !isCellphoneFormatted(phone) {
		return formatToCellphone(phone)
	}
	return phone, nil
}

func isCellphoneFormatted(s string) bool {
	re := regexp.MustCompile(`^\(\d{2}\)\d{5}-\d{4}$`)
	return re.MatchString(s)
}

func formatToCellphone(phone string) (string, error) {
	raw := removeEmptySpaces(phone)
	if hasNonDigits(raw) {
		return "", errors.New("failed to format phone: must have only digits")
	}

	if len(raw) != 11 {
		return "", errors.New("failed to format phone: invalid length")
	}
	formatted := fmt.Sprintf("(%s)%s-%s", raw[0:2], raw[2:7], raw[7:11])
	return formatted, nil
}
