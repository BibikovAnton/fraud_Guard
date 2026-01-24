package di

import (
	"context"

	"github.com/jackc/pgx/v5"
	"solution/internal/repository/stats"
	"solution/internal/service/stats"
)

func (d *diContainer) StatsRepository(ctx context.Context) stats.Repository {
	if d.statsRepo == nil {
		db := d.DB(ctx)
		d.statsRepo = stats.NewRepository(db)
	}
	return d.statsRepo
}

func (d *diContainer) StatsService(ctx context.Context) stats.Service {
	if d.statsService == nil {
		statsRepo := d.StatsRepository(ctx)
		d.statsService = stats.NewService(statsRepo)
	}
	return d.statsService
}
