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
			CASE 
				WHEN COUNT(*) > 0 THEN COUNT(*) FILTER (WHERE status = 'APPROVED')::float / COUNT(*)
				ELSE 0 
			END as approval_rate,
			CASE 
				WHEN COUNT(*) > 0 THEN COUNT(*) FILTER (WHERE status = 'DECLINED')::float / COUNT(*)
				ELSE 0 
			END as decline_rate
		FROM transactions 
		WHERE timestamp BETWEEN $1 AND $2
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

	
	merchantQuery := `
		SELECT 
			merchant_id,
			merchant_category_code,
			COUNT(*) as tx_count,
			COALESCE(SUM(amount), 0) as gmv,
			COUNT(*) FILTER (WHERE status = 'DECLINED')::float / COUNT(*) as decline_rate
		FROM transactions 
		WHERE timestamp BETWEEN $1 AND $2 
			AND merchant_id IS NOT NULL
		GROUP BY merchant_id, merchant_category_code
		ORDER BY decline_rate DESC, tx_count DESC
		LIMIT 10
	`

	fmt.Printf("DEBUG: Executing merchant query with from=%v, to=%v\n", from, to)
	rows, err := r.db.Query(ctx, merchantQuery, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get merchant risk stats: %w", err)
	}
	defer rows.Close()

	var topRiskMerchants []MerchantRiskStat
	for rows.Next() {
		var merchant MerchantRiskStat
		var merchantCategoryCode sql.NullString
		
		err := rows.Scan(
			&merchant.MerchantID,
			&merchantCategoryCode,
			&merchant.TxCount,
			&merchant.GMV,
			&merchant.DeclineRate,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan merchant risk stat: %w", err)
		}
		
		// Handle NULL values
		if merchantCategoryCode.Valid {
			merchant.MerchantCategoryCode = merchantCategoryCode.String
		} else {
			merchant.MerchantCategoryCode = ""
		}
		
		topRiskMerchants = append(topRiskMerchants, merchant)
	}

	stats.From = from
	stats.To = to
	stats.TopRiskMerchants = topRiskMerchants
	return &stats, nil
}

func (r *repository) GetTransactionsTimeSeries(ctx context.Context, from, to time.Time, interval string) ([]TimeSeriesPoint, error) {
	query := fmt.Sprintf(`
		SELECT 
			date_trunc('%s', timestamp) as bucket_start,
			COUNT(*) as tx_count,
			COALESCE(SUM(amount), 0) as gmv,
			COUNT(*) FILTER (WHERE status = 'APPROVED')::float / COUNT(*) as approval_rate,
			COUNT(*) FILTER (WHERE status = 'DECLINED')::float / COUNT(*) as decline_rate
		FROM transactions 
		WHERE timestamp BETWEEN $1 AND $2
		GROUP BY date_trunc('%s', timestamp)
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
			fr.name,
			COUNT(*) as matches,
			COUNT(*) FILTER (WHERE trr.matched = true) * 1.0 / COUNT(*) as share_of_declines,
			COUNT(DISTINCT trr.transaction_id) as unique_users
		FROM transaction_rule_results trr
		JOIN transactions t ON trr.transaction_id = t.id
		JOIN fraud_rules fr ON trr.rule_id = fr.id
		WHERE t.timestamp BETWEEN $1 AND $2
			AND t.status = 'DECLINED'
			AND trr.rule_id != '00000000-0000-0000-0000-000000000000'
		GROUP BY trr.rule_id, fr.name
		HAVING COUNT(*) FILTER (WHERE trr.matched = true) > 0
		ORDER BY matches DESC
	`

	fmt.Printf("DEBUG: Rule matches query: %s\n", query)
	fmt.Printf("DEBUG: Query params: from=%v, to=%v\n", from, to)
	
	// First, let's see what's in the tables
	debugQuery1 := `SELECT COUNT(*) as total_results FROM transaction_rule_results WHERE rule_id != '00000000-0000-0000-0000-000000000000'`
	var totalResults int
	r.db.QueryRow(ctx, debugQuery1).Scan(&totalResults)
	fmt.Printf("DEBUG: Total non-nil rule_results: %d\n", totalResults)
	
	debugQuery2 := `SELECT rule_id, matched, COUNT(*) FROM transaction_rule_results WHERE rule_id != '00000000-0000-0000-0000-000000000000' GROUP BY rule_id, matched ORDER BY rule_id, matched`
	rows2, _ := r.db.Query(ctx, debugQuery2)
	fmt.Printf("DEBUG: Rule results breakdown:\n")
	for rows2.Next() {
		var ruleID string
		var matched bool
		var count int
		rows2.Scan(&ruleID, &matched, &count)
		fmt.Printf("  - %s: matched=%v, count=%d\n", ruleID, matched, count)
	}
	rows2.Close()

	rows, err := r.db.Query(ctx, query, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get rule matches stats: %w", err)
	}
	defer rows.Close()

	fmt.Printf("DEBUG: Rule matches query executed, checking results...\n")

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
		fmt.Printf("DEBUG: Found rule match: %s - %s (%d matches)\n", stat.RuleID, stat.RuleName, stat.Matches)
	}
	
	fmt.Printf("DEBUG: Total rule matches found: %d\n", len(stats))

	return stats, nil
}

func (r *repository) GetMerchantRiskStats(ctx context.Context, from, to time.Time, limit int) ([]MerchantRiskStat, error) {
	query := `
		SELECT 
			COALESCE(merchant_id, 'unknown') as merchant_id,
			COALESCE(merchant_category_code, '') as merchant_category_code,
			COUNT(*) as tx_count,
			COALESCE(SUM(amount), 0) as gmv,
			COUNT(*) FILTER (WHERE status = 'DECLINED')::float / COUNT(*) as decline_rate
		FROM transactions 
		WHERE timestamp BETWEEN $1 AND $2
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
				COUNT(DISTINCT CASE WHEN location IS NOT NULL THEN location->>'city' END) as distinct_cities_24h
			FROM transactions 
			WHERE user_id = $1 
			AND timestamp >= NOW() - INTERVAL '24 hours'
		),
		user_stats_30d AS (
			SELECT 
				CASE 
					WHEN COUNT(*) > 0 THEN COUNT(*) FILTER (WHERE status = 'DECLINED')::float / COUNT(*)
					ELSE 0 
				END as decline_rate_30d
			FROM transactions 
			WHERE user_id = $1 
			AND timestamp >= NOW() - INTERVAL '30 days'
		)
		SELECT 
			$1::uuid as user_id,
			COALESCE(s24.tx_count_24h, 0) as tx_count_24h,
			COALESCE(s24.gmv_24h, 0) as gmv_24h,
			COALESCE(s24.distinct_devices_24h, 0) as distinct_devices_24h,
			COALESCE(s24.distinct_ips_24h, 0) as distinct_ips_24h,
			COALESCE(s24.distinct_cities_24h, 0) as distinct_cities_24h,
			COALESCE(s30.decline_rate_30d, 0) as decline_rate_30d,
			COALESCE((SELECT MAX(timestamp) FROM transactions WHERE user_id = $1), NOW() - INTERVAL '1 day') as last_seen_at
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
