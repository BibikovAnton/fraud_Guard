package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "PENDING"
	TransactionStatusApproved  TransactionStatus = "APPROVED"
	TransactionStatusDeclined TransactionStatus = "DECLINED"
)

type TransactionChannel string

const (
	TransactionChannelWeb   TransactionChannel = "WEB"
	TransactionChannelMobile TransactionChannel = "MOBILE"
	TransactionChannelPOS    TransactionChannel = "POS"
	TransactionChannelOther TransactionChannel = "OTHER"
)

type CurrencyCode string

const (
	CurrencyUSD CurrencyCode = "USD"
	CurrencyEUR CurrencyCode = "EUR"
	CurrencyRUB CurrencyCode = "RUB"
)

type MCCCode string

type TransactionLocation struct {
	Country   string  `json:"country"`
	City      string  `json:"city"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

type TransactionMetadata map[string]interface{}

type Transaction struct {
	ID                   uuid.UUID            `json:"id" db:"id"`
	UserID               uuid.UUID            `json:"user_id" db:"user_id"`
	Amount               float64              `json:"amount" db:"amount"`
	Currency             CurrencyCode         `json:"currency" db:"currency"`
	Status               TransactionStatus    `json:"status" db:"status"`
	MerchantID           *string              `json:"merchant_id,omitempty" db:"merchant_id"`
	MerchantCategoryCode *MCCCode             `json:"merchant_category_code,omitempty" db:"merchant_category_code"`
	Timestamp            time.Time            `json:"timestamp" db:"timestamp"`
	IPAddress            *string              `json:"ip_address,omitempty" db:"ip_address"`
	DeviceID             *string              `json:"device_id,omitempty" db:"device_id"`
	Channel              *TransactionChannel  `json:"channel,omitempty" db:"channel"`
	Location             *TransactionLocation `json:"location,omitempty" db:"location"`
	IsFraud              bool                 `json:"is_fraud" db:"is_fraud"`
	Metadata             *TransactionMetadata `json:"metadata,omitempty" db:"metadata"`
	CreatedAt            time.Time            `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time            `json:"updated_at" db:"updated_at"`
}

type TransactionCreateRequest struct {
	UserID               uuid.UUID            `json:"user_id"`
	Amount               float64              `json:"amount"`
	Currency             CurrencyCode         `json:"currency"`
	MerchantID           *string              `json:"merchant_id,omitempty"`
	MerchantCategoryCode *MCCCode             `json:"merchant_category_code,omitempty"`
	Timestamp            time.Time            `json:"timestamp"`
	IPAddress            *string              `json:"ip_address,omitempty"`
	DeviceID             *string              `json:"device_id,omitempty"`
	Channel              *TransactionChannel  `json:"channel,omitempty"`
	Location             *TransactionLocation `json:"location,omitempty"`
	Metadata             *TransactionMetadata `json:"metadata,omitempty"`
}

type TransactionUpdateRequest struct {
	Status   *TransactionStatus   `json:"status,omitempty"`
	IsFraud  *bool                `json:"is_fraud,omitempty"`
	Metadata *TransactionMetadata `json:"metadata,omitempty"`
}

const (
	MaxTransactionAmount = 999999999.99 // максимум из OpenAPI
	MinTransactionAmount = 0.01         // минимум из OpenAPI
	MaxBatchSize         = 500          // максимум из OpenAPI
)

type RuleResult struct {
	RuleID      string `json:"ruleId"`
	RuleName    string `json:"ruleName"`
	Priority    int    `json:"priority"`
	Enabled     bool   `json:"enabled"`
	Matched     bool   `json:"matched"`
	Description string `json:"description"`
}

type TransactionDecision struct {
	Transaction *Transaction   `json:"transaction"`
	RuleResults []RuleResult  `json:"ruleResults"`
}

type ApiError struct {
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	TraceID   uuid.UUID              `json:"traceId"`
	Timestamp time.Time              `json:"timestamp"`
	Path      string                 `json:"path"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

type TransactionBatchCreateRequest struct {
	Items []TransactionCreateRequest `json:"items"`
}

type TransactionBatchResult struct {
	Items []TransactionBatchItem `json:"items"`
}

type TransactionBatchItem struct {
	Index    int                    `json:"index"`
	Decision *TransactionDecision   `json:"decision,omitempty"`
	Error    *ApiError              `json:"error,omitempty"`
}
