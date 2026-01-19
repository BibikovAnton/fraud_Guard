package transaction

import (
	"solution/internal/repository"

	"github.com/jackc/pgx/v5"
)

var _ repository.TransactionRepository = (*repo)(nil)

type repo struct {
	db *pgx.Conn
}

// NewRepository создает новый репозиторий транзакций
func NewRepository(db *pgx.Conn) repository.TransactionRepository {
	return &repo{db: db}
}
