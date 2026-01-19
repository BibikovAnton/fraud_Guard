package model

import (
	"time"

	"github.com/google/uuid"
)

// TransactionStatus - статус транзакции
type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "PENDING"
	TransactionStatusApproved  TransactionStatus = "APPROVED"
	TransactionStatusRejected  TransactionStatus = "REJECTED"
	TransactionStatusProcessed TransactionStatus = "PROCESSED"
)

// TransactionChannel - канал транзакции
type TransactionChannel string

const (
	TransactionChannelOnline TransactionChannel = "ONLINE"
	TransactionChannelPOS    TransactionChannel = "POS"
	TransactionChannelMobile TransactionChannel = "MOBILE"
	TransactionChannelATM    TransactionChannel = "ATM"
)

// CurrencyCode - код валюты
type CurrencyCode string

const (
	CurrencyUSD CurrencyCode = "USD"
	CurrencyEUR CurrencyCode = "EUR"
	CurrencyRUB CurrencyCode = "RUB"
)

// MCCCode - код категории мерчанта
type MCCCode string

// TransactionLocation - геолокация
type TransactionLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
}

// TransactionMetadata - метаданные транзакции
type TransactionMetadata map[string]interface{}

// Transaction - модель транзакции
// Из прошлого проекта: важна производительность на 10k RPS
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

// TransactionCreateRequest - запрос на создание транзакции
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

// TransactionBatchCreateRequest - запрос на создание пачки транзакций
type TransactionBatchCreateRequest struct {
	Transactions []TransactionCreateRequest `json:"transactions"`
}

// TransactionBatchResult - результат обработки пачки
type TransactionBatchResult struct {
	Processed int                    `json:"processed"`
	Success   int                    `json:"success"`
	Failed    int                    `json:"failed"`
	Results   []TransactionBatchItem `json:"results"`
}

// TransactionBatchItem - элемент результата пачки
type TransactionBatchItem struct {
	Index   int          `json:"index"`
	Success bool         `json:"success"`
	Error   *string      `json:"error,omitempty"`
	Data    *Transaction `json:"data,omitempty"`
}

// TransactionUpdateRequest - запрос на обновление транзакции
type TransactionUpdateRequest struct {
	Status   *TransactionStatus   `json:"status,omitempty"`
	IsFraud  *bool                `json:"is_fraud,omitempty"`
	Metadata *TransactionMetadata `json:"metadata,omitempty"`
}

// Constants
const (
	MaxTransactionAmount = 1000000.0 // 1 миллион
	MinTransactionAmount = 0.01      // 1 цент
	MaxBatchSize         = 1000      // максимум 1000 транзакций в пачке
)
