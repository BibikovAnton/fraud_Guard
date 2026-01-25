package v1

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"solution/internal/api/antifraud/v1/convertor"
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

	createReq := convertor.ConvertTransactionCreateRequest(req, userID)
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

	apiTransaction := convertor.ConvertTransactionToAPI(decision.Transaction)

	ruleResults := make([]antifraud_v1.FraudRuleEvaluationResult, len(decision.RuleResults))
	for i, rule := range decision.RuleResults {
		var ruleUUID uuid.UUID
		if rule.RuleID != "" {
			ruleUUID = uuid.MustParse(rule.RuleID)
		}
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

		serviceReq := convertor.ConvertTransactionCreateRequest(&item, userID)

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
				var ruleUUID uuid.UUID
				if rule.RuleID != "" {
					ruleUUID = uuid.MustParse(rule.RuleID)
				}
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
						Transaction: convertor.ConvertTransactionToAPI(decision.Transaction),
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
	for i, tx := range pagedTransactions.Items {
		apiTransactions[i] = convertor.ConvertTransactionToAPI(tx)
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

	if userRole != "ADMIN" && txDecision.Transaction.UserID.String() != userID {
		return nil, fmt.Errorf("access denied: users can only view their own transactions")
	}

	apiTransaction := convertor.ConvertTransactionToAPI(txDecision.Transaction)

	ruleResults := make([]antifraud_v1.FraudRuleEvaluationResult, len(txDecision.RuleResults))
	for i, rule := range txDecision.RuleResults {
		var ruleUUID uuid.UUID
		if rule.RuleID != "" {
			ruleUUID = uuid.MustParse(rule.RuleID)
		}
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
