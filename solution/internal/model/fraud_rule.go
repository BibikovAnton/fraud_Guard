package model

import (
	"time"
)

type FraudRule struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	DslExpression string  `json:"dslExpression" db:"dsl"`
	Enabled     bool      `json:"enabled" db:"is_active"`
	Priority    int       `json:"priority" db:"priority"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type FraudRuleCreateRequest struct {
	Name         string `json:"name" validate:"required,min=3,max=120"`
	Description  string `json:"description" validate:"required,max=500"`
	DslExpression string `json:"dslExpression" validate:"required,min=3,max=2000"`
	Enabled      *bool  `json:"enabled"`
	Priority     *int   `json:"priority" validate:"omitempty,min=1"`
}

type FraudRuleUpdateRequest struct {
	Name         *string `json:"name" validate:"omitempty,min=3,max=120"`
	Description  *string `json:"description" validate:"omitempty,max=500"`
	DslExpression *string `json:"dslExpression" validate:"omitempty,min=3,max=2000"`
	Enabled      *bool   `json:"enabled"`
	Priority     *int    `json:"priority" validate:"omitempty,min=1"`
}

type DslValidateRequest struct {
	DslExpression string `json:"dslExpression" validate:"required,min=3,max=2000"`
}

type DslValidateResponse struct {
	IsValid            bool                    `json:"isValid"`
	NormalizedExpression *string               `json:"normalizedExpression,omitempty"`
	Errors             []DSLError              `json:"errors,omitempty"`
}

type DSLError struct {
	Code      string  `json:"code"`
	Message   string  `json:"message"`
	Position  *int    `json:"position,omitempty"`
	Near      *string `json:"near,omitempty"`
}

type DSLValidationError struct {
	Line   int    `json:"line"`
	Column int    `json:"column"`
	Error  string `json:"error"`
}

const (
	DefaultPriority = 100
	MaxDSLSize       = 2000
	MaxRulesCount    = 1000
	MaxASTNodes      = 100 // лимит сложности DSL
)
