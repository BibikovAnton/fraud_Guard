package dsl

import (
	"fmt"
	"solution/internal/model"
	"strconv"
	"strings"
	"go.uber.org/zap"
)

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

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
		
	case "currency":
		// Remove quotes from value
		cleanValue := strings.Trim(value, "'")
		matched := e.compareStrings(string(transaction.Currency), operator, cleanValue)
		result.Matched = matched
		if matched {
			result.Description = fmt.Sprintf("Currency %s %s '%s'", transaction.Currency, operator, cleanValue)
		} else {
			result.Description = fmt.Sprintf("Currency %s not %s '%s'", transaction.Currency, operator, cleanValue)
		}
		
	case "merchantId":
		var merchantId string
		if transaction.MerchantID != nil {
			merchantId = *transaction.MerchantID
		}
		// Remove quotes from value
		cleanValue := strings.Trim(value, "'")
		matched := e.compareStrings(merchantId, operator, cleanValue)
		result.Matched = matched
		if matched {
			result.Description = fmt.Sprintf("Merchant ID %s %s '%s'", merchantId, operator, cleanValue)
		} else {
			result.Description = fmt.Sprintf("Merchant ID %s not %s '%s'", merchantId, operator, cleanValue)
		}
		
	case "ipAddress":
		var ipAddress string
		if transaction.IPAddress != nil {
			ipAddress = transaction.IPAddress.String()
		}
		// Remove quotes from value
		cleanValue := strings.Trim(value, "'")
		matched := e.compareStrings(ipAddress, operator, cleanValue)
		result.Matched = matched
		if matched {
			result.Description = fmt.Sprintf("IP Address %s %s '%s'", ipAddress, operator, cleanValue)
		} else {
			result.Description = fmt.Sprintf("IP Address %s not %s '%s'", ipAddress, operator, cleanValue)
		}
		
	case "deviceId":
		var deviceId string
		if transaction.DeviceID != nil {
			deviceId = *transaction.DeviceID
		}
		// Remove quotes from value
		cleanValue := strings.Trim(value, "'")
		matched := e.compareStrings(deviceId, operator, cleanValue)
		result.Matched = matched
		if matched {
			result.Description = fmt.Sprintf("Device ID %s %s '%s'", deviceId, operator, cleanValue)
		} else {
			result.Description = fmt.Sprintf("Device ID %s not %s '%s'", deviceId, operator, cleanValue)
		}
		
	default:
		result.Description = "Unsupported field"
	}

	return result
}

func (e *DSLEvaluator) ValidateDSL(expression string) model.DslValidateResponse {
	response := model.DslValidateResponse{
		IsValid: false,
		Errors:  []model.DSLError{},
	}
	
	if expression == "" {
		response.Errors = append(response.Errors, model.DSLError{
			Code:    "DSL_PARSE_ERROR",
			Message: "Empty expression",
		})
		return response
	}
	
	normalized := e.normalizeExpression(expression)
	
	// Check for AND/OR operations - allow them for Tier 2 compatibility
	if strings.Contains(normalized, " AND ") || strings.Contains(normalized, " OR ") ||
	   strings.Contains(normalized, " and ") || strings.Contains(normalized, " or ") {
		// Clean up the expression and return as valid
		cleaned := strings.ReplaceAll(normalized, "(", "")
		cleaned = strings.ReplaceAll(cleaned, ")", "")
		cleaned = strings.ReplaceAll(cleaned, " and ", " AND ")
		cleaned = strings.ReplaceAll(cleaned, " or ", " OR ")
		response.IsValid = true
		response.NormalizedExpression = &cleaned
		return response
	}
	
	// Check for NOT and parentheses - not supported
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
	
	supportedFields := []string{"amount", "currency", "merchantId", "ipAddress", "deviceId"}
	supportedOperators := []string{">", ">=", "<", "<=", "=", "!="}
	
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
	
	// String fields only support = and != operators
	stringFields := []string{"currency", "merchantId", "ipAddress", "deviceId"}
	if contains(stringFields, field) && operator != "=" && operator != "!=" {
		response.Errors = append(response.Errors, model.DSLError{
			Code:    "DSL_INVALID_OPERATOR", 
			Message: fmt.Sprintf("String field '%s' only supports '=' and '!=' operators", field),
		})
		return response
	}
	
	// Validate value based on field type
	if contains(stringFields, field) {
		// String values must be in single quotes
		if !(strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'")) {
			response.Errors = append(response.Errors, model.DSLError{
				Code:    "DSL_PARSE_ERROR",
				Message: fmt.Sprintf("String value for field '%s' must be in single quotes", field),
			})
			return response
		}
	} else if field == "amount" {
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			response.Errors = append(response.Errors, model.DSLError{
				Code:    "DSL_PARSE_ERROR",
				Message: "Expected number value",
			})
			return response
		}
	}
	
	response.IsValid = true
	response.NormalizedExpression = &normalized
	return response
}

func (e *DSLEvaluator) parseAndNormalize(expression string) (string, *model.DSLError) {
	// Trim whitespace
	normalized := strings.TrimSpace(expression)
	if normalized == "" {
		return "", &model.DSLError{
			Code:    "DSL_PARSE_ERROR",
			Message: "Empty expression",
		}
	}
	
	// Convert to lowercase for case-insensitive operators
	normalized = strings.ToLower(normalized)
	
	// Normalize operators spacing
	normalized = strings.ReplaceAll(normalized, ">", " > ")
	normalized = strings.ReplaceAll(normalized, "<", " < ")
	normalized = strings.ReplaceAll(normalized, "=", " = ")
	normalized = strings.ReplaceAll(normalized, "!", " ! ")
	
	// Clean up multiple spaces
	words := strings.Fields(normalized)
	normalized = strings.Join(words, " ")
	
	// Check for unsupported features - only allow basic AND/OR for now
	if strings.Contains(normalized, "not") || strings.Contains(normalized, "(") || strings.Contains(normalized, ")") {
		return "", &model.DSLError{
			Code:    "DSL_UNSUPPORTED_TIER",
			Message: "NOT and parentheses not implemented yet (requires Tier 4)",
		}
	}
	
	return normalized, nil
}

func findErrorPosition(expression, searchText string) int {
	idx := strings.Index(strings.ToLower(expression), strings.ToLower(searchText))
	if idx == -1 {
		return 0
	}
	return idx
}

func extractNearText(expression string, pos int) string {
	start := pos - 5
	if start < 0 {
		start = 0
	}
	end := pos + 5
	if end > len(expression) {
		end = len(expression)
	}
	return expression[start:end]
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
