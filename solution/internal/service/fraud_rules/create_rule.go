package fraud_rules

import (
	"context"
	"fmt"
	"solution/internal/model"
	"strings"
	"time"
)

func (s *service) Create(ctx context.Context, req model.FraudRuleCreateRequest) (*model.FraudRule, error) {
	validation, err := s.ValidateDSL(ctx, req.DslExpression)
	if err != nil {
		return nil, fmt.Errorf("DSL validation failed: %w", err)
	}
	if !validation.IsValid {
		return nil, fmt.Errorf("invalid DSL: %s", validation.Errors)
	}

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

	now := time.Now()
	rule := model.FraudRule{
		Name:          strings.TrimSpace(req.Name),
		Description:   strings.TrimSpace(req.Description),
		DslExpression: strings.TrimSpace(req.DslExpression),
		Priority:      priority,
		Enabled:       enabled,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := s.fraudRuleRepo.Create(ctx, rule); err != nil {
		return nil, fmt.Errorf("failed to create fraud rule: %w", err)
	}

	created, err := s.fraudRuleRepo.GetByName(ctx, rule.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created rule: %w", err)
	}

	return created, nil
}
