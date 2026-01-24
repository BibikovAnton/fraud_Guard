package v1

import (
	"context"
	"time"

	"github.com/google/uuid"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
	"solution/internal/service/stats"
)

type statsHandlerAdapter struct {
	statsService stats.Service
}

func NewStatsHandlerAdapter(statsService stats.Service) *statsHandlerAdapter {
	return &statsHandlerAdapter{
		statsService: statsService,
	}
}

func (h *statsHandlerAdapter) APIV1StatsOverviewGet(ctx context.Context, params antifraud_v1.APIV1StatsOverviewGetParams) (antifraud_v1.APIV1StatsOverviewGetRes, error) {
	// Parse query parameters
	from := time.Now().AddDate(0, -1, 0) // Default: 1 month ago
	to := time.Now()
	
	if params.From.Set {
		from = params.From.Value
	}
	if params.To.Set {
		to = params.To.Value
	}

	// Get stats from service
	result, err := h.statsService.GetOverview(ctx, from, to)
	if err != nil {
		return nil, err
	}

	// Convert top risk merchants
	topRiskMerchants := make([]antifraud_v1.MerchantRiskRow, len(result.TopRiskMerchants))
	for i, merchant := range result.TopRiskMerchants {
		topRiskMerchants[i] = antifraud_v1.MerchantRiskRow{
			MerchantId:           merchant.MerchantID,
			MerchantCategoryCode: merchant.MerchantCategoryCode,
			TxCount:              merchant.TxCount,
			GMV:                  merchant.GMV,
			DeclineRate:          merchant.DeclineRate,
		}
	}

	return &antifraud_v1.StatsOverview{
		From:            result.From.Format(time.RFC3339),
		To:              result.To.Format(time.RFC3339),
		Volume:          result.Volume,
		Gmv:             result.GMV,
		ApprovalRate:    result.ApprovalRate,
		DeclineRate:     result.DeclineRate,
		TopRiskMerchants: topRiskMerchants,
	}, nil
}

func (h *statsHandlerAdapter) APIV1StatsTransactionsTimeseriesGet(ctx context.Context, params antifraud_v1.APIV1StatsTransactionsTimeseriesGetParams) (antifraud_v1.APIV1StatsTransactionsTimeseriesGetRes, error) {
	// Parse query parameters
	from := time.Now().AddDate(0, -1, 0) // Default: 1 month ago
	to := time.Now()
	interval := "day" // Default interval
	
	if params.From.Set {
		from = params.From.Value
	}
	if params.To.Set {
		to = params.To.Value
	}
	if params.Interval.Set {
		interval = params.Interval.Value
	}

	// Get time series from service
	result, err := h.statsService.GetTransactionsTimeSeries(ctx, from, to, interval)
	if err != nil {
		return nil, err
	}

	// Convert points
	points := make([]antifraud_v1.TransactionsTimePoint, len(result.Points))
	for i, point := range result.Points {
		points[i] = antifraud_v1.TransactionsTimePoint{
			BucketStart:  point.BucketStart.Format(time.RFC3339),
			TxCount:      point.TxCount,
			Gmv:          point.GMV,
			ApprovalRate: point.ApprovalRate,
			DeclineRate:  point.DeclineRate,
		}
	}

	return &antifraud_v1.TransactionsTimeSeries{
		Points: points,
	}, nil
}

func (h *statsHandlerAdapter) APIV1StatsRulesMatchesGet(ctx context.Context, params antifraud_v1.APIV1StatsRulesMatchesGetParams) (antifraud_v1.APIV1StatsRulesMatchesGetRes, error) {
	// Parse query parameters
	from := time.Now().AddDate(0, -1, 0) // Default: 1 month ago
	to := time.Now()
	
	if params.From.Set {
		from = params.From.Value
	}
	if params.To.Set {
		to = params.To.Value
	}

	// Get rule matches from service
	result, err := h.statsService.GetRuleMatches(ctx, from, to)
	if err != nil {
		return nil, err
	}

	// Convert items
	items := make([]antifraud_v1.RuleMatchRow, len(result.Items))
	for i, item := range result.Items {
		items[i] = antifraud_v1.RuleMatchRow{
			RuleId:          item.RuleID,
			RuleName:        item.RuleName,
			Matches:         item.Matches,
			ShareOfDeclines: item.ShareOfDeclines,
			UniqueUsers:     item.UniqueUsers,
		}
	}

	return &antifraud_v1.StatsRuleMatches{
		Items: items,
	}, nil
}

func (h *statsHandlerAdapter) APIV1StatsMerchantsRiskGet(ctx context.Context, params antifraud_v1.APIV1StatsMerchantsRiskGetParams) (antifraud_v1.APIV1StatsMerchantsRiskGetRes, error) {
	// Parse query parameters
	from := time.Now().AddDate(0, -1, 0) // Default: 1 month ago
	to := time.Now()
	limit := 100 // Default limit
	
	if params.From.Set {
		from = params.From.Value
	}
	if params.To.Set {
		to = params.To.Value
	}
	if params.Limit.Set {
		limit = params.Limit.Value
	}

	// Get merchant risk from service
	result, err := h.statsService.GetMerchantRisk(ctx, from, to, limit)
	if err != nil {
		return nil, err
	}

	// Convert items
	items := make([]antifraud_v1.MerchantRiskRow, len(result.Items))
	for i, item := range result.Items {
		items[i] = antifraud_v1.MerchantRiskRow{
			MerchantId:           item.MerchantID,
			MerchantCategoryCode: item.MerchantCategoryCode,
			TxCount:              item.TxCount,
			Gmv:                  item.GMV,
			DeclineRate:          item.DeclineRate,
		}
	}

	return &antifraud_v1.StatsMerchantRisk{
		Items: items,
	}, nil
}

func (h *statsHandlerAdapter) APIV1StatsUsersIDRiskProfileGet(ctx context.Context, params antifraud_v1.APIV1StatsUsersIDRiskProfileGetParams) (antifraud_v1.APIV1StatsUsersIDRiskProfileGetRes, error) {
	userID := params.ID

	// Get user risk profile from service
	result, err := h.statsService.GetUserRiskProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &antifraud_v1.UserRiskProfile{
		UserId:              result.UserID,
		TxCount24h:          result.TxCount24h,
		Gmv24h:              result.GMV24h,
		DistinctDevices24h:  result.DistinctDevices24h,
		DistinctIps24h:      result.DistinctIPs24h,
		DistinctCities24h:    result.DistinctCities24h,
		DeclineRate30d:      result.DeclineRate30d,
		LastSeenAt:          result.LastSeenAt.Format(time.RFC3339),
	}, nil
}
