package v1

import (
	"context"
	"strings"

	"github.com/ogen-go/ogen/ogenerrors"
	"solution/internal/config"
	"solution/pkg/jwt"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
)

type SecurityHandler struct {
}

func NewSecurityHandlerAdapter() antifraud_v1.SecurityHandler {
	return &SecurityHandler{}
}

func (s *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName antifraud_v1.OperationName, t antifraud_v1.BearerAuth) (context.Context, error) {
	if t.Token == "" {
		return ctx, ogenerrors.ErrSkipServerSecurity
	}

	parts := strings.Split(t.Token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	token := parts[1]
	isValid, data, err := jwt.NewJWT(config.AppConfig().RandomSecret.RANDOM_SECRET()).Parse(token)
	if err != nil || !isValid || data == nil {
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	if data.UserID == "" {
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	if data.Role != "USER" && data.Role != "ADMIN" {
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	ctx = context.WithValue(ctx, ContextUserIDKey, data.UserID)
	ctx = context.WithValue(ctx, ContextRoleKey, data.Role)
	ctx = context.WithValue(ctx, ContextJWTDataKey, data)

	return ctx, nil
}

type ContextKey string

const (
	ContextUserIDKey  ContextKey = "user_id"
	ContextRoleKey    ContextKey = "user_role"
	ContextJWTDataKey ContextKey = "jwt_data"
)
