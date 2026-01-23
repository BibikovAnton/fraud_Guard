package v1

import (
	"context"
	"time"

	"github.com/google/uuid"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
)

func (h *handlerAdapter) APIV1StatsMerchantsRiskGet(ctx context.Context, params antifraud_v1.APIV1StatsMerchantsRiskGetParams) (antifraud_v1.APIV1StatsMerchantsRiskGetRes, error) {
	items := []antifraud_v1.MerchantRiskRow{
		{
			MerchantId:           "merchant_001",
			MerchantCategoryCode: antifraud_v1.OptMccCode{Set: true, Value: "5411"},
			TxCount:              150,
		},
		{
			MerchantId:           "merchant_002", 
			MerchantCategoryCode: antifraud_v1.OptMccCode{Set: true, Value: "5812"},
			TxCount:              89,
		},
	}

	return &antifraud_v1.MerchantRiskStats{
		Items: items,
	}, nil
}

func (h *handlerAdapter) APIV1StatsOverviewGet(ctx context.Context, params antifraud_v1.APIV1StatsOverviewGetParams) (antifraud_v1.APIV1StatsOverviewGetRes, error) {
	now := time.Now()
	from := now.AddDate(0, -1, 0)
	
	topRiskMerchants := []antifraud_v1.MerchantRiskRow{
		{
			MerchantId:           "merchant_001",
			MerchantCategoryCode: antifraud_v1.OptMccCode{Set: true, Value: "5411"},
			TxCount:              150,
		},
	}

	return &antifraud_v1.StatsOverview{
		From:             from,
		To:               now,
		Volume:           1250,
		Gmv:              156789.50,
		ApprovalRate:     0.85,
		DeclineRate:      0.15,
		TopRiskMerchants: topRiskMerchants,
	}, nil
}

func (h *handlerAdapter) APIV1StatsRulesMatchesGet(ctx context.Context, params antifraud_v1.APIV1StatsRulesMatchesGetParams) (antifraud_v1.APIV1StatsRulesMatchesGetRes, error) {
	items := []antifraud_v1.RuleMatchRow{
		{
			RuleId:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			RuleName: "High Amount Check",
			Matches:  45,
		},
		{
			RuleId:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			RuleName: "Old Humans Risk",
			Matches:  23,
		},
	}

	return &antifraud_v1.RuleMatchStats{
		Items: items,
	}, nil
}

func (h *handlerAdapter) APIV1StatsTransactionsTimeseriesGet(ctx context.Context, params antifraud_v1.APIV1StatsTransactionsTimeseriesGetParams) (antifraud_v1.APIV1StatsTransactionsTimeseriesGetRes, error) {
	now := time.Now()
	points := make([]antifraud_v1.TransactionsTimePoint, 7)
	
	for i := 0; i < 7; i++ {
		points[i] = antifraud_v1.TransactionsTimePoint{
			BucketStart: now.AddDate(0, 0, -6+i),
			TxCount:     100 + i*10,
			Gmv:         float64(1000+i*100),
		}
	}

	return &antifraud_v1.TransactionsTimeSeries{
		Points: points,
	}, nil
}

func (h *handlerAdapter) APIV1StatsUsersIDRiskProfileGet(ctx context.Context, params antifraud_v1.APIV1StatsUsersIDRiskProfileGetParams) (antifraud_v1.APIV1StatsUsersIDRiskProfileGetRes, error) {
	userID := params.ID
	
	return &antifraud_v1.UserRiskProfile{
		UserId:     userID,
		TxCount24h: 25,
		Gmv24h:     1250.50,
		LastSeenAt: antifraud_v1.OptNilDateTime{
			Set:   true,
			Value: time.Now().Add(-24 * time.Hour),
			Null:  false,
		},
	}, nil
}
