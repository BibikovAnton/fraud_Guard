package transaction

import (
	"context"
	"solution/internal/model"
	"time"
)

// TransactionRepository - интерфейс репозитория транзакций
type TransactionRepository interface {
	// CRUD операции
	Create(ctx context.Context, transaction model.Transaction) error
	CreateBatch(ctx context.Context, transactions []model.Transaction) error
	GetByID(ctx context.Context, id string) (*model.Transaction, error)
	GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*model.Transaction, error)
	Update(ctx context.Context, id string, req model.TransactionUpdateRequest) (*model.Transaction, error)
	
	// Поиск и фильтрация
	GetFraudTransactions(ctx context.Context, limit, offset int) ([]*model.Transaction, error)
	GetByMerchantID(ctx context.Context, merchantID string, limit, offset int) ([]*model.Transaction, error)
	GetByTimeRange(ctx context.Context, start, end time.Time, limit, offset int) ([]*model.Transaction, error)
	
	// Статистика
	CountByStatus(ctx context.Context) (map[model.TransactionStatus]int64, error)
	CountFraudTransactions(ctx context.Context, timeRange time.Duration) (int64, error)
	GetTotalAmount(ctx context.Context, currency model.CurrencyCode) (float64, error)
}
