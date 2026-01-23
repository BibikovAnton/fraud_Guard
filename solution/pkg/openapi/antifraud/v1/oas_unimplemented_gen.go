
package antifraud_v1

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

func (UnimplementedHandler) APIV1AuthLoginPost(ctx context.Context, req *LoginRequest) (r APIV1AuthLoginPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1AuthRegisterPost(ctx context.Context, req *RegisterRequest) (r APIV1AuthRegisterPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1FraudRulesGet(ctx context.Context) (r APIV1FraudRulesGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1FraudRulesIDDelete(ctx context.Context, params APIV1FraudRulesIDDeleteParams) (r APIV1FraudRulesIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1FraudRulesIDGet(ctx context.Context, params APIV1FraudRulesIDGetParams) (r APIV1FraudRulesIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1FraudRulesIDPut(ctx context.Context, req *FraudRuleUpdateRequest, params APIV1FraudRulesIDPutParams) (r APIV1FraudRulesIDPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1FraudRulesPost(ctx context.Context, req *FraudRuleCreateRequest) (r APIV1FraudRulesPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1FraudRulesValidatePost(ctx context.Context, req *DslValidateRequest) (r APIV1FraudRulesValidatePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1PingGet(ctx context.Context) (r *APIV1PingGetOK, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1StatsMerchantsRiskGet(ctx context.Context, params APIV1StatsMerchantsRiskGetParams) (r APIV1StatsMerchantsRiskGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1StatsOverviewGet(ctx context.Context, params APIV1StatsOverviewGetParams) (r APIV1StatsOverviewGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1StatsRulesMatchesGet(ctx context.Context, params APIV1StatsRulesMatchesGetParams) (r APIV1StatsRulesMatchesGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1StatsTransactionsTimeseriesGet(ctx context.Context, params APIV1StatsTransactionsTimeseriesGetParams) (r APIV1StatsTransactionsTimeseriesGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1StatsUsersIDRiskProfileGet(ctx context.Context, params APIV1StatsUsersIDRiskProfileGetParams) (r APIV1StatsUsersIDRiskProfileGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1TransactionsBatchPost(ctx context.Context, req *TransactionBatchCreateRequest) (r APIV1TransactionsBatchPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1TransactionsGet(ctx context.Context, params APIV1TransactionsGetParams) (r APIV1TransactionsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1TransactionsIDGet(ctx context.Context, params APIV1TransactionsIDGetParams) (r APIV1TransactionsIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1TransactionsPost(ctx context.Context, req *TransactionCreateRequest) (r APIV1TransactionsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1UsersGet(ctx context.Context, params APIV1UsersGetParams) (r APIV1UsersGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1UsersIDDelete(ctx context.Context, params APIV1UsersIDDeleteParams) (r APIV1UsersIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1UsersIDGet(ctx context.Context, params APIV1UsersIDGetParams) (r APIV1UsersIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1UsersIDPut(ctx context.Context, req *UserUpdateRequest, params APIV1UsersIDPutParams) (r APIV1UsersIDPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1UsersMeGet(ctx context.Context) (r APIV1UsersMeGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1UsersMePut(ctx context.Context, req *UserUpdateRequest) (r APIV1UsersMePutRes, _ error) {
	return r, ht.ErrNotImplemented
}

func (UnimplementedHandler) APIV1UsersPost(ctx context.Context, req *UserCreateRequest) (r APIV1UsersPostRes, _ error) {
	return r, ht.ErrNotImplemented
}
