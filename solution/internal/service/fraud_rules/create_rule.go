package fraud_rules

import (
	"context"
	"fmt"
	"solution/internal/model"
	"strings"
	"time"
)

func (s *service) Create(ctx context.Context, req model.FraudRuleCreateRequest) (*model.FraudRule, error) {
	validation, err := s.ValidateDSL(ctx, req.DSL)
	if err != nil {
		return nil, fmt.Errorf("DSL validation failed: %w", err)
	}
	if !validation.IsValid {
		return nil, fmt.Errorf("invalid DSL: %s", validation.Error)
	}

	if req.Priority == nil {
		defaultPriority := 100
		req.Priority = &defaultPriority
	}
	priority := model.DefaultPriority
	if req.Priority != nil {
		priority = *req.Priority
		if priority < 1 {
			return nil, fmt.Errorf("priority must be >= 1")
		}
	}

	now := time.Now()
	rule := model.FraudRule{
		Name:        strings.TrimSpace(req.Name),
		Description: strings.TrimSpace(req.Description),
		DSL:         strings.TrimSpace(req.DSL),
		Priority:    priority,
		IsActive:    true,
		CreatedAt:   now,
		UpdatedAt:   now,
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
