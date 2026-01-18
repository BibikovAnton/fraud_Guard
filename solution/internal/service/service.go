package service

import (
	"context"
	"solution/internal/model"
)

// AntifraudService - основной сервис антифрода
// TODO: добавить методы для работы с правилами и транзакциями
type AntifraudService interface {
}

// UserService - сервис для работы с пользователями
// Содержит всю бизнес-логику регистрации, авторизации и управления профилями
type UserService interface {
	// Регистрация и авторизация
	Register(ctx context.Context, req model.RegisterRequest) (model.AuthResponse, error)
	Login(ctx context.Context, req model.LoginRequest) (model.AuthResponse, error)
	
	// Управление профилем
	GetMe(ctx context.Context, userID string) (model.User, error)
	UpdateMe(ctx context.Context, userID string, req model.UserUpdateRequest) (model.User, error)
	
	// TODO: добавить методы для управления пользователями (админские функции)
	// CreateByAdmin(ctx context.Context, req model.UserCreateRequest) (model.User, error)
	// GetByID(ctx context.Context, userID string) (model.User, error)
	// UpdateByAdmin(ctx context.Context, userID string, req model.UserUpdateRequest) (model.User, error)
	// GetAll(ctx context.Context, page, size int) ([]model.User, int, error)
	// SoftDelete(ctx context.Context, userID string) error
}
