package fraud_rules

import (
	"context"
	"solution/internal/model"
	"solution/internal/service/dsl"
)

func (s *service) ValidateDSL(ctx context.Context, dslStr string) (*model.DslValidateResponse, error) {
	evaluator := dsl.NewEvaluator(1) // Уровень 1 - базовый парсер
	result, err := evaluator.Validate(ctx, dslStr)
	if err != nil {
		return nil, err
	}
	return result, nil
}
