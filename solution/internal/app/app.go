package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-faster/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	v1 "solution/internal/api/antifraud/v1"
	"solution/internal/config"
	"solution/internal/migrator"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
	"solution/platform/pkg/closer"
	"solution/platform/pkg/logger"
	"strings"
)

type App struct {
	diContainer *diContainer
	httpServer  *http.Server
	listener    net.Listener
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
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
		a.createAdminUser,
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

	handlerAdapter := v1.NewHandlerAdapter(a.diContainer.UserService(ctx), a.diContainer.FraudRuleService(ctx), a.diContainer.TransactionService(ctx), a.diContainer.StatsService(ctx))
	secHandlerAdapter := v1.NewSecurityHandlerAdapter()

	transactionHandler := v1.NewTransactionHandler(a.diContainer.UserService(ctx), a.diContainer.TransactionService(ctx))

	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(os.Stderr, "=== HTTP REQUEST ===\n")
			fmt.Fprintf(os.Stderr, "METHOD: %s\n", r.Method)
			fmt.Fprintf(os.Stderr, "PATH: %s\n", r.URL.Path)
			fmt.Fprintf(os.Stderr, "QUERY: %s\n", r.URL.RawQuery)
			fmt.Fprintf(os.Stderr, "HEADERS: %+v\n", r.Header)
			fmt.Fprintf(os.Stderr, "=== END REQUEST ===\n")
			
			next.ServeHTTP(w, r)
		})
	}

	antifraudServer, err := antifraud_v1.NewServer(handlerAdapter, secHandlerAdapter)
	if err != nil {
		logger.Error(ctx, "Error creating OpenAPI antifraudServer", zap.Error(err))
		return err
	}

	logger.Info(ctx, "Creating custom HTTP handlers without OpenAPI validation")
	
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(loggingMiddleware)
		
		r.Post("/transactions", transactionHandler.CreateTransaction)
		r.Post("/transactions/batch", transactionHandler.CreateBatchTransactions)
		r.Get("/transactions/{id}", transactionHandler.GetTransaction)
		r.Get("/transactions", transactionHandler.GetTransactions)
		
		r.Mount("/", antifraudServer)
	})

	finalHandler := loggingMiddleware(r)

	a.httpServer = &http.Server{
		Addr:              ":" + config.AppConfig().Http.Address(),
		Handler:           finalHandler,
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

func (a *App) createAdminUser(ctx context.Context) error {
	userService := a.diContainer.UserService(ctx)

	email := config.AppConfig().Admin.ADMIN_EMAIL()
	fullName := config.AppConfig().Admin.ADMIN_FULLNAME()
	password := config.AppConfig().Admin.ADMIN_PASSWORD()

	if email == "" || fullName == "" || password == "" {
		logger.Warn(ctx, "Admin credentials not fully configured, skipping admin creation")
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(ctx, "Failed to hash admin password", zap.Error(err))
		return err
	}

	err = userService.CreateAdmin(ctx, email, string(hashedPassword), fullName)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			logger.Info(ctx, "Admin user already exists", zap.String("email", email))
			return nil
		}

		logger.Error(ctx, "Failed to create admin user", zap.Error(err))
		return err
	}

	logger.Info(ctx, "Admin user created successfully", zap.String("email", email))
	return nil
}
