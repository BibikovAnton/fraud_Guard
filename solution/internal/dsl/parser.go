package dsl

import (
	"fmt"
	"go.uber.org/zap"
	"solution/internal/model"
	"strconv"
	"strings"
)

type DSLEvaluator struct {
	logger *zap.Logger
}

func NewDSLEvaluator(logger *zap.Logger) *DSLEvaluator {
	return &DSLEvaluator{
		logger: logger,
	}
}

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

	dsl := strings.ToLower(strings.TrimSpace(rule.DslExpression))
	
	// Parse simple comparison: field operator value
	field, operator, value := e.parseSimpleComparison(dsl)
	if field == "" || operator == "" || value == "" {
		result.Description = "Invalid DSL expression"
		return result
	}
	
	switch field {
	case "amount":
		var threshold float64
		if _, err := fmt.Sscanf(value, "%f", &threshold); err == nil {
			matched := e.compareFloats(transaction.Amount, operator, threshold)
			result.Matched = matched
			if matched {
				result.Description = fmt.Sprintf("Amount %.2f %s %.2f", transaction.Amount, operator, threshold)
			} else {
				result.Description = fmt.Sprintf("Amount %.2f not %s %.2f", transaction.Amount, operator, threshold)
			}
		} else {
			result.Description = "Invalid amount value"
		}
		
	case "user.age":
		var threshold float64
		if _, err := fmt.Sscanf(value, "%f", &threshold); err == nil {
			if user != nil && user.Age != nil {
				age := float64(*user.Age)
				matched := e.compareFloats(age, operator, threshold)
				result.Matched = matched
				if matched {
					result.Description = fmt.Sprintf("User age %d %s %.0f", *user.Age, operator, threshold)
				} else {
					result.Description = fmt.Sprintf("User age %d not %s %.0f", *user.Age, operator, threshold)
				}
			} else {
				result.Description = "User age not available"
			}
		} else {
			result.Description = "Invalid age value"
		}
		
	case "user.score":
		var threshold float64
		if _, err := fmt.Sscanf(value, "%f", &threshold); err == nil {
			if user != nil && user.Score != nil {
				score := float64(*user.Score)
				matched := e.compareFloats(score, operator, threshold)
				result.Matched = matched
				if matched {
					result.Description = fmt.Sprintf("User score %.2f %s %.2f", *user.Score, operator, threshold)
				} else {
					result.Description = fmt.Sprintf("User score %.2f not %s %.2f", *user.Score, operator, threshold)
				}
			} else {
				result.Description = "User score not available"
			}
		} else {
			result.Description = "Invalid score value"
		}
		
	default:
		result.Description = "Unsupported field"
	}

	return result
}

func (e *DSLEvaluator) ValidateDSL(expression string) model.DslValidateResponse {
	response := model.DslValidateResponse{
		IsValid:              false,
		NormalizedExpression: nil,
		Errors:               []model.DSLError{},
	}
	
	if len(expression) < 3 || len(expression) > 2000 {
		response.Errors = append(response.Errors, model.DSLError{
			Code:    "DSL_PARSE_ERROR",
			Message: "Expression length must be between 3 and 2000 characters",
		})
		return response
	}
	
	normalized := e.normalizeExpression(expression)
	
	if normalized == "" {
		response.Errors = append(response.Errors, model.DSLError{
			Code:    "DSL_PARSE_ERROR", 
			Message: "Empty expression",
		})
		return response
	}
	
	if strings.Contains(normalized, " AND ") || strings.Contains(normalized, " OR ") {
		response.Errors = append(response.Errors, model.DSLError{
			Code:    "DSL_UNSUPPORTED_TIER",
			Message: "AND/OR not implemented yet (requires Tier 3)",
		})
		return response
	}
	
	if strings.Contains(normalized, " NOT ") || strings.Contains(normalized, "(") || strings.Contains(normalized, ")") {
		response.Errors = append(response.Errors, model.DSLError{
			Code:    "DSL_UNSUPPORTED_TIER", 
			Message: "NOT and parentheses not implemented yet (requires Tier 4)",
		})
		return response
	}
	
	field, operator, value := e.parseSimpleComparison(normalized)
	if field == "" || operator == "" || value == "" {
		response.Errors = append(response.Errors, model.DSLError{
			Code:    "DSL_PARSE_ERROR",
			Message: "Expected expression format: field operator value",
		})
		return response
	}
	
	supportedFields := []string{"amount", "user.age", "user.score"}
	supportedOperators := []string{">", ">=", "<", "<=", "=", "!="}
	stringFields := []string{}
	
	fieldSupported := false
	for _, f := range supportedFields {
		if field == f {
			fieldSupported = true
			break
		}
	}
	if !fieldSupported {
		response.Errors = append(response.Errors, model.DSLError{
			Code:    "DSL_INVALID_FIELD",
			Message: fmt.Sprintf("Field '%s' not supported. Supported: %s", field, strings.Join(supportedFields, ", ")),
		})
		return response
	}
	
	operatorSupported := false
	for _, op := range supportedOperators {
		if operator == op {
			operatorSupported = true
			break
		}
	}
	if !operatorSupported {
		response.Errors = append(response.Errors, model.DSLError{
			Code:    "DSL_INVALID_OPERATOR",
			Message: fmt.Sprintf("Operator '%s' not supported. Supported: %s", operator, strings.Join(supportedOperators, ", ")),
		})
		return response
	}
	
	isStringField := false
	for _, sf := range stringFields {
		if field == sf {
			isStringField = true
			break
		}
	}
	
	if isStringField {
		if operator != "=" && operator != "!=" {
			response.Errors = append(response.Errors, model.DSLError{
				Code:    "DSL_INVALID_OPERATOR",
				Message: fmt.Sprintf("String fields only support '=' and '!='. Field '%s' doesn't support operator '%s'", field, operator),
			})
			return response
		}
		
		if !strings.HasPrefix(value, "'") || !strings.HasSuffix(value, "'") {
			response.Errors = append(response.Errors, model.DSLError{
				Code:    "DSL_PARSE_ERROR",
				Message: fmt.Sprintf("String values must be in single quotes. Expected: '%s'", value),
			})
			return response
		}
	} else {
		// Validate numeric fields
		if field == "amount" || field == "user.age" || field == "user.score" {
			if _, err := strconv.ParseFloat(value, 64); err != nil {
				response.Errors = append(response.Errors, model.DSLError{
					Code:    "DSL_PARSE_ERROR",
					Message: fmt.Sprintf("Expected number after '%s', got '%s'", operator, value),
				})
				return response
			}
		}
	}
	
	response.IsValid = true
	response.NormalizedExpression = &normalized
	return response
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

func (e *DSLEvaluator) normalizeExpression(expression string) string {
	// Remove extra whitespace and trim
	normalized := strings.Join(strings.Fields(expression), " ")
	
	// Add spaces around operators
	for _, op := range []string{">=", "<=", "!=", ">", "<", "="} {
		normalized = strings.ReplaceAll(normalized, op, " "+op+" ")
	}
	
	// Clean up multiple spaces again
	normalized = strings.Join(strings.Fields(normalized), " ")
	
	return strings.TrimSpace(normalized)
}

func (e *DSLEvaluator) parseSimpleComparison(expression string) (string, string, string) {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return "", "", ""
	}
	
	return parts[0], parts[1], parts[2]
}

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
		return false, nil
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
		return false, nil
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
