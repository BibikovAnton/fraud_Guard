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
					Code:    "DSL_PARSE_ERROR",
					Message: "DSL cannot be empty",
				},
			},
		}, nil
	}

	if len(dsl) > model.MaxDSLSize {
		return &model.DslValidateResponse{
			IsValid: false,
			Errors: []model.DSLError{
				{
					Code:    "DSL_TOO_COMPLEX",
					Message: fmt.Sprintf("DSL too large (max %d characters)", model.MaxDSLSize),
				},
			},
		}, nil
	}

	if e.tier == 0 {
		return &model.DslValidateResponse{
			IsValid: false,
			Errors: []model.DSLError{
				{
					Code:    "DSL_UNSUPPORTED_TIER",
					Message: "DSL not supported at current tier",
				},
			},
		}, nil
	}

	normalized := e.normalize(dsl)
	
	if e.tier >= 1 {
		if err := e.validateSyntax(normalized); err != nil {
			return &model.DslValidateResponse{
				IsValid:            false,
				NormalizedExpression: &normalized,
				Errors:             []model.DSLError{*err},
			}, nil
		}
	}

	if e.tier >= 1 {
		if nodeCount := e.countNodes(normalized); nodeCount > model.MaxASTNodes {
			return &model.DslValidateResponse{
				IsValid:            false,
				NormalizedExpression: &normalized,
				Errors: []model.DSLError{
					{
						Code:    "DSL_TOO_COMPLEX",
						Message: fmt.Sprintf("Expression too complex (max %d nodes, got %d)", model.MaxASTNodes, nodeCount),
					},
				},
			}, nil
		}
	}

	return &model.DslValidateResponse{
		IsValid:            true,
		NormalizedExpression: &normalized,
		Errors:             []model.DSLError{},
	}, nil
}

func (e *evaluator) normalize(dsl string) string {
	// Convert to lowercase first
	dsl = strings.ToLower(dsl)
	
	// Then convert operators to uppercase
	dsl = strings.ReplaceAll(dsl, "and", " AND ")
	dsl = strings.ReplaceAll(dsl, "or", " OR ")
	dsl = strings.ReplaceAll(dsl, "not", " NOT ")
	
	// Clean up extra spaces
	dsl = strings.Join(strings.Fields(dsl), " ")
	
	return dsl
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
