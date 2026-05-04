package user

import (
	"context"
	"database/sql"
	"fmt"
	"solution/internal/model"
	"strings"
	"time"
)

func (r *repo) Update(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error) {
	now := time.Now().UTC()

	var updates []string
	var args []interface{}
	argIndex := 1

	if req.FullName != nil {
		updates = append(updates, fmt.Sprintf("full_name = $%d", argIndex))
		args = append(args, *req.FullName)
		argIndex++
	}

	var age sql.NullInt32
	if req.Age != nil {
		age = sql.NullInt32{Int32: int32(*req.Age), Valid: true}
	}
	updates = append(updates, fmt.Sprintf("age = $%d", argIndex))
	args = append(args, age)
	argIndex++

	var region sql.NullString
	if req.Region != nil {
		region = sql.NullString{String: *req.Region, Valid: true}
	}
	updates = append(updates, fmt.Sprintf("region = $%d", argIndex))
	args = append(args, region)
	argIndex++

	var gender sql.NullString
	if req.Gender != nil {
		gender = sql.NullString{String: string(*req.Gender), Valid: true}
	}
	updates = append(updates, fmt.Sprintf("gender = $%d", argIndex))
	args = append(args, gender)
	argIndex++

	var maritalStatus sql.NullString
	if req.MaritalStatus != nil {
		maritalStatus = sql.NullString{String: string(*req.MaritalStatus), Valid: true}
	}
	updates = append(updates, fmt.Sprintf("marital_status = $%d", argIndex))
	args = append(args, maritalStatus)
	argIndex++

	updates = append(updates, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, now)

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d AND is_active = true",
		strings.Join(updates, ", "), argIndex+1)

	args = append(args, userID)

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return r.GetByID(ctx, userID)
}
