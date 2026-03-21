package common

import (
	"fmt"
	"regexp"
	"time"
)

var onlyDigitsRegex = regexp.MustCompile(`^[0-9]+$`)

func ValidateBirthDate(birthDate string) error {
	if _, err := time.Parse("2006-01-02", birthDate); err != nil {
		return fmt.Errorf("birthdate must match format AAAA-MM-DD")
	}

	return nil
}

func ValidateDNI(dni string) error {
	if !onlyDigitsRegex.MatchString(dni) {
		return fmt.Errorf("dni must contain only digits")
	}

	return nil
}

func ValidateBetNumber(number string) error {
	if !onlyDigitsRegex.MatchString(number) {
		return fmt.Errorf("bet number must contain only digits")
	}

	return nil
}
