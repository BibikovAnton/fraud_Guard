package fraud_rules

import (
	"context"
	"fmt"
	"solution/internal/model"
)

func (s *service) GetByID(ctx context.Context, id string) (*model.FraudRule, error) {
	rule, err := s.fraudRuleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get fraud rule: %w", err)
	}
	return rule, nil
}
