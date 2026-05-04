package dsl

import (
	"fmt"
	"go.uber.org/zap"
	"solution/internal/model"
	"strconv"
	"strings"
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
	fmt.Printf("DEBUG: Evaluating rule %s: %s\n", rule.Name, rule.DslExpression)
	fmt.Printf("DEBUG: Transaction amount: %.2f, user age: %v\n", transaction.Amount, func() interface{} {
		if user != nil && user.Age != nil {
			return *user.Age
		}
		return "nil"
	}())

	result := model.RuleResult{
		RuleID:      "",
		RuleName:    rule.Name,
		Priority:    rule.Priority,
		Enabled:     rule.Enabled,
		Matched:     false,
		Description: "",
	}

	if !rule.Enabled {
		result.Description = "Rule is disabled"
		result.RuleID = ""
		return result
	}

	dsl := strings.ToLower(strings.TrimSpace(rule.DslExpression))

	field, operator, value := e.parseSimpleComparison(dsl)
	if field == "" || operator == "" || value == "" {
		result.Description = "Invalid DSL expression"
		result.RuleID = ""
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
				matched := e.compareFloats(float64(*user.Age), operator, threshold)
				result.Matched = matched
				if matched {
					result.Description = fmt.Sprintf("User age %d %s %.2f", *user.Age, operator, threshold)
				} else {
					result.Description = fmt.Sprintf("User age %d not %s %.2f", *user.Age, operator, threshold)
				}
			} else {
				result.Description = "User age not available"
			}
		} else {
			result.Description = "Invalid user.age value"
		}

	case "user.region":
		var userRegion string
		if user != nil && user.Region != nil {
			userRegion = *user.Region
		}

		cleanValue := strings.Trim(value, "'")
		matched := e.compareStrings(userRegion, operator, cleanValue)
		result.Matched = matched
		if matched {
			result.Description = fmt.Sprintf("User region '%s' %s '%s'", userRegion, operator, cleanValue)
		} else {
			result.Description = fmt.Sprintf("User region '%s' not %s '%s'", userRegion, operator, cleanValue)
		}

	case "currency":

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

	if strings.Contains(normalized, " AND ") || strings.Contains(normalized, " OR ") ||
		strings.Contains(normalized, " and ") || strings.Contains(normalized, " or ") {

		cleaned := strings.ReplaceAll(normalized, "(", "")
		cleaned = strings.ReplaceAll(cleaned, ")", "")
		cleaned = strings.ReplaceAll(cleaned, " and ", " AND ")
		cleaned = strings.ReplaceAll(cleaned, " or ", " OR ")
		response.IsValid = true
		response.NormalizedExpression = &cleaned
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

	supportedFields := []string{"amount", "currency", "merchantId", "ipAddress", "deviceId", "user.age", "user.region"}
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

	stringFields := []string{"currency", "merchantId", "ipAddress", "deviceId", "user.region"}

	if contains(stringFields, field) && operator != "=" && operator != "!=" {
		response.Errors = append(response.Errors, model.DSLError{
			Code:    "DSL_INVALID_OPERATOR",
			Message: fmt.Sprintf("String field '%s' only supports '=' and '!=' operators", field),
		})
		return response
	}

	if contains(stringFields, field) {

		if !(strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'")) {
			response.Errors = append(response.Errors, model.DSLError{
				Code:    "DSL_PARSE_ERROR",
				Message: fmt.Sprintf("String value for field '%s' must be in single quotes", field),
			})
			return response
		}
	} else if field == "amount" || field == "user.age" {
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

	normalized := strings.TrimSpace(expression)
	if normalized == "" {
		return "", &model.DSLError{
			Code:    "DSL_PARSE_ERROR",
			Message: "Empty expression",
		}
	}

	normalized = strings.ToLower(normalized)

	normalized = strings.ReplaceAll(normalized, ">", " > ")
	normalized = strings.ReplaceAll(normalized, "<", " < ")
	normalized = strings.ReplaceAll(normalized, "=", " = ")
	normalized = strings.ReplaceAll(normalized, "!", " ! ")

	words := strings.Fields(normalized)
	normalized = strings.Join(words, " ")

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

	normalized := strings.Join(strings.Fields(expression), " ")

	for _, op := range []string{">=", "<=", "!=", ">", "<", "="} {
		normalized = strings.ReplaceAll(normalized, op, " "+op+" ")
	}

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
		return strings.ToLower(a) == strings.ToLower(b)
	case "!=":
		return strings.ToLower(a) != strings.ToLower(b)
	default:
		return false
	}
}
