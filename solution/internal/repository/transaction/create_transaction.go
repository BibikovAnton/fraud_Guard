package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"solution/internal/model"
	"strings"
)

// Create сохраняет новую транзакцию в базу данных
// Из прошлого проекта: batch insert критически важен для производительности
func (r *repo) Create(ctx context.Context, transaction model.Transaction) error {
	query := `
		INSERT INTO transactions (
			id, user_id, amount, currency, status, merchant_id, merchant_category_code,
			timestamp, ip_address, device_id, channel, location, is_fraud, metadata, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
		)
	`

	// Подготовка location JSON
	var locationJSON []byte
	if transaction.Location != nil {
		var err error
		locationJSON, err = json.Marshal(transaction.Location)
		if err != nil {
			return fmt.Errorf("failed to marshal location: %w", err)
		}
	}

	// Подготовка metadata JSON
	var metadataJSON []byte
	if transaction.Metadata != nil {
		var err error
		metadataJSON, err = json.Marshal(transaction.Metadata)
		if err != nil {
			return fmt.Errorf("failed to marshal metadata: %w", err)
		}
	}

	_, err := r.db.Exec(ctx, query,
		transaction.ID,                    // $1
		transaction.UserID,                // $2
		transaction.Amount,                 // $3
		transaction.Currency,               // $4
		transaction.Status,                // $5
		transaction.MerchantID,            // $6
		transaction.MerchantCategoryCode,   // $7
		transaction.Timestamp,              // $8
		transaction.IPAddress,              // $9
		transaction.DeviceID,               // $10
		transaction.Channel,                // $11
		locationJSON,                      // $12
		transaction.IsFraud,               // $13
		metadataJSON,                      // $14
		transaction.CreatedAt,              // $15
		transaction.UpdatedAt,              // $16
	)

	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	return nil
}

// CreateBatch сохраняет пачку транзакций
// Из прошлого проекта: COPY для максимальной производительности
func (r *repo) CreateBatch(ctx context.Context, transactions []model.Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	// Для больших пачек используем COPY
	if len(transactions) > 100 {
		return r.createBatchCopy(ctx, transactions)
	}

	// Для маленьких пачек используем обычный INSERT
	return r.createBatchInsert(ctx, transactions)
}

// createBatchInsert - обычный INSERT для небольших пачек
func (r *repo) createBatchInsert(ctx context.Context, transactions []model.Transaction) error {
	query := `
		INSERT INTO transactions (
			id, user_id, amount, currency, status, merchant_id, merchant_category_code,
			timestamp, ip_address, device_id, channel, location, is_fraud, metadata, created_at, updated_at
		) VALUES `

	// Строим VALUES часть
	values := make([]string, len(transactions))
	args := make([]interface{}, 0, len(transactions)*16)

	for i, tx := range transactions {
		values[i] = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			i*16+1, i*16+2, i*16+3, i*16+4, i*16+5, i*16+6, i*16+7, i*16+8,
			i*16+9, i*16+10, i*16+11, i*16+12, i*16+13, i*16+14, i*16+15, i*16+16)

		// Подготовка JSON полей
		var locationJSON, metadataJSON []byte
		var err error

		if tx.Location != nil {
			locationJSON, err = json.Marshal(tx.Location)
			if err != nil {
				return fmt.Errorf("failed to marshal location for transaction %d: %w", i, err)
			}
		}

		if tx.Metadata != nil {
			metadataJSON, err = json.Marshal(tx.Metadata)
			if err != nil {
				return fmt.Errorf("failed to marshal metadata for transaction %d: %w", i, err)
			}
		}

		args = append(args,
			tx.ID, tx.UserID, tx.Amount, tx.Currency, tx.Status,
			tx.MerchantID, tx.MerchantCategoryCode, tx.Timestamp,
			tx.IPAddress, tx.DeviceID, tx.Channel,
			locationJSON, tx.IsFraud, metadataJSON,
			tx.CreatedAt, tx.UpdatedAt,
		)
	}

	query += fmt.Sprintf("%s", fmt.Sprintf(strings.Join(values, ", ")))

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to create batch transactions: %w", err)
	}

	return nil
}

// createBatchCopy - COPY для больших пачек (максимальная производительность)
func (r *repo) createBatchCopy(ctx context.Context, transactions []model.Transaction) error {
	// TODO: реализовать COPY для максимальной производительности
	// Из прошлого проекта: COPY в 10 раз быстрее INSERT для больших объемов
	return r.createBatchInsert(ctx, transactions)
}
