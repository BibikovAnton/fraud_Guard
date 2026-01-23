
package antifraud_v1

import (
	"context"
)

type Handler interface {
	APIV1AuthLoginPost(ctx context.Context, req *LoginRequest) (APIV1AuthLoginPostRes, error)
	APIV1AuthRegisterPost(ctx context.Context, req *RegisterRequest) (APIV1AuthRegisterPostRes, error)
	APIV1FraudRulesGet(ctx context.Context) (APIV1FraudRulesGetRes, error)
	APIV1FraudRulesIDDelete(ctx context.Context, params APIV1FraudRulesIDDeleteParams) (APIV1FraudRulesIDDeleteRes, error)
	APIV1FraudRulesIDGet(ctx context.Context, params APIV1FraudRulesIDGetParams) (APIV1FraudRulesIDGetRes, error)
	APIV1FraudRulesIDPut(ctx context.Context, req *FraudRuleUpdateRequest, params APIV1FraudRulesIDPutParams) (APIV1FraudRulesIDPutRes, error)
	APIV1FraudRulesPost(ctx context.Context, req *FraudRuleCreateRequest) (APIV1FraudRulesPostRes, error)
	APIV1FraudRulesValidatePost(ctx context.Context, req *DslValidateRequest) (APIV1FraudRulesValidatePostRes, error)
	APIV1PingGet(ctx context.Context) (*APIV1PingGetOK, error)
	APIV1StatsMerchantsRiskGet(ctx context.Context, params APIV1StatsMerchantsRiskGetParams) (APIV1StatsMerchantsRiskGetRes, error)
	APIV1StatsOverviewGet(ctx context.Context, params APIV1StatsOverviewGetParams) (APIV1StatsOverviewGetRes, error)
	APIV1StatsRulesMatchesGet(ctx context.Context, params APIV1StatsRulesMatchesGetParams) (APIV1StatsRulesMatchesGetRes, error)
	APIV1StatsTransactionsTimeseriesGet(ctx context.Context, params APIV1StatsTransactionsTimeseriesGetParams) (APIV1StatsTransactionsTimeseriesGetRes, error)
	APIV1StatsUsersIDRiskProfileGet(ctx context.Context, params APIV1StatsUsersIDRiskProfileGetParams) (APIV1StatsUsersIDRiskProfileGetRes, error)
	APIV1TransactionsBatchPost(ctx context.Context, req *TransactionBatchCreateRequest) (APIV1TransactionsBatchPostRes, error)
	APIV1TransactionsGet(ctx context.Context, params APIV1TransactionsGetParams) (APIV1TransactionsGetRes, error)
	APIV1TransactionsIDGet(ctx context.Context, params APIV1TransactionsIDGetParams) (APIV1TransactionsIDGetRes, error)
	APIV1TransactionsPost(ctx context.Context, req *TransactionCreateRequest) (APIV1TransactionsPostRes, error)
	APIV1UsersGet(ctx context.Context, params APIV1UsersGetParams) (APIV1UsersGetRes, error)
	APIV1UsersIDDelete(ctx context.Context, params APIV1UsersIDDeleteParams) (APIV1UsersIDDeleteRes, error)
	APIV1UsersIDGet(ctx context.Context, params APIV1UsersIDGetParams) (APIV1UsersIDGetRes, error)
	APIV1UsersIDPut(ctx context.Context, req *UserUpdateRequest, params APIV1UsersIDPutParams) (APIV1UsersIDPutRes, error)
	APIV1UsersMeGet(ctx context.Context) (APIV1UsersMeGetRes, error)
	APIV1UsersMePut(ctx context.Context, req *UserUpdateRequest) (APIV1UsersMePutRes, error)
	APIV1UsersPost(ctx context.Context, req *UserCreateRequest) (APIV1UsersPostRes, error)
}

type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
