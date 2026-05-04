package fraud_rules

import (
	"context"
	"fmt"
)

func (s *service) Delete(ctx context.Context, id string) error {
	existing, err := s.fraudRuleRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to check rule existence: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("fraud rule not found")
	}

	if err := s.fraudRuleRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete fraud rule: %w", err)
	}

	return nil
}
