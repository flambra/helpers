package hValidate

import (
	"time"

	"github.com/flambra/helpers/hError"
)

// BirthDate checks if a given birth date is valid.
// It returns an error if the birth date is after 18 years ago.
func BirthDate(birthDate time.Time) error {
	eighteenYearsAgo := time.Now().AddDate(-18, 0, 0)

	if birthDate.After(eighteenYearsAgo) {
		return hError.New("user must be at least 18 years old")
	}

	return nil
}
