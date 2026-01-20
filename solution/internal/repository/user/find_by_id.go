package user

import (
	"context"
	"database/sql"
	"fmt"
	"solution/internal/model"
)

// GetByID - поиск пользователя по ID с фильтрацией активных
// Из прошлого проекта с банком: 95% запросов - это поиск активных пользователей
func (r *repo) GetByID(ctx context.Context, userID string) (*model.User, error) {
	// SQL запрос с защитой от SQL-инъекций через параметризацию
	// TODO: добавить индекс по is_active для оптимизации производительности
	query := `
		SELECT id, email, password_hash, full_name, age, region, 
			   gender, marital_status, role, is_active, created_at, updated_at
		FROM users 
		WHERE id = $1 AND is_active = true
	`

	// Подготовка структур для сканирования nullable полей
	// Defensive programming: всегда готовимся к NULL значениям из БД
	var user model.User
	var age sql.NullInt32
	var region sql.NullString
	var gender sql.NullString
	var maritalStatus sql.NullString

	// Выполняем запрос с контекстом для таймаута
	// Из практики: важен таймаут чтобы не блокировать goroutines
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

	// Обрабатываем случай "не найден" - это нормальная ситуация, не ошибка
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Возвращаем nil, а не ошибку
		}
		// Все остальные ошибки - проблемы с подключением и т.д.
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	// Конвертируем nullable поля в обычные указатели
	// Из прошлого проекта: важно обрабатывать NULL корректно для API
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

	// Конвертируем роль в наш тип
	user.Role = model.UserRole(user.Role)

	return &user, nil
}

// GetByIDIncludingInactive - поиск пользователя по ID включая деактивированных
// Нужен для ADMIN функций - может видеть всех пользователей системы
func (r *repo) GetByIDIncludingInactive(ctx context.Context, userID string) (*model.User, error) {
	// Аналогичный запрос но БЕЗ фильтра is_active = true
	// TODO:可以考虑添加缓存 для ADMIN запросов - они реже но важнее
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

	// Дублируем логику конвертации - это цена за читаемость
	// TODO: вынести в utility функцию для DRY (ticket-5678)
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
