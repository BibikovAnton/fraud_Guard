package stats

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	// Stats Overview
	GetOverviewStats(ctx context.Context, from, to time.Time) (*OverviewStats, error)
	
	// Time Series
	GetTransactionsTimeSeries(ctx context.Context, from, to time.Time, interval string) ([]TimeSeriesPoint, error)
	
	// Rule Matches
	GetRuleMatchesStats(ctx context.Context, from, to time.Time) ([]RuleMatchStat, error)
	
	// Merchant Risk
	GetMerchantRiskStats(ctx context.Context, from, to time.Time, limit int) ([]MerchantRiskStat, error)
	
	// User Risk Profile
	GetUserRiskProfile(ctx context.Context, userID uuid.UUID) (*UserRiskProfile, error)
}

type OverviewStats struct {
	Volume       int64   `json:"volume"`
	GMV          float64 `json:"gmv"`
	ApprovalRate float64 `json:"approvalRate"`
	DeclineRate  float64 `json:"declineRate"`
	From         time.Time `json:"from"`
	To           time.Time `json:"to"`
}

type TimeSeriesPoint struct {
	BucketStart  time.Time `json:"bucketStart"`
	TxCount      int64     `json:"txCount"`
	GMV          float64   `json:"gmv"`
	ApprovalRate float64   `json:"approvalRate"`
	DeclineRate  float64   `json:"declineRate"`
}

type RuleMatchStat struct {
	RuleID            uuid.UUID `json:"ruleId"`
	RuleName          string    `json:"ruleName"`
	Matches           int64     `json:"matches"`
	ShareOfDeclines   float64   `json:"shareOfDeclines"`
	UniqueUsers       int64     `json:"uniqueUsers"`
}

type MerchantRiskStat struct {
	MerchantID           string  `json:"merchantId"`
	MerchantCategoryCode string  `json:"merchantCategoryCode"`
	TxCount              int64   `json:"txCount"`
	GMV                  float64 `json:"gmv"`
	DeclineRate          float64 `json:"declineRate"`
}

type UserRiskProfile struct {
	UserID              uuid.UUID `json:"userId"`
	TxCount24h          int64     `json:"txCount_24h"`
	GMV24h              float64   `json:"gmv_24h"`
	DistinctDevices24h  int64     `json:"distinctDevices_24h"`
	DistinctIPs24h      int64     `json:"distinctIps_24h"`
	DistinctCities24h    int64     `json:"distinctCities_24h"`
	DeclineRate30d      float64   `json:"declineRate_30d"`
	LastSeenAt          time.Time `json:"lastSeenAt"`
}
