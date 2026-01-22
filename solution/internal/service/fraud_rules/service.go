package fraud_rules

import "solution/internal/repository"

type service struct {
	fraudRuleRepo repository.FraudRuleRepository
}

func NewService(fraudRuleRepo repository.FraudRuleRepository) *service {
	return &service{
		fraudRuleRepo: fraudRuleRepo,
	}
}
