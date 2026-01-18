package user

import (
	"context"
	"database/sql"
	"fmt"
	"solution/internal/model"
	"time"
)

// Create сохраняет нового пользователя в базу данных
// Используем named parameters для читаемости и защиты от SQL инъекций
func (r *repository) Create(ctx context.Context, user model.User) error {
	// Воспроизводимость: всегда используем UTC время
	now := time.Now().UTC()
	
	query := `
		INSERT INTO users (
			id, email, password_hash, full_name, age, region, 
			gender, marital_status, role, is_active, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		)`

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

	_, err := r.db.ExecContext(ctx, query,
		user.ID,                    // $1
		user.Email,                 // $2
		user.PasswordHash,          // $3
		user.FullName,              // $4
		age,                        // $5
		region,                     // $6
		gender,                     // $7
		maritalStatus,              // $8
		string(user.Role),          // $9
		user.IsActive,              // $10
		now,                        // $11
		now,                        // $12
	)

	if err != nil {
		// Проверяем на уникальность email - частая ошибка при регистрации
		if isUniqueViolation(err) {
			return fmt.Errorf("user with email %s already exists: %w", user.Email, err)
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// isUniqueViolation проверяет, что ошибка связана с нарушением уникальности
// В PostgreSQL это код 23505
func isUniqueViolation(err error) bool {
	if err != nil && err.Error() == "pq: duplicate key value violates unique constraint" {
		return true
	}
	return false
}

// CreateAdmin создает администратора при старте приложения
// Используем в DI контейнере для инициализации
func (r *repository) CreateAdmin(ctx context.Context, email, passwordHash, fullName string) error {
	admin := model.NewUser(email, passwordHash, fullName, model.AdminRole)
	return r.Create(ctx, admin)
}
