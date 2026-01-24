package dsl

import (
	"fmt"
	"go.uber.org/zap"
	"solution/internal/model"
	"strings"
)

// DSL Parser implementation for fraud rules
// Supports Tier 0 (no parsing) and basic Tier 1-5 functionality

type DSLEvaluator struct {
	logger *zap.Logger
}

func NewDSLEvaluator(logger *zap.Logger) *DSLEvaluator {
	return &DSLEvaluator{
		logger: logger,
	}
}

// EvaluateRule evaluates a DSL rule against transaction and user data
func (e *DSLEvaluator) EvaluateRule(rule *model.FraudRule, transaction *model.Transaction, user *model.User) model.RuleResult {
	result := model.RuleResult{
		RuleID:      rule.ID,
		RuleName:    rule.Name,
		Priority:    rule.Priority,
		Enabled:     rule.Enabled,
		Matched:     false,
		Description: "",
	}

	if !rule.Enabled {
		result.Description = "Rule is disabled"
		return result
	}

	// For now, implement Tier 0 - always return false
	// This allows the system to work while we implement the full parser
	result.Description = "Tier 0: Rule evaluation not implemented yet"
	
	// TODO: Implement full DSL parsing for Tier 1-5
	// This would include:
	// - Tier 1: amount field with numeric operators
	// - Tier 2: string fields (currency, merchantId, etc.)
	// - Tier 3: AND/OR logic
	// - Tier 4: NOT and parentheses
	// - Tier 5: user.age and user.region fields

	return result
}

// ValidateDSL validates DSL expression syntax
func (e *DSLEvaluator) ValidateDSL(expression string) model.DslValidateResponse {
	response := model.DslValidateResponse{
		IsValid:              false,
		NormalizedExpression: nil,
		Errors:               []model.DSLError{},
	}

	// For now, implement Tier 0 validation - always return unsupported tier
	response.Errors = append(response.Errors, model.DSLError{
		Code:    "DSL_UNSUPPORTED_TIER",
		Message: "DSL parsing not implemented yet - Tier 0 mode",
		Position: nil,
		Near:     nil,
	})

	return response
}

// NormalizeExpression normalizes DSL expression (uppercase operators, spacing)
func (e *DSLEvaluator) NormalizeExpression(expression string) string {
	// Basic normalization - convert to uppercase and normalize spacing
	normalized := strings.ToUpper(expression)
	normalized = strings.ReplaceAll(normalized, "AND", " AND ")
	normalized = strings.ReplaceAll(normalized, "OR", " OR ")
	normalized = strings.ReplaceAll(normalized, "NOT", " NOT ")
	
	// Clean up extra spaces
	normalized = strings.Join(strings.Fields(normalized), " ")
	
	return normalized
}

// EvaluateComparison evaluates a single comparison expression
func (e *DSLEvaluator) EvaluateComparison(field, operator string, value interface{}, transaction *model.Transaction, user *model.User) (bool, error) {
	switch field {
	case "amount":
		return e.evaluateAmount(operator, value, transaction)
	case "currency":
		return e.evaluateCurrency(operator, value, transaction)
	case "merchantId":
		return e.evaluateMerchantId(operator, value, transaction)
	case "ipAddress":
		return e.evaluateIPAddress(operator, value, transaction)
	case "deviceId":
		return e.evaluateDeviceId(operator, value, transaction)
	case "user.age":
		return e.evaluateUserAge(operator, value, user)
	case "user.region":
		return e.evaluateUserRegion(operator, value, user)
	default:
		return false, fmt.Errorf("unsupported field: %s", field)
	}
}

func (e *DSLEvaluator) evaluateAmount(operator string, value interface{}, transaction *model.Transaction) (bool, error) {
	compareValue, ok := value.(float64)
	if !ok {
		return false, fmt.Errorf("amount comparison requires numeric value")
	}

	switch operator {
	case ">", ">=", "<", "<=":
		return e.compareFloats(transaction.Amount, operator, compareValue), nil
	case "=", "!=":
		return e.compareFloats(transaction.Amount, operator, compareValue), nil
	default:
		return false, fmt.Errorf("unsupported operator for amount: %s", operator)
	}
}

func (e *DSLEvaluator) evaluateCurrency(operator string, value interface{}, transaction *model.Transaction) (bool, error) {
	compareValue, ok := value.(string)
	if !ok {
		return false, fmt.Errorf("currency comparison requires string value")
	}

	if operator != "=" && operator != "!=" {
		return false, fmt.Errorf("currency only supports = and != operators")
	}

	currencyStr := string(transaction.Currency)
	return e.compareStrings(currencyStr, operator, compareValue), nil
}

func (e *DSLEvaluator) evaluateMerchantId(operator string, value interface{}, transaction *model.Transaction) (bool, error) {
	if transaction.MerchantID == nil {
		return false, nil
	}

	compareValue, ok := value.(string)
	if !ok {
		return false, fmt.Errorf("merchantId comparison requires string value")
	}

	if operator != "=" && operator != "!=" {
		return false, fmt.Errorf("merchantId only supports = and != operators")
	}

	return e.compareStrings(*transaction.MerchantID, operator, compareValue), nil
}

func (e *DSLEvaluator) evaluateIPAddress(operator string, value interface{}, transaction *model.Transaction) (bool, error) {
	if transaction.IPAddress == nil {
		return false, nil
	}

	compareValue, ok := value.(string)
	if !ok {
		return false, fmt.Errorf("ipAddress comparison requires string value")
	}

	if operator != "=" && operator != "!=" {
		return false, fmt.Errorf("ipAddress only supports = and != operators")
	}

	return e.compareStrings(transaction.IPAddress.String(), operator, compareValue), nil
}

func (e *DSLEvaluator) evaluateDeviceId(operator string, value interface{}, transaction *model.Transaction) (bool, error) {
	if transaction.DeviceID == nil {
		return false, nil
	}

	compareValue, ok := value.(string)
	if !ok {
		return false, fmt.Errorf("deviceId comparison requires string value")
	}

	if operator != "=" && operator != "!=" {
		return false, fmt.Errorf("deviceId only supports = and != operators")
	}

	return e.compareStrings(*transaction.DeviceID, operator, compareValue), nil
}

func (e *DSLEvaluator) evaluateUserAge(operator string, value interface{}, user *model.User) (bool, error) {
	if user == nil || user.Age == nil {
		return false, nil // null user.age returns false
	}

	compareValue, ok := value.(float64)
	if !ok {
		return false, fmt.Errorf("user.age comparison requires numeric value")
	}

	userAge := float64(*user.Age)

	switch operator {
	case ">", ">=", "<", "<=":
		return e.compareFloats(userAge, operator, compareValue), nil
	case "=", "!=":
		return e.compareFloats(userAge, operator, compareValue), nil
	default:
		return false, fmt.Errorf("unsupported operator for user.age: %s", operator)
	}
}

func (e *DSLEvaluator) evaluateUserRegion(operator string, value interface{}, user *model.User) (bool, error) {
	if user == nil || user.Region == nil {
		return false, nil // null user.region returns false
	}

	compareValue, ok := value.(string)
	if !ok {
		return false, fmt.Errorf("user.region comparison requires string value")
	}

	if operator != "=" && operator != "!=" {
		return false, fmt.Errorf("user.region only supports = and != operators")
	}

	return e.compareStrings(*user.Region, operator, compareValue), nil
}

func (e *DSLEvaluator) compareFloats(a float64, operator string, b float64) bool {
	switch operator {
	case ">":
		return a > b
	case ">=":
		return a >= b
	case "<":
		return a < b
	case "<=":
		return a <= b
	case "=":
		return a == b
	case "!=":
		return a != b
	default:
		return false
	}
}

func (e *DSLEvaluator) compareStrings(a, operator, b string) bool {
	switch operator {
	case "=":
		return a == b
	case "!=":
		return a != b
	default:
		return false
	}
}
