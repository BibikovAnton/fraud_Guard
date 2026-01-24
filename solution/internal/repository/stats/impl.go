package stats

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) Repository {
	return &repository{db: db}
}

func (r *repository) GetOverviewStats(ctx context.Context, from, to time.Time) (*OverviewStats, error) {
	query := `
		SELECT 
			COUNT(*) as volume,
			COALESCE(SUM(amount), 0) as gmv,
			COUNT(*) FILTER (WHERE status = 'APPROVED')::float / COUNT(*) as approval_rate,
			COUNT(*) FILTER (WHERE status = 'DECLINED')::float / COUNT(*) as decline_rate
		FROM transactions 
		WHERE created_at BETWEEN $1 AND $2
	`

	var stats OverviewStats
	err := r.db.QueryRow(ctx, query, from, to).Scan(
		&stats.Volume,
		&stats.GMV,
		&stats.ApprovalRate,
		&stats.DeclineRate,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get overview stats: %w", err)
	}

	stats.From = from
	stats.To = to
	return &stats, nil
}

func (r *repository) GetTransactionsTimeSeries(ctx context.Context, from, to time.Time, interval string) ([]TimeSeriesPoint, error) {
	// For now, implement daily buckets
	bucketFormat := "YYYY-MM-DD"
	if interval == "hour" {
		bucketFormat = "YYYY-MM-DD HH24:00:00"
	}

	query := fmt.Sprintf(`
		SELECT 
			date_trunc('%s', created_at) as bucket_start,
			COUNT(*) as tx_count,
			COALESCE(SUM(amount), 0) as gmv,
			COUNT(*) FILTER (WHERE status = 'APPROVED')::float / COUNT(*) as approval_rate,
			COUNT(*) FILTER (WHERE status = 'DECLINED')::float / COUNT(*) as decline_rate
		FROM transactions 
		WHERE created_at BETWEEN $1 AND $2
		GROUP BY date_trunc('%s', created_at)
		ORDER BY bucket_start
	`, interval, interval)

	rows, err := r.db.Query(ctx, query, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get time series: %w", err)
	}
	defer rows.Close()

	var points []TimeSeriesPoint
	for rows.Next() {
		var point TimeSeriesPoint
		err := rows.Scan(
			&point.BucketStart,
			&point.TxCount,
			&point.GMV,
			&point.ApprovalRate,
			&point.DeclineRate,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan time series point: %w", err)
		}
		points = append(points, point)
	}

	return points, nil
}

func (r *repository) GetRuleMatchesStats(ctx context.Context, from, to time.Time) ([]RuleMatchStat, error) {
	query := `
		SELECT 
			trr.rule_id,
			fr.name as rule_name,
			COUNT(*) as matches,
			COUNT(*) FILTER (WHERE trr.matched = true)::float / COUNT(*) FILTER (WHERE t.status = 'DECLINED') as share_of_declines,
			COUNT(DISTINCT t.user_id) as unique_users
		FROM transaction_rule_results trr
		JOIN transactions t ON trr.transaction_id = t.id
		JOIN fraud_rules fr ON trr.rule_id = fr.id
		WHERE t.created_at BETWEEN $1 AND $2
		GROUP BY trr.rule_id, fr.name
		ORDER BY matches DESC
	`

	rows, err := r.db.Query(ctx, query, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get rule matches stats: %w", err)
	}
	defer rows.Close()

	var stats []RuleMatchStat
	for rows.Next() {
		var stat RuleMatchStat
		err := rows.Scan(
			&stat.RuleID,
			&stat.RuleName,
			&stat.Matches,
			&stat.ShareOfDeclines,
			&stat.UniqueUsers,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rule match stat: %w", err)
		}
		stats = append(stats, stat)
	}

	return stats, nil
}

func (r *repository) GetMerchantRiskStats(ctx context.Context, from, to time.Time, limit int) ([]MerchantRiskStat, error) {
	query := `
		SELECT 
			merchant_id,
			merchant_category_code,
			COUNT(*) as tx_count,
			COALESCE(SUM(amount), 0) as gmv,
			COUNT(*) FILTER (WHERE status = 'DECLINED')::float / COUNT(*) as decline_rate
		FROM transactions 
		WHERE created_at BETWEEN $1 AND $2
		GROUP BY merchant_id, merchant_category_code
		ORDER BY decline_rate DESC, tx_count DESC
		LIMIT $3
	`

	rows, err := r.db.Query(ctx, query, from, to, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get merchant risk stats: %w", err)
	}
	defer rows.Close()

	var stats []MerchantRiskStat
	for rows.Next() {
		var stat MerchantRiskStat
		err := rows.Scan(
			&stat.MerchantID,
			&stat.MerchantCategoryCode,
			&stat.TxCount,
			&stat.GMV,
			&stat.DeclineRate,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan merchant risk stat: %w", err)
		}
		stats = append(stats, stat)
	}

	return stats, nil
}

func (r *repository) GetUserRiskProfile(ctx context.Context, userID uuid.UUID) (*UserRiskProfile, error) {
	query := `
		WITH user_stats_24h AS (
			SELECT 
				COUNT(*) as tx_count_24h,
				COALESCE(SUM(amount), 0) as gmv_24h,
				COUNT(DISTINCT device_id) as distinct_devices_24h,
				COUNT(DISTINCT ip_address) as distinct_ips_24h,
				COUNT(DISTINCT location->>'city') as distinct_cities_24h
			FROM transactions 
			WHERE user_id = $1 
			AND created_at >= NOW() - INTERVAL '24 hours'
		),
		user_stats_30d AS (
			SELECT 
				COUNT(*) FILTER (WHERE status = 'DECLINED')::float / COUNT(*) as decline_rate_30d
			FROM transactions 
			WHERE user_id = $1 
			AND created_at >= NOW() - INTERVAL '30 days'
		)
		SELECT 
			$1::uuid as user_id,
			COALESCE(s24.tx_count_24h, 0) as tx_count_24h,
			COALESCE(s24.gmv_24h, 0) as gmv_24h,
			COALESCE(s24.distinct_devices_24h, 0) as distinct_devices_24h,
			COALESCE(s24.distinct_ips_24h, 0) as distinct_ips_24h,
			COALESCE(s24.distinct_cities_24h, 0) as distinct_cities_24h,
			COALESCE(s30.decline_rate_30d, 0) as decline_rate_30d,
			COALESCE(MAX(created_at), NOW() - INTERVAL '1 day') as last_seen_at
		FROM user_stats_24h s24
		CROSS JOIN user_stats_30d s30
	`

	var profile UserRiskProfile
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&profile.UserID,
		&profile.TxCount24h,
		&profile.GMV24h,
		&profile.DistinctDevices24h,
		&profile.DistinctIPs24h,
		&profile.DistinctCities24h,
		&profile.DeclineRate30d,
		&profile.LastSeenAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user risk profile: %w", err)
	}

	return &profile, nil
}
