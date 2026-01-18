package v1

import (
	"context"
	"solution/internal/service"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
)

type handlerAdapter struct {
	antifraudService service.AntifraudService
}

func NewHandlerAdapter(antifraudService service.AntifraudService) antifraud_v1.Handler {
	return &handlerAdapter{
		antifraudService: antifraudService,
	}
}

func (h *handlerAdapter) PingGet(ctx context.Context) (*antifraud_v1.PingGetOK, error) {
	opt := antifraud_v1.OptString{
		Value: "ok",
	}
	return &antifraud_v1.PingGetOK{
		Status: opt,
	}, nil

}

func (h *handlerAdapter) AuthLoginPost(ctx context.Context, req *antifraud_v1.LoginRequest) (antifraud_v1.AuthLoginPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) AuthRegisterPost(ctx context.Context, req *antifraud_v1.RegisterRequest) (antifraud_v1.AuthRegisterPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) FraudRulesGet(ctx context.Context) (antifraud_v1.FraudRulesGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) FraudRulesIDDelete(ctx context.Context, params antifraud_v1.FraudRulesIDDeleteParams) (antifraud_v1.FraudRulesIDDeleteRes, error) {
	return nil, nil
}

func (h *handlerAdapter) FraudRulesIDGet(ctx context.Context, params antifraud_v1.FraudRulesIDGetParams) (antifraud_v1.FraudRulesIDGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) FraudRulesIDPut(ctx context.Context, req *antifraud_v1.FraudRuleUpdateRequest, params antifraud_v1.FraudRulesIDPutParams) (antifraud_v1.FraudRulesIDPutRes, error) {
	return nil, nil
}

func (h *handlerAdapter) FraudRulesPost(ctx context.Context, req *antifraud_v1.FraudRuleCreateRequest) (antifraud_v1.FraudRulesPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) FraudRulesValidatePost(ctx context.Context, req *antifraud_v1.DslValidateRequest) (antifraud_v1.FraudRulesValidatePostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) StatsMerchantsRiskGet(ctx context.Context, params antifraud_v1.StatsMerchantsRiskGetParams) (antifraud_v1.StatsMerchantsRiskGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) StatsOverviewGet(ctx context.Context, params antifraud_v1.StatsOverviewGetParams) (antifraud_v1.StatsOverviewGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) StatsRulesMatchesGet(ctx context.Context, params antifraud_v1.StatsRulesMatchesGetParams) (antifraud_v1.StatsRulesMatchesGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) StatsTransactionsTimeseriesGet(ctx context.Context, params antifraud_v1.StatsTransactionsTimeseriesGetParams) (antifraud_v1.StatsTransactionsTimeseriesGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) StatsUsersIDRiskProfileGet(ctx context.Context, params antifraud_v1.StatsUsersIDRiskProfileGetParams) (antifraud_v1.StatsUsersIDRiskProfileGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) TransactionsBatchPost(ctx context.Context, req *antifraud_v1.TransactionBatchCreateRequest) (antifraud_v1.TransactionsBatchPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) TransactionsGet(ctx context.Context, params antifraud_v1.TransactionsGetParams) (antifraud_v1.TransactionsGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) TransactionsIDGet(ctx context.Context, params antifraud_v1.TransactionsIDGetParams) (antifraud_v1.TransactionsIDGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) TransactionsPost(ctx context.Context, req *antifraud_v1.TransactionCreateRequest) (antifraud_v1.TransactionsPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) UsersGet(ctx context.Context, params antifraud_v1.UsersGetParams) (antifraud_v1.UsersGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) UsersIDDelete(ctx context.Context, params antifraud_v1.UsersIDDeleteParams) (antifraud_v1.UsersIDDeleteRes, error) {
	return nil, nil
}

func (h *handlerAdapter) UsersIDGet(ctx context.Context, params antifraud_v1.UsersIDGetParams) (antifraud_v1.UsersIDGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) UsersIDPut(ctx context.Context, req *antifraud_v1.UserUpdateRequest, params antifraud_v1.UsersIDPutParams) (antifraud_v1.UsersIDPutRes, error) {
	return nil, nil
}

func (h *handlerAdapter) UsersMeGet(ctx context.Context) (antifraud_v1.UsersMeGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) UsersMePut(ctx context.Context, req *antifraud_v1.UserUpdateRequest) (antifraud_v1.UsersMePutRes, error) {
	return nil, nil
}

func (h *handlerAdapter) UsersPost(ctx context.Context, req *antifraud_v1.UserCreateRequest) (antifraud_v1.UsersPostRes, error) {
	return nil, nil
}
