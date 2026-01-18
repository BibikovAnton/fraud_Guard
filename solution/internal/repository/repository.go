package repository

import (
	"context"
	"solution/internal/model"
)

// AntifraudRepository - репозиторий для работы с правилами антифрода
// TODO: добавить методы для CRUD операций с правилами
type AntifraudRepository interface {
}

// UserRepository - полный интерфейс для работы с пользователями
// Включает все необходимые методы для бизнес-логики
type UserRepository interface {
	// Базовые операции
	Create(ctx context.Context, user model.User) error
	FindByEmail(ctx context.Context, email string) (model.User, error)
	FindByID(ctx context.Context, id string) (model.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Update(ctx context.Context, user model.User) error
	
	// Дополнительные методы для авторизации и управления
	FindByEmailIncludingInactive(ctx context.Context, email string) (model.User, error)
	FindByIDIncludingInactive(ctx context.Context, id string) (model.User, error)
	ExistsByEmailAndActive(ctx context.Context, email string) (bool, error)
	
	// Административные функции
	UpdateByAdmin(ctx context.Context, user model.User) error
	SoftDelete(ctx context.Context, id string) error
	
	// TODO: добавить методы для списка пользователей
	// FindAll(ctx context.Context, offset, limit int) ([]model.User, int, error)
}
