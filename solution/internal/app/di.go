package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"solution/internal/config"
	"solution/internal/repository"
	repositoryFraudRules "solution/internal/repository/fraud_rules"
	repositoryStats "solution/internal/repository/stats"
	repositoryTransactions "solution/internal/repository/transactions"
	repositoryUser "solution/internal/repository/user"
	"solution/internal/service"
	serviceDSL "solution/internal/service/dsl"
	serviceFraudRules "solution/internal/service/fraud_rules"
	serviceStats "solution/internal/service/stats"
	serviceTransactions "solution/internal/service/transactions"
	serviceUser "solution/internal/service/user"
	"solution/platform/pkg/closer"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
	"solution/internal/api/antifraud/v1"
)

type diContainer struct {
	userService       service.UserService
	fraudRuleService  service.FraudRuleService
	transactionService service.TransactionService
	statsService       serviceStats.Service
	dslEvaluator      serviceDSL.Evaluator

	userRepository        repository.UserRepository
	fraudRuleRepository   repositoryFraudRules.Repository
	statsRepository       repositoryStats.Repository
	transactionRepository repositoryTransactions.Repository

	postgresConn     *pgx.Conn
	postgresDBHandle *sql.DB
}

func NewDIContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) V1API(ctx context.Context) antifraud_v1.Handler {
	return v1.NewHandlerAdapter(d.UserService(ctx), d.FraudRuleService(ctx), d.TransactionService(ctx), d.StatsService(ctx))
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

func (d *diContainer) TransactionService(ctx context.Context) service.TransactionService {
	if d.transactionService == nil {
		d.transactionService = serviceTransactions.NewService(
			d.TransactionRepository(ctx),
			d.UserRepository(ctx),
			d.FraudRuleRepository(ctx),
			d.DSLEvaluator(ctx),
		)
	}
	return d.transactionService
}

func (d *diContainer) DSLEvaluator(ctx context.Context) serviceDSL.Evaluator {
	if d.dslEvaluator == nil {
		d.dslEvaluator = serviceDSL.NewEvaluator(2) // Возвращаем tier 2 - поддерживает AND/OR
	}
	return d.dslEvaluator
}

func (d *diContainer) UserRepository(ctx context.Context) repository.UserRepository {
	if d.userRepository == nil {
		d.userRepository = repositoryUser.NewRepository(d.PostgresDBClient(ctx))
	}
	return d.userRepository
}

func (d *diContainer) FraudRuleRepository(ctx context.Context) repositoryFraudRules.Repository {
	if d.fraudRuleRepository == nil {
		d.fraudRuleRepository = repositoryFraudRules.NewRepository(d.PostgresDBClient(ctx))
	}
	return d.fraudRuleRepository
}

func (d *diContainer) TransactionRepository(ctx context.Context) repositoryTransactions.Repository {
	if d.transactionRepository == nil {
		d.transactionRepository = repositoryTransactions.NewRepository(d.PostgresDBClient(ctx))
	}
	return d.transactionRepository
}

func (d *diContainer) StatsRepository(ctx context.Context) repositoryStats.Repository {
	if d.statsRepository == nil {
		d.statsRepository = repositoryStats.NewRepository(d.PostgresDBClient(ctx))
	}
	return d.statsRepository
}

func (d *diContainer) StatsService(ctx context.Context) serviceStats.Service {
	if d.statsService == nil {
		d.statsService = serviceStats.NewService(d.StatsRepository(ctx))
	}
	return d.statsService
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
