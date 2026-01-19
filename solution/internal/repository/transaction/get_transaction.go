package transaction

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"solution/internal/model"
	"time"

	"github.com/jackc/pgx/v5"
)

// GetByID получает транзакцию по ID
func (r *repo) GetByID(ctx context.Context, id string) (*model.Transaction, error) {
	query := `
		SELECT id, user_id, amount, currency, status, merchant_id, merchant_category_code,
			   timestamp, ip_address, device_id, channel, location, is_fraud, metadata, created_at, updated_at
		FROM transactions
		WHERE id = $1
	`

	var transaction model.Transaction
	var locationJSON, metadataJSON []byte

	err := r.db.QueryRow(ctx, query, id).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Amount,
		&transaction.Currency,
		&transaction.Status,
		&transaction.MerchantID,
		&transaction.MerchantCategoryCode,
		&transaction.Timestamp,
		&transaction.IPAddress,
		&transaction.DeviceID,
		&transaction.Channel,
		&locationJSON,
		&transaction.IsFraud,
		&metadataJSON,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // транзакция не найдена
		}
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	// Десериализация JSON полей
	if len(locationJSON) > 0 {
		var location model.TransactionLocation
		if err := json.Unmarshal(locationJSON, &location); err != nil {
			return nil, fmt.Errorf("failed to unmarshal location: %w", err)
		}
		transaction.Location = &location
	}

	if len(metadataJSON) > 0 {
		var metadata model.TransactionMetadata
		if err := json.Unmarshal(metadataJSON, &metadata); err != nil {
			return nil, fmt.Errorf("failed to unmarshal metadata: %w", err)
		}
		transaction.Metadata = &metadata
	}

	return &transaction, nil
}

// GetByUserID получает транзакции пользователя с пагинацией
func (r *repo) GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*model.Transaction, error) {
	query := `
		SELECT id, user_id, amount, currency, status, merchant_id, merchant_category_code,
			   timestamp, ip_address, device_id, channel, location, is_fraud, metadata, created_at, updated_at
		FROM transactions
		WHERE user_id = $1
		ORDER BY timestamp DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query user transactions: %w", err)
	}
	defer rows.Close()

	var transactions []*model.Transaction
	for rows.Next() {
		var transaction model.Transaction
		var locationJSON, metadataJSON []byte

		err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.Amount,
			&transaction.Currency,
			&transaction.Status,
			&transaction.MerchantID,
			&transaction.MerchantCategoryCode,
			&transaction.Timestamp,
			&transaction.IPAddress,
			&transaction.DeviceID,
			&transaction.Channel,
			&locationJSON,
			&transaction.IsFraud,
			&metadataJSON,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %w", err)
		}

		// Десериализация JSON полей
		if len(locationJSON) > 0 {
			var location model.TransactionLocation
			if err := json.Unmarshal(locationJSON, &location); err != nil {
				return nil, fmt.Errorf("failed to unmarshal location: %w", err)
			}
			transaction.Location = &location
		}

		if len(metadataJSON) > 0 {
			var metadata model.TransactionMetadata
			if err := json.Unmarshal(metadataJSON, &metadata); err != nil {
				return nil, fmt.Errorf("failed to unmarshal metadata: %w", err)
			}
			transaction.Metadata = &metadata
		}

		transactions = append(transactions, &transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transactions: %w", err)
	}

	return transactions, nil
}

// GetFraudTransactions получает фродовые транзакции
func (r *repo) GetFraudTransactions(ctx context.Context, limit, offset int) ([]*model.Transaction, error) {
	query := `
		SELECT id, user_id, amount, currency, status, merchant_id, merchant_category_code,
			   timestamp, ip_address, device_id, channel, location, is_fraud, metadata, created_at, updated_at
		FROM transactions
		WHERE is_fraud = true
		ORDER BY timestamp DESC
		LIMIT $1 OFFSET $2
	`

	return r.scanTransactions(ctx, query, limit, offset)
}

// GetByMerchantID получает транзакции мерчанта
func (r *repo) GetByMerchantID(ctx context.Context, merchantID string, limit, offset int) ([]*model.Transaction, error) {
	query := `
		SELECT id, user_id, amount, currency, status, merchant_id, merchant_category_code,
			   timestamp, ip_address, device_id, channel, location, is_fraud, metadata, created_at, updated_at
		FROM transactions
		WHERE merchant_id = $1
		ORDER BY timestamp DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(ctx, query, merchantID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query merchant transactions: %w", err)
	}
	defer rows.Close()

	return r.scanTransactionRows(rows)
}

// GetByTimeRange получает транзакции за период
func (r *repo) GetByTimeRange(ctx context.Context, start, end time.Time, limit, offset int) ([]*model.Transaction, error) {
	query := `
		SELECT id, user_id, amount, currency, status, merchant_id, merchant_category_code,
			   timestamp, ip_address, device_id, channel, location, is_fraud, metadata, created_at, updated_at
		FROM transactions
		WHERE timestamp BETWEEN $1 AND $2
		ORDER BY timestamp DESC
		LIMIT $3 OFFSET $4
	`

	return r.scanTransactions(ctx, query, start, end, limit, offset)
}

// scanTransactions - вспомогательный метод для сканирования транзакций
func (r *repo) scanTransactions(ctx context.Context, query string, args ...interface{}) ([]*model.Transaction, error) {
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions: %w", err)
	}
	defer rows.Close()

	return r.scanTransactionRows(rows)
}

// scanTransactionRows - сканирует строки в транзакции
func (r *repo) scanTransactionRows(rows pgx.Rows) ([]*model.Transaction, error) {
	var transactions []*model.Transaction

	for rows.Next() {
		var transaction model.Transaction
		var locationJSON, metadataJSON []byte

		err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.Amount,
			&transaction.Currency,
			&transaction.Status,
			&transaction.MerchantID,
			&transaction.MerchantCategoryCode,
			&transaction.Timestamp,
			&transaction.IPAddress,
			&transaction.DeviceID,
			&transaction.Channel,
			&locationJSON,
			&transaction.IsFraud,
			&metadataJSON,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %w", err)
		}

		// Десериализация JSON полей
		if len(locationJSON) > 0 {
			var location model.TransactionLocation
			if err := json.Unmarshal(locationJSON, &location); err != nil {
				return nil, fmt.Errorf("failed to unmarshal location: %w", err)
			}
			transaction.Location = &location
		}

		if len(metadataJSON) > 0 {
			var metadata model.TransactionMetadata
			if err := json.Unmarshal(metadataJSON, &metadata); err != nil {
				return nil, fmt.Errorf("failed to unmarshal metadata: %w", err)
			}
			transaction.Metadata = &metadata
		}

		transactions = append(transactions, &transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transactions: %w", err)
	}

	return transactions, nil
}
