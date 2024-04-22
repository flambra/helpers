package validate

import (
	"errors"
	"fmt"
	"regexp"
)

func CPF(cpf string) (string, error) {
	if !isCpfFormatted(cpf) {
		return formatToCpf(cpf)
	}
	return cpf, nil
}

func isCpfFormatted(s string) bool {
	re := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)
	return re.MatchString(s)
}

func formatToCpf(cpf string) (string, error) {
	raw := removeEmptySpaces(cpf)
	if hasNonDigits(raw) {
		return "", errors.New("failed to format cpf: must have only digits")
	}

	if len(raw) != 11 {
		return "", errors.New("failed to format cpf: invalid length")
	}
	formatted := fmt.Sprintf("%s.%s.%s-%s", raw[0:3], raw[3:6], raw[6:9], raw[9:11])
	return formatted, nil
}
