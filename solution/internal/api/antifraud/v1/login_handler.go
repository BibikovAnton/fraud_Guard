package v1

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"solution/internal/api/antifraud/v1/convertor"
	"solution/internal/model"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
	"strings"
	"time"
)

func (h *handlerAdapter) APIV1AuthLoginPost(ctx context.Context, req *antifraud_v1.LoginRequest) (antifraud_v1.APIV1AuthLoginPostRes, error) {
	if req == nil || ctx == nil {
		return &antifraud_v1.APIV1AuthLoginPostBadRequest{
			Code:      antifraud_v1.ErrorCodeBADREQUEST,
			Message:   "Invalid request or context",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/auth/login",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

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
		Email:    strings.TrimSpace(strings.ToLower(req.Email)),
		Password: req.Password,
	}

	userData, err := h.userService.Login(ctx, loginReq)
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

		return &antifraud_v1.APIV1AuthLoginPostUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Ошибка аутентификации",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/auth/login",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	apiResponse := antifraud_v1.AuthResponse{
		AccessToken: userData.AccessToken,
		ExpiresIn:   userData.ExpiresIn,
		User:        convertor.ConvertUserToAPI(&userData.User),
	}

	return &apiResponse, nil
}

func (h *handlerAdapter) validateLoginRequest(req *antifraud_v1.LoginRequest) error {
	if req.Email == "" {
		return fmt.Errorf("email обязателен")
	}

	if len(req.Email) > 254 {
		return fmt.Errorf("email слишком длинный (максимум 254 символа)")
	}

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
