package fraud_rules

import (
	"context"
	"solution/internal/model"
)

func (s *service) ValidateDSL(ctx context.Context, req model.DslValidateRequest) model.DslValidateResponse {
	return s.dslEvaluator.ValidateDSL(req.DslExpression)
}
