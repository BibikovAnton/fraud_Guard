package user

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"solution/internal/model"
)

func (s *userService) Register(ctx context.Context, req model.RegisterRequest) (*model.RegisterResponse, error) {

	emailExists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("email check failed: %w", err)
	}
	if emailExists {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return nil, fmt.Errorf("password hashing failed: %w", hashErr)
	}

	newUser := model.NewUser(req.Email, string(hashedPassword), req.FullName, model.UserRoleConst, req.Age, req.Region, req.Gender, req.MaritalStatus)

	if err := s.userRepo.Create(ctx, newUser); err != nil {
		return nil, fmt.Errorf("user creation failed: %w", err)
	}

	token, tokenErr := s.generateJWT(newUser)
	if tokenErr != nil {
		return nil, fmt.Errorf("JWT generation failed: %w", tokenErr)
	}

	return &model.RegisterResponse{
		User:        newUser,
		AccessToken: token,
		ExpiresIn:   3600,
	}, nil
}
