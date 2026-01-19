package fraud_rules

import (
	"context"
	"fmt"
	"solution/internal/model"

	"github.com/google/uuid"
)

func (r *repository) Create(ctx context.Context, rule model.FraudRule) error {
	if rule.ID == "" {
		rule.ID = uuid.New().String()
	}
	
	if rule.Priority == 0 {
		rule.Priority = model.DefaultPriority
	}

	exists, err := r.ExistsByName(ctx, rule.Name, "")
	if err != nil {
		return fmt.Errorf("failed to check name uniqueness: %w", err)
	}
	if exists {
		return fmt.Errorf("rule with name '%s' already exists", rule.Name)
	}

	count, err := r.GetActiveRulesCount(ctx)
	if err != nil {
		return fmt.Errorf("failed to check rules count: %w", err)
	}
	if count >= model.MaxRulesCount {
		return fmt.Errorf("maximum rules count (%d) exceeded", model.MaxRulesCount)
	}

	query := `
		INSERT INTO fraud_rules (id, name, description, dsl, priority, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err = r.db.Exec(ctx, query,
		rule.ID,
		rule.Name,
		rule.Description,
		rule.DSL,
		rule.Priority,
		rule.IsActive,
		rule.CreatedAt,
		rule.UpdatedAt,
	)
	
	if err != nil {
		return fmt.Errorf("failed to create fraud rule: %w", err)
	}

	return nil
}

func (r *repository) ExistsByName(ctx context.Context, name string, excludeID string) (bool, error) {
	var exists bool
	var query string
	var args []interface{}

	if excludeID != "" {
		query = `SELECT EXISTS(SELECT 1 FROM fraud_rules WHERE name = $1 AND id != $2)`
		args = []interface{}{name, excludeID}
	} else {
		query = `SELECT EXISTS(SELECT 1 FROM fraud_rules WHERE name = $1)`
		args = []interface{}{name}
	}

	err := r.db.QueryRow(ctx, query, args...).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check name existence: %w", err)
	}

	return exists, nil
}

func (r *repository) GetActiveRulesCount(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM fraud_rules WHERE is_active = true`
	
	err := r.db.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count active rules: %w", err)
	}

	return count, nil
}
