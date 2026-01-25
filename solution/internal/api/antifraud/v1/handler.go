package v1

import (
	"context"
	"fmt"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"net/netip"
	"solution/internal/model"
	"solution/internal/service"
	"solution/internal/service/stats"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
	"strings"
	"time"
)

const (
	ContextRoleKey   = "user_role"
	ContextUserIDKey = "user_id"
)

type handlerAdapter struct {
	userService        service.UserService
	fraudRuleService   service.FraudRuleService
	transactionService service.TransactionService
	statsHandler       *statsHandlerAdapter
}

func NewHandlerAdapter(userService service.UserService, fraudRuleService service.FraudRuleService, transactionService service.TransactionService, statsService stats.Service) antifraud_v1.Handler {
	return &handlerAdapter{
		userService:        userService,
		fraudRuleService:   fraudRuleService,
		transactionService: transactionService,
		statsHandler:       NewStatsHandlerAdapter(statsService),
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

func (h *handlerAdapter) APIV1StatsOverviewGet(ctx context.Context, params antifraud_v1.APIV1StatsOverviewGetParams) (antifraud_v1.APIV1StatsOverviewGetRes, error) {
	return h.statsHandler.APIV1StatsOverviewGet(ctx, params)
}

func (h *handlerAdapter) APIV1StatsTransactionsTimeseriesGet(ctx context.Context, params antifraud_v1.APIV1StatsTransactionsTimeseriesGetParams) (antifraud_v1.APIV1StatsTransactionsTimeseriesGetRes, error) {
	return h.statsHandler.APIV1StatsTransactionsTimeseriesGet(ctx, params)
}

func (h *handlerAdapter) APIV1StatsRulesMatchesGet(ctx context.Context, params antifraud_v1.APIV1StatsRulesMatchesGetParams) (antifraud_v1.APIV1StatsRulesMatchesGetRes, error) {
	return h.statsHandler.APIV1StatsRulesMatchesGet(ctx, params)
}

func (h *handlerAdapter) APIV1StatsMerchantsRiskGet(ctx context.Context, params antifraud_v1.APIV1StatsMerchantsRiskGetParams) (antifraud_v1.APIV1StatsMerchantsRiskGetRes, error) {
	return h.statsHandler.APIV1StatsMerchantsRiskGet(ctx, params)
}

func (h *handlerAdapter) APIV1StatsUsersIDRiskProfileGet(ctx context.Context, params antifraud_v1.APIV1StatsUsersIDRiskProfileGetParams) (antifraud_v1.APIV1StatsUsersIDRiskProfileGetRes, error) {
	return h.statsHandler.APIV1StatsUsersIDRiskProfileGet(ctx, params)
}

func (h *handlerAdapter) APIV1TransactionsPost(ctx context.Context, req *antifraud_v1.TransactionCreateRequest) (antifraud_v1.APIV1TransactionsPostRes, error) {
	if ctx == nil {
		return &antifraud_v1.APIV1TransactionsPostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Context is required",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/transactions",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	userRole, ok := ctx.Value(ContextRoleKey).(string)
	if !ok {
		return &antifraud_v1.APIV1TransactionsPostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Access denied: authentication required",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/transactions",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	userID, ok := ctx.Value(ContextUserIDKey).(string)
	if !ok {
		return &antifraud_v1.APIV1TransactionsPostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Access denied: user ID not found",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/transactions",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	if userRole != "ADMIN" && userID != req.UserId.String() {
		return &antifraud_v1.APIV1TransactionsPostForbidden{
			Code:      antifraud_v1.ErrorCodeFORBIDDEN,
			Message:   "Access denied: users can only create transactions for themselves",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/transactions",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	createReq := convertTransactionCreateRequest(req, userID)
	decision, err := h.transactionService.Create(ctx, createReq)
	if err != nil {

		if strings.Contains(err.Error(), "failed to get user by ID") {
			return &antifraud_v1.APIV1TransactionsPostNotFound{
				Code:      antifraud_v1.ErrorCodeNOTFOUND,
				Message:   "User not found",
				TraceId:   uuid.New(),
				Timestamp: time.Now().UTC(),
				Path:      "/api/v1/transactions",
				Details:   antifraud_v1.OptApiErrorDetails{},
			}, nil
		}
		if strings.Contains(err.Error(), "user is deactivated") {
			return &antifraud_v1.APIV1TransactionsPostForbidden{
				Code:      antifraud_v1.ErrorCodeFORBIDDEN,
				Message:   "User is deactivated",
				TraceId:   uuid.New(),
				Timestamp: time.Now().UTC(),
				Path:      "/api/v1/transactions",
				Details:   antifraud_v1.OptApiErrorDetails{},
			}, nil
		}
		return &antifraud_v1.APIV1TransactionsPostBadRequest{
			Code:      antifraud_v1.ErrorCodeVALIDATIONFAILED,
			Message:   err.Error(),
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/transactions",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	apiTransaction := convertTransactionToAPI(decision.Transaction)

	ruleResults := make([]antifraud_v1.FraudRuleEvaluationResult, len(decision.RuleResults))
	for i, rule := range decision.RuleResults {
		ruleUUID, _ := uuid.Parse(rule.RuleID)
		ruleResults[i] = antifraud_v1.FraudRuleEvaluationResult{
			RuleId:      ruleUUID,
			RuleName:    rule.RuleName,
			Priority:    rule.Priority,
			Matched:     rule.Matched,
			Description: rule.Description,
		}
	}

	transactionDecision := antifraud_v1.TransactionDecision{
		Transaction: apiTransaction,
		RuleResults: ruleResults,
	}
	return &transactionDecision, nil
}

func (h *handlerAdapter) APIV1TransactionsBatchPost(ctx context.Context, req *antifraud_v1.TransactionBatchCreateRequest) (antifraud_v1.APIV1TransactionsBatchPostRes, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is required")
	}

	userID, ok := ctx.Value(ContextUserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("access denied: user ID not found")
	}

	if req == nil || len(req.Items) == 0 {
		return nil, fmt.Errorf("batch request cannot be empty")
	}

	results := make([]antifraud_v1.TransactionBatchResultItem, len(req.Items))
	hasErrors := false

	for i, item := range req.Items {

		serviceReq := convertTransactionCreateRequest(&item, userID)

		decision, err := h.transactionService.Create(ctx, serviceReq)
		if err != nil {
			hasErrors = true

			if strings.Contains(err.Error(), "failed to get user by ID") {
				results[i] = antifraud_v1.TransactionBatchResultItem{
					Index: i,
					Error: antifraud_v1.OptApiError{
						Value: antifraud_v1.ApiError{
							Code:    antifraud_v1.ErrorCodeNOTFOUND,
							Message: "User not found",
						},
						Set: true,
					},
				}
				continue
			} else if strings.Contains(err.Error(), "user is deactivated") {
				results[i] = antifraud_v1.TransactionBatchResultItem{
					Index: i,
					Error: antifraud_v1.OptApiError{
						Value: antifraud_v1.ApiError{
							Code:    antifraud_v1.ErrorCodeFORBIDDEN,
							Message: "User is deactivated",
						},
						Set: true,
					},
				}
				continue
			}

			results[i] = antifraud_v1.TransactionBatchResultItem{
				Index: i,
				Error: antifraud_v1.OptApiError{
					Value: antifraud_v1.ApiError{
						Code:    antifraud_v1.ErrorCodeVALIDATIONFAILED,
						Message: err.Error(),
					},
					Set: true,
				},
			}
		} else {

			ruleResults := make([]antifraud_v1.FraudRuleEvaluationResult, len(decision.RuleResults))
			for j, rule := range decision.RuleResults {
				ruleUUID, _ := uuid.Parse(rule.RuleID)
				ruleResults[j] = antifraud_v1.FraudRuleEvaluationResult{
					RuleId:      ruleUUID,
					RuleName:    rule.RuleName,
					Priority:    rule.Priority,
					Matched:     rule.Matched,
					Description: rule.Description,
				}
			}

			results[i] = antifraud_v1.TransactionBatchResultItem{
				Index: i,
				Decision: antifraud_v1.OptTransactionDecision{
					Value: antifraud_v1.TransactionDecision{
						Transaction: convertTransactionToAPI(decision.Transaction),
						RuleResults: ruleResults,
					},
					Set: true,
				},
			}
		}
	}

	if hasErrors {
		return &antifraud_v1.APIV1TransactionsBatchPostMultiStatus{
			Items: results,
		}, nil
	}

	return &antifraud_v1.APIV1TransactionsBatchPostCreated{
		Items: results,
	}, nil
}

func (h *handlerAdapter) APIV1TransactionsGet(ctx context.Context, params antifraud_v1.APIV1TransactionsGetParams) (antifraud_v1.APIV1TransactionsGetRes, error) {
	userRole, ok := ctx.Value(ContextRoleKey).(string)
	if !ok || userRole != "ADMIN" {
		return nil, fmt.Errorf("access denied: only ADMIN can view all transactions")
	}

	var filterUserID *string
	if params.UserId.Set {
		userIDStr := params.UserId.Value.String()
		filterUserID = &userIDStr
	}

	var filterStatus *model.TransactionStatus
	if params.Status.Set {
		status := model.TransactionStatus(params.Status.Value)
		filterStatus = &status
	}

	var filterIsFraud *bool
	if params.IsFraud.Set {
		filterIsFraud = &params.IsFraud.Value
	}

	serviceParams := service.TransactionListParams{
		UserID:  filterUserID,
		Status:  filterStatus,
		IsFraud: filterIsFraud,
		Page:    0,
		Size:    100,
	}

	pagedTransactions, err := h.transactionService.GetList(ctx, serviceParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}

	apiTransactions := make([]antifraud_v1.Transaction, len(pagedTransactions.Items))
	for i, txDecision := range pagedTransactions.Items {
		apiTransactions[i] = convertTransactionToAPI(txDecision.Transaction)
	}

	result := antifraud_v1.PagedTransactions{
		Items: apiTransactions,
		Total: int(pagedTransactions.Total),
		Page:  pagedTransactions.Page,
	}

	return &result, nil
}

func (h *handlerAdapter) APIV1TransactionsIDGet(ctx context.Context, params antifraud_v1.APIV1TransactionsIDGetParams) (antifraud_v1.APIV1TransactionsIDGetRes, error) {
	userRole, ok := ctx.Value(ContextRoleKey).(string)
	if !ok {
		return nil, fmt.Errorf("access denied: authentication required")
	}

	userID, ok := ctx.Value(ContextUserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("access denied: user ID not found")
	}

	txDecision, err := h.transactionService.GetByID(ctx, params.ID.String())
	if err != nil {
		return nil, fmt.Errorf("transaction not found: %w", err)
	}

	// Check access: users can only see their own transactions, admins can see all
	if userRole != "ADMIN" && txDecision.Transaction.UserID.String() != userID {
		return nil, fmt.Errorf("access denied: users can only view their own transactions")
	}

	apiTransaction := convertTransactionToAPI(txDecision.Transaction)

	// Convert rule results
	ruleResults := make([]antifraud_v1.FraudRuleEvaluationResult, len(txDecision.RuleResults))
	for i, rule := range txDecision.RuleResults {
		ruleUUID, _ := uuid.Parse(rule.RuleID)
		ruleResults[i] = antifraud_v1.FraudRuleEvaluationResult{
			RuleId:      ruleUUID,
			RuleName:    rule.RuleName,
			Priority:    rule.Priority,
			Matched:     rule.Matched,
			Description: rule.Description,
		}
	}

	transactionDecision := antifraud_v1.TransactionDecision{
		Transaction: apiTransaction,
		RuleResults: ruleResults,
	}
	return &transactionDecision, nil
}

func convertTransactionCreateRequest(req *antifraud_v1.TransactionCreateRequest, userID string) model.TransactionCreateRequest {
	createReq := model.TransactionCreateRequest{
		UserID:    &req.UserId,
		Amount:    req.Amount,
		Currency:  model.CurrencyCode(req.Currency),
		Timestamp: req.Timestamp,
	}

	if req.MerchantId.Set {
		createReq.MerchantID = &req.MerchantId.Value
	}

	if req.MerchantCategoryCode.Set {
		mcc := model.MCCCode(req.MerchantCategoryCode.Value)
		createReq.MerchantCategoryCode = &mcc
	}

	if req.IpAddress.Set {
		if ip, err := netip.ParseAddr(req.IpAddress.Value); err == nil {
			createReq.IPAddress = &ip
		}
	}

	if req.DeviceId.Set {
		createReq.DeviceID = &req.DeviceId.Value
	}

	if req.Channel.Set {
		channel := model.TransactionChannel(req.Channel.Value)
		createReq.Channel = &channel
	}

	if req.Location.Set {
		if req.Location.Value.Latitude.Set && req.Location.Value.Longitude.Set {
			createReq.Location = &model.TransactionLocation{
				Latitude:  &req.Location.Value.Latitude.Value,
				Longitude: &req.Location.Value.Longitude.Value,
			}
			if req.Location.Value.Country.Set {
				createReq.Location.Country = req.Location.Value.Country.Value
			}
		}
	}

	if req.Metadata.Set {
		metadata := make(model.TransactionMetadata)
		for k, v := range req.Metadata.Value {
			metadata[k] = v
		}
		createReq.Metadata = &metadata
	}

	return createReq
}

func convertTransactionToAPI(t *model.Transaction) antifraud_v1.Transaction {
	transaction := antifraud_v1.Transaction{
		ID:        t.ID,
		UserId:    t.UserID,
		Amount:    t.Amount,
		Currency:  antifraud_v1.CurrencyCode(t.Currency),
		Status:    antifraud_v1.TransactionStatus(t.Status),
		Timestamp: t.Timestamp,
		Channel:   antifraud_v1.OptTransactionChannel{},
		IsFraud:   t.IsFraud,
		CreatedAt: t.CreatedAt,
	}

	if t.MerchantID != nil {
		transaction.MerchantId = antifraud_v1.OptString{Set: true, Value: *t.MerchantID}
	}

	if t.MerchantCategoryCode != nil {
		transaction.MerchantCategoryCode = antifraud_v1.OptMccCode{Set: true, Value: antifraud_v1.MccCode(*t.MerchantCategoryCode)}
	}

	if t.IPAddress != nil {
		transaction.IpAddress = antifraud_v1.OptString{Set: true, Value: t.IPAddress.String()}
	}

	if t.DeviceID != nil {
		transaction.DeviceId = antifraud_v1.OptString{Set: true, Value: *t.DeviceID}
	}

	if t.Channel != nil {
		transaction.Channel = antifraud_v1.OptTransactionChannel{Set: true, Value: antifraud_v1.TransactionChannel(*t.Channel)}
	}

	if t.Location != nil {
		transaction.Location = antifraud_v1.OptTransactionLocation{
			Set: true,
			Value: antifraud_v1.TransactionLocation{
				Latitude:  antifraud_v1.OptFloat64{Set: true, Value: *t.Location.Latitude},
				Longitude: antifraud_v1.OptFloat64{Set: true, Value: *t.Location.Longitude},
			},
		}
		if t.Location.Country != "" {
			transaction.Location.Value.Country = antifraud_v1.OptString{Set: true, Value: t.Location.Country}
		}
	}

	if t.Metadata != nil {
		metadata := make(antifraud_v1.TransactionMetadata)
		for k, v := range *t.Metadata {
			if str, ok := v.(string); ok {
				metadata[k] = jx.Raw(str)
			}
		}
		transaction.Metadata = antifraud_v1.OptTransactionMetadata{Set: true, Value: metadata}
	}

	return transaction
}
