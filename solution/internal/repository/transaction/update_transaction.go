package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"solution/internal/model"
	"strings"
	"time"
)

// Update обновляет транзакцию
func (r *repo) Update(ctx context.Context, id string, req model.TransactionUpdateRequest) (*model.Transaction, error) {
	// Строим динамический UPDATE запрос
	var updates []string
	var args []interface{}
	argIndex := 1

	if req.Status != nil {
		updates = append(updates, fmt.Sprintf("status = $%d", argIndex))
		args = append(args, *req.Status)
		argIndex++
	}

	if req.IsFraud != nil {
		updates = append(updates, fmt.Sprintf("is_fraud = $%d", argIndex))
		args = append(args, *req.IsFraud)
		argIndex++
	}

	if req.Metadata != nil {
		metadataJSON, err := json.Marshal(req.Metadata)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal metadata: %w", err)
		}
		updates = append(updates, fmt.Sprintf("metadata = $%d", argIndex))
		args = append(args, metadataJSON)
		argIndex++
	}

	if len(updates) == 0 {
		// Нечего обновлять
		return r.GetByID(ctx, id)
	}

	// Всегда обновляем updated_at
	updates = append(updates, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, time.Now())

	// Собираем запрос
	query := fmt.Sprintf("UPDATE transactions SET %s WHERE id = $%d",
		strings.Join(updates, ", "), argIndex+1)

	// Добавляем ID в конец
	args = append(args, id)

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update transaction: %w", err)
	}

	// Возвращаем обновленную транзакцию
	return r.GetByID(ctx, id)
}

// CountByStatus считает количество транзакций по статусам
func (r *repo) CountByStatus(ctx context.Context) (map[model.TransactionStatus]int64, error) {
	query := `
		SELECT status, COUNT(*)
		FROM transactions
		GROUP BY status
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to count transactions by status: %w", err)
	}
	defer rows.Close()

	result := make(map[model.TransactionStatus]int64)
	for rows.Next() {
		var status model.TransactionStatus
		var count int64

		err := rows.Scan(&status, &count)
		if err != nil {
			return nil, fmt.Errorf("failed to scan status count: %w", err)
		}

		result[status] = count
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating status counts: %w", err)
	}

	return result, nil
}

// CountFraudTransactions считает количество фродовых транзакций за период
func (r *repo) CountFraudTransactions(ctx context.Context, timeRange time.Duration) (int64, error) {
	query := `
		SELECT COUNT(*)
		FROM transactions
		WHERE is_fraud = true AND timestamp >= NOW() - INTERVAL '1 second' * $1
	`

	var count int64
	err := r.db.QueryRow(ctx, query, int64(timeRange.Seconds())).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count fraud transactions: %w", err)
	}

	return count, nil
}

// GetTotalAmount получает общую сумму транзакций по валюте
func (r *repo) GetTotalAmount(ctx context.Context, currency model.CurrencyCode) (float64, error) {
	query := `
		SELECT COALESCE(SUM(amount), 0)
		FROM transactions
		WHERE currency = $1 AND status IN ('APPROVED', 'PROCESSED')
	`

	var total float64
	err := r.db.QueryRow(ctx, query, currency).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total amount: %w", err)
	}

	return total, nil
}
