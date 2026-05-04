package user

import (
	"context"
	"database/sql"
	"fmt"
	"solution/internal/model"
)

func (r *repo) GetByID(ctx context.Context, userID string) (*model.User, error) {
	query := `
		SELECT id, email, password_hash, full_name, age, region, 
			   gender, marital_status, role, is_active, created_at, updated_at
		FROM users 
		WHERE id = $1 AND is_active = true
	`

	var user model.User
	var age sql.NullInt32
	var region sql.NullString
	var gender sql.NullString
	var maritalStatus sql.NullString

	err := r.db.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&age,
		&region,
		&gender,
		&maritalStatus,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	if age.Valid {
		ageInt := int(age.Int32)
		user.Age = &ageInt
	}

	if region.Valid {
		user.Region = &region.String
	}

	if gender.Valid {
		g := model.Gender(gender.String)
		user.Gender = &g
	}

	if maritalStatus.Valid {
		ms := model.MaritalStatus(maritalStatus.String)
		user.MaritalStatus = &ms
	}

	user.Role = model.UserRole(user.Role)

	return &user, nil
}

func (r *repo) GetByIDIncludingInactive(ctx context.Context, userID string) (*model.User, error) {
	query := `
		SELECT id, email, password_hash, full_name, age, region, 
			   gender, marital_status, role, is_active, created_at, updated_at
		FROM users 
		WHERE id = $1
	`

	var user model.User
	var age sql.NullInt32
	var region sql.NullString
	var gender sql.NullString
	var maritalStatus sql.NullString

	err := r.db.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&age,
		&region,
		&gender,
		&maritalStatus,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	if age.Valid {
		ageInt := int(age.Int32)
		user.Age = &ageInt
	}

	if region.Valid {
		user.Region = &region.String
	}

	if gender.Valid {
		g := model.Gender(gender.String)
		user.Gender = &g
	}

	if maritalStatus.Valid {
		ms := model.MaritalStatus(maritalStatus.String)
		user.MaritalStatus = &ms
	}

	user.Role = model.UserRole(user.Role)

	return &user, nil
}
