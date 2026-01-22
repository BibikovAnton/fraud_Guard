package user

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"solution/internal/model"
)

func (s *userService) Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error) {

	user, err := s.userRepo.FindByEmailIncludingInactive(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("user lookup failed: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if !user.IsActive {
		return nil, fmt.Errorf("account is deactivated")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	token, tokenErr := s.generateJWT(*user)
	if tokenErr != nil {
		return nil, fmt.Errorf("JWT generation failed: %w", tokenErr)
	}

	return &model.LoginResponse{
		User:  *user,
		Token: token,
	}, nil
}
