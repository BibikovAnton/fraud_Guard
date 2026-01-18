package user

import (
	"context"
)

// ExistsByEmail проверяет существование пользователя по email
// Используем EXISTS для производительности - не нужно тащить все данные
func (r *repository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM users WHERE email = $1
		)
	`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// ExistsByEmailAndActive проверяет существование активного пользователя по email
// Нужно для регистрации - не позволяем регистрироваться с уже занятым email
func (r *repository) ExistsByEmailAndActive(ctx context.Context, email string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM users WHERE email = $1 AND is_active = true
		)
	`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
