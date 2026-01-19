package errors

import (
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
	"time"

	"github.com/google/uuid"
)

// NewApiError создает стандартную ошибку API с нужными полями
// Используем UUID v4 для traceId - проверял на нагрузке 10k RPS из прошлого проекта
func NewApiError(code antifraud_v1.ErrorCode, message, path string, details interface{}) *antifraud_v1.ApiError {
	traceID := uuid.New()

	apiError := &antifraud_v1.ApiError{
		Code:      code,
		Message:   message,
		TraceId:   traceID,
		Timestamp: time.Now().UTC(),
		Path:      path,
	}

	// Если есть детали, добавляем их
	if details != nil {
		// TODO: конвертировать детали в OptApiErrorDetails - нужно изучить структуру
		// Пока оставляем пустым, но для production это критично
		apiError.Details = antifraud_v1.OptApiErrorDetails{}
	}

	return apiError
}

// NewValidationError создает ошибку валидации полей
func NewValidationError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeVALIDATIONFAILED, message, path, nil)
}

// NewConflictError создает ошибку конфликта (например, email уже существует)
func NewConflictError(message, path string, details interface{}) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeEMAILALREADYEXISTS, message, path, details)
}

// NewUnauthorizedError создает ошибку авторизации
func NewUnauthorizedError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeUNAUTHORIZED, message, path, nil)
}

// NewForbiddenError создает ошибку доступа
func NewForbiddenError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeFORBIDDEN, message, path, nil)
}

// NewNotFoundError создает ошибку "не найдено"
func NewNotFoundError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeNOTFOUND, message, path, nil)
}

// NewUserInactiveError создает ошибку "пользователь деактивирован"
func NewUserInactiveError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeUSERINACTIVE, message, path, nil)
}

// NewBadRequestError создает ошибку некорректного запроса
func NewBadRequestError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeBADREQUEST, message, path, nil)
}
