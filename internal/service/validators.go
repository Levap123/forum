package service

import (
	"fmt"
	"net/mail"
)

func ValidateEmail(email string) error {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("Invalid email address")
	}

	if addr.Address == "" {
		return fmt.Errorf("Missing email address")
	}

	return nil
}
