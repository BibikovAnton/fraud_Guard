package transactions

import (
	"context"
	"fmt"
	"time"
	"solution/internal/model"
	"solution/internal/service"
	"solution/internal/repository"
	"github.com/google/uuid"
	transactionsRepo "solution/internal/repository/transactions"
	fraudRulesRepo "solution/internal/repository/fraud_rules"
	dslService "solution/internal/service/dsl"
)

type Service struct {
	txRepo       transactionsRepo.Repository
	userRepo     repository.UserRepository
	fraudRuleRepo fraudRulesRepo.Repository
	dslEvaluator dslService.Evaluator
}

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
}

type FraudRuleRepository interface {
	GetActiveRules(ctx context.Context) ([]*model.FraudRule, error)
}

type DSLEvaluator interface {
	Evaluate(ctx context.Context, dsl string, transaction *model.Transaction, user *model.User) (bool, string, error)
}

type ListParams struct {
	UserID  *string
	Status  *model.TransactionStatus
	IsFraud *bool
	From    *time.Time
	To      *time.Time
	Page    int
	Size    int
}

func NewService(txRepo transactionsRepo.Repository, userRepo repository.UserRepository, fraudRuleRepo fraudRulesRepo.Repository, dslEvaluator dslService.Evaluator) *Service {
	return &Service{
		txRepo:       txRepo,
		userRepo:     userRepo,
		fraudRuleRepo: fraudRuleRepo,
		dslEvaluator: dslEvaluator,
	}
}

func (s *Service) Create(ctx context.Context, req model.TransactionCreateRequest) (*model.TransactionDecision, error) {
	fmt.Printf("DEBUG: Creating transaction for user %s, amount %.2f\n", req.UserID.String(), req.Amount)
	
	if err := s.validateCreateRequest(ctx, req); err != nil {
		fmt.Printf("DEBUG: Validation failed: %v\n", err)
		return nil, err
	}

	user, err := s.userRepo.GetByID(ctx, req.UserID.String())
	if err != nil {
		fmt.Printf("DEBUG: Failed to get user: %v\n", err)
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if !user.IsActive {
		fmt.Printf("DEBUG: User is deactivated\n")
		return nil, fmt.Errorf("user is deactivated")
	}

	transaction := &model.Transaction{
		ID:                   uuid.New(),
		UserID:               req.UserID,
		Amount:               req.Amount,
		Currency:             req.Currency,
		Status:               model.TransactionStatusPending,
		MerchantID:           req.MerchantID,
		MerchantCategoryCode: req.MerchantCategoryCode,
		Timestamp:            req.Timestamp,
		IPAddress:            req.IPAddress,
		DeviceID:             req.DeviceID,
		Channel:              req.Channel,
		Location:             req.Location,
		IsFraud:              false,
		Metadata:             req.Metadata,
		CreatedAt:            time.Now().UTC(),
		UpdatedAt:            time.Now().UTC(),
	}

	rules, err := s.fraudRuleRepo.GetActiveRules(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get fraud rules: %w", err)
	}
	
	// Debug: log number of rules loaded
	fmt.Printf("DEBUG: Loaded %d active rules\n", len(rules))

	ruleResults := s.applyFraudRules(ctx, transaction, user, rules)

	isFraud := false
	for _, result := range ruleResults {
		if result.Matched {
			isFraud = true
			break
		}
	}

	if isFraud {
		transaction.Status = model.TransactionStatusDeclined
		transaction.IsFraud = true
	} else {
		transaction.Status = model.TransactionStatusApproved
		transaction.IsFraud = false
	}

	if err := s.txRepo.CreateWithResults(ctx, transaction, ruleResults); err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return &model.TransactionDecision{
		Transaction: transaction,
		RuleResults: ruleResults,
	}, nil
}

func (s *Service) CreateBatch(ctx context.Context, req model.TransactionBatchCreateRequest) (*model.TransactionBatchResult, error) {
	if len(req.Items) == 0 {
		return nil, fmt.Errorf("batch cannot be empty")
	}
	if len(req.Items) > model.MaxBatchSize {
		return nil, fmt.Errorf("batch too large (max %d)", model.MaxBatchSize)
	}

	result := &model.TransactionBatchResult{
		Items: make([]model.TransactionBatchItem, len(req.Items)),
	}

	for i, item := range req.Items {
		decision, err := s.Create(ctx, item)
		if err != nil {
			result.Items[i] = model.TransactionBatchItem{
				Index: i,
				Error: &model.ApiError{
					Code:      "VALIDATION_FAILED",
					Message:   err.Error(),
					TraceID:   uuid.New(),
					Timestamp: time.Now().UTC(),
					Path:      "/api/v1/transactions/batch",
				},
			}
		} else {
			result.Items[i] = model.TransactionBatchItem{
				Index:    i,
				Decision: decision,
			}
		}
	}

	return result, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*model.TransactionDecision, error) {
	return s.txRepo.GetByIDWithResults(ctx, id)
}

func (s *Service) GetList(ctx context.Context, params service.TransactionListParams) (*service.PagedTransactions, error) {
	repoParams := transactionsRepo.ListParams{
		UserID:  params.UserID,
		Status:  params.Status,
		IsFraud: params.IsFraud,
		From:    params.From,
		To:      params.To,
		Page:    params.Page,
		Size:    params.Size,
	}

	transactions, total, err := s.txRepo.GetList(ctx, repoParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}

	decisions := make([]model.TransactionDecision, len(transactions))
	for i, tx := range transactions {
		decisions[i] = model.TransactionDecision{
			Transaction: tx,
			RuleResults: []model.RuleResult{}, // Пусто для списка
		}
	}

	return &service.PagedTransactions{
		Transactions: decisions,
		Total:        total,
		Page:         params.Page,
		Size:         params.Size,
	}, nil
}

func (s *Service) validateCreateRequest(ctx context.Context, req model.TransactionCreateRequest) error {
	if req.Amount < model.MinTransactionAmount || req.Amount > model.MaxTransactionAmount {
		return fmt.Errorf("amount must be between %.2f and %.2f", model.MinTransactionAmount, model.MaxTransactionAmount)
	}

	if req.Currency == "" {
		return fmt.Errorf("currency is required")
	}

	if req.Timestamp.IsZero() {
		return fmt.Errorf("timestamp is required")
	}

	if req.Timestamp.After(time.Now().Add(5 * time.Minute)) {
		return fmt.Errorf("timestamp cannot be more than 5 minutes in the future")
	}

	if req.Location != nil {
		if req.Location.Country == "" {
			return fmt.Errorf("location.country is required when location is provided")
		}
		if (req.Location.Latitude != nil && req.Location.Longitude == nil) ||
			(req.Location.Latitude == nil && req.Location.Longitude != nil) {
			return fmt.Errorf("both latitude and longitude must be provided together")
		}
		if req.Location.Latitude != nil {
			lat := *req.Location.Latitude
			if lat < -90 || lat > 90 {
				return fmt.Errorf("latitude must be between -90 and 90")
			}
		}
		if req.Location.Longitude != nil {
			lon := *req.Location.Longitude
			if lon < -180 || lon > 180 {
				return fmt.Errorf("longitude must be between -180 and 180")
			}
		}
	}

	return nil
}

func (s *Service) applyFraudRules(ctx context.Context, transaction *model.Transaction, user *model.User, rules []*model.FraudRule) []model.RuleResult {
	results := make([]model.RuleResult, 0, len(rules))

	sortedRules := make([]*model.FraudRule, len(rules))
	copy(sortedRules, rules)
	
	for i := 0; i < len(sortedRules)-1; i++ {
		for j := i + 1; j < len(sortedRules); j++ {
			if sortedRules[i].Priority > sortedRules[j].Priority ||
				(sortedRules[i].Priority == sortedRules[j].Priority && sortedRules[i].ID > sortedRules[j].ID) {
				sortedRules[i], sortedRules[j] = sortedRules[j], sortedRules[i]
			}
		}
	}

	// Debug: log all rules
	fmt.Printf("DEBUG: Processing %d rules for transaction %s\n", len(sortedRules), transaction.ID.String())
	for i, rule := range sortedRules {
		fmt.Printf("DEBUG: Rule[%d]: ID=%s, Name=%s, DSL=%s, Priority=%d, Active=%v\n", 
			i, rule.ID, rule.Name, rule.DslExpression, rule.Priority, rule.Enabled)
	}

	for _, rule := range sortedRules {
		matched, description, err := s.dslEvaluator.Evaluate(ctx, rule.DslExpression, transaction, user)
		if err != nil {
			matched = false
			description = fmt.Sprintf("Error evaluating rule: %s", err.Error())
		}

		// Debug: log rule evaluation
		fmt.Printf("DEBUG: Rule %s: matched=%v, description=%s\n", rule.Name, matched, description)

		result := model.RuleResult{
			RuleID:      rule.ID,
			RuleName:    rule.Name,
			Priority:    rule.Priority,
			Enabled:     rule.Enabled,
			Matched:     matched,
			Description: description,
		}

		results = append(results, result)
	}

	return results
}
