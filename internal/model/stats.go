package model

import (
	"time"
)

type StatsOverview struct {
	From            time.Time              `json:"from"`
	To              time.Time              `json:"to"`
	Volume          int64                  `json:"volume"`
	GMV             float64                `json:"gmv"`
	ApprovalRate    float64                `json:"approvalRate"`
	DeclineRate     float64                `json:"declineRate"`
	TopRiskMerchants []MerchantRiskMetrics `json:"topRiskMerchants"`
}

type MerchantRiskMetrics struct {
	MerchantID           string  `json:"merchantId"`
	MerchantCategoryCode *string `json:"merchantCategoryCode,omitempty"`
	TxCount              int64   `json:"txCount"`
	GMV                  float64 `json:"gmv"`
	DeclineRate          float64 `json:"declineRate"`
}

type TransactionsTimeSeries struct {
	Points []TimeSeriesPoint `json:"points"`
}

type TimeSeriesPoint struct {
	BucketStart time.Time `json:"bucketStart"`
	TxCount     int64     `json:"txCount"`
	GMV         float64   `json:"gmv"`
	ApprovalRate float64  `json:"approvalRate"`
	DeclineRate  float64  `json:"declineRate"`
}

type RuleMatchStats struct {
	Items []RuleMatchMetrics `json:"items"`
}

type RuleMatchMetrics struct {
	RuleID          string  `json:"ruleId"`
	RuleName        string  `json:"ruleName"`
	Matches         int64   `json:"matches"`
	UniqueUsers     int64   `json:"uniqueUsers"`
	UniqueMerchants *int64  `json:"uniqueMerchants,omitempty"`
	ShareOfDeclines float64 `json:"shareOfDeclines"`
}

type MerchantRiskStats struct {
	Items []MerchantRiskMetrics `json:"items"`
}

type UserRiskProfile struct {
	UserID             string     `json:"userId"`
	TxCount24h         int        `json:"txCount_24h"`
	GMV24h             float64    `json:"gmv_24h"`
	DistinctDevices24h int        `json:"distinctDevices_24h"`
	DistinctIps24h     int        `json:"distinctIps_24h"`
	DistinctCities24h  int        `json:"distinctCities_24h"`
	DeclineRate30d     float64    `json:"declineRate_30d"`
	LastSeenAt         *time.Time `json:"lastSeenAt,omitempty"`
}

type GroupByType string

const (
	GroupByHour  GroupByType = "hour"
	GroupByDay   GroupByType = "day"
	GroupByWeek  GroupByType = "week"
)

type StatsOverviewParams struct {
	From     *time.Time
	To       *time.Time
	Timezone string
}

type TimeSeriesParams struct {
	From     *time.Time
	To       *time.Time
	GroupBy  GroupByType
	Timezone string
	Channel  *TransactionChannel
}

type RuleMatchesParams struct {
	From *time.Time
	To   *time.Time
	Top  int
}

type MerchantRiskParams struct {
	From                *time.Time
	To                  *time.Time
	MerchantCategoryCode *string
	Top                 int
}

type Time = time.Time
