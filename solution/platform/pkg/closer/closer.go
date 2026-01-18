package closer

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"solution/platform/pkg/logger"
	"sync"
	"time"

	"go.uber.org/zap"
)

const shutdownTimeout = 5 * time.Second

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
}

type Closer struct {
	mu     sync.Mutex
	once   sync.Once
	done   chan struct{}
	funcs  []func(context.Context) error
	logger Logger
}

var globalCloser = NewWithLogger(&logger.NoopLogger{})

func AddNamed(name string, f func(context.Context) error) {
	globalCloser.AddNamed(name, f)
}

func Add(f ...func(context.Context) error) {
	globalCloser.Add(f...)
}

func CloseAll(ctx context.Context) error {
	return globalCloser.CloseAll(ctx)
}

func SetLogger(l Logger) {
	globalCloser.SetLogger(l)
}

func Configure(signals ...os.Signal) {
	go globalCloser.handleSignals(signals...)
}

func New(signals ...os.Signal) *Closer {
	return NewWithLogger(logger.Logger(), signals...)
}

func NewWithLogger(logger Logger, signals ...os.Signal) *Closer {
	c := &Closer{
		done:   make(chan struct{}),
		logger: logger,
	}

	if len(signals) > 0 {
		go c.handleSignals(signals...)
	}

	return c
}

func (c *Closer) SetLogger(l Logger) {
	c.logger = l
}

func (c *Closer) handleSignals(signals ...os.Signal) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	defer signal.Stop(ch)

	select {
	case <-ch:
		c.logger.Info(context.Background(), "🛑 Получен системный сигнал, начинаем graceful shutdown...")

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer shutdownCancel()

		if err := c.CloseAll(shutdownCtx); err != nil {
			c.logger.Error(context.Background(), "❌ Ошибка при закрытии ресурсов: %v", zap.Error(err))
		}

	case <-c.done:

	}
}

func (c *Closer) AddNamed(name string, f func(context.Context) error) {
	c.Add(func(ctx context.Context) error {
		start := time.Now()
		c.logger.Info(ctx, fmt.Sprintf("🧩 Закрываем %s...", name))

		err := f(ctx)

		duration := time.Since(start)
		if err != nil {
			c.logger.Error(ctx, fmt.Sprintf("❌ Ошибка при закрытии %s: %v (заняло %s)", name, err, duration))
		} else {
			c.logger.Info(ctx, fmt.Sprintf("✅ %s успешно закрыт за %s", name, duration))
		}
		return err
	})
}

func (c *Closer) Add(f ...func(context.Context) error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.funcs = append(c.funcs, f...)
}

func (c *Closer) CloseAll(ctx context.Context) error {
	var result error

	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()
		funcs := c.funcs
		c.funcs = nil
		c.mu.Unlock()

		if len(funcs) == 0 {
			c.logger.Info(ctx, "ℹ️ Нет функций для закрытия.")
			return
		}

		c.logger.Info(ctx, "🚦 Начинаем процесс graceful shutdown...")

		errCh := make(chan error, len(funcs))
		var wg sync.WaitGroup

		for i := len(funcs) - 1; i >= 0; i-- {
			f := funcs[i]
			wg.Add(1)
			go func(f func(context.Context) error) {
				defer wg.Done()

				defer func() {
					if r := recover(); r != nil {
						errCh <- errors.New("panic recovered in closer")
						c.logger.Error(ctx, "⚠️ Panic в функции закрытия", zap.Any("error", r))
					}
				}()

				if err := f(ctx); err != nil {
					errCh <- err
				}
			}(f)
		}

		go func() {
			wg.Wait()
			close(errCh)
		}()

		for {
			select {
			case <-ctx.Done():
				c.logger.Info(ctx, "⚠️ Контекст отменён во время закрытия", zap.Error(ctx.Err()))
				if result == nil {
					result = ctx.Err()
				}
				return
			case err, ok := <-errCh:
				if !ok {
					c.logger.Info(ctx, "✅ Все ресурсы успешно закрыты")
					return
				}
				c.logger.Error(ctx, "❌ Ошибка при закрытии", zap.Error(err))
				if result == nil {
					result = err
				}
			}
		}
	})

	return result
}
