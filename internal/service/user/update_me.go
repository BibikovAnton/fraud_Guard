package user

import (
	"context"
	"fmt"
	"solution/internal/model"
)

func (s *userService) UpdateMe(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error) {
	existingUser, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing user: %w", err)
	}
	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	if req.FullName != nil {
		existingUser.FullName = *req.FullName
	}

	if req.Age != nil {
		existingUser.Age = req.Age
	}
	if req.Region != nil {
		existingUser.Region = req.Region
	}
	if req.Gender != nil {
		existingUser.Gender = req.Gender
	}
	if req.MaritalStatus != nil {
		existingUser.MaritalStatus = req.MaritalStatus
	}

	updatedUser, err := s.userRepo.Update(ctx, userID, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return updatedUser, nil
}
