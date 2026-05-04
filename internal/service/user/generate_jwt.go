package user

import (
	"fmt"
	"solution/internal/model"
	"solution/pkg/jwt"
	"time"
)

func (s *userService) generateJWT(user model.User) (string, error) {

	if user.ID == "" || user.Role == "" {
		return "", fmt.Errorf("invalid user data for JWT generation")
	}

	return jwt.GenerateToken(user.ID, string(user.Role), s.jwtSecret, time.Hour)
}
