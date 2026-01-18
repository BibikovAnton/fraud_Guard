package v1

import (
	"context"
	"fmt"
	"solution/internal/model"
	"solution/internal/service"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"

	"github.com/google/uuid"
)

type handlerAdapter struct {
	antifraudService service.AntifraudService
	userService      service.UserService
}

func NewHandlerAdapter(antifraudService service.AntifraudService, userService service.UserService) antifraud_v1.Handler {
	return &handlerAdapter{
		antifraudService: antifraudService,
		userService:      userService,
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

func (h *handlerAdapter) APIV1AuthLoginPost(ctx context.Context, req *antifraud_v1.LoginRequest) (antifraud_v1.APIV1AuthLoginPostRes, error) {
	// Конвертируем запрос в нашу модель
	loginReq := model.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	// Вызываем сервис
	authResp, err := h.userService.Login(ctx, loginReq)
	if err != nil {
		// TODO: добавить обработку разных типов ошибок (401, 423)
		return nil, fmt.Errorf("login failed: %w", err)
	}

	// Конвертируем ответ в OpenAPI модель
	apiResp := antifraud_v1.AuthResponse{
		AccessToken: authResp.AccessToken,
		ExpiresIn:   authResp.ExpiresIn,
		User:        convertUserToAPI(authResp.User),
	}

	return &apiResp, nil
}

func (h *handlerAdapter) APIV1AuthRegisterPost(ctx context.Context, req *antifraud_v1.RegisterRequest) (antifraud_v1.APIV1AuthRegisterPostRes, error) {
	// Конвертируем запрос в нашу модель
	registerReq := model.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
	}

	// Обрабатываем опциональные поля
	if req.Age.Set {
		age := req.Age.Value
		registerReq.Age = &age
	}
	if req.Region.Set {
		region := req.Region.Value
		registerReq.Region = &region
	}
	if req.Gender.Set {
		gender := model.Gender(req.Gender.Value)
		registerReq.Gender = &gender
	}
	if req.MaritalStatus.Set {
		maritalStatus := model.MaritalStatus(req.MaritalStatus.Value)
		registerReq.MaritalStatus = &maritalStatus
	}

	// Вызываем сервис
	authResp, err := h.userService.Register(ctx, registerReq)
	if err != nil {
		// TODO: добавить обработку разных типов ошибок (409, 422)
		return nil, fmt.Errorf("registration failed: %w", err)
	}

	// Конвертируем ответ в OpenAPI модель
	apiResp := antifraud_v1.AuthResponse{
		AccessToken: authResp.AccessToken,
		ExpiresIn:   authResp.ExpiresIn,
		User:        convertUserToAPI(authResp.User),
	}

	return &apiResp, nil
}

func (h *handlerAdapter) APIV1FraudRulesGet(ctx context.Context) (antifraud_v1.APIV1FraudRulesGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1FraudRulesIDDelete(ctx context.Context, params antifraud_v1.APIV1FraudRulesIDDeleteParams) (antifraud_v1.APIV1FraudRulesIDDeleteRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1FraudRulesIDGet(ctx context.Context, params antifraud_v1.APIV1FraudRulesIDGetParams) (antifraud_v1.APIV1FraudRulesIDGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1FraudRulesIDPut(ctx context.Context, req *antifraud_v1.FraudRuleUpdateRequest, params antifraud_v1.APIV1FraudRulesIDPutParams) (antifraud_v1.APIV1FraudRulesIDPutRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1FraudRulesPost(ctx context.Context, req *antifraud_v1.FraudRuleCreateRequest) (antifraud_v1.APIV1FraudRulesPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1FraudRulesValidatePost(ctx context.Context, req *antifraud_v1.DslValidateRequest) (antifraud_v1.APIV1FraudRulesValidatePostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1StatsMerchantsRiskGet(ctx context.Context, params antifraud_v1.APIV1StatsMerchantsRiskGetParams) (antifraud_v1.APIV1StatsMerchantsRiskGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1StatsOverviewGet(ctx context.Context, params antifraud_v1.APIV1StatsOverviewGetParams) (antifraud_v1.APIV1StatsOverviewGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1StatsRulesMatchesGet(ctx context.Context, params antifraud_v1.APIV1StatsRulesMatchesGetParams) (antifraud_v1.APIV1StatsRulesMatchesGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1StatsTransactionsTimeseriesGet(ctx context.Context, params antifraud_v1.APIV1StatsTransactionsTimeseriesGetParams) (antifraud_v1.APIV1StatsTransactionsTimeseriesGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1StatsUsersIDRiskProfileGet(ctx context.Context, params antifraud_v1.APIV1StatsUsersIDRiskProfileGetParams) (antifraud_v1.APIV1StatsUsersIDRiskProfileGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1TransactionsBatchPost(ctx context.Context, req *antifraud_v1.TransactionBatchCreateRequest) (antifraud_v1.APIV1TransactionsBatchPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1TransactionsGet(ctx context.Context, params antifraud_v1.APIV1TransactionsGetParams) (antifraud_v1.APIV1TransactionsGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1TransactionsIDGet(ctx context.Context, params antifraud_v1.APIV1TransactionsIDGetParams) (antifraud_v1.APIV1TransactionsIDGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1TransactionsPost(ctx context.Context, req *antifraud_v1.TransactionCreateRequest) (antifraud_v1.APIV1TransactionsPostRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersGet(ctx context.Context, params antifraud_v1.APIV1UsersGetParams) (antifraud_v1.APIV1UsersGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersIDDelete(ctx context.Context, params antifraud_v1.APIV1UsersIDDeleteParams) (antifraud_v1.APIV1UsersIDDeleteRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersIDGet(ctx context.Context, params antifraud_v1.APIV1UsersIDGetParams) (antifraud_v1.APIV1UsersIDGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersIDPut(ctx context.Context, req *antifraud_v1.UserUpdateRequest, params antifraud_v1.APIV1UsersIDPutParams) (antifraud_v1.APIV1UsersIDPutRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersMeGet(ctx context.Context) (antifraud_v1.APIV1UsersMeGetRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersMePut(ctx context.Context, req *antifraud_v1.UserUpdateRequest) (antifraud_v1.APIV1UsersMePutRes, error) {
	return nil, nil
}

func (h *handlerAdapter) APIV1UsersPost(ctx context.Context, req *antifraud_v1.UserCreateRequest) (antifraud_v1.APIV1UsersPostRes, error) {
	return nil, nil
}

// convertUserToAPI конвертирует нашу модель User в OpenAPI модель
// Важно: правильно обрабатывать nullable поля для соответствия спецификации
func convertUserToAPI(user model.User) antifraud_v1.User {
	apiUser := antifraud_v1.User{
		Email:     user.Email,
		FullName:  user.FullName,
		Role:      antifraud_v1.UserRole(user.Role),
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	// Конвертируем ID из string в uuid.UUID
	id, err := uuid.Parse(user.ID)
	if err == nil {
		apiUser.ID = id
	}

	// Обрабатываем nullable поля
	if user.Age != nil {
		apiUser.Age = antifraud_v1.OptInt{Value: *user.Age, Set: true}
	}
	
	if user.Region != nil {
		apiUser.Region = antifraud_v1.OptString{Value: *user.Region, Set: true}
	}
	
	if user.Gender != nil {
		apiUser.Gender = antifraud_v1.OptGender{Value: antifraud_v1.Gender(*user.Gender), Set: true}
	}
	
	if user.MaritalStatus != nil {
		apiUser.MaritalStatus = antifraud_v1.OptMaritalStatus{Value: antifraud_v1.MaritalStatus(*user.MaritalStatus), Set: true}
	}

	return apiUser
}
