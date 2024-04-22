package validate

import (
	"errors"
	"fmt"
	"regexp"
)

func CNPJ(cnpj string) (string, error) {
	if !isCnpjFormatted(cnpj) {
		return formatToCnpj(cnpj)
	}
	return cnpj, nil
}

func isCnpjFormatted(s string) bool {
	re := regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`)
	return re.MatchString(s)
}

func formatToCnpj(cpf string) (string, error) {
	raw := removeEmptySpaces(cpf)
	if hasNonDigits(raw) {
		return "", errors.New("failed to format cnpj: must have only digits")
	}

	if len(raw) != 14 {
		return "", errors.New("failed to format cnpj: invalid length")
	}
	formatted := fmt.Sprintf("%s.%s.%s/%s-%s", raw[0:2], raw[2:5], raw[5:8], raw[8:12], raw[12:14])
	return formatted, nil
}
