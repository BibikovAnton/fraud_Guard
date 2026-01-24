package main

import (
	"context"
	"go.uber.org/zap"
	"os/signal"
	"solution/internal/app"
	"solution/internal/config"
	"solution/platform/pkg/closer"
	"solution/platform/pkg/logger"
	"syscall"
	"time"
)

func main() {
	// CI build fix - force new image build
	if err := config.Load(); err != nil {
		panic(err)
	}

	appCtx, appCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer appCancel()
	defer gracefulShutdown()

	closer.Configure(syscall.SIGINT, syscall.SIGTERM)

	a, err := app.New(appCtx)
	if err != nil {
		logger.Error(appCtx, "Failed to create application", zap.Error(err))
		return
	}

	err = a.Run(appCtx)
	if err != nil {
		logger.Error(appCtx, "Application Error", zap.Error(err))
		return
	}
}

func gracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := closer.CloseAll(ctx); err != nil {
		logger.Error(ctx, "Shutdown error", zap.Error(err))
	}
}
