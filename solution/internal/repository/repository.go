package repository

import (
	"context"
	"solution/internal/model"
)

// AntifraudRepository - репозиторий для работы с антифродом
// TODO: добавить методы для правил и транзакций
type AntifraudRepository interface {
	// TODO: Добавить методы для работы с антифродом
}

// UserRepository - репозиторий для работы с пользователями
type UserRepository interface {
	// Базовые операции
	Create(ctx context.Context, user model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	FindByEmailIncludingInactive(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, userID string) (*model.User, error)
	GetByIDIncludingInactive(ctx context.Context, userID string) (*model.User, error)

	// Обновление и удаление
	Update(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error)
	UpdateByAdmin(ctx context.Context, userID string, req model.UserUpdateRequest) (*model.User, error)
	SoftDelete(ctx context.Context, userID string) error

	// Получение пользователей
	GetAll(ctx context.Context, page, size int) ([]*model.User, int, error)
}

// FraudRuleRepository - репозиторий для работы с правилами антифрода
type FraudRuleRepository interface {
	// CRUD операции
	Create(ctx context.Context, rule model.FraudRule) error
	GetByID(ctx context.Context, id string) (*model.FraudRule, error)
	GetByName(ctx context.Context, name string) (*model.FraudRule, error)
	GetAll(ctx context.Context, activeOnly bool) ([]*model.FraudRule, error)
	Update(ctx context.Context, id string, req model.FraudRuleUpdateRequest) (*model.FraudRule, error)
	Delete(ctx context.Context, id string) error

	// Вспомогательные методы
	ExistsByName(ctx context.Context, name string, excludeID string) (bool, error)
	GetActiveRulesCount(ctx context.Context) (int, error)
}
