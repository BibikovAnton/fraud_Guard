package user

import (
	"context"
	"database/sql"
	"fmt"
	"solution/internal/model"
	"time"
)

func (r *repo) Create(ctx context.Context, user model.User) error {
	now := time.Now().UTC()

	query := `
		INSERT INTO users (
			id, email, password_hash, full_name, age, region, 
			gender, marital_status, role, is_active, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		)`

	var age sql.NullInt32
	if user.Age != nil {
		age = sql.NullInt32{Int32: int32(*user.Age), Valid: true}
	}

	var region sql.NullString
	if user.Region != nil {
		region = sql.NullString{String: *user.Region, Valid: true}
	}

	var gender sql.NullString
	if user.Gender != nil {
		gender = sql.NullString{String: string(*user.Gender), Valid: true}
	}

	var maritalStatus sql.NullString
	if user.MaritalStatus != nil {
		maritalStatus = sql.NullString{String: string(*user.MaritalStatus), Valid: true}
	}

	_, err := r.db.Exec(ctx, query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.FullName,
		age,
		region,
		gender,
		maritalStatus,
		string(user.Role),
		user.IsActive,
		now,
		now,
	)

	if err != nil {
		if isUniqueViolation(err) {
			return fmt.Errorf("user with email %s already exists: %w", user.Email, err)
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func isUniqueViolation(err error) bool {
	if err != nil && err.Error() == "pq: duplicate key value violates unique constraint" {
		return true
	}
	return false
}
