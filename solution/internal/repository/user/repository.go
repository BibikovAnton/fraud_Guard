package user

import (
	"context"
	"solution/internal/model"
	"solution/internal/repository"

	"github.com/jackc/pgx/v5"
)

var _ repository.UserRepository = (*repo)(nil)

type repo struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) GetByIDIncludingInactive(ctx context.Context, userID string) (*model.User, error) {
	return r.GetByID(ctx, userID)
}
