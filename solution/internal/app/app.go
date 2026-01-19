package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-faster/errors"
	"go.uber.org/zap"
	"net"
	"net/http"
	v1 "solution/internal/api/antifraud/v1"
	"solution/internal/config"
	"solution/internal/migrator"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
	"solution/platform/pkg/closer"
	"solution/platform/pkg/logger"
	"time"
)

type App struct {
	diContainer *diContainer
	httpServer  *http.Server
	listener    net.Listener
}

func New(ctx context.Context) (*App, error) {
	a := &App{}
	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	if err := a.runHTTPServer(ctx); err != nil {
		return err
	}
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.initLogger,
		a.initCloser,
		a.initListener,
		a.initHTTPServer,
		a.runMigrations,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initDI(_ context.Context) error {
	a.diContainer = NewDIContainer()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		config.AppConfig().Logger.Level(),
		config.AppConfig().Logger.AsJson(),
	)
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) initListener(_ context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.AppConfig().Http.Address()))
	if err != nil {
		return err
	}

	closer.AddNamed("TCP listener", func(ctx context.Context) error {
		lerr := lis.Close()
		if lerr != nil && !errors.Is(lerr, net.ErrClosed) {
			return lerr
		}
		return nil
	})

	a.listener = lis
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	handlerAdapter := v1.NewHandlerAdapter(a.diContainer.AntifraudService(ctx), a.diContainer.UserService(ctx), a.diContainer.FraudRuleService(ctx))
	secHandlerAdapter := v1.NewSecurityHandlerAdapter()

	antifraudServer, err := antifraud_v1.NewServer(handlerAdapter, secHandlerAdapter)
	if err != nil {
		logger.Error(ctx, "Error creating OpenAPI antifraudServer", zap.Error(err))
		return err
	}

	logger.Info(ctx, "OpenAPI server created successfully", zap.Any("server", antifraudServer != nil))

	if antifraudServer != nil {
		logger.Info(ctx, "Mounting OpenAPI server on /api/v1")
		r.Mount("/api/v1", antifraudServer)

	}

	a.httpServer = &http.Server{
		Addr:              ":" + config.AppConfig().Http.Address(),
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
	}

	closer.AddNamed("HTTP server", func(ctx context.Context) error {
		if err := a.httpServer.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	return nil
}

func (a *App) runMigrations(ctx context.Context) error {
	migratorRunner := migrator.NewMigrator(a.diContainer.PostgresDBHandle(ctx), config.AppConfig().Database.MigrationsPath())
	err := migratorRunner.Up()
	if err != nil {
		logger.Error(ctx, "Database migration error", zap.Error(err))
		return err
	}
	logger.Info(ctx, "Database migrations completed successfully")
	return nil
}

func (a *App) runHTTPServer(ctx context.Context) error {
	logger.Info(ctx, "Starting AntiFraud HTTP server", zap.String("port", config.AppConfig().Http.Address()))

	err := a.httpServer.Serve(a.listener)
	if err != nil && !errors.Is(http.ErrServerClosed, err) {
		logger.Error(ctx, "Failed to serve HTTP", zap.Error(err))
		return err
	}

	return nil
}
