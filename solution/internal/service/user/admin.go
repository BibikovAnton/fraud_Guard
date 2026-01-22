package user

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"solution/internal/model"
)

func (s *userService) CreateAdmin(ctx context.Context, email, passwordHash, fullName string) error {
	exists, err := s.userRepo.ExistsByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("failed to check email existence: %w", err)
	}
	if exists {
		return nil // admin already exists, skip creation
	}

	return s.userRepo.CreateAdmin(ctx, email, passwordHash, fullName)
}

func (s *userService) CreateByAdmin(ctx context.Context, req model.UserCreateRequest) (*model.User, error) {
	exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email existence: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := model.NewUser(req.Email, string(hashedPassword), req.FullName, req.Role)

	user.Age = req.Age
	user.Region = req.Region
	user.Gender = req.Gender
	user.MaritalStatus = req.MaritalStatus

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

func (s *userService) UpdateByAdmin(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error) {
	return s.userRepo.UpdateByAdmin(ctx, userID, req)
}
