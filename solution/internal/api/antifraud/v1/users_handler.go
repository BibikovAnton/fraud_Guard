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

	userRole, ok := ctx.Value(ContextRoleKey).(string)
	if !ok {
		return &antifraud_v1.APIV1UsersGetUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Access denied: authentication required",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	if userRole != "ADMIN" {
		return &antifraud_v1.APIV1UsersGetForbidden{
			Code:      antifraud_v1.ErrorCodeFORBIDDEN,
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
		apiUsers[i] = convertor.ConvertUserToAPI(user)
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

	userRole, ok := ctx.Value(ContextRoleKey).(string)
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

	userRole, ok := ctx.Value(ContextRoleKey).(string)
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

	userID, ok := ctx.Value(ContextUserIDKey).(string)
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

	var user *model.User
	var err error

	if userRole == "ADMIN" {
		user, err = h.userService.GetByIDIncludingInactive(ctx, params.ID.String())
	} else {
		user, err = h.userService.GetByID(ctx, params.ID.String())
	}

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

	apiUser := convertor.ConvertUserToAPI(user)
	return &apiUser, nil
}

func (h *handlerAdapter) APIV1UsersIDPut(ctx context.Context, req *antifraud_v1.UserUpdateRequest, params antifraud_v1.APIV1UsersIDPutParams) (antifraud_v1.APIV1UsersIDPutRes, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}

	userRole, ok := ctx.Value(ContextRoleKey).(string)
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

	userID, ok := ctx.Value(ContextUserIDKey).(string)
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

	if userRole != "ADMIN" {
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

	apiUser := convertor.ConvertUserToAPI(user)
	return &apiUser, nil
}

func (h *handlerAdapter) APIV1UsersMeGet(ctx context.Context) (antifraud_v1.APIV1UsersMeGetRes, error) {
	if ctx == nil {
		return &antifraud_v1.APIV1UsersMeGetUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Context is required",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/me",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	userID, ok := ctx.Value(ContextUserIDKey).(string)
	if !ok {
		return &antifraud_v1.APIV1UsersMeGetUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "User ID not found in context",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/me",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	user, err := h.userService.GetMe(ctx, userID)
	if err != nil {
		return &antifraud_v1.APIV1UsersMeGetUnauthorized{
			Code:      antifraud_v1.ErrorCodeUNAUTHORIZED,
			Message:   "Failed to get user profile",
			TraceId:   uuid.New(),
			Timestamp: time.Now().UTC(),
			Path:      "/api/v1/users/me",
			Details:   antifraud_v1.OptApiErrorDetails{},
		}, nil
	}

	apiUser := convertor.ConvertUserToAPI(user)
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

	userID, ok := ctx.Value(ContextUserIDKey).(string)
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
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to update user profile: %w", err)
	}

	apiUser := convertor.ConvertUserToAPI(user)
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

	userRole, ok := ctx.Value(ContextRoleKey).(string)
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

	apiUser := convertor.ConvertUserToAPI(user)
	return &apiUser, nil
}
