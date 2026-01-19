package fraud_rules

import (
	"context"
	"solution/internal/model"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Create(ctx context.Context, rule model.FraudRule) error
	GetByID(ctx context.Context, id string) (*model.FraudRule, error)
	GetByName(ctx context.Context, name string) (*model.FraudRule, error)
	GetAll(ctx context.Context, activeOnly bool) ([]*model.FraudRule, error)
	Update(ctx context.Context, id string, req model.FraudRuleUpdateRequest) (*model.FraudRule, error)
	Delete(ctx context.Context, id string) error
	
	ExistsByName(ctx context.Context, name string, excludeID string) (bool, error)
	GetActiveRulesCount(ctx context.Context) (int, error)
}

type repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) Repository {
	return &repository{
		db: db,
	}
}
