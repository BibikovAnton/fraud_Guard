package user

import (
	"context"
	"fmt"
	"solution/internal/model"
)

func (s *userService) GetByID(ctx context.Context, userID string) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *userService) GetByIDIncludingInactive(ctx context.Context, userID string) (*model.User, error) {
	user, err := s.userRepo.GetByIDIncludingInactive(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}
