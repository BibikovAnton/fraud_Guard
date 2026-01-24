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

func (h *handlerAdapter) APIV1AuthRegisterPost(ctx context.Context, req *antifraud_v1.RegisterRequest) (antifraud_v1.APIV1AuthRegisterPostRes, error) {
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
		Email:    strings.TrimSpace(strings.ToLower(req.Email)),
		Password: req.Password,
		FullName: strings.TrimSpace(req.FullName),
	}

	authResp, err := h.userService.Register(ctx, registerReq)
	if err != nil {
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

		return nil, fmt.Errorf("registration failed: %w", err)
	}

	apiResp := antifraud_v1.AuthResponse{
		AccessToken: authResp.AccessToken,
		ExpiresIn:   authResp.ExpiresIn,
		User:        convertor.ConvertUserToAPI(&authResp.User),
	}

	return &apiResp, nil
}

func (h *handlerAdapter) validateRegisterRequest(req *antifraud_v1.RegisterRequest) error {
	if req.Email == "" {
		return fmt.Errorf("email обязателен")
	}

	if len(req.Email) > 254 {
		return fmt.Errorf("email слишком длинный (максимум 254 символа)")
	}

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
