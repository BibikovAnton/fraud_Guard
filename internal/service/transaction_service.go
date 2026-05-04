package service

import (
	"context"
	"solution/internal/model"
	"time"
)

type TransactionService interface {
	Create(ctx context.Context, req model.TransactionCreateRequest) (*model.TransactionDecision, error)
	
	CreateBatch(ctx context.Context, req model.TransactionBatchCreateRequest) (*model.TransactionBatchResult, error)
	
	GetByID(ctx context.Context, id string) (*model.TransactionDecision, error)
	
	GetList(ctx context.Context, params TransactionListParams) (*PagedTransactions, error)
}

type TransactionListParams struct {
	UserID   *string
	Status   *model.TransactionStatus
	IsFraud  *bool
	From     *time.Time
	To       *time.Time
	Page     int
	Size     int
}

type PagedTransactions struct {
	Items []*model.Transaction `json:"items"`
	Total int64                `json:"total"`
	Page  int                  `json:"page"`
	Size  int                  `json:"size"`
}
