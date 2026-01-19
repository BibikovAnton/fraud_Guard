package model

import (
	"time"

	"github.com/google/uuid"
)

// TransactionStats - статистика по транзакциям
type TransactionStats struct {
	StatusCounts   map[TransactionStatus]int64 `json:"status_counts"`
	FraudLast24h   int64                       `json:"fraud_last_24h"`
	TotalAmountUSD float64                     `json:"total_amount_usd"`
	TotalAmountEUR float64                     `json:"total_amount_eur"`
	TotalAmountRUB float64                     `json:"total_amount_rub"`
	GeneratedAt    time.Time                   `json:"generated_at"`
}

// MerchantRiskStats - статистика рисков по мерчантам
type MerchantRiskStats struct {
	MerchantID       string    `json:"merchant_id"`
	TransactionCount int64     `json:"transaction_count"`
	FraudCount       int64     `json:"fraud_count"`
	RiskScore        float64   `json:"risk_score"`
	TotalAmount      float64   `json:"total_amount"`
	FraudAmount      float64   `json:"fraud_amount"`
	LastTransaction  time.Time `json:"last_transaction"`
}

// UserRiskProfile - профиль риска пользователя
type UserRiskProfile struct {
	UserID           uuid.UUID `json:"user_id"`
	TransactionCount int64     `json:"transaction_count"`
	FraudCount       int64     `json:"fraud_count"`
	RiskScore        float64   `json:"risk_score"`
	AverageAmount    float64   `json:"average_amount"`
	MaxAmount        float64   `json:"max_amount"`
	LastTransaction  time.Time `json:"last_transaction"`
}

// FraudRuleMatch - результат срабатывания правила
type FraudRuleMatch struct {
	RuleID        uuid.UUID `json:"rule_id"`
	RuleName      string    `json:"rule_name"`
	TransactionID uuid.UUID `json:"transaction_id"`
	MatchedAt     time.Time `json:"matched_at"`
	Description   string    `json:"description"`
}
