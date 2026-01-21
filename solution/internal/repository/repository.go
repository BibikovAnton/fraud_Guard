package repository

import (
	"context"
	"solution/internal/model"
)


type AntifraudRepository interface {
	
}


type UserRepository interface {
	Create(ctx context.Context, user model.User) error
	CreateAdmin(ctx context.Context, email, passwordHash, fullName string) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	FindByEmailIncludingInactive(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, userID string) (*model.User, error)
	GetByIDIncludingInactive(ctx context.Context, userID string) (*model.User, error)

	
	Update(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error)
	UpdateByAdmin(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error)
	SoftDelete(ctx context.Context, userID string) error


	GetAll(ctx context.Context, page, size int) ([]*model.User, int, error)
}


type FraudRuleRepository interface {
	
	Create(ctx context.Context, rule model.FraudRule) error
	GetByID(ctx context.Context, id string) (*model.FraudRule, error)
	GetByName(ctx context.Context, name string) (*model.FraudRule, error)
	GetAll(ctx context.Context, activeOnly bool) ([]*model.FraudRule, error)
	Update(ctx context.Context, id string, req model.FraudRuleUpdateRequest) (*model.FraudRule, error)
	Delete(ctx context.Context, id string) error

	
	ExistsByName(ctx context.Context, name string, excludeID string) (bool, error)
	GetActiveRulesCount(ctx context.Context) (int, error)
}
