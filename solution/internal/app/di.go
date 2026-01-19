package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"solution/internal/api/antifraud/v1"
	"solution/internal/config"
	"solution/internal/repository"
	repositoryAntifraud "solution/internal/repository/antifraud"
	repositoryFraudRules "solution/internal/repository/fraud_rules"
	repositoryUser "solution/internal/repository/user"
	"solution/internal/service"
	serviceAntifraud "solution/internal/service/antifraurd"
	serviceFraudRules "solution/internal/service/fraud_rules"
	serviceUser "solution/internal/service/user"
	"solution/platform/pkg/closer"
)

type diContainer struct {
	antifraudService    service.AntifraudService
	userService        service.UserService
	fraudRuleService   service.FraudRuleService
	
	antifraudRepository repository.AntifraudRepository
	userRepository      repository.UserRepository
	fraudRuleRepository repository.FraudRuleRepository

	postgresConn     *pgx.Conn
	postgresDBHandle *sql.DB
}

func NewDIContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) AntifraudV1API(ctx context.Context) {
	v1.NewAPI(d.AntifraudService(ctx))
}

func (d *diContainer) AntifraudService(ctx context.Context) service.AntifraudService {
	if d.antifraudService == nil {
		d.antifraudService = serviceAntifraud.NewService(d.AntifraudRepository(ctx))
	}
	return d.antifraudService
}

func (d *diContainer) UserService(ctx context.Context) service.UserService {
	if d.userService == nil {
		d.userService = serviceUser.NewUserService(d.UserRepository(ctx), config.AppConfig().RandomSecret.RANDOM_SECRET())
	}
	return d.userService
}

func (d *diContainer) FraudRuleService(ctx context.Context) service.FraudRuleService {
	if d.fraudRuleService == nil {
		d.fraudRuleService = serviceFraudRules.NewService(d.FraudRuleRepository(ctx))
	}
	return d.fraudRuleService
}

func (d *diContainer) AntifraudRepository(ctx context.Context) repository.AntifraudRepository {
	if d.antifraudRepository == nil {
		d.antifraudRepository = repositoryAntifraud.NewRepository(d.PostgresDBHandle(ctx))
	}
	return d.antifraudRepository
}

func (d *diContainer) UserRepository(ctx context.Context) repository.UserRepository {
	if d.userRepository == nil {
		d.userRepository = repositoryUser.NewRepository(d.PostgresDBClient(ctx))
	}
	return d.userRepository
}

func (d *diContainer) FraudRuleRepository(ctx context.Context) repository.FraudRuleRepository {
	if d.fraudRuleRepository == nil {
		d.fraudRuleRepository = repositoryFraudRules.NewRepository(d.PostgresDBClient(ctx))
	}
	return d.fraudRuleRepository
}

func (d *diContainer) PostgresDBClient(ctx context.Context) *pgx.Conn {

	if d.postgresConn == nil {
		client, err := pgx.Connect(ctx, config.AppConfig().Database.URI()) //собрать URI
		if err != nil {
			panic(fmt.Sprintf("failed to connect to PostgresDB: %s\n", err.Error()))
		}

		err = client.Ping(ctx)
		if err != nil {
			panic(fmt.Sprintf("failed to ping PostgresDB: %v\n", err))
		}

		closer.AddNamed("PostgresDB client", func(ctx context.Context) error {
			return client.Close(ctx)

		})

		d.postgresConn = client
	}
	return d.postgresConn
}

func (d *diContainer) PostgresDBHandle(ctx context.Context) *sql.DB {
	if d.postgresDBHandle == nil {
		d.postgresDBHandle = stdlib.OpenDB(*d.PostgresDBClient(ctx).Config().Copy())
	}
	return d.postgresDBHandle
}
