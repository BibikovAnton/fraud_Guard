package errors

import (
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
	"time"

	"github.com/google/uuid"
)

func NewApiError(code antifraud_v1.ErrorCode, message, path string, details interface{}) *antifraud_v1.ApiError {
	traceID := uuid.New()

	apiError := &antifraud_v1.ApiError{
		Code:      code,
		Message:   message,
		TraceId:   traceID,
		Timestamp: time.Now().UTC(),
		Path:      path,
	}

	if details != nil {
		apiError.Details = antifraud_v1.OptApiErrorDetails{}
	}

	return apiError
}

func NewValidationError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeVALIDATIONFAILED, message, path, nil)
}

func NewConflictError(message, path string, details interface{}) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeEMAILALREADYEXISTS, message, path, details)
}

func NewUnauthorizedError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeUNAUTHORIZED, message, path, nil)
}

func NewForbiddenError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeFORBIDDEN, message, path, nil)
}

func NewNotFoundError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeNOTFOUND, message, path, nil)
}

func NewUserInactiveError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeUSERINACTIVE, message, path, nil)
}

func NewBadRequestError(message, path string) *antifraud_v1.ApiError {
	return NewApiError(antifraud_v1.ErrorCodeBADREQUEST, message, path, nil)
}
