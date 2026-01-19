package model

import (
	"time"
)

type FraudRule struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	DSL         string    `json:"dsl" db:"dsl"`
	Priority    int       `json:"priority" db:"priority"`
	IsActive    bool      `json:"isActive" db:"is_active"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type FraudRuleCreateRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"required,min=1,max=500"`
	DSL         string `json:"dsl" validate:"required,min=1,max=10000"`
	Priority    *int   `json:"priority"`
}

type FraudRuleUpdateRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,min=1,max=500"`
	DSL         *string `json:"dsl" validate:"omitempty,min=1,max=10000"`
	Priority    *int    `json:"priority"`
	IsActive    *bool   `json:"isActive"`
}

type FraudRuleValidateRequest struct {
	DSL string `json:"dsl" validate:"required,min=1,max=10000"`
}

type FraudRuleValidateResponse struct {
	IsValid bool   `json:"isValid"`
	Error   string `json:"error,omitempty"`
	AST     string `json:"ast,omitempty"`
}

type DSLValidationError struct {
	Line   int    `json:"line"`
	Column int    `json:"column"`
	Error  string `json:"error"`
}

// Константы для бизнес-логики
const (
	DefaultPriority = 100

	MaxDSLSize = 10000

	MaxRulesCount = 1000
)
