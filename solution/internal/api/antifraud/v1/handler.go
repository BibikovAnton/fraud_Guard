package v1

import (
	"context"
	"fmt"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
	"solution/internal/service"
)

const (
	ContextRoleKey   = "user_role"
	ContextUserIDKey = "user_id"
)

type handlerAdapter struct {
	userService       service.UserService
	fraudRuleService  service.FraudRuleService
	transactionService service.TransactionService
}

func NewHandlerAdapter(userService service.UserService, fraudRuleService service.FraudRuleService, transactionService service.TransactionService) antifraud_v1.Handler {
	return &handlerAdapter{
		userService:       userService,
		fraudRuleService:  fraudRuleService,
		transactionService: transactionService,
	}
}

func (h *handlerAdapter) APIV1PingGet(ctx context.Context) (*antifraud_v1.APIV1PingGetOK, error) {
	opt := antifraud_v1.OptString{
		Value: "ok",
	}
	return &antifraud_v1.APIV1PingGetOK{
		Status: opt,
	}, nil
}

func (h *handlerAdapter) APIV1TransactionsPost(ctx context.Context, req *antifraud_v1.TransactionCreateRequest) (antifraud_v1.APIV1TransactionsPostRes, error) {
	return nil, fmt.Errorf("not implemented yet")
}

func (h *handlerAdapter) APIV1TransactionsBatchPost(ctx context.Context, req *antifraud_v1.TransactionBatchCreateRequest) (antifraud_v1.APIV1TransactionsBatchPostRes, error) {
	return nil, fmt.Errorf("not implemented yet")
}

func (h *handlerAdapter) APIV1TransactionsGet(ctx context.Context, params antifraud_v1.APIV1TransactionsGetParams) (antifraud_v1.APIV1TransactionsGetRes, error) {
	return nil, fmt.Errorf("not implemented yet")
}

func (h *handlerAdapter) APIV1TransactionsIDGet(ctx context.Context, params antifraud_v1.APIV1TransactionsIDGetParams) (antifraud_v1.APIV1TransactionsIDGetRes, error) {
	return nil, fmt.Errorf("not implemented yet")
}
