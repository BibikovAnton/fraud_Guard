package v1

import (
	"context"
	"fmt"
	"solution/internal/middleware"
	"solution/internal/model"
	"solution/internal/service"
	"strings"
	"time"

	"github.com/google/uuid"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
)

type handlerAdapter struct {
	antifraudService service.AntifraudService
	userService      service.UserService
	fraudRuleService service.FraudRuleService
}

func NewHandlerAdapter(antifraudService service.AntifraudService, userService service.UserService, fraudRuleService service.FraudRuleService) antifraud_v1.Handler {
	return &handlerAdapter{
		antifraudService: antifraudService,
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

func (h *handlerAdapter) APIV1AuthLoginPost(ctx context.Context, req *antifraud_v1.LoginRequest) (antifraud_v1.APIV1AuthLoginPostRes, error) {
	// Валидация полей входа
	if err := h.validateLoginRequest(req); err != nil {
		return &antifraud_v1.APIV1AuthLoginPostBadRequest{
			Code:      antifraud_v1.ErrorCodeVALIDATIONFAILED,
			Message:   err.Error(),
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/auth/login",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	loginReq := model.LoginRequest{
		Email:    strings.TrimSpace(strings.ToLower(req.Email)), // нормализация email
		Password: req.Password,
	}

	authResp, err := h.userService.Login(ctx, loginReq)
	if err != nil {
		if strings.Contains(err.Error(), "invalid credentials") || strings.Contains(err.Error(), "user not found") {
			return &antifraud_v1.APIV1AuthLoginPostUnauthorized{
				Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
				Message:   "Неверные учетные данные",
				TraceId:   uuid.New(),
				Timestamp: time.Now().UTC(),
				Path:      "/api/v1/auth/login",
				Details:   antifraud_v1.OptApiErrorDetails{},
			}, nil
		}

		if strings.Contains(err.Error(), "deactivated") || strings.Contains(err.Error(), "USER_INACTIVE") {
			return &antifraud_v1.APIV1AuthLoginPostLocked{
				Code:      antifraud_v1.ErrorCodeUSERINACTIVE,
				Message:   "Пользователь деактивирован",
				TraceId:   uuid.New(),
				Timestamp: time.Now().UTC(),
				Path:      "/api/v1/auth/login",
				Details:   antifraud_v1.OptApiErrorDetails{},
			}, nil
		}

		// Для любых других ошибок также возвращаем 401 чтобы избежать 500
		return &antifraud_v1.APIV1AuthLoginPostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Ошибка аутентификации",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/auth/login",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	apiResp := antifraud_v1.AuthResponse{
		AccessToken: authResp.Token,
		ExpiresIn:   3600,
		User:        convertUserToAPI(&authResp.User),
	}

	return &apiResp, nil
}

// validateLoginRequest валидирует поля входа
// Упрощенная валидация по сравнению с регистрацией
func (h *handlerAdapter) validateLoginRequest(req *antifraud_v1.LoginRequest) error {
	if req.Email == "" {
		return fmt.Errorf("email обязателен")
	}

	if len(req.Email) > 254 {
		return fmt.Errorf("email слишком длинный (максимум 254 символа)")
	}

	// Простая проверка email
	if !strings.Contains(req.Email, "@") {
		return fmt.Errorf("неверный формат email")
	}

	if req.Password == "" {
		return fmt.Errorf("пароль обязателен")
	}

	if len(req.Password) > 72 {
		return fmt.Errorf("пароль слишком длинный (максимум 72 символа)")
	}

	return nil
}

func (h *handlerAdapter) APIV1AuthRegisterPost(ctx context.Context, req *antifraud_v1.RegisterRequest) (antifraud_v1.APIV1AuthRegisterPostRes, error) {
	// Валидация полей - критично для безопасности из прошлого проекта с банком
	if err := h.validateRegisterRequest(req); err != nil {
		return &antifraud_v1.APIV1AuthRegisterPostBadRequest{
			Code:      antifraud_v1.ErrorCodeVALIDATIONFAILED,
			Message:   err.Error(),
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/auth/register",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	registerReq := model.RegisterRequest{
		Email:    strings.TrimSpace(strings.ToLower(req.Email)), // нормализация email
		Password: req.Password,
		FullName: strings.TrimSpace(req.FullName),
	}

	authResp, err := h.userService.Register(ctx, registerReq)
	if err != nil {
		// Разные типы ошибок для безопасности - не раскрываем существование email
		if strings.Contains(err.Error(), "already exists") {
			return &antifraud_v1.APIV1AuthRegisterPostConflict{
				Code:      antifraud_v1.ErrorCodeEMAILALREADYEXISTS,
				Message:   "Пользователь с таким email уже существует",
				TraceId:   uuid.New(),
				Timestamp: time.Now().UTC(),
				Path:      "/api/v1/auth/register",
				Details:   antifraud_v1.OptApiErrorDetails{},
			}, nil
		}

		// Другие ошибки считаем внутренними
		return nil, fmt.Errorf("registration failed: %w", err)
	}

	apiResp := antifraud_v1.AuthResponse{
		AccessToken: authResp.Token,
		ExpiresIn:   3600,
		User:        convertUserToAPI(&authResp.User),
	}

	return &apiResp, nil
}

// validateRegisterRequest валидирует поля регистрации
// Добавил строгую проверку из опыта финтех проектов
func (h *handlerAdapter) validateRegisterRequest(req *antifraud_v1.RegisterRequest) error {
	if req.Email == "" {
		return fmt.Errorf("email обязателен")
	}

	if len(req.Email) > 254 {
		return fmt.Errorf("email слишком длинный (максимум 254 символа)")
	}

	// Простая проверка email - для production нужна сложнее
	if !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, ".") {
		return fmt.Errorf("неверный формат email")
	}

	if req.Password == "" {
		return fmt.Errorf("пароль обязателен")
	}

	if len(req.Password) < 8 {
		return fmt.Errorf("пароль должен содержать минимум 8 символов")
	}

	if len(req.Password) > 72 {
		return fmt.Errorf("пароль слишком длинный (максимум 72 символа)")
	}

	// Проверка на цифры и буквы - базовая защита
	hasDigit := false
	hasLetter := false
	for _, char := range req.Password {
		if char >= '0' && char <= '9' {
			hasDigit = true
		}
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			hasLetter = true
		}
	}

	if !hasDigit || !hasLetter {
		return fmt.Errorf("пароль должен содержать минимум одну цифру и одну букву")
	}

	if req.FullName == "" {
		return fmt.Errorf("имя обязательно")
	}

	if len(req.FullName) < 2 {
		return fmt.Errorf("имя слишком короткое (минимум 2 символа)")
	}

	if len(req.FullName) > 200 {
		return fmt.Errorf("имя слишком длинное (максимум 200 символов)")
	}

	return nil
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
	if ctx == nil {
		return &antifraud_v1.APIV1FraudRulesPostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Context is nil",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/fraud-rules",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}
	
	userRole, ok := ctx.Value(middleware.ContextRoleKey).(string)
	if !ok || userRole != "ADMIN" {
		return &antifraud_v1.APIV1FraudRulesPostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Access denied: admin rights required",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/fraud-rules",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	createReq := model.FraudRuleCreateRequest{
		Name:        req.Name,
		Description: req.Description.Value,
		DSL:         req.DslExpression,
	}

	if req.Priority.Set {
		createReq.Priority = &req.Priority.Value
	}

	rule, err := h.fraudRuleService.Create(ctx, createReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create fraud rule: %w", err)
	}

	apiRule := convertFraudRuleToAPI(*rule)
	return &apiRule, nil
}

func (h *handlerAdapter) APIV1FraudRulesValidatePost(ctx context.Context, req *antifraud_v1.DslValidateRequest) (antifraud_v1.APIV1FraudRulesValidatePostRes, error) {
	if ctx == nil {
		return &antifraud_v1.APIV1FraudRulesValidatePostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Context is nil",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/fraud-rules/validate",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}
	
	userRole, ok := ctx.Value(middleware.ContextRoleKey).(string)
	if !ok || userRole != "ADMIN" {
		return &antifraud_v1.APIV1FraudRulesValidatePostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Access denied: admin rights required",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/fraud-rules/validate",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	validation, err := h.fraudRuleService.ValidateDSL(ctx, req.DslExpression)
	if err != nil {
		return nil, fmt.Errorf("failed to validate DSL: %w", err)
	}

	return &antifraud_v1.DslValidateResponse{
		IsValid: validation.IsValid,
		Errors:  []antifraud_v1.DslError{}, // TODO: конвертация ошибок - нужно парсить позиции токенов для UI
	}, nil
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
	if ctx == nil {
		return &antifraud_v1.APIV1UsersGetUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Context is nil",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}
	
	userRole, ok := ctx.Value(middleware.ContextRoleKey).(string)
	if !ok || userRole != "ADMIN" {
		return &antifraud_v1.APIV1UsersGetUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Access denied: admin rights required",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	page := 0
	size := 20

	if params.Page.Set {
		page = params.Page.Value
	}
	if params.Size.Set {
		size = params.Size.Value
	}

	users, total, err := h.userService.GetAll(ctx, page, size)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	apiUsers := make([]antifraud_v1.User, len(users))
	for i, user := range users {
		apiUsers[i] = convertUserToAPI(user)
	}

	return &antifraud_v1.PagedUsers{
		Items: apiUsers,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}

func (h *handlerAdapter) APIV1UsersIDDelete(ctx context.Context, params antifraud_v1.APIV1UsersIDDeleteParams) (antifraud_v1.APIV1UsersIDDeleteRes, error) {
	if ctx == nil {
		return &antifraud_v1.APIV1UsersIDDeleteUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Context is nil",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/" + params.ID.String(),
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}
	
	userRole, ok := ctx.Value(middleware.ContextRoleKey).(string)
	if !ok || userRole != "ADMIN" {
		return &antifraud_v1.APIV1UsersIDDeleteUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Access denied: admin rights required",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/" + params.ID.String(),
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	if err := h.userService.SoftDelete(ctx, params.ID.String()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return &antifraud_v1.APIV1UsersIDDeleteNotFound{
				Code:      antifraud_v1.ErrorCodeNOTFOUND,
				Message:   "User not found",
				TraceId:   uuid.New(),
				Timestamp: time.Now().UTC(),
				Path:      "/api/v1/users/" + params.ID.String(),
				Details:   antifraud_v1.OptApiErrorDetails{},
			}, nil
		}
		
		return &antifraud_v1.APIV1UsersIDDeleteForbidden{
			Code:      antifraud_v1.ErrorCodeVALIDATIONFAILED,
			Message:   "Failed to delete user: " + err.Error(),
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/" + params.ID.String(),
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	return &antifraud_v1.APIV1UsersIDDeleteNoContent{}, nil
}

func (h *handlerAdapter) APIV1UsersIDGet(ctx context.Context, params antifraud_v1.APIV1UsersIDGetParams) (antifraud_v1.APIV1UsersIDGetRes, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}
	
	userRole, ok := ctx.Value(middleware.ContextRoleKey).(string)
	if !ok {
		return &antifraud_v1.APIV1UsersIDGetUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Не удалось определить роль пользователя",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/" + params.ID.String(),
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	userID, ok := ctx.Value(middleware.ContextUserIDKey).(string)
	if !ok {
		return &antifraud_v1.APIV1UsersIDGetUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Не удалось определить ID пользователя",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/" + params.ID.String(),
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	// Проверка прав доступа: ADMIN может смотреть любой профиль, USER только свой
	if userRole != "ADMIN" && userID != params.ID.String() {
		return &antifraud_v1.APIV1UsersIDGetForbidden{
			Code:      antifraud_v1.ErrorCodeFORBIDDEN,
			Message:   "Доступ запрещен: можно просматривать только свой профиль",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/" + params.ID.String(),
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	user, err := h.userService.GetByID(ctx, params.ID.String())
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return &antifraud_v1.APIV1UsersIDGetNotFound{
				Code:      antifraud_v1.ErrorCodeNOTFOUND,
				Message:   "Пользователь не найден",
				TraceId:   uuid.New(),
				Timestamp: time.Now().UTC(),
				Path:      "/api/v1/users/" + params.ID.String(),
				Details:   antifraud_v1.OptApiErrorDetails{},
			}, nil
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	apiUser := convertUserToAPI(user)
	return &apiUser, nil
}

func (h *handlerAdapter) APIV1UsersIDPut(ctx context.Context, req *antifraud_v1.UserUpdateRequest, params antifraud_v1.APIV1UsersIDPutParams) (antifraud_v1.APIV1UsersIDPutRes, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}
	
	userRole, ok := ctx.Value(middleware.ContextRoleKey).(string)
	if !ok {
		return &antifraud_v1.APIV1UsersIDPutUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Не удалось определить роль пользователя",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/" + params.ID.String(),
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	userID, ok := ctx.Value(middleware.ContextUserIDKey).(string)
	if !ok {
		return &antifraud_v1.APIV1UsersIDPutUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Не удалось определить ID пользователя",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/" + params.ID.String(),
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	// Проверка прав доступа: ADMIN может обновлять любой профиль, USER только свой
	if userRole != "ADMIN" && userID != params.ID.String() {
		return &antifraud_v1.APIV1UsersIDPutForbidden{
			Code:      antifraud_v1.ErrorCodeFORBIDDEN,
			Message:   "Доступ запрещен: можно обновлять только свой профиль",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/" + params.ID.String(),
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	// Дополнительные ограничения для USER
	if userRole != "ADMIN" {
		// USER не может менять роль
		if req.Role.Set {
			return &antifraud_v1.APIV1UsersIDPutForbidden{
				Code:      antifraud_v1.ErrorCodeFORBIDDEN,
				Message:   "USER не может изменять роль",
				TraceId:   uuid.New(),
				Timestamp: time.Now().UTC(),
				Path:      "/api/v1/users/" + params.ID.String(),
				Details:   antifraud_v1.OptApiErrorDetails{},
			}, nil
		}

		// USER не может менять статус активности
		if req.IsActive.Set {
			return &antifraud_v1.APIV1UsersIDPutForbidden{
				Code:      antifraud_v1.ErrorCodeFORBIDDEN,
				Message:   "USER не может изменять статус активности",
				TraceId:   uuid.New(),
				Timestamp: time.Now().UTC(),
				Path:      "/api/v1/users/" + params.ID.String(),
				Details:   antifraud_v1.OptApiErrorDetails{},
			}, nil
		}
	}

	updateReq := model.UserUpdateRequest{
		FullName: &req.FullName,
	}

	if req.Region.Null {
		updateReq.Region = nil
	} else {
		region := req.Region.Value
		updateReq.Region = &region
	}

	if req.Gender.Null {
		updateReq.Gender = nil
	} else {
		gender := model.Gender(req.Gender.Value)
		updateReq.Gender = &gender
	}

	if req.Age.Null {
		updateReq.Age = nil
	} else {
		age := req.Age.Value
		updateReq.Age = &age
	}

	if req.MaritalStatus.Null {
		updateReq.MaritalStatus = nil
	} else {
		maritalStatus := model.MaritalStatus(req.MaritalStatus.Value)
		updateReq.MaritalStatus = &maritalStatus
	}

	// Только ADMIN может менять роль и статус активности
	if userRole == "ADMIN" {
		if req.Role.Set {
			role := model.UserRole(req.Role.Value)
			updateReq.Role = &role
		}
		if req.IsActive.Set {
			isActive := req.IsActive.Value
			updateReq.IsActive = &isActive
		}
	}

	var user *model.User
	var err error

	if userRole == "ADMIN" {
		user, err = h.userService.UpdateByAdmin(ctx, params.ID.String(), updateReq)
	} else {
		user, err = h.userService.UpdateMe(ctx, params.ID.String(), updateReq)
	}

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return &antifraud_v1.APIV1UsersIDPutNotFound{
				Code:      antifraud_v1.ErrorCodeNOTFOUND,
				Message:   "Пользователь не найден",
				TraceId:   uuid.New(),
				Timestamp: time.Now().UTC(),
				Path:      "/api/v1/users/" + params.ID.String(),
				Details:   antifraud_v1.OptApiErrorDetails{},
			}, nil
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	apiUser := convertUserToAPI(user)
	return &apiUser, nil
}

func (h *handlerAdapter) APIV1UsersMeGet(ctx context.Context) (antifraud_v1.APIV1UsersMeGetRes, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}
	
	userID, ok := ctx.Value(middleware.ContextUserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("user ID not found in context")
	}

	user, err := h.userService.GetMe(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}

	apiUser := convertUserToAPI(user)
	return &apiUser, nil
}

func (h *handlerAdapter) APIV1UsersMePut(ctx context.Context, req *antifraud_v1.UserUpdateRequest) (antifraud_v1.APIV1UsersMePutRes, error) {
	if ctx == nil {
		return &antifraud_v1.APIV1UsersMePutUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Context is nil",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/me",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}
	
	userID, ok := ctx.Value(middleware.ContextUserIDKey).(string)
	if !ok {
		return &antifraud_v1.APIV1UsersMePutUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Не удалось определить ID пользователя",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/me",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	// USER не может менять роль - проверяем на уровне API
	if req.Role.Set {
		return &antifraud_v1.APIV1UsersMePutForbidden{
			Code:      antifraud_v1.ErrorCodeFORBIDDEN,
			Message:   "USER не может изменять роль",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/me",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	// USER не может менять статус активности - проверяем на уровне API
	if req.IsActive.Set {
		return &antifraud_v1.APIV1UsersMePutForbidden{
			Code:      antifraud_v1.ErrorCodeFORBIDDEN,
			Message:   "USER не может изменять статус активности",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/me",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	updateReq := model.UserUpdateRequest{
		FullName: &req.FullName,
	}

	if req.Region.Null {
		updateReq.Region = nil
	} else {
		region := req.Region.Value
		updateReq.Region = &region
	}
	if req.Gender.Null {
		updateReq.Gender = nil
	} else {
		gender := model.Gender(req.Gender.Value)
		updateReq.Gender = &gender
	}
	if req.Age.Null {
		updateReq.Age = nil
	} else {
		age := req.Age.Value
		updateReq.Age = &age
	}
	if req.MaritalStatus.Null {
		updateReq.MaritalStatus = nil
	} else {
		maritalStatus := model.MaritalStatus(req.MaritalStatus.Value)
		updateReq.MaritalStatus = &maritalStatus
	}

	user, err := h.userService.UpdateMe(ctx, userID, updateReq)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			// Для UsersMePut нет NotFound типа, используем общую ошибку
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to update user profile: %w", err)
	}

	apiUser := convertUserToAPI(user)
	return &apiUser, nil
}

func (h *handlerAdapter) APIV1UsersPost(ctx context.Context, req *antifraud_v1.UserCreateRequest) (antifraud_v1.APIV1UsersPostRes, error) {
	if ctx == nil {
		return &antifraud_v1.APIV1UsersPostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Context is nil",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}
	
	userRole, ok := ctx.Value(middleware.ContextRoleKey).(string)
	if !ok || userRole != "ADMIN" {
		return &antifraud_v1.APIV1UsersPostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Access denied: admin rights required",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	createReq := model.UserCreateRequest{
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
		Role:     model.UserRole(req.Role),
	}

	if req.Region.Set {
		region := req.Region.Value
		createReq.Region = &region
	}
	if req.Gender.Set {
		gender := model.Gender(req.Gender.Value)
		createReq.Gender = &gender
	}
	if req.Age.Set {
		age := req.Age.Value
		createReq.Age = &age
	}
	if req.MaritalStatus.Set {
		maritalStatus := model.MaritalStatus(req.MaritalStatus.Value)
		createReq.MaritalStatus = &maritalStatus
	}

	user, err := h.userService.CreateByAdmin(ctx, createReq)
	if err != nil {
		if strings.Contains(err.Error(), "email already exists") || strings.Contains(err.Error(), "duplicate") {
			return &antifraud_v1.APIV1UsersPostConflict{
				Code:      antifraud_v1.ErrorCodeEMAILALREADYEXISTS,
				Message:   "User with this email already exists",
				TraceId:   uuid.New(),
				Timestamp: time.Now().UTC(),
				Path:      "/api/v1/users",
				Details:   antifraud_v1.OptApiErrorDetails{},
			}, nil
		}
		
		return &antifraud_v1.APIV1UsersPostForbidden{
			Code:      antifraud_v1.ErrorCodeVALIDATIONFAILED,
			Message:   "Failed to create user: " + err.Error(),
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	apiUser := convertUserToAPI(user)
	return &apiUser, nil
}

func convertUserToAPI(user *model.User) antifraud_v1.User {
	apiUser := antifraud_v1.User{
		Email:     user.Email,
		FullName:  user.FullName,
		Role:      antifraud_v1.UserRole(user.Role),
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	id, err := uuid.Parse(user.ID)
	if err == nil {
		apiUser.ID = id
	}

	if user.Region != nil {
		apiUser.Region = antifraud_v1.OptString{
			Value: *user.Region,
			Set:   true,
		}
	}
	if user.Gender != nil {
		apiUser.Gender = antifraud_v1.OptGender{
			Value: antifraud_v1.Gender(*user.Gender),
			Set:   true,
		}
	}
	if user.Age != nil {
		apiUser.Age = antifraud_v1.OptInt{
			Value: *user.Age,
			Set:   true,
		}
	}
	if user.MaritalStatus != nil {
		apiUser.MaritalStatus = antifraud_v1.OptMaritalStatus{
			Value: antifraud_v1.MaritalStatus(*user.MaritalStatus),
			Set:   true,
		}
	}

	return apiUser
}

func convertFraudRuleToAPI(rule model.FraudRule) antifraud_v1.FraudRule {
	id, err := uuid.Parse(rule.ID)
	if err != nil {
		id = uuid.New()
	}

	return antifraud_v1.FraudRule{
		ID:            id,
		Name:          rule.Name,
		Description:   antifraud_v1.OptString{Value: rule.Description, Set: rule.Description != ""},
		DslExpression: rule.DSL,
		Enabled:       rule.IsActive,
		Priority:      rule.Priority,
		CreatedAt:     rule.CreatedAt,
		UpdatedAt:     rule.UpdatedAt,
	}
}
