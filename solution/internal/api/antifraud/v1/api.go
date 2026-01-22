package v1

import (
	"solution/internal/service"
)

type API struct {
	userService      service.UserService
	fraudRuleService service.FraudRuleService
}

func NewAPI(
	userService service.UserService,
	fraudRuleService service.FraudRuleService,
) *API {
	return &API{
		userService:      userService,
		fraudRuleService: fraudRuleService,
	}
}
