package user

import (
	"context"
	"solution/internal/model"
)

func (r *repo) GetAll(ctx context.Context, page, size int) ([]*model.User, int, error) {
	offset := page * size

	var total int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM users WHERE is_active = true").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, email, password_hash, full_name, age, region, gender, marital_status, role, is_active, created_at, updated_at
		FROM users 
		WHERE is_active = true
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, size, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.PasswordHash,
			&user.FullName,
			&user.Age,
			&user.Region,
			&user.Gender,
			&user.MaritalStatus,
			&user.Role,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
