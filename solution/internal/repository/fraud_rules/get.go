package fraud_rules

import (
	"context"
	"fmt"
	"solution/internal/model"
)

func (r *repository) GetByID(ctx context.Context, id string) (*model.FraudRule, error) {
	query := `
		SELECT id, name, description, dsl, priority, is_active, created_at, updated_at
		FROM fraud_rules
		WHERE id = $1
	`

	rule := &model.FraudRule{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&rule.ID,
		&rule.Name,
		&rule.Description,
		&rule.DSL,
		&rule.Priority,
		&rule.IsActive,
		&rule.CreatedAt,
		&rule.UpdatedAt,
	)

	if err != nil {
		if isNotFoundErr(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get fraud rule by ID: %w", err)
	}

	return rule, nil
}

func (r *repository) GetByName(ctx context.Context, name string) (*model.FraudRule, error) {
	query := `
		SELECT id, name, description, dsl, priority, is_active, created_at, updated_at
		FROM fraud_rules
		WHERE name = $1
	`

	rule := &model.FraudRule{}
	err := r.db.QueryRow(ctx, query, name).Scan(
		&rule.ID,
		&rule.Name,
		&rule.Description,
		&rule.DSL,
		&rule.Priority,
		&rule.IsActive,
		&rule.CreatedAt,
		&rule.UpdatedAt,
	)

	if err != nil {
		if isNotFoundErr(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get fraud rule by name: %w", err)
	}

	return rule, nil
}

func (r *repository) GetAll(ctx context.Context, activeOnly bool) ([]*model.FraudRule, error) {
	query := `
		SELECT id, name, description, dsl, priority, is_active, created_at, updated_at
		FROM fraud_rules
	`

	if activeOnly {
		query += ` WHERE is_active = true`
	}
	
	query += ` ORDER BY priority ASC, created_at DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query fraud rules: %w", err)
	}
	defer rows.Close()

	var rules []*model.FraudRule
	for rows.Next() {
		rule := &model.FraudRule{}
		err := rows.Scan(
			&rule.ID,
			&rule.Name,
			&rule.Description,
			&rule.DSL,
			&rule.Priority,
			&rule.IsActive,
			&rule.CreatedAt,
			&rule.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan fraud rule: %w", err)
		}
		rules = append(rules, rule)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating fraud rules: %w", err)
	}

	return rules, nil
}

func isNotFoundErr(err error) bool {
	// TODO: добавить более точную проверку для pgx
	return err != nil
}
