package transaction

import (
	"context"
	"fmt"
	"solution/internal/model"
	"solution/internal/repository"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Service - бизнес-логика для транзакций
type Service interface {
	// CRUD операции
	Create(ctx context.Context, req model.TransactionCreateRequest) (*model.Transaction, error)
	CreateBatch(ctx context.Context, req model.TransactionBatchCreateRequest) (*model.TransactionBatchResult, error)
	GetByID(ctx context.Context, id string) (*model.Transaction, error)
	GetByUserID(ctx context.Context, userID string, page, size int) ([]*model.Transaction, error)
	Update(ctx context.Context, id string, req model.TransactionUpdateRequest) (*model.Transaction, error)
	
	// Поиск и фильтрация
	GetFraudTransactions(ctx context.Context, page, size int) ([]*model.Transaction, error)
	GetByMerchantID(ctx context.Context, merchantID string, page, size int) ([]*model.Transaction, error)
	
	// Статистика
	GetTransactionStats(ctx context.Context) (*model.TransactionStats, error)
}

// service - имплементация с человеческой логикой
type service struct {
	txRepo       repository.TransactionRepository
	fraudRuleRepo repository.FraudRuleRepository
}

// NewService создает новый сервис транзакций
func NewService(txRepo repository.TransactionRepository, fraudRuleRepo repository.FraudRuleRepository) Service {
	return &service{
		txRepo:       txRepo,
		fraudRuleRepo: fraudRuleRepo,
	}
}

// Create создает новую транзакцию с проверкой на фрод
func (s *service) Create(ctx context.Context, req model.TransactionCreateRequest) (*model.Transaction, error) {
	// Валидация запроса
	if err := s.validateCreateRequest(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Создаем модель транзакции
	now := time.Now()
	transaction := model.Transaction{
		ID:                   uuid.New(),
		UserID:               req.UserID,
		Amount:               req.Amount,
		Currency:             req.Currency,
		Status:               model.TransactionStatusPending,
		MerchantID:           req.MerchantID,
		MerchantCategoryCode:   req.MerchantCategoryCode,
		Timestamp:            req.Timestamp,
		IPAddress:            req.IPAddress,
		DeviceID:             req.DeviceID,
		Channel:              req.Channel,
		Location:             req.Location,
		IsFraud:              false, // будет определено после проверки правилами
		Metadata:             req.Metadata,
		CreatedAt:            now,
		UpdatedAt:            now,
	}

	// Проверяем транзакцию на фрод
	isFraud, err := s.evaluateFraudRules(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("fraud evaluation failed: %w", err)
	}
	transaction.IsFraud = isFraud

	// Устанавливаем статус на основе результата проверки фрода
	if isFraud {
		transaction.Status = model.TransactionStatusRejected
	} else {
		transaction.Status = model.TransactionStatusApproved
	}

	// Сохраняем в БД
	if err := s.txRepo.Create(ctx, transaction); err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return &transaction, nil
}

// CreateBatch создает пачку транзакций
func (s *service) CreateBatch(ctx context.Context, req model.TransactionBatchCreateRequest) (*model.TransactionBatchResult, error) {
	if len(req.Transactions) == 0 {
		return &model.TransactionBatchResult{
			Processed: 0,
			Success:   0,
			Failed:    0,
			Results:   []model.TransactionBatchItem{},
		}, nil
	}

	// Проверяем размер пачки
	if len(req.Transactions) > model.MaxBatchSize {
		return nil, fmt.Errorf("batch size exceeds maximum of %d", model.MaxBatchSize)
	}

	result := &model.TransactionBatchResult{
		Processed: len(req.Transactions),
		Results:   make([]model.TransactionBatchItem, len(req.Transactions)),
	}

	transactions := make([]model.Transaction, 0, len(req.Transactions))

	// Валидация и подготовка транзакций
	for i, txReq := range req.Transactions {
		if err := s.validateCreateRequest(txReq); err != nil {
			result.Results[i] = model.TransactionBatchItem{
				Index:   i,
				Success: false,
				Error:    stringPtr(err.Error()),
			}
			result.Failed++
			continue
		}

		now := time.Now()
		transaction := model.Transaction{
			ID:                   uuid.New(),
			UserID:               txReq.UserID,
			Amount:               txReq.Amount,
			Currency:             txReq.Currency,
			Status:               model.TransactionStatusPending,
			MerchantID:           txReq.MerchantID,
			MerchantCategoryCode:   txReq.MerchantCategoryCode,
			Timestamp:            txReq.Timestamp,
			IPAddress:            txReq.IPAddress,
			DeviceID:             txReq.DeviceID,
			Channel:              txReq.Channel,
			Location:             txReq.Location,
			IsFraud:              false,
			Metadata:             txReq.Metadata,
			CreatedAt:            now,
			UpdatedAt:            now,
		}

		// Проверяем на фрод
		isFraud, err := s.evaluateFraudRules(ctx, transaction)
		if err != nil {
			result.Results[i] = model.TransactionBatchItem{
				Index:   i,
				Success: false,
				Error:    stringPtr(fmt.Sprintf("fraud evaluation failed: %v", err)),
			}
			result.Failed++
			continue
		}

		transaction.IsFraud = isFraud
		if isFraud {
			transaction.Status = model.TransactionStatusRejected
		} else {
			transaction.Status = model.TransactionStatusApproved
		}

		transactions = append(transactions, transaction)
		result.Results[i] = model.TransactionBatchItem{
			Index:   i,
			Success: true,
			Data:    &transaction,
		}
		result.Success++
	}

	// Сохраняем в БД пачкой
	if len(transactions) > 0 {
		if err := s.txRepo.CreateBatch(ctx, transactions); err != nil {
			return nil, fmt.Errorf("failed to create batch transactions: %w", err)
		}
	}

	return result, nil
}

// GetByID получает транзакцию по ID
func (s *service) GetByID(ctx context.Context, id string) (*model.Transaction, error) {
	transaction, err := s.txRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}
	return transaction, nil
}

// GetByUserID получает транзакции пользователя с пагинацией
func (s *service) GetByUserID(ctx context.Context, userID string, page, size int) ([]*model.Transaction, error) {
	offset := (page - 1) * size
	transactions, err := s.txRepo.GetByUserID(ctx, userID, size, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get user transactions: %w", err)
	}
	return transactions, nil
}

// Update обновляет транзакцию
func (s *service) Update(ctx context.Context, id string, req model.TransactionUpdateRequest) (*model.Transaction, error) {
	transaction, err := s.txRepo.Update(ctx, id, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update transaction: %w", err)
	}
	return transaction, nil
}

// GetFraudTransactions получает фродовые транзакции
func (s *service) GetFraudTransactions(ctx context.Context, page, size int) ([]*model.Transaction, error) {
	offset := (page - 1) * size
	transactions, err := s.txRepo.GetFraudTransactions(ctx, size, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get fraud transactions: %w", err)
	}
	return transactions, nil
}

// GetByMerchantID получает транзакции мерчанта
func (s *service) GetByMerchantID(ctx context.Context, merchantID string, page, size int) ([]*model.Transaction, error) {
	offset := (page - 1) * size
	transactions, err := s.txRepo.GetByMerchantID(ctx, merchantID, size, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get merchant transactions: %w", err)
	}
	return transactions, nil
}

// GetTransactionStats получает статистику по транзакциям
func (s *service) GetTransactionStats(ctx context.Context) (*model.TransactionStats, error) {
	statusCounts, err := s.txRepo.CountByStatus(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get status counts: %w", err)
	}

	fraudLast24h, err := s.txRepo.CountFraudTransactions(ctx, 24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to get fraud count: %w", err)
	}

	totalUSD, err := s.txRepo.GetTotalAmount(ctx, model.CurrencyUSD)
	if err != nil {
		return nil, fmt.Errorf("failed to get total USD amount: %w", err)
	}

	totalEUR, err := s.txRepo.GetTotalAmount(ctx, model.CurrencyEUR)
	if err != nil {
		return nil, fmt.Errorf("failed to get total EUR amount: %w", err)
	}

	totalRUB, err := s.txRepo.GetTotalAmount(ctx, model.CurrencyRUB)
	if err != nil {
		return nil, fmt.Errorf("failed to get total RUB amount: %w", err)
	}

	return &model.TransactionStats{
		StatusCounts:      statusCounts,
		FraudLast24h:     fraudLast24h,
		TotalAmountUSD:    totalUSD,
		TotalAmountEUR:    totalEUR,
		TotalAmountRUB:    totalRUB,
		GeneratedAt:       time.Now(),
	}, nil
}

// Вспомогательные методы

// validateCreateRequest валидирует запрос на создание транзакции
func (s *service) validateCreateRequest(req model.TransactionCreateRequest) error {
	if req.UserID == uuid.Nil {
		return fmt.Errorf("user_id is required")
	}
	if req.Amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	if req.Amount > model.MaxTransactionAmount {
		return fmt.Errorf("amount exceeds maximum of %f", model.MaxTransactionAmount)
	}
	if req.Amount < model.MinTransactionAmount {
		return fmt.Errorf("amount is below minimum of %f", model.MinTransactionAmount)
	}
	if req.Currency == "" {
		return fmt.Errorf("currency is required")
	}
	if req.Timestamp.IsZero() {
		return fmt.Errorf("timestamp is required")
	}
	return nil
}

// evaluateFraudRules проверяет транзакцию по правилам фрода
func (s *service) evaluateFraudRules(ctx context.Context, transaction model.Transaction) (bool, error) {
	// Получаем активные правила
	rules, err := s.fraudRuleRepo.GetAll(ctx, true)
	if err != nil {
		return false, fmt.Errorf("failed to get fraud rules: %w", err)
	}

	// TODO: реализовать полноценный DSL парсер
	// Из прошлого проекта: antlr4 для сложных правил, но для начала хватит базовой логики
	
	for _, rule := range rules {
		// Базовая эвристика на основе DSL
		if s.matchesRule(transaction, *rule) {
			return true, nil // хотя бы одно правило сработало - это фрод
		}
	}

	return false, nil
}

// matchesRule проверяет, соответствует ли транзакция правилу
func (s *service) matchesRule(transaction model.Transaction, rule model.FraudRule) bool {
	// TODO: реализовать полноценный парсер DSL
	// Из прошлого проекта: простые эвристики для демонстрации
	
	dsl := rule.DSL
	
	// Простые проверки на основе ключевых слов в DSL
	if len(dsl) == 0 {
		return false
	}
	
	// Проверка на большую сумму
	if transaction.Amount > 10000 && containsIgnoreCase(dsl, "amount") && containsIgnoreCase(dsl, ">") {
		return true
	}
	
	// Проверка на подозрительные IP
	if transaction.IPAddress != nil && containsIgnoreCase(dsl, "ip") {
		// TODO: добавить реальную проверку IP
		return false
	}
	
	// Проверка на подозрительные устройства
	if transaction.DeviceID != nil && containsIgnoreCase(dsl, "device") {
		// TODO: добавить реальную проверку device
		return false
	}
	
	return false
}

// containsIgnoreCase проверяет наличие подстроки без учета регистра
func containsIgnoreCase(s, substr string) bool {
	return len(s) >= len(substr) && 
		   (s == substr || 
		    len(s) > len(substr) && 
		    (s[:len(substr)] == substr || 
		     s[len(s)-len(substr):] == substr ||
		     findSubstringIgnoreCase(s, substr)))
}

// findSubstringIgnoreCase ищет подстроку без учета регистра
func findSubstringIgnoreCase(s, substr string) bool {
	sLower := strings.ToLower(s)
	substrLower := strings.ToLower(substr)
	return strings.Contains(sLower, substrLower)
}

// stringPtr возвращает указатель на строку
func stringPtr(s string) *string {
	return &s
}
