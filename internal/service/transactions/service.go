package transactions

import (
	"context"
	"fmt"
	"time"
	"solution/internal/dsl"
	"solution/internal/model"
	"solution/internal/service"
	"solution/internal/repository"
	"github.com/google/uuid"
	"go.uber.org/zap"
	transactionsRepo "solution/internal/repository/transactions"
	fraudRulesRepo "solution/internal/repository/fraud_rules"
)

type Service struct {
	txRepo       transactionsRepo.Repository
	userRepo     repository.UserRepository
	fraudRuleRepo fraudRulesRepo.Repository
	dslEvaluator *dsl.DSLEvaluator
	logger        *zap.Logger
}

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByIDIncludingInactive(ctx context.Context, id string) (*model.User, error)
}

type FraudRuleRepository interface {
	GetActiveRules(ctx context.Context) ([]*model.FraudRule, error)
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

func NewService(txRepo transactionsRepo.Repository, userRepo repository.UserRepository, fraudRuleRepo fraudRulesRepo.Repository, logger *zap.Logger) *Service {
	return &Service{
		txRepo:       txRepo,
		userRepo:     userRepo,
		fraudRuleRepo: fraudRuleRepo,
		dslEvaluator: dsl.NewDSLEvaluator(logger),
		logger:        logger,
	}
}

func (s *Service) Create(ctx context.Context, req model.TransactionCreateRequest) (*model.TransactionDecision, error) {
	fmt.Printf("DEBUG: Creating transaction for user %v, amount %.2f\n", req.UserID, req.Amount)
	
	if err := s.validateCreateRequest(ctx, req); err != nil {
		fmt.Printf("DEBUG: Validation failed: %v\n", err)
		return nil, err
	}

	var user *model.User
	var err error
	
	
	if req.UserID != nil {
		user, err = s.userRepo.GetByIDIncludingInactive(ctx, req.UserID.String())
		if err != nil {
			return nil, fmt.Errorf("failed to get user by ID: %w", err)
		}
		
		if user == nil {
			return nil, fmt.Errorf("failed to get user by ID: no rows in result set")
		}
		
		if !user.IsActive {
			return nil, fmt.Errorf("user is deactivated")
		}
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

	return &service.PagedTransactions{
		Items: transactions,
		Total: total,
		Page:  params.Page,
		Size:  params.Size,
	}, nil
}

func (s *Service) validateCreateRequest(ctx context.Context, req model.TransactionCreateRequest) error {
	if req.Amount <= 0 {
		return fmt.Errorf("must be greater > 0")
	}
	
	if req.Amount > model.MaxTransactionAmount {
		return fmt.Errorf("must be less than or equal to %.2f", model.MaxTransactionAmount)
	}

	if req.Currency == "" {
		return fmt.Errorf("currency is required")
	}

	validCurrencies := map[model.CurrencyCode]bool{
		model.CurrencyUSD: true,
		model.CurrencyEUR: true,
		model.CurrencyRUB: true,
	}
	if !validCurrencies[req.Currency] {
		return fmt.Errorf("invalid currency code: %s", req.Currency)
	}

	if req.Timestamp.IsZero() {
		return fmt.Errorf("timestamp is required")
	}

	if req.Timestamp.After(time.Now().Add(5 * time.Minute)) {
		return fmt.Errorf("timestamp cannot be more than 5 minutes in the future")
	}

	if req.Location != nil {
		
		if req.Location.Country != "" {
			if len(req.Location.Country) > 2 {
				return fmt.Errorf("location.country must be at most 2 characters")
			}
		}
		
		if (req.Location.Latitude != nil && req.Location.Longitude == nil) ||
			(req.Location.Latitude == nil && req.Location.Longitude != nil) {
			return fmt.Errorf("longitude and latitude must be provided together")
		}
		
		if req.Location.Latitude != nil && (*req.Location.Latitude < -90 || *req.Location.Latitude > 90) {
			return fmt.Errorf("location.latitude must be between -90 and 90")
		}
		
		if req.Location.Longitude != nil && (*req.Location.Longitude < -180 || *req.Location.Longitude > 180) {
			return fmt.Errorf("location.longitude must be between -180 and 180")
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

	fmt.Printf("DEBUG: Processing %d rules for transaction %s\n", len(sortedRules), transaction.ID.String())
	for i, rule := range sortedRules {
		fmt.Printf("DEBUG: Rule[%d]: ID=%s, Name=%s, DSL=%s, Priority=%d, Active=%v\n", 
			i, rule.ID, rule.Name, rule.DslExpression, rule.Priority, rule.Enabled)
	}

	for _, rule := range sortedRules {
		ruleResult := s.dslEvaluator.EvaluateRule(rule, transaction, user)
		
		fmt.Printf("DEBUG: Rule %s: matched=%v, description=%s\n", rule.Name, ruleResult.Matched, ruleResult.Description)

		results = append(results, ruleResult)
	}

	return results
}
