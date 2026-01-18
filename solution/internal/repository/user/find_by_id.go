package user

import (
	"context"
	"database/sql"
	"fmt"
	"solution/internal/model"
)

// FindByID ищет пользователя по ID
// Используется в GetMe и других операциях где мы уже знаем ID из JWT
func (r *repository) FindByID(ctx context.Context, id string) (model.User, error) {
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

	err := r.db.QueryRowContext(ctx, query, id).Scan(
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
			return model.User{}, fmt.Errorf("user with id %s not found: %w", id, err)
		}
		return model.User{}, fmt.Errorf("failed to find user by id: %w", err)
	}

	// Конвертируем nullable поля в указатели
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

	// Конвертируем роль в строгий тип
	user.Role = model.UserRole(user.Role)

	return user, nil
}

// FindByIDIncludingInactive ищет пользователя по ID включая неактивных
// Нужен для административных функций
func (r *repository) FindByIDIncludingInactive(ctx context.Context, id string) (model.User, error) {
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

	err := r.db.QueryRowContext(ctx, query, id).Scan(
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
			return model.User{}, fmt.Errorf("user with id %s not found: %w", id, err)
		}
		return model.User{}, fmt.Errorf("failed to find user by id: %w", err)
	}

	// Конвертируем nullable поля в указатели
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

	// Конвертируем роль в строгий тип
	user.Role = model.UserRole(user.Role)

	return user, nil
}
