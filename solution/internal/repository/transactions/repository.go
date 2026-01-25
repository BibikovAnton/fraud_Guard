package transactions

import (
	"context"
	"fmt"
	"time"
	"solution/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repository interface {
	Create(ctx context.Context, tx *model.Transaction) error
	CreateWithResults(ctx context.Context, tx *model.Transaction, ruleResults []model.RuleResult) error
	GetByID(ctx context.Context, id string) (*model.Transaction, error)
	GetByIDWithResults(ctx context.Context, id string) (*model.TransactionDecision, error)
	GetList(ctx context.Context, params ListParams) ([]*model.Transaction, int64, error)
	Update(ctx context.Context, tx *model.Transaction) error
}

type ListParams struct {
	UserID   *string
	Status   *model.TransactionStatus
	IsFraud  *bool
	From     *time.Time
	To       *time.Time
	Page     int
	Size     int
}

type repository struct {
	db DBTX
}

type DBTX interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, tx *model.Transaction) error {
	query := `
		INSERT INTO transactions (
			id, user_id, amount, currency, status, merchant_id, merchant_category_code,
			timestamp, ip_address, device_id, channel, location, is_fraud, metadata
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`
	
	_, err := r.db.Exec(ctx, query,
		tx.ID, tx.UserID, tx.Amount, tx.Currency, tx.Status,
		tx.MerchantID, tx.MerchantCategoryCode, tx.Timestamp,
		tx.IPAddress, tx.DeviceID, tx.Channel, tx.Location,
		tx.IsFraud, tx.Metadata,
	)
	return err
}

func (r *repository) CreateWithResults(ctx context.Context, tx *model.Transaction, ruleResults []model.RuleResult) error {
	txPg, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer txPg.Rollback(ctx)

	if err := r.createTx(ctx, txPg, tx); err != nil {
		return err
	}

	for _, result := range ruleResults {
		if err := r.createRuleResult(ctx, txPg, tx.ID, result); err != nil {
			return err
		}
	}

	return txPg.Commit(ctx)
}

func (r *repository) createTx(ctx context.Context, txPg pgx.Tx, tx *model.Transaction) error {
	query := `
		INSERT INTO transactions (
			id, user_id, amount, currency, status, merchant_id, merchant_category_code,
			timestamp, ip_address, device_id, channel, location, is_fraud, metadata
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`
	
	_, err := txPg.Exec(ctx, query,
		tx.ID, tx.UserID, tx.Amount, tx.Currency, tx.Status,
		tx.MerchantID, tx.MerchantCategoryCode, tx.Timestamp,
		tx.IPAddress, tx.DeviceID, tx.Channel, tx.Location,
		tx.IsFraud, tx.Metadata,
	)
	return err
}

func (r *repository) createRuleResult(ctx context.Context, txPg pgx.Tx, txID uuid.UUID, result model.RuleResult) error {
	query := `
		INSERT INTO transaction_rule_results (
			transaction_id, rule_id, rule_name, priority, enabled, matched, description
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	
	_, err := txPg.Exec(ctx, query,
		txID, result.RuleID, result.RuleName, result.Priority,
		result.Enabled, result.Matched, result.Description,
	)
	return err
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.Transaction, error) {
	query := `
		SELECT id, user_id, amount, currency, status, merchant_id, merchant_category_code,
			   timestamp, ip_address, device_id, channel, location, is_fraud, metadata,
			   created_at, updated_at
		FROM transactions
		WHERE id = $1
	`
	
	tx := &model.Transaction{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&tx.ID, &tx.UserID, &tx.Amount, &tx.Currency, &tx.Status,
		&tx.MerchantID, &tx.MerchantCategoryCode, &tx.Timestamp,
		&tx.IPAddress, &tx.DeviceID, &tx.Channel, &tx.Location,
		&tx.IsFraud, &tx.Metadata, &tx.CreatedAt, &tx.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (r *repository) GetByIDWithResults(ctx context.Context, id string) (*model.TransactionDecision, error) {
	tx, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	ruleResults, err := r.getRuleResults(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.TransactionDecision{
		Transaction: tx,
		RuleResults: ruleResults,
	}, nil
}

func (r *repository) getRuleResults(ctx context.Context, txID string) ([]model.RuleResult, error) {
	query := `
		SELECT rule_id, rule_name, priority, enabled, matched, description
		FROM transaction_rule_results
		WHERE transaction_id = $1
		ORDER BY priority ASC, rule_id ASC
	`
	
	rows, err := r.db.Query(ctx, query, txID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.RuleResult
	for rows.Next() {
		var result model.RuleResult
		if err := rows.Scan(
			&result.RuleID, &result.RuleName, &result.Priority,
			&result.Enabled, &result.Matched, &result.Description,
		); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, rows.Err()
}

func (r *repository) GetList(ctx context.Context, params ListParams) ([]*model.Transaction, int64, error) {
	baseQuery := "FROM transactions WHERE 1=1"
	countQuery := "SELECT COUNT(*) " + baseQuery
	selectQuery := `
		SELECT id, user_id, amount, currency, status, merchant_id, merchant_category_code,
			   timestamp, ip_address, device_id, channel, location, is_fraud, metadata,
			   created_at, updated_at
		` + baseQuery

	args := []interface{}{}
	argIndex := 1

	if params.UserID != nil {
		arg := fmt.Sprintf(" AND user_id = $%d", argIndex)
		baseQuery += arg
		countQuery += arg
		selectQuery += arg
		args = append(args, *params.UserID)
		argIndex++
	}

	if params.Status != nil {
		arg := fmt.Sprintf(" AND status = $%d", argIndex)
		baseQuery += arg
		countQuery += arg
		selectQuery += arg
		args = append(args, *params.Status)
		argIndex++
	}

	if params.IsFraud != nil {
		arg := fmt.Sprintf(" AND is_fraud = $%d", argIndex)
		baseQuery += arg
		countQuery += arg
		selectQuery += arg
		args = append(args, *params.IsFraud)
		argIndex++
	}

	if params.From != nil {
		arg := fmt.Sprintf(" AND timestamp >= $%d", argIndex)
		baseQuery += arg
		countQuery += arg
		selectQuery += arg
		args = append(args, *params.From)
		argIndex++
	}

	if params.To != nil {
		arg := fmt.Sprintf(" AND timestamp <= $%d", argIndex)
		baseQuery += arg
		countQuery += arg
		selectQuery += arg
		args = append(args, *params.To)
		argIndex++
	}

	var total int64
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	selectQuery += " ORDER BY timestamp DESC"
	selectQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, params.Size, (params.Page-1)*params.Size)

	rows, err := r.db.Query(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var transactions []*model.Transaction
	for rows.Next() {
		tx := &model.Transaction{}
		if err := rows.Scan(
			&tx.ID, &tx.UserID, &tx.Amount, &tx.Currency, &tx.Status,
			&tx.MerchantID, &tx.MerchantCategoryCode, &tx.Timestamp,
			&tx.IPAddress, &tx.DeviceID, &tx.Channel, &tx.Location,
			&tx.IsFraud, &tx.Metadata, &tx.CreatedAt, &tx.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		transactions = append(transactions, tx)
	}

	return transactions, total, rows.Err()
}

func (r *repository) Update(ctx context.Context, tx *model.Transaction) error {
	query := `
		UPDATE transactions
		SET status = $2, is_fraud = $3, metadata = $4, updated_at = NOW()
		WHERE id = $1
	`
	
	_, err := r.db.Exec(ctx, query, tx.ID, tx.Status, tx.IsFraud, tx.Metadata)
	return err
}
