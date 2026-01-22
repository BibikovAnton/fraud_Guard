package user

import (
	"context"
	"fmt"
	"time"
)

func (r *repo) SoftDelete(ctx context.Context, userID string) error {
	query := `
		UPDATE users 
		SET is_active = false, updated_at = $2
		WHERE id = $1
	`

	now := time.Now().UTC()

	result, err := r.db.Exec(ctx, query, userID, now)
	if err != nil {
		return fmt.Errorf("failed to soft delete user: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", userID)
	}

	return nil
}
