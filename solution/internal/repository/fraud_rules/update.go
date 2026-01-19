package fraud_rules

import (
	"context"
	"fmt"
	"solution/internal/model"
	"time"
)

func (r *repository) Update(ctx context.Context, id string, req model.FraudRuleUpdateRequest) (*model.FraudRule, error) {
	existing, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing rule: %w", err)
	}
	if existing == nil {
		return nil, fmt.Errorf("fraud rule not found")
	}

	if req.Name != nil && *req.Name != existing.Name {
		exists, err := r.ExistsByName(ctx, *req.Name, id)
		if err != nil {
			return nil, fmt.Errorf("failed to check name uniqueness: %w", err)
		}
		if exists {
			return nil, fmt.Errorf("rule with name '%s' already exists", *req.Name)
		}
	}

	query := `UPDATE fraud_rules SET `
	var updates []string
	var args []interface{}
	argIndex := 1

	if req.Name != nil {
		updates = append(updates, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, *req.Name)
		argIndex++
	}

	if req.Description != nil {
		updates = append(updates, fmt.Sprintf("description = $%d", argIndex))
		args = append(args, *req.Description)
		argIndex++
	}

	if req.DSL != nil {
		updates = append(updates, fmt.Sprintf("dsl = $%d", argIndex))
		args = append(args, *req.DSL)
		argIndex++
	}

	if req.Priority != nil {
		updates = append(updates, fmt.Sprintf("priority = $%d", argIndex))
		args = append(args, *req.Priority)
		argIndex++
	}

	if req.IsActive != nil {
		updates = append(updates, fmt.Sprintf("is_active = $%d", argIndex))
		args = append(args, *req.IsActive)
		argIndex++
	}

	updates = append(updates, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, time.Now())
	argIndex++

	if len(updates) == 0 {
		return existing, nil
	}

	query += fmt.Sprintf("%s WHERE id = $%d", fmt.Sprintf("%s", updates), argIndex)
	args = append(args, id)

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update fraud rule: %w", err)
	}

	return r.GetByID(ctx, id)
}

func (r *repository) Delete(ctx context.Context, id string) error {
	exists, err := r.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to check rule existence: %w", err)
	}
	if exists == nil {
		return fmt.Errorf("fraud rule not found")
	}

	query := `UPDATE fraud_rules SET is_active = false, updated_at = $1 WHERE id = $2`

	_, err = r.db.Exec(ctx, query, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to delete fraud rule: %w", err)
	}

	return nil
}
