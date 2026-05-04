package user

import (
	"fmt"
)

func (s *userService) ValidateTokenAndGetUserID(token string) (string, error) {
	valid, jwtData, err := s.jwtService.Parse(token)
	if err != nil {
		return "", fmt.Errorf("invalid token: %w", err)
	}
	if !valid {
		return "", fmt.Errorf("invalid token")
	}
	return jwtData.UserID, nil
}

func (s *userService) ValidateTokenAndGetUserRole(token string) (string, error) {
	valid, jwtData, err := s.jwtService.Parse(token)
	if err != nil {
		return "", fmt.Errorf("invalid token: %w", err)
	}
	if !valid {
		return "", fmt.Errorf("invalid token")
	}
	return jwtData.Role, nil
}
