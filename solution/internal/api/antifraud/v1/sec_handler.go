package v1

import (
	"context"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
)

type SecurityHandler struct {
}

func NewSecurityHandlerAdapter() antifraud_v1.SecurityHandler {
	return &SecurityHandler{}
}

// Реализация
func (s *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName antifraud_v1.OperationName, t antifraud_v1.BearerAuth) (context.Context, error) {
	return nil, nil
}
