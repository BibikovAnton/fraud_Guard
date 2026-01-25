package v1

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/netip"
	"solution/internal/model"
	"solution/internal/service"
	"strings"
	"time"
)

type TransactionHandler struct {
	userService        service.UserService
	transactionService service.TransactionService
	fraudRuleService   service.FraudRuleService
}

func NewTransactionHandler(userService service.UserService, transactionService service.TransactionService, fraudRuleService service.FraudRuleService) *TransactionHandler {
	return &TransactionHandler{
		userService:        userService,
		transactionService: transactionService,
		fraudRuleService:   fraudRuleService,
	}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	userID, err := h.extractUserIDFromToken(r)
	if err != nil {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: authentication required")
		return
	}

	userRole, err := h.extractUserRoleFromToken(r)
	if err != nil {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: authentication required")
		return
	}

	var rawRequest map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&rawRequest); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid JSON")
		return
	}

	transactionReq, err := h.validateAndConvertTransaction(rawRequest, userID, userRole)
	if err != nil {
		// Check if it's a validation error that should return 422 with fieldErrors
		fieldErrors := h.extractFieldErrors(err.Error(), rawRequest)
		if len(fieldErrors) > 0 {
			writeValidationErrorResponse(w, "/api/v1/transactions", fieldErrors)
		} else {
			if strings.Contains(err.Error(), "failed to get user by ID") || strings.Contains(err.Error(), "no rows in result set") {
				writeErrorResponse(w, http.StatusNotFound, "USER_NOT_FOUND", "User not found")
			} else if strings.Contains(err.Error(), "user is deactivated") {
				writeErrorResponse(w, http.StatusForbidden, "USER_INACTIVE", "User is deactivated")
			} else if strings.Contains(err.Error(), "invalid userId format") {
				writeErrorResponse(w, http.StatusUnprocessableEntity, "VALIDATION_FAILED", err.Error())
			} else {
				writeErrorResponse(w, http.StatusUnprocessableEntity, "VALIDATION_FAILED", err.Error())
			}
		}
		return
	}

	decision, err := h.transactionService.Create(ctx, *transactionReq)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "VALIDATION_FAILED", err.Error())
		return
	}

	response := h.convertDecisionToResponse(decision)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *TransactionHandler) CreateBatchTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	userID, err := h.extractUserIDFromToken(r)
	if err != nil {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: authentication required")
		return
	}

	userRole, err := h.extractUserRoleFromToken(r)
	if err != nil {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: authentication required")
		return
	}

	var rawRequest map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&rawRequest); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid JSON")
		return
	}

	items, ok := rawRequest["items"].([]interface{})
	if !ok || len(items) == 0 {
		writeErrorResponse(w, http.StatusBadRequest, "VALIDATION_ERROR", "Batch request cannot be empty")
		return
	}

	results := make([]map[string]interface{}, len(items))
	hasErrors := false

	for i, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			hasErrors = true
			results[i] = map[string]interface{}{
				"index": i,
				"error": map[string]interface{}{
					"code":    "VALIDATION_FAILED",
					"message": "Invalid item format",
				},
			}
			continue
		}

		transactionReq, err := h.validateAndConvertTransaction(itemMap, userID, userRole)
		if err != nil {
			hasErrors = true
			errorCode := "VALIDATION_FAILED"
			errorMessage := err.Error()
			
			if strings.Contains(err.Error(), "failed to get user by ID") {
				errorCode = "NOT_FOUND"
				errorMessage = "User not found"
			} else if strings.Contains(err.Error(), "user is deactivated") {
				errorCode = "FORBIDDEN"
				errorMessage = "User is deactivated"
			}
			
			results[i] = map[string]interface{}{
				"index": i,
				"error": map[string]interface{}{
					"code":    errorCode,
					"message": errorMessage,
				},
			}
			continue
		}

		decision, err := h.transactionService.Create(ctx, *transactionReq)
		if err != nil {
			hasErrors = true
			results[i] = map[string]interface{}{
				"index": i,
				"error": map[string]interface{}{
					"code":    "VALIDATION_FAILED",
					"message": err.Error(),
				},
			}
			continue
		}

		results[i] = map[string]interface{}{
			"index": i,
			"decision": h.convertDecisionToResponse(decision),
		}
	}

	statusCode := http.StatusCreated
	if hasErrors {
		statusCode = http.StatusMultiStatus
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": results,
	})
}

func (h *TransactionHandler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	_, err := h.extractUserIDFromToken(r)
	if err != nil {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: authentication required")
		return
	}

	_, err = h.extractUserRoleFromToken(r)
	if err != nil {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: authentication required")
		return
	}

	transactionID := strings.TrimPrefix(r.URL.Path, "/api/v1/transactions/")
	if transactionID == "" {
		writeErrorResponse(w, http.StatusNotFound, "NOT_FOUND", "Transaction not found")
		return
	}

	decision, err := h.transactionService.GetByID(ctx, transactionID)
	if err != nil {
		writeErrorResponse(w, http.StatusNotFound, "NOT_FOUND", "Transaction not found")
		return
	}

	response := h.convertDecisionToResponse(decision)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	userID, err := h.extractUserIDFromToken(r)
	if err != nil {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: authentication required")
		return
	}

	userRole, err := h.extractUserRoleFromToken(r)
	if err != nil {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: authentication required")
		return
	}

	userIDStr := r.URL.Query().Get("userId")
	statusStr := r.URL.Query().Get("status")
	isFraudStr := r.URL.Query().Get("isFraud")

	// USER can only view their own transactions
	if userRole != "ADMIN" {
		if userIDStr != "" && userIDStr != userID {
			writeErrorResponse(w, http.StatusForbidden, "FORBIDDEN", "USER can only view their own transactions")
			return
		}
		// Force filter to current user's ID for non-admin
		userIDStr = userID
	}

	params := service.TransactionListParams{
		Page: 1,
		Size: 10,
	}
	
	if userIDStr != "" {
		params.UserID = &userIDStr
	}
	if statusStr != "" {
		status := model.TransactionStatus(statusStr)
		params.Status = &status
	}
	if isFraudStr != "" {
		isFraud := isFraudStr == "true"
		params.IsFraud = &isFraud
	}

	result, err := h.transactionService.GetList(ctx, params)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to get transactions")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *TransactionHandler) validateAndConvertTransaction(raw map[string]interface{}, userID, userRole string) (*model.TransactionCreateRequest, error) {
	if userRole == "ADMIN" {
		userIdStr, ok := raw["userId"].(string)
		if !ok || userIdStr == "" {
			return nil, fmt.Errorf("userId is required for admin")
		}
		userID = userIdStr
	}

	amountRaw, ok := raw["amount"].(float64)
	if !ok {
		return nil, fmt.Errorf("must be greater > 0")
	}
	if amountRaw <= 0 {
		return nil, fmt.Errorf("must be greater > 0")
	}
	if amountRaw > model.MaxTransactionAmount {
		return nil, fmt.Errorf("must be less than or equal to %.2f", model.MaxTransactionAmount)
	}

	currencyRaw, ok := raw["currency"].(string)
	if !ok || currencyRaw == "" {
		return nil, fmt.Errorf("currency is required")
	}
	validCurrencies := map[model.CurrencyCode]bool{
		model.CurrencyUSD: true,
		model.CurrencyEUR: true,
		model.CurrencyRUB: true,
	}
	if !validCurrencies[model.CurrencyCode(currencyRaw)] {
		return nil, fmt.Errorf("invalid currency code: %s", currencyRaw)
	}

	timestampRaw, ok := raw["timestamp"].(string)
	if !ok || timestampRaw == "" {
		return nil, fmt.Errorf("timestamp is required")
	}
	timestamp, err := time.Parse(time.RFC3339, timestampRaw)
	if err != nil {
		return nil, fmt.Errorf("timestamp is required")
	}
	if timestamp.After(time.Now().Add(5 * time.Minute)) {
		return nil, fmt.Errorf("timestamp cannot be more than 5 minutes in the future")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid userId format")
	}

	req := &model.TransactionCreateRequest{
		UserID:   userUUID,
		Amount:   amountRaw,
		Currency: model.CurrencyCode(currencyRaw),
		Timestamp: timestamp,
	}

	if merchantId, ok := raw["merchantId"].(string); ok {
		req.MerchantID = &merchantId
	}
	if mcc, ok := raw["merchantCategoryCode"].(string); ok {
		mccCode := model.MCCCode(mcc)
		req.MerchantCategoryCode = &mccCode
	}
	if ipStr, ok := raw["ipAddress"].(string); ok {
		if ip, err := netip.ParseAddr(ipStr); err == nil {
			req.IPAddress = &ip
		}
	}
	if device, ok := raw["deviceId"].(string); ok {
		req.DeviceID = &device
	}
	if channel, ok := raw["channel"].(string); ok {
		channelCode := model.TransactionChannel(channel)
		req.Channel = &channelCode
	}

	if locationRaw, ok := raw["location"].(map[string]interface{}); ok {
		location := &model.TransactionLocation{}
		if country, ok := locationRaw["country"].(string); ok {
			if len(country) > 2 {
				return nil, fmt.Errorf("location.country must be at most 2 characters")
			}
			location.Country = country
		} else {
			return nil, fmt.Errorf("location.country is required")
		}

		if lat, ok := locationRaw["latitude"].(float64); ok {
			if lat < -90 || lat > 90 {
				return nil, fmt.Errorf("must be between -90 and 90")
			}
			location.Latitude = &lat
		}
		if lon, ok := locationRaw["longitude"].(float64); ok {
			if lon < -180 || lon > 180 {
				return nil, fmt.Errorf("must be between -180 and 180")
			}
			location.Longitude = &lon
		}

		if (location.Latitude != nil && location.Longitude == nil) ||
			(location.Latitude == nil && location.Longitude != nil) {
			return nil, fmt.Errorf("longitude and latitude must be provided together")
		}

		req.Location = location
	}

	return req, nil
}

func (h *TransactionHandler) convertDecisionToResponse(decision *model.TransactionDecision) map[string]interface{} {
	transaction := map[string]interface{}{
		"id":                   decision.Transaction.ID.String(),
		"userId":               decision.Transaction.UserID.String(),
		"amount":               decision.Transaction.Amount,
		"currency":             decision.Transaction.Currency,
		"status":               decision.Transaction.Status,
		"merchantId":           decision.Transaction.MerchantID,
		"merchantCategoryCode": decision.Transaction.MerchantCategoryCode,
		"timestamp":            decision.Transaction.Timestamp.Format(time.RFC3339),
		"ipAddress":            decision.Transaction.IPAddress,
		"deviceId":             decision.Transaction.DeviceID,
		"channel":              decision.Transaction.Channel,
		"isFraud":              decision.Transaction.IsFraud,
		"location":             decision.Transaction.Location,
		"metadata":             decision.Transaction.Metadata,
		"createdAt":            decision.Transaction.CreatedAt.Format(time.RFC3339),
		"updatedAt":            decision.Transaction.UpdatedAt.Format(time.RFC3339),
	}

	ruleResults := make([]map[string]interface{}, len(decision.RuleResults))
	for i, rule := range decision.RuleResults {
		ruleUUID, _ := uuid.Parse(rule.RuleID)
		ruleResults[i] = map[string]interface{}{
			"ruleId":      ruleUUID.String(),
			"ruleName":    rule.RuleName,
			"priority":    rule.Priority,
			"matched":     rule.Matched,
			"description": rule.Description,
		}
	}

	return map[string]interface{}{
		"transaction": transaction,
		"ruleResults": ruleResults,
	}
}

func (h *TransactionHandler) extractUserIDFromToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no authorization header")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("invalid authorization header format")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	
	userID, err := h.userService.ValidateTokenAndGetUserID(token)
	if err != nil {
		return "", fmt.Errorf("invalid token")
	}

	return userID, nil
}

func (h *TransactionHandler) extractUserRoleFromToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no authorization header")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	
	userRole, err := h.userService.ValidateTokenAndGetUserRole(token)
	if err != nil {
		return "", fmt.Errorf("invalid token")
	}

	return userRole, nil
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, code, message string) {
	writeErrorResponseWithPath(w, statusCode, code, message, "/api/v1/transactions")
}

func writeErrorResponseWithPath(w http.ResponseWriter, statusCode int, code, message, path string) {
	response := map[string]interface{}{
		"code":      code,
		"message":   message,
		"traceId":   uuid.New().String(),
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"path":      path,
	}
	
	// Add details for specific error codes
	if code == "USER_NOT_FOUND" {
		response["details"] = map[string]interface{}{
			"userId": "unknown",
		}
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func (h *TransactionHandler) extractFieldErrors(errMsg string, rawRequest map[string]interface{}) []map[string]interface{} {
	var fieldErrors []map[string]interface{}
	
	// Amount validation errors
	if strings.Contains(errMsg, "must be greater > 0") {
		if val, ok := rawRequest["amount"]; ok {
			fieldErrors = append(fieldErrors, map[string]interface{}{
				"field":        "amount",
				"issue":        "must be greater > 0",
				"rejectedValue": val,
			})
		}
	}
	
	if strings.Contains(errMsg, "must be less than or equal to") {
		if val, ok := rawRequest["amount"]; ok {
			fieldErrors = append(fieldErrors, map[string]interface{}{
				"field":        "amount", 
				"issue":        errMsg,
				"rejectedValue": val,
			})
		}
	}
	
	// Currency validation errors
	if strings.Contains(errMsg, "currency is required") {
		fieldErrors = append(fieldErrors, map[string]interface{}{
			"field":        "currency",
			"issue":        "currency is required",
			"rejectedValue": rawRequest["currency"],
		})
	}
	
	if strings.Contains(errMsg, "invalid currency code") {
		if val, ok := rawRequest["currency"]; ok {
			fieldErrors = append(fieldErrors, map[string]interface{}{
				"field":        "currency",
				"issue":        "invalid currency code",
				"rejectedValue": val,
			})
		}
	}
	
	// Timestamp validation errors
	if strings.Contains(errMsg, "timestamp is required") {
		fieldErrors = append(fieldErrors, map[string]interface{}{
			"field":        "timestamp",
			"issue":        "timestamp is required", 
			"rejectedValue": rawRequest["timestamp"],
		})
	}
	
	// Location validation errors
	if strings.Contains(errMsg, "location.country is required") {
		fieldErrors = append(fieldErrors, map[string]interface{}{
			"field":        "location.country",
			"issue":        "location.country is required",
			"rejectedValue": nil,
		})
	}
	
	if strings.Contains(errMsg, "location.country must be at most 2 characters") {
		var rejectedValue interface{}
		if location, ok := rawRequest["location"].(map[string]interface{}); ok {
			if val, countryOk := location["country"]; countryOk {
				rejectedValue = val
			}
		}
		fieldErrors = append(fieldErrors, map[string]interface{}{
			"field":        "location.country",
			"issue":        "location.country must be at most 2 characters",
			"rejectedValue": rejectedValue,
		})
	}
	
	if strings.Contains(errMsg, "longitude and latitude must be provided together") {
		fieldErrors = append(fieldErrors, map[string]interface{}{
			"field":        "location",
			"issue":        "longitude and latitude must be provided together",
			"rejectedValue": rawRequest["location"],
		})
	}
	
	if strings.Contains(errMsg, "must be between -90 and 90") {
		var rejectedValue interface{}
		if location, ok := rawRequest["location"].(map[string]interface{}); ok {
			if val, latOk := location["latitude"]; latOk {
				rejectedValue = val
			}
		}
		fieldErrors = append(fieldErrors, map[string]interface{}{
			"field":        "location.latitude",
			"issue":        "must be between -90 and 90",
			"rejectedValue": rejectedValue,
		})
	}
	
	if strings.Contains(errMsg, "must be between -180 and 180") {
		var rejectedValue interface{}
		if location, ok := rawRequest["location"].(map[string]interface{}); ok {
			if val, lonOk := location["longitude"]; lonOk {
				rejectedValue = val
			}
		}
		fieldErrors = append(fieldErrors, map[string]interface{}{
			"field":        "location.longitude",
			"issue":        "must be between -180 and 180",
			"rejectedValue": rejectedValue,
		})
	}
	
	// User ID validation errors
	if strings.Contains(errMsg, "userId is required for admin") {
		fieldErrors = append(fieldErrors, map[string]interface{}{
			"field":        "userId",
			"issue":        "userId is required for admin",
			"rejectedValue": rawRequest["userId"],
		})
	}
	
	return fieldErrors
}

func (h *TransactionHandler) ValidateDSL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var req struct {
		DslExpression string `json:"dslExpression"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":      "VALIDATION_ERROR",
			"message":   "Invalid JSON",
			"traceId":   uuid.New().String(),
			"timestamp": time.Now().UTC().Format(time.RFC3339),
			"path":      "/api/v1/fraud-rules/validate",
		})
		return
	}
	
	validateReq := model.DslValidateRequest{
		DslExpression: req.DslExpression,
	}
	
	result := h.fraudRuleService.ValidateDSL(r.Context(), validateReq)
	
	errors := make([]map[string]interface{}, len(result.Errors))
	for i, err := range result.Errors {
		errorObj := map[string]interface{}{
			"code":    err.Code,
			"message": err.Message,
		}
		if err.Position != nil && *err.Position > 0 {
			errorObj["position"] = *err.Position
		}
		if err.Near != nil && *err.Near != "" {
			errorObj["near"] = *err.Near
		}
		errors[i] = errorObj
	}
	
	response := map[string]interface{}{
		"isValid":              result.IsValid,
		"normalizedExpression": result.NormalizedExpression,
		"errors":               errors,
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func writeValidationErrorResponse(w http.ResponseWriter, path string, fieldErrors []map[string]interface{}) {
	response := map[string]interface{}{
		"code":       "VALIDATION_FAILED",
		"message":    "Some fields did not pass validation",
		"traceId":    uuid.New().String(),
		"timestamp":  time.Now().UTC().Format(time.RFC3339),
		"path":       path,
		"fieldErrors": fieldErrors,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(response)
}
