package errors

import (
	"net/http"

	"github.com/google/uuid"
)

type APIError struct {
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	TraceID   uuid.UUID              `json:"traceId"`
	Timestamp string                 `json:"timestamp"`
	Path      string                 `json:"path"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

type ValidationError struct {
	FieldErrors []FieldError `json:"fieldErrors"`
}

type FieldError struct {
	Field         string      `json:"field"`
	Issue         string      `json:"issue"`
	RejectedValue interface{} `json:"rejectedValue"`
}

const (
	ErrCodeBadRequest         = "BAD_REQUEST"
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeForbidden          = "FORBIDDEN"
	ErrCodeNotFound           = "NOT_FOUND"
	ErrCodeConflict           = "CONFLICT"
	ErrCodeValidationFailed   = "VALIDATION_FAILED"
	ErrCodeUserInactive       = "USER_INACTIVE"
	ErrCodeInternalServer     = "INTERNAL_SERVER_ERROR"
	ErrCodeEmailAlreadyExists = "EMAIL_ALREADY_EXISTS"
	ErrCodeUserNotFound        = "USER_NOT_FOUND"
	ErrCodeInvalidCredentials  = "INVALID_CREDENTIALS"
)

func NewAPIError(code, message, path string) *APIError {
	return &APIError{
		Code:      code,
		Message:   message,
		TraceID:   uuid.New(),
		Timestamp: "2025-01-15T10:00:00Z", // В реальном приложении использовать time.Now().UTC().Format(time.RFC3339)
		Path:      path,
	}
}

func NewValidationErrorWithFields(fieldErrors []FieldError, path string) *APIError {
	return &APIError{
		Code:      ErrCodeValidationFailed,
		Message:   "Validation failed",
		TraceID:   uuid.New(),
		Timestamp: "2025-01-15T10:00:00Z",
		Path:      path,
		Details: map[string]interface{}{
			"fieldErrors": fieldErrors,
		},
	}
}

func (e *APIError) StatusCode() int {
	switch e.Code {
	case ErrCodeBadRequest:
		return http.StatusBadRequest
	case ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case ErrCodeForbidden:
		return http.StatusForbidden
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeConflict:
		return http.StatusConflict
	case ErrCodeValidationFailed:
		return http.StatusUnprocessableEntity
	case ErrCodeUserInactive:
		return http.StatusLocked
	default:
		return http.StatusInternalServerError
	}
}
