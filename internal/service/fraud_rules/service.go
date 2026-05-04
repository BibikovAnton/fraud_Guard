package fraud_rules

import (
	"go.uber.org/zap"
	"solution/internal/dsl"
	"solution/internal/repository"
)

type service struct {
	fraudRuleRepo repository.FraudRuleRepository
	dslEvaluator *dsl.DSLEvaluator
	logger        *zap.Logger
}

func NewService(fraudRuleRepo repository.FraudRuleRepository, logger *zap.Logger) *service {
	return &service{
		fraudRuleRepo: fraudRuleRepo,
		dslEvaluator: dsl.NewDSLEvaluator(logger),
		logger:        logger,
	}
}
