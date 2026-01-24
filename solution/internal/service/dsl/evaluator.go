package dsl

import (
	"context"
	"fmt"
	"solution/internal/model"
	"strings"
)

type Evaluator interface {
	Evaluate(ctx context.Context, dsl string, transaction *model.Transaction, user *model.User) (bool, string, error)
	Validate(ctx context.Context, dsl string) (*model.DslValidateResponse, error)
}

type evaluator struct {
	tier int // Уровень поддержки DSL (0-5)
}

func NewEvaluator(tier int) Evaluator {
	return &evaluator{tier: tier}
}

func (e *evaluator) Evaluate(ctx context.Context, dsl string, transaction *model.Transaction, user *model.User) (bool, string, error) {
	if e.tier == 0 {
		return false, "DSL not supported (Tier 0)", nil
	}

	normalized := e.normalize(dsl)

	result, err := e.evaluateExpression(normalized, transaction, user)
	if err != nil {
		return false, fmt.Sprintf("DSL evaluation error: %s", err.Error()), nil
	}

	description := e.buildDescription(normalized, result)
	return result, description, nil
}

func (e *evaluator) Validate(ctx context.Context, dsl string) (*model.DslValidateResponse, error) {
	dsl = strings.TrimSpace(dsl)
	
	if len(dsl) == 0 {
		return &model.DslValidateResponse{
			IsValid: false,
			Errors: []model.DSLError{
				{
					Code:     "DSL_PARSE_ERROR",
					Message:  "Empty expression",
					Position: &[]int{0}[0],
					Near:     &[]string{""}[0],
				},
			},
		}, nil
	}

	// Normalize expression like Python version
	normalized := e.normalize(dsl)
	
	if len(normalized) == 0 {
		return &model.DslValidateResponse{
			IsValid: false,
			Errors: []model.DSLError{
				{
					Code:     "DSL_PARSE_ERROR",
					Message:  "Empty expression",
					Position: &[]int{0}[0],
					Near:     &[]string{""}[0],
				},
			},
		}, nil
	}

	// Check tier-specific restrictions
	if e.tier == 0 {
		return &model.DslValidateResponse{
			IsValid: false,
			Errors: []model.DSLError{
				{
					Code:    "DSL_UNSUPPORTED_TIER",
					Message: "DSL tier 0 - expressions not supported in current implementation",
				},
			},
		}, nil
	}
	
	if e.tier == 1 {
		return e.validateTier1(normalized)
	}
	
	if e.tier == 2 {
		return e.validateTier2(normalized)
	}
	
	// Higher tiers not implemented
	return &model.DslValidateResponse{
		IsValid: false,
		Errors: []model.DSLError{
			{
				Code:    "DSL_UNSUPPORTED_TIER",
				Message: fmt.Sprintf("DSL tier %d not implemented yet", e.tier),
			},
		},
	}, nil
}

func (e *evaluator) validateTier1(normalized string) (*model.DslValidateResponse, error) {
	// Tier 1: only amount field supported
	field, operator, value := e.parseSimpleComparison(normalized)
	if field == "" || operator == "" || value == "" {
		return &model.DslValidateResponse{
			IsValid: false,
			Errors: []model.DSLError{
				{
					Code:     "DSL_PARSE_ERROR",
					Message:  "Expected expression format: field operator value",
					Position: &[]int{0}[0],
					Near:     &[]string{normalized}[0],
				},
			},
		}, nil
	}

	// Only amount field supported in tier 1
	if field != "amount" {
		return &model.DslValidateResponse{
			IsValid: false,
			Errors: []model.DSLError{
				{
					Code:     "DSL_INVALID_FIELD",
					Message:  "Field 'amount' not supported. Supported: amount",
					Position: &[]int{0}[0],
					Near:     &[]string{field}[0],
				},
			},
		}, nil
	}

	return e.validateComparison(field, operator, value, normalized)
}

func (e *evaluator) validateTier2(normalized string) (*model.DslValidateResponse, error) {
	// Tier 2: supports AND/OR without parentheses
	if strings.Contains(normalized, " AND ") || strings.Contains(normalized, " OR ") {
		// Python version: just clean parentheses and return true
		cleaned := strings.ReplaceAll(normalized, "(", "")
		cleaned = strings.ReplaceAll(cleaned, ")", "")
		cleaned = strings.ReplaceAll(cleaned, " and ", " AND ")
		cleaned = strings.ReplaceAll(cleaned, " or ", " OR ")
		
		return &model.DslValidateResponse{
			IsValid:            true,
			NormalizedExpression: &cleaned,
			Errors:             []model.DSLError{},
		}, nil
	}

	// For simple expressions without AND/OR, use tier 1 logic
	// But first check if it has parentheses - if so, clean them and try again
	simple := strings.ReplaceAll(normalized, "(", "")
	simple = strings.ReplaceAll(simple, ")", "")
	
	if strings.Contains(simple, " AND ") || strings.Contains(simple, " OR ") {
		// After removing parentheses we see AND/OR
		cleaned := strings.ReplaceAll(simple, " and ", " AND ")
		cleaned = strings.ReplaceAll(cleaned, " or ", " OR ")
		
		return &model.DslValidateResponse{
			IsValid:            true,
			NormalizedExpression: &cleaned,
			Errors:             []model.DSLError{},
		}, nil
	}

	// For truly simple expressions, use tier 1 logic
	return e.validateTier1(normalized)
}

func (e *evaluator) validateComparison(field, operator, value, normalized string) (*model.DslValidateResponse, error) {
	supportedFields := []string{"amount", "currency", "merchantId", "ipAddress", "deviceId"}
	supportedOperators := []string{">", ">=", "<", "<=", "=", "!="}
	stringFields := []string{"currency", "merchantId", "ipAddress", "deviceId"}

	// Check field
	fieldSupported := false
	for _, f := range supportedFields {
		if field == f {
			fieldSupported = true
			break
		}
	}
	if !fieldSupported {
		return &model.DslValidateResponse{
			IsValid: false,
			Errors: []model.DSLError{
				{
					Code:     "DSL_INVALID_FIELD",
					Message:  fmt.Sprintf("Field '%s' not supported. Supported: %s", field, strings.Join(supportedFields, ", ")),
					Position: &[]int{0}[0],
					Near:     &[]string{field}[0],
				},
			},
		}, nil
	}

	// Check operator
	operatorSupported := false
	for _, op := range supportedOperators {
		if operator == op {
			operatorSupported = true
			break
		}
	}
	if !operatorSupported {
		return &model.DslValidateResponse{
			IsValid: false,
			Errors: []model.DSLError{
				{
					Code:     "DSL_INVALID_OPERATOR",
					Message:  fmt.Sprintf("Operator '%s' not supported. Supported: %s", operator, strings.Join(supportedOperators, ", ")),
					Position: &[]int{len(field) + len(operator) + 1}[0],
					Near:     &[]string{operator}[0],
				},
			},
		}, nil
	}

	// Check string field restrictions
	isStringField := false
	for _, sf := range stringFields {
		if field == sf {
			isStringField = true
			break
		}
	}
	
	if isStringField {
		if operator != "=" && operator != "!=" {
			return &model.DslValidateResponse{
				IsValid: false,
				Errors: []model.DSLError{
					{
						Code:     "DSL_INVALID_OPERATOR",
						Message:  fmt.Sprintf("String fields only support '=' and '!='. Field '%s' doesn't support operator '%s'", field, operator),
						Position: &[]int{len(field) + len(operator) + 1}[0],
						Near:     &[]string{operator}[0],
					},
				},
			}, nil
		}

		// String values must be in single quotes
		if !(strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'")) {
			return &model.DslValidateResponse{
				IsValid: false,
				Errors: []model.DSLError{
					{
						Code:     "DSL_PARSE_ERROR",
						Message:  fmt.Sprintf("String values must be in single quotes. Expected: '%s'", value),
						Position: &[]int{len(field) + len(operator) + 1}[0],
						Near:     &[]string{value}[0],
					},
				},
			}, nil
		}
	} else if field == "amount" {
		// Amount must be numeric
		_, err := fmt.Sscanf(value, "%f", new(float64))
		if err != nil {
			return &model.DslValidateResponse{
				IsValid: false,
				Errors: []model.DSLError{
					{
						Code:     "DSL_PARSE_ERROR",
						Message:  fmt.Sprintf("Expected number after '%s', got '%s'", operator, value),
						Position: &[]int{len(field) + len(operator) + 1}[0],
						Near:     &[]string{value}[0],
					},
				},
			}, nil
		}
	}

	// Check complexity (nodes count)
	nodeCount := e.countNodes(normalized)
	if nodeCount > 100 {
		return &model.DslValidateResponse{
			IsValid: false,
			Errors: []model.DSLError{
				{
					Code:     "DSL_TOO_COMPLEX",
					Message:  "Expression too complex (exceeds 100 nodes)",
					Position: &[]int{0}[0],
					Near:     &[]string{normalized}[0],
				},
			},
		}, nil
	}

	return &model.DslValidateResponse{
		IsValid:            true,
		NormalizedExpression: &normalized,
		Errors:             []model.DSLError{},
	}, nil
}

func (e *evaluator) normalize(dsl string) string {
	// Python version: just normalize spaces and add spaces around operators
	normalized := strings.Join(strings.Fields(dsl), " ") // Remove extra spaces
	
	// Add spaces around operators
	operators := []string{">=", "<=", "!=", ">", "<", "="}
	for _, op := range operators {
		// Replace operator with space-padded version
		normalized = strings.ReplaceAll(normalized, op, " "+op+" ")
	}
	
	// Clean up extra spaces again
	normalized = strings.Join(strings.Fields(normalized), " ")
	
	return normalized
}

func (e *evaluator) parseSimpleComparison(expression string) (string, string, string) {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return "", "", ""
	}
	
	field := parts[0]
	operator := parts[1]
	value := parts[2]
	
	return field, operator, value
}

func (e *evaluator) evaluateExpression(expr string, transaction *model.Transaction, user *model.User) (bool, error) {
	if e.tier == 1 {
		return e.evaluateComparison(expr, transaction, user)
	}

	return e.evaluateFullExpression(expr, transaction, user)
}

func (e *evaluator) evaluateComparison(expr string, transaction *model.Transaction, user *model.User) (bool, error) {
	parts := strings.Fields(expr)
	if len(parts) != 3 {
		return false, fmt.Errorf("invalid comparison format")
	}

	field := parts[0]
	operator := parts[1]
	value := parts[2]

	return e.compareValues(field, operator, value, transaction, user)
}

func (e *evaluator) compareValues(field, operator, value string, transaction *model.Transaction, user *model.User) (bool, error) {
	fieldValue, err := e.getFieldValue(field, transaction, user)
	if err != nil {
		return false, err
	}

	dslValue, err := e.parseValue(value)
	if err != nil {
		return false, err
	}

	switch operator {
	case ">":
		return e.compareGreaterThan(fieldValue, dslValue)
	case ">=":
		return e.compareGreaterThanOrEqual(fieldValue, dslValue)
	case "<":
		return e.compareLessThan(fieldValue, dslValue)
	case "<=":
		return e.compareLessThanOrEqual(fieldValue, dslValue)
	case "=":
		return e.compareEqual(fieldValue, dslValue), nil
	case "!=":
		return e.compareNotEqual(fieldValue, dslValue)
	default:
		return false, fmt.Errorf("unsupported operator: %s", operator)
	}
}

func (e *evaluator) getFieldValue(field string, transaction *model.Transaction, user *model.User) (interface{}, error) {
	switch field {
	case "amount", "AMOUNT":
		return transaction.Amount, nil
	case "currency", "CURRENCY":
		return string(transaction.Currency), nil
	case "merchantId", "MERCHANTID", "MERCHANT_ID":
		if transaction.MerchantID == nil {
			return "", nil
		}
		return *transaction.MerchantID, nil
	case "ipAddress", "IPADDRESS", "IP_ADDRESS":
		if transaction.IPAddress == nil {
			return "", nil
		}
		return transaction.IPAddress.String(), nil
	case "deviceId", "DEVICEID", "DEVICE_ID":
		if transaction.DeviceID == nil {
			return "", nil
		}
		return *transaction.DeviceID, nil
	case "user.age", "user.AGE", "USER.age", "USER.AGE":
		if user.Age == nil {
			return nil, nil // null - правило не сработает
		}
		return float64(*user.Age), nil
	case "user.region", "user.REGION", "USER.region", "USER.REGION":
		if user.Region == nil {
			return nil, nil // null - правило не сработает
		}
		return *user.Region, nil
	default:
		return nil, fmt.Errorf("unknown field: %s", field)
	}
}

func (e *evaluator) parseValue(value string) (interface{}, error) {
	if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
		return strings.Trim(value, "'"), nil
	}

	var num float64
	_, err := fmt.Sscanf(value, "%f", &num)
	if err != nil {
		return nil, fmt.Errorf("invalid value: %s", value)
	}
	return num, nil
}

func (e *evaluator) compareGreaterThan(a, b interface{}) (bool, error) {
	af, ok1 := a.(float64)
	bf, ok2 := b.(float64)
	if !ok1 || !ok2 {
		return false, fmt.Errorf("comparison requires numeric values")
	}
	return af > bf, nil
}

func (e *evaluator) compareGreaterThanOrEqual(a, b interface{}) (bool, error) {
	af, ok1 := a.(float64)
	bf, ok2 := b.(float64)
	if !ok1 || !ok2 {
		return false, fmt.Errorf("comparison requires numeric values")
	}
	return af >= bf, nil
}

func (e *evaluator) compareLessThan(a, b interface{}) (bool, error) {
	af, ok1 := a.(float64)
	bf, ok2 := b.(float64)
	if !ok1 || !ok2 {
		return false, fmt.Errorf("comparison requires numeric values")
	}
	return af < bf, nil
}

func (e *evaluator) compareLessThanOrEqual(a, b interface{}) (bool, error) {
	af, ok1 := a.(float64)
	bf, ok2 := b.(float64)
	if !ok1 || !ok2 {
		return false, fmt.Errorf("comparison requires numeric values")
	}
	return af <= bf, nil
}

func (e *evaluator) compareEqual(a, b interface{}) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func (e *evaluator) compareNotEqual(a, b interface{}) (bool, error) {
	return !e.compareEqual(a, b), nil
}

func (e *evaluator) validateSyntax(expr string) *model.DSLError {
	dangerous := []string{"DROP", "DELETE", "TRUNCATE", "INSERT", "UPDATE"}
	exprUpper := strings.ToUpper(expr)
	for _, word := range dangerous {
		if strings.Contains(exprUpper, word) {
			return &model.DSLError{
				Code:    "DSL_PARSE_ERROR",
				Message: fmt.Sprintf("Dangerous keyword: %s", word),
			}
		}
	}

	if !strings.Contains(expr, ">") && !strings.Contains(expr, "<") && !strings.Contains(expr, "=") {
		return &model.DSLError{
			Code:    "DSL_PARSE_ERROR",
			Message: "No comparison operator found",
		}
	}

	return nil
}

func (e *evaluator) countNodes(expr string) int {
	tokens := strings.Fields(expr)
	return len(tokens)
}

func (e *evaluator) buildDescription(expr string, matched bool) string {
	if matched {
		return fmt.Sprintf("%s, rule matched", expr)
	}
	return fmt.Sprintf("%s, rule not matched", expr)
}

func (e *evaluator) evaluateFullExpression(expr string, transaction *model.Transaction, user *model.User) (bool, error) {
	return e.evaluateComparison(expr, transaction, user)
}
