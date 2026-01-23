package v1

import (
	"context"
	"fmt"
	"strings"
	"time"
	"github.com/google/uuid"
	"github.com/go-faster/jx"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
	"solution/internal/model"
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

	if req.Amount <= 0 {
		return &antifraud_v1.APIV1TransactionsPostBadRequest{
			Code:      antifraud_v1.ErrorCodeVALIDATIONFAILED,
			Message:   "Amount must be positive",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/transactions",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	if strings.TrimSpace(string(req.Currency)) == "" {
		return &antifraud_v1.APIV1TransactionsPostBadRequest{
			Code:      antifraud_v1.ErrorCodeVALIDATIONFAILED,
			Message:   "Currency is required",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/transactions",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	createReq := convertTransactionCreateRequest(req, userID)
	decision, err := h.transactionService.Create(ctx, createReq)
	if err != nil {
		return &antifraud_v1.APIV1TransactionsPostBadRequest{
			Code:      antifraud_v1.ErrorCodeINTERNALSERVERERROR,
			Message:   "Failed to create transaction",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/transactions",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	apiTransaction := convertTransactionToAPI(decision.Transaction)
	transactionDecision := antifraud_v1.TransactionDecision{
		Transaction: apiTransaction,
		RuleResults: []antifraud_v1.FraudRuleEvaluationResult{},
	}
	return &transactionDecision, nil
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

func convertTransactionCreateRequest(req *antifraud_v1.TransactionCreateRequest, userID string) model.TransactionCreateRequest {
	userUUID, _ := uuid.Parse(userID)
	
	createReq := model.TransactionCreateRequest{
		UserID:               userUUID,
		Amount:               req.Amount,
		Currency:             model.CurrencyCode(req.Currency),
		Timestamp:            req.Timestamp,
	}

	if req.MerchantId.Set {
		createReq.MerchantID = &req.MerchantId.Value
	}

	if req.MerchantCategoryCode.Set {
		mcc := model.MCCCode(req.MerchantCategoryCode.Value)
		createReq.MerchantCategoryCode = &mcc
	}

	if req.IpAddress.Set {
		createReq.IPAddress = &req.IpAddress.Value
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
		ID:       t.ID,
		UserId:   t.UserID,
		Amount:   t.Amount,
		Currency: antifraud_v1.CurrencyCode(t.Currency),
		Status:   antifraud_v1.TransactionStatus(t.Status),
		Timestamp: t.Timestamp,
		Channel:  antifraud_v1.OptTransactionChannel{},
		IsFraud:  t.IsFraud,
		CreatedAt: t.CreatedAt,
	}

	if t.MerchantID != nil {
		transaction.MerchantId = antifraud_v1.OptString{Set: true, Value: *t.MerchantID}
	}

	if t.MerchantCategoryCode != nil {
		transaction.MerchantCategoryCode = antifraud_v1.OptMccCode{Set: true, Value: antifraud_v1.MccCode(*t.MerchantCategoryCode)}
	}

	if t.IPAddress != nil {
		transaction.IpAddress = antifraud_v1.OptString{Set: true, Value: *t.IPAddress}
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
