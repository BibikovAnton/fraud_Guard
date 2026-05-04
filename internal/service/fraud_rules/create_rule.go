package fraud_rules

import (
	"context"
	"fmt"
	"solution/internal/model"
	"strings"
)

func (s *service) Create(ctx context.Context, req model.FraudRuleCreateRequest) (*model.FraudRule, error) {
	

	priority := model.DefaultPriority
	if req.Priority != nil {
		priority = *req.Priority
		if priority < 1 {
			return nil, fmt.Errorf("priority must be >= 1")
		}
	}

	enabled := true
	if req.Enabled != nil {
		enabled = *req.Enabled
	}

	rule := model.NewFraudRule(
		strings.TrimSpace(req.Name),
		strings.TrimSpace(req.Description),
		strings.TrimSpace(req.DslExpression),
		priority,
		enabled,
	)

	if err := s.fraudRuleRepo.Create(ctx, rule); err != nil {
		return nil, fmt.Errorf("failed to create fraud rule: %w", err)
	}

	created, err := s.fraudRuleRepo.GetByName(ctx, rule.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created rule: %w", err)
	}

	return created, nil
}
