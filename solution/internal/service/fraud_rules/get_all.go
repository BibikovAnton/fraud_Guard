package fraud_rules

import (
	"context"
	"fmt"
	"solution/internal/model"
)

func (s *service) GetAll(ctx context.Context, activeOnly bool) ([]*model.FraudRule, error) {
	rules, err := s.fraudRuleRepo.GetAll(ctx, activeOnly)
	if err != nil {
		return nil, fmt.Errorf("failed to get fraud rules: %w", err)
	}
	return rules, nil
}
