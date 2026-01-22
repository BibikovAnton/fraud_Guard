package v1

import (
	"context"
	"solution/internal/service"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
)

type handlerAdapter struct {
	userService      service.UserService
	fraudRuleService service.FraudRuleService
}

func NewHandlerAdapter(userService service.UserService, fraudRuleService service.FraudRuleService) antifraud_v1.Handler {
	return &handlerAdapter{
		userService:      userService,
		fraudRuleService: fraudRuleService,
	}
}

func (h *handlerAdapter) APIV1PingGet(ctx context.Context) (*antifraud_v1.APIV1PingGetOK, error) {
	opt := antifraud_v1.OptString{
		Value: "ok",
	}
	return &antifraud_v1.APIV1PingGetOK{
		Status: opt,
	}, nil

}
