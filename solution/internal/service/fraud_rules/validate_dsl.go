package fraud_rules

import (
	"context"
	"fmt"
	"solution/internal/model"
	"strings"
)

func (s *service) ValidateDSL(ctx context.Context, dsl string) (*model.FraudRuleValidateResponse, error) {
	dsl = strings.TrimSpace(dsl)
	if len(dsl) == 0 {
		return &model.FraudRuleValidateResponse{
			IsValid: false,
			Error:   "DSL cannot be empty",
		}, nil
	}

	if len(dsl) > model.MaxDSLSize {
		return &model.FraudRuleValidateResponse{
			IsValid: false,
			Error:   fmt.Sprintf("DSL too large (max %d characters)", model.MaxDSLSize),
		}, nil
	}

	dangerousWords := []string{"DROP", "DELETE", "TRUNCATE", "INSERT", "UPDATE"}
	dslUpper := strings.ToUpper(dsl)
	for _, word := range dangerousWords {
		if strings.Contains(dslUpper, word) {
			return &model.FraudRuleValidateResponse{
				IsValid: false,
				Error:   fmt.Sprintf("DSL contains dangerous keyword: %s", word),
			}, nil
		}
	}

	return &model.FraudRuleValidateResponse{
		IsValid: true,
		AST:     "AST parsing not implemented yet",
	}, nil
}
