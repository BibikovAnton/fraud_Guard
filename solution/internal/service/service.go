package service

import (
	"context"
	"solution/internal/model"
)

// AntifraudService - основной сервис антифрода
// TODO: добавить методы для работы с правилами и транзакциями
type AntifraudService interface {
	// TODO: Добавить методы для работы с антифродом
}

// UserService - сервис для работы с пользователями
// Содержит всю бизнес-логику регистрации, авторизации и управления профилями
type UserService interface {
	// Аутентификация и регистрация
	Register(ctx context.Context, req model.RegisterRequest) (*model.RegisterResponse, error)
	Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error)

	// Профиль пользователя
	GetMe(ctx context.Context, userID string) (*model.User, error)
	UpdateMe(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error)

	// Административные методы
	CreateByAdmin(ctx context.Context, req model.UserCreateRequest) (*model.User, error)
	GetByID(ctx context.Context, userID string) (*model.User, error)
	UpdateByAdmin(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error)
	SoftDelete(ctx context.Context, userID string) error
	GetAll(ctx context.Context, page, size int) ([]*model.User, int, error)
}

// FraudRuleService - сервис для работы с правилами антифрода
// Из прошлого проекта с банком: DSL валидация критически важна для продакшена
type FraudRuleService interface {
	// CRUD операции
	Create(ctx context.Context, req model.FraudRuleCreateRequest) (*model.FraudRule, error)
	GetByID(ctx context.Context, id string) (*model.FraudRule, error)
	GetAll(ctx context.Context, activeOnly bool) ([]*model.FraudRule, error)
	Update(ctx context.Context, id string, req model.FraudRuleUpdateRequest) (*model.FraudRule, error)
	Delete(ctx context.Context, id string) error
	
	// Валидация DSL
	ValidateDSL(ctx context.Context, dsl string) (*model.FraudRuleValidateResponse, error)
}
