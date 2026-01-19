package fraud_rules

import (
	"context"
	"fmt"
	"solution/internal/model"
	"solution/internal/repository"
	"strings"
	"time"
)

type Service interface {
	Create(ctx context.Context, req model.FraudRuleCreateRequest) (*model.FraudRule, error)
	GetByID(ctx context.Context, id string) (*model.FraudRule, error)
	GetAll(ctx context.Context, activeOnly bool) ([]*model.FraudRule, error)
	Update(ctx context.Context, id string, req model.FraudRuleUpdateRequest) (*model.FraudRule, error)
	Delete(ctx context.Context, id string) error
	
	ValidateDSL(ctx context.Context, dsl string) (*model.FraudRuleValidateResponse, error)
}

type service struct {
	fraudRuleRepo repository.FraudRuleRepository
}

func NewService(fraudRuleRepo repository.FraudRuleRepository) Service {
	return &service{
		fraudRuleRepo: fraudRuleRepo,
	}
}

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

func (s *service) GetByID(ctx context.Context, id string) (*model.FraudRule, error) {
	rule, err := s.fraudRuleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get fraud rule: %w", err)
	}
	return rule, nil
}

func (s *service) GetAll(ctx context.Context, activeOnly bool) ([]*model.FraudRule, error) {
	rules, err := s.fraudRuleRepo.GetAll(ctx, activeOnly)
	if err != nil {
		return nil, fmt.Errorf("failed to get fraud rules: %w", err)
	}
	return rules, nil
}

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

func (s *service) ValidateDSL(ctx context.Context, dsl string) (*model.FraudRuleValidateResponse, error) {
	dsl = strings.TrimSpace(dsl)
	if len(dsl) == 0 {
		return &model.FraudRuleValidateResponse{
			IsValid: false,
			Error:   "DSL cannot be empty",
		}, nil
	}

	if len(dsl) > model.MaxDSLSize {
		return &model.FraudRuleValidateResponse{
			IsValid: false,
			Error:   fmt.Sprintf("DSL too large (max %d characters)", model.MaxDSLSize),
		}, nil
	}

	dangerousWords := []string{"DROP", "DELETE", "TRUNCATE", "INSERT", "UPDATE"}
	dslUpper := strings.ToUpper(dsl)
	for _, word := range dangerousWords {
		if strings.Contains(dslUpper, word) {
			return &model.FraudRuleValidateResponse{
				IsValid: false,
				Error:   fmt.Sprintf("DSL contains dangerous keyword: %s", word),
			}, nil
		}
	}

	return &model.FraudRuleValidateResponse{
		IsValid: true,
		AST:     "AST parsing not implemented yet",
	}, nil
}
