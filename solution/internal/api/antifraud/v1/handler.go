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

func (h *handlerAdapter) APIV1PingGet(ctx context.Context) (*antifraud_v1.APIV1PingGetOK, error) {
	opt := antifraud_v1.OptString{
		Value: "ok",
	}
	return &antifraud_v1.APIV1PingGetOK{
		Status: opt,
	}, nil

}

func (h *handlerAdapter) APIV1AuthLoginPost(ctx context.Context, req *antifraud_v1.LoginRequest) (antifraud_v1.APIV1AuthLoginPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1AuthRegisterPost(ctx context.Context, req *antifraud_v1.RegisterRequest) (antifraud_v1.APIV1AuthRegisterPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1FraudRulesGet(ctx context.Context) (antifraud_v1.APIV1FraudRulesGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1FraudRulesIDDelete(ctx context.Context, params antifraud_v1.APIV1FraudRulesIDDeleteParams) (antifraud_v1.APIV1FraudRulesIDDeleteRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1FraudRulesIDGet(ctx context.Context, params antifraud_v1.APIV1FraudRulesIDGetParams) (antifraud_v1.APIV1FraudRulesIDGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1FraudRulesIDPut(ctx context.Context, req *antifraud_v1.FraudRuleUpdateRequest, params antifraud_v1.APIV1FraudRulesIDPutParams) (antifraud_v1.APIV1FraudRulesIDPutRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1FraudRulesPost(ctx context.Context, req *antifraud_v1.FraudRuleCreateRequest) (antifraud_v1.APIV1FraudRulesPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1FraudRulesValidatePost(ctx context.Context, req *antifraud_v1.DslValidateRequest) (antifraud_v1.APIV1FraudRulesValidatePostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1StatsMerchantsRiskGet(ctx context.Context, params antifraud_v1.APIV1StatsMerchantsRiskGetParams) (antifraud_v1.APIV1StatsMerchantsRiskGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1StatsOverviewGet(ctx context.Context, params antifraud_v1.APIV1StatsOverviewGetParams) (antifraud_v1.APIV1StatsOverviewGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1StatsRulesMatchesGet(ctx context.Context, params antifraud_v1.APIV1StatsRulesMatchesGetParams) (antifraud_v1.APIV1StatsRulesMatchesGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1StatsTransactionsTimeseriesGet(ctx context.Context, params antifraud_v1.APIV1StatsTransactionsTimeseriesGetParams) (antifraud_v1.APIV1StatsTransactionsTimeseriesGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1StatsUsersIDRiskProfileGet(ctx context.Context, params antifraud_v1.APIV1StatsUsersIDRiskProfileGetParams) (antifraud_v1.APIV1StatsUsersIDRiskProfileGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1TransactionsBatchPost(ctx context.Context, req *antifraud_v1.TransactionBatchCreateRequest) (antifraud_v1.APIV1TransactionsBatchPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1TransactionsGet(ctx context.Context, params antifraud_v1.APIV1TransactionsGetParams) (antifraud_v1.APIV1TransactionsGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1TransactionsIDGet(ctx context.Context, params antifraud_v1.APIV1TransactionsIDGetParams) (antifraud_v1.APIV1TransactionsIDGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1TransactionsPost(ctx context.Context, req *antifraud_v1.TransactionCreateRequest) (antifraud_v1.APIV1TransactionsPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersGet(ctx context.Context, params antifraud_v1.APIV1UsersGetParams) (antifraud_v1.APIV1UsersGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersIDDelete(ctx context.Context, params antifraud_v1.APIV1UsersIDDeleteParams) (antifraud_v1.APIV1UsersIDDeleteRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersIDGet(ctx context.Context, params antifraud_v1.APIV1UsersIDGetParams) (antifraud_v1.APIV1UsersIDGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersIDPut(ctx context.Context, req *antifraud_v1.UserUpdateRequest, params antifraud_v1.APIV1UsersIDPutParams) (antifraud_v1.APIV1UsersIDPutRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersMeGet(ctx context.Context) (antifraud_v1.APIV1UsersMeGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersMePut(ctx context.Context, req *antifraud_v1.UserUpdateRequest) (antifraud_v1.APIV1UsersMePutRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersPost(ctx context.Context, req *antifraud_v1.UserCreateRequest) (antifraud_v1.APIV1UsersPostRes, error) {
	return nil, nil
}
