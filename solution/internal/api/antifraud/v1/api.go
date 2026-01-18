package v1

import "solution/internal/service"

type api struct {
	antiFraudService service.AntifraudService
}

func NewAPI(antiFraudService service.AntifraudService) *api {
	return &api{
		antiFraudService: antiFraudService,
	}
}
