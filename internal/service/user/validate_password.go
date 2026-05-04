package user

import "fmt"

func (s *userService) validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if len(password) > 72 {
		return fmt.Errorf("password must be at most 72 characters long")
	}

	hasDigit := false
	hasLetter := false
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasDigit = true
		}
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			hasLetter = true
		}
		if hasDigit && hasLetter {
			break
		}
	}

	if !hasDigit || !hasLetter {
		return fmt.Errorf("password must contain at least one digit and one letter")
	}

	return nil
}
