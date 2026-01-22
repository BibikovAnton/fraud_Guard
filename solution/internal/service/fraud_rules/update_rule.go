package fraud_rules

import (
	"context"
	"fmt"
	"solution/internal/model"
)

func (s *service) Update(ctx context.Context, id string, req model.FraudRuleUpdateRequest) (*model.FraudRule, error) {
	existing, err := s.fraudRuleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing rule: %w", err)
	}
	if existing == nil {
		return nil, fmt.Errorf("fraud rule not found")
	}

	if req.DSL != nil && *req.DSL != existing.DSL {
		validation, err := s.ValidateDSL(ctx, *req.DSL)
		if err != nil {
			return nil, fmt.Errorf("DSL validation failed: %w", err)
		}
		if !validation.IsValid {
			return nil, fmt.Errorf("invalid DSL: %s", validation.Error)
		}
	}

	if req.Priority != nil && *req.Priority < 1 {
		return nil, fmt.Errorf("priority must be >= 1")
	}

	updated, err := s.fraudRuleRepo.Update(ctx, id, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update fraud rule: %w", err)
	}

	return updated, nil
}
