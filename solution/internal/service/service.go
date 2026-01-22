package service

import (
	"context"
	"solution/internal/model"
)

type UserService interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.RegisterResponse, error)
	Login(ctx context.Context, req model.LoginRequest) (*model.LoginResponse, error)
	CreateAdmin(ctx context.Context, email, passwordHash, fullName string) error

	GetMe(ctx context.Context, userID string) (*model.User, error)
	UpdateMe(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error)

	CreateByAdmin(ctx context.Context, req model.UserCreateRequest) (*model.User, error)
	GetByID(ctx context.Context, userID string) (*model.User, error)
	GetByIDIncludingInactive(ctx context.Context, userID string) (*model.User, error)
	UpdateByAdmin(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error)
	SoftDelete(ctx context.Context, userID string) error
	GetAll(ctx context.Context, page, size int) ([]*model.User, int, error)
}

type FraudRuleService interface {
	Create(ctx context.Context, req model.FraudRuleCreateRequest) (*model.FraudRule, error)
	GetByID(ctx context.Context, id string) (*model.FraudRule, error)
	GetAll(ctx context.Context, activeOnly bool) ([]*model.FraudRule, error)
	Update(ctx context.Context, id string, req model.FraudRuleUpdateRequest) (*model.FraudRule, error)
	Delete(ctx context.Context, id string) error
	ValidateDSL(ctx context.Context, dsl string) (*model.FraudRuleValidateResponse, error)
}
