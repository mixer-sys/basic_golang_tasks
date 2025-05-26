package main

import (
	"errors"
	"fmt"
	"unicode"
)

func ValidatePassword(pass string) error {
	if len(pass) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}
	hasDigit := false
	hasUpper := false

	for _, char := range pass {
		if unicode.IsDigit(char) {
			hasDigit = true
		}
		if unicode.IsUpper(char) {
			hasUpper = true
		}
	}

	if !hasDigit {
		return errors.New("Password must contain at least one digit")
	}
	if !hasUpper {
		return errors.New("Password must contain at least one uppercase letter")
	}
	return nil
}

func main() {
	err := ValidatePassword("qwerty")
	fmt.Println(err)

	err = ValidatePassword("Qwertyqwerty1")
	fmt.Println(err)

}
