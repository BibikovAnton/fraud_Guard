package v1

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"solution/internal/api/antifraud/v1/convertor"
	"solution/internal/model"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
	"strings"
	"time"
)

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
		return &antifraud_v1.ValidationError{
			Code:      string(antifraud_v1.ErrorCodeVALIDATIONFAILED),
			Message:   "Amount must be positive",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/transactions",
			FieldErrors: []antifraud_v1.FieldError{
				{
					Field: "amount",
					Issue: "Amount must be positive",
				},
			},
		}, nil
	}

	if strings.TrimSpace(req.Currency) == "" {
		return &antifraud_v1.ValidationError{
			Code:      string(antifraud_v1.ErrorCodeVALIDATIONFAILED),
			Message:   "Currency is required",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/transactions",
			FieldErrors: []antifraud_v1.FieldError{
				{
					Field: "currency",
					Issue: "Currency is required",
				},
			},
		}, nil
	}

	createReq := model.TransactionCreateRequest{
		UserID:               req.UserId,
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
		createReq.Location = &model.TransactionLocation{
			Latitude:  req.Location.Value.Latitude,
			Longitude: req.Location.Value.Longitude,
			Country:   req.Location.Value.Country,
			City:      req.Location.Value.City,
		}
	}
	if len(req.Metadata) > 0 {
		metadata := make(model.TransactionMetadata)
		for k, v := range req.Metadata {
			metadata[k] = v
		}
		createReq.Metadata = &metadata
	}

	transaction, err := h.transactionService.Create(ctx, createReq)
	if err != nil {
		return &antifraud_v1.APIV1TransactionsPostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Failed to create transaction",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/transactions",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	apiTransaction := convertor.ConvertTransactionToAPI(transaction)
	return &apiTransaction, nil
}
