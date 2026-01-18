package user

import (
	"context"
	"database/sql"
	"fmt"
	"solution/internal/model"
	"time"
)

// Update обновляет данные пользователя в базе данных
// Важно: обновляем только разрешенные поля для обычных пользователей
func (r *repository) Update(ctx context.Context, user model.User) error {
	// Воспроизводимость: всегда используем UTC время
	now := time.Now().UTC()
	
	query := `
		UPDATE users 
		SET full_name = $2, age = $3, region = $4, gender = $5, 
			marital_status = $6, updated_at = $7
		WHERE id = $1 AND is_active = true
	`

	// Подготовка значений с учетом nullable полей
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

	result, err := r.db.ExecContext(ctx, query,
		user.ID,                    // $1
		user.FullName,              // $2
		age,                        // $3
		region,                     // $4
		gender,                     // $5
		maritalStatus,              // $6
		now,                        // $7
	)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	// Проверяем, что хотя бы одна строка была обновлена
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found or inactive", user.ID)
	}

	return nil
}

// UpdateByAdmin обновляет данные пользователя с возможностью менять роль и статус
// Используется административными функциями
func (r *repository) UpdateByAdmin(ctx context.Context, user model.User) error {
	// Воспроизводимость: всегда используем UTC время
	now := time.Now().UTC()
	
	query := `
		UPDATE users 
		SET full_name = $2, age = $3, region = $4, gender = $5, 
			marital_status = $6, role = $7, is_active = $8, updated_at = $9
		WHERE id = $1
	`

	// Подготовка значений с учетом nullable полей
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

	result, err := r.db.ExecContext(ctx, query,
		user.ID,                    // $1
		user.FullName,              // $2
		age,                        // $3
		region,                     // $4
		gender,                     // $5
		maritalStatus,              // $6
		string(user.Role),          // $7
		user.IsActive,              // $8
		now,                        // $9
	)

	if err != nil {
		return fmt.Errorf("failed to update user by admin: %w", err)
	}

	// Проверяем, что хотя бы одна строка была обновлена
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", user.ID)
	}

	return nil
}

// SoftDelete деактивирует пользователя вместо физического удаления
// Следуем принципу soft-delete из readme.txt
func (r *repository) SoftDelete(ctx context.Context, id string) error {
	query := `
		UPDATE users 
		SET is_active = false, updated_at = $2
		WHERE id = $1
	`

	now := time.Now().UTC()
	
	result, err := r.db.ExecContext(ctx, query, id, now)
	if err != nil {
		return fmt.Errorf("failed to soft delete user: %w", err)
	}

	// Проверяем, что хотя бы одна строка была обновлена
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	return nil
}
