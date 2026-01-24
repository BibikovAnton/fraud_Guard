package stats

import (
	"context"
	"time"

	"github.com/google/uuid"
	"solution/internal/repository/stats"
)

type Service interface {
	// Stats Overview
	GetOverview(ctx context.Context, from, to time.Time) (*OverviewResult, error)
	
	// Time Series
	GetTransactionsTimeSeries(ctx context.Context, from, to time.Time, interval string) (*TimeSeriesResult, error)
	
	// Rule Matches
	GetRuleMatches(ctx context.Context, from, to time.Time) (*RuleMatchesResult, error)
	
	// Merchant Risk
	GetMerchantRisk(ctx context.Context, from, to time.Time, limit int) (*MerchantRiskResult, error)
	
	// User Risk Profile
	GetUserRiskProfile(ctx context.Context, userID uuid.UUID) (*UserRiskProfileResult, error)
}

type service struct {
	statsRepo stats.Repository
}

func NewService(statsRepo stats.Repository) Service {
	return &service{
		statsRepo: statsRepo,
	}
}

type OverviewResult struct {
	From         time.Time `json:"from"`
	To           time.Time `json:"to"`
	Volume       int64     `json:"volume"`
	GMV          float64   `json:"gmv"`
	ApprovalRate float64   `json:"approvalRate"`
	DeclineRate  float64   `json:"declineRate"`
	TopRiskMerchants []MerchantRiskItem `json:"topRiskMerchants"`
}

type MerchantRiskItem struct {
	MerchantID           string  `json:"merchantId"`
	MerchantCategoryCode string  `json:"merchantCategoryCode"`
	TxCount              int64   `json:"txCount"`
	GMV                  float64 `json:"gmv"`
	DeclineRate          float64 `json:"declineRate"`
}

type TimeSeriesResult struct {
	Points []TimeSeriesPoint `json:"points"`
}

type TimeSeriesPoint struct {
	BucketStart  time.Time `json:"bucketStart"`
	TxCount      int64     `json:"txCount"`
	GMV          float64   `json:"gmv"`
	ApprovalRate float64   `json:"approvalRate"`
	DeclineRate  float64   `json:"declineRate"`
}

type RuleMatchesResult struct {
	Items []RuleMatchItem `json:"items"`
}

type RuleMatchItem struct {
	RuleID          uuid.UUID `json:"ruleId"`
	RuleName        string    `json:"ruleName"`
	Matches         int64     `json:"matches"`
	ShareOfDeclines float64   `json:"shareOfDeclines"`
	UniqueUsers     int64     `json:"uniqueUsers"`
}

type MerchantRiskResult struct {
	Items []MerchantRiskItem `json:"items"`
}

type UserRiskProfileResult struct {
	UserID              uuid.UUID `json:"userId"`
	TxCount24h          int64     `json:"txCount_24h"`
	GMV24h              float64   `json:"gmv_24h"`
	DistinctDevices24h  int64     `json:"distinctDevices_24h"`
	DistinctIPs24h      int64     `json:"distinctIps_24h"`
	DistinctCities24h    int64     `json:"distinctCities_24h"`
	DeclineRate30d      float64   `json:"declineRate_30d"`
	LastSeenAt          time.Time `json:"lastSeenAt"`
}

func (s *service) GetOverview(ctx context.Context, from, to time.Time) (*OverviewResult, error) {
	// Get overview stats
	stats, err := s.statsRepo.GetOverviewStats(ctx, from, to)
	if err != nil {
		return nil, err
	}

	// Get top risk merchants
	merchantStats, err := s.statsRepo.GetMerchantRiskStats(ctx, from, to, 5)
	if err != nil {
		return nil, err
	}

	// Convert to result format
	topRiskMerchants := make([]MerchantRiskItem, len(merchantStats))
	for i, m := range merchantStats {
		topRiskMerchants[i] = MerchantRiskItem{
			MerchantID:           m.MerchantID,
			MerchantCategoryCode: m.MerchantCategoryCode,
			TxCount:              m.TxCount,
			GMV:                  m.GMV,
			DeclineRate:          m.DeclineRate,
		}
	}

	return &OverviewResult{
		From:            stats.From,
		To:              stats.To,
		Volume:          stats.Volume,
		GMV:             stats.GMV,
		ApprovalRate:    stats.ApprovalRate,
		DeclineRate:     stats.DeclineRate,
		TopRiskMerchants: topRiskMerchants,
	}, nil
}

func (s *service) GetTransactionsTimeSeries(ctx context.Context, from, to time.Time, interval string) (*TimeSeriesResult, error) {
	points, err := s.statsRepo.GetTransactionsTimeSeries(ctx, from, to, interval)
	if err != nil {
		return nil, err
	}

	// Convert to result format
	result := make([]TimeSeriesPoint, len(points))
	for i, p := range points {
		result[i] = TimeSeriesPoint{
			BucketStart:  p.BucketStart,
			TxCount:      p.TxCount,
			GMV:          p.GMV,
			ApprovalRate: p.ApprovalRate,
			DeclineRate:  p.DeclineRate,
		}
	}

	return &TimeSeriesResult{
		Points: result,
	}, nil
}

func (s *service) GetRuleMatches(ctx context.Context, from, to time.Time) (*RuleMatchesResult, error) {
	stats, err := s.statsRepo.GetRuleMatchesStats(ctx, from, to)
	if err != nil {
		return nil, err
	}

	// Convert to result format
	items := make([]RuleMatchItem, len(stats))
	for i, s := range stats {
		items[i] = RuleMatchItem{
			RuleID:          s.RuleID,
			RuleName:        s.RuleName,
			Matches:         s.Matches,
			ShareOfDeclines: s.ShareOfDeclines,
			UniqueUsers:     s.UniqueUsers,
		}
	}

	return &RuleMatchesResult{
		Items: items,
	}, nil
}

func (s *service) GetMerchantRisk(ctx context.Context, from, to time.Time, limit int) (*MerchantRiskResult, error) {
	stats, err := s.statsRepo.GetMerchantRiskStats(ctx, from, to, limit)
	if err != nil {
		return nil, err
	}

	// Convert to result format
	items := make([]MerchantRiskItem, len(stats))
	for i, m := range stats {
		items[i] = MerchantRiskItem{
			MerchantID:           m.MerchantID,
			MerchantCategoryCode: m.MerchantCategoryCode,
			TxCount:              m.TxCount,
			GMV:                  m.GMV,
			DeclineRate:          m.DeclineRate,
		}
	}

	return &MerchantRiskResult{
		Items: items,
	}, nil
}

func (s *service) GetUserRiskProfile(ctx context.Context, userID uuid.UUID) (*UserRiskProfileResult, error) {
	profile, err := s.statsRepo.GetUserRiskProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &UserRiskProfileResult{
		UserID:              profile.UserID,
		TxCount24h:          profile.TxCount24h,
		GMV24h:              profile.GMV24h,
		DistinctDevices24h:  profile.DistinctDevices24h,
		DistinctIPs24h:      profile.DistinctIPs24h,
		DistinctCities24h:    profile.DistinctCities24h,
		DeclineRate30d:      profile.DeclineRate30d,
		LastSeenAt:          profile.LastSeenAt,
	}, nil
}
