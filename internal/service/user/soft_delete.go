package user

import (
	"context"
	"fmt"
)

func (s *userService) SoftDelete(ctx context.Context, userID string) error {
	if err := s.userRepo.SoftDelete(ctx, userID); err != nil {
		return fmt.Errorf("failed to soft delete user: %w", err)
	}

	return nil
}
