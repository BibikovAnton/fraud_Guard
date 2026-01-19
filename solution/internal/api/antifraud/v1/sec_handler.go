package v1

import (
	"context"
	"solution/internal/config"
	"solution/internal/middleware"
	"solution/pkg/jwt"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
)

type SecurityHandler struct {
}

func NewSecurityHandlerAdapter() antifraud_v1.SecurityHandler {
	return &SecurityHandler{}
}

// Реализация
func (s *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName antifraud_v1.OperationName, t antifraud_v1.BearerAuth) (context.Context, error) {
	// Валидация JWT токена
	isValid, data, err := jwt.NewJWT(config.AppConfig().RandomSecret.RANDOM_SECRET()).Parse(t.Token)
	if err != nil {
		return ctx, err
	}

	if !isValid || data == nil {
		return ctx, nil
	}

	if data.UserID == "" {
		return ctx, nil
	}

	if data.Role != "USER" && data.Role != "ADMIN" {
		return ctx, nil
	}

	// Устанавливаем данные пользователя в контекст
	ctx = context.WithValue(ctx, middleware.ContextUserIDKey, data.UserID)
	ctx = context.WithValue(ctx, middleware.ContextRoleKey, data.Role)
	ctx = context.WithValue(ctx, middleware.ContextJWTDataKey, data)

	return ctx, nil
}
