package v1

import (
	"context"

	"github.com/ogen-go/ogen/ogenerrors"
	"solution/internal/config"
	"solution/pkg/jwt"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
)

type SecurityHandler struct{}

func NewSecurityHandlerAdapter() antifraud_v1.SecurityHandler {
	return &SecurityHandler{}
}

func (s *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName antifraud_v1.OperationName, t antifraud_v1.BearerAuth) (context.Context, error) {

	if t.Token == "" {
		return ctx, ogenerrors.ErrSkipServerSecurity
	}

	jwtValidator := jwt.NewJWT(config.AppConfig().RandomSecret.RANDOM_SECRET())
	isValid, jwtData, parseErr := jwtValidator.Parse(t.Token)

	if parseErr != nil || !isValid || jwtData == nil {
		// Логируем неудачные попытки для мониторинга атак
		// TODO: добавить метрики для rate limiting по IP
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	if jwtData.UserID == "" {
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	if jwtData.Role != "USER" && jwtData.Role != "ADMIN" {
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	authCtx := context.WithValue(ctx, ContextUserIDKey, jwtData.UserID)
	authCtx = context.WithValue(authCtx, ContextRoleKey, jwtData.Role)
	authCtx = context.WithValue(authCtx, ContextJWTDataKey, jwtData)

	return authCtx, nil
}

type ContextKey string

const (
	ContextUserIDKey  ContextKey = "user_id"
	ContextRoleKey    ContextKey = "user_role"
	ContextJWTDataKey ContextKey = "jwt_data"
)
