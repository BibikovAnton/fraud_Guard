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

type RawTransactionHandler struct {
	transactionService service.TransactionService
}

func NewRawTransactionHandler(transactionService service.TransactionService) *RawTransactionHandler {
	return &RawTransactionHandler{
		transactionService: transactionService,
	}
}

func (h *RawTransactionHandler) CreateTransactionRaw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	userRole, ok := ctx.Value(ContextRoleKey).(string)
	if !ok {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: authentication required")
		return
	}

	userID, ok := ctx.Value(ContextUserIDKey).(string)
	if !ok {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: user ID not found")
		return
	}

	var rawRequest map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&rawRequest); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid JSON")
		return
	}

	transactionReq, err := h.validateAndConvertTransaction(rawRequest, userID, userRole)
	if err != nil {
		if strings.Contains(err.Error(), "failed to get user by ID") {
			writeErrorResponse(w, http.StatusNotFound, "NOT_FOUND", "User not found")
		} else if strings.Contains(err.Error(), "user is deactivated") {
			writeErrorResponse(w, http.StatusForbidden, "FORBIDDEN", "User is deactivated")
		} else {
			writeErrorResponse(w, http.StatusUnprocessableEntity, "VALIDATION_FAILED", err.Error())
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

func (h *RawTransactionHandler) CreateBatchTransactionRaw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	userRole, ok := ctx.Value(ContextRoleKey).(string)
	if !ok {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: authentication required")
		return
	}

	userID, ok := ctx.Value(ContextUserIDKey).(string)
	if !ok {
		writeErrorResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", "Access denied: user ID not found")
		return
	}

	// Читаем raw batch request
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

func (h *RawTransactionHandler) validateAndConvertTransaction(raw map[string]interface{}, userID, userRole string) (*model.TransactionCreateRequest, error) {
	if userRole == "ADMIN" {
		userIdStr, ok := raw["userId"].(string)
		if !ok || userIdStr == "" {
			return nil, fmt.Errorf("userId is required for admin")
		}
		userID = userIdStr
	}

	amountRaw, ok := raw["amount"].(float64)
	if !ok {
		return nil, fmt.Errorf("amount is required and must be number")
	}
	if amountRaw <= 0 {
		return nil, fmt.Errorf("must be greater > 0")
	}
	if amountRaw > 999999999.99 {
		return nil, fmt.Errorf("must be less than or equal to 999999999.99")
	}

	currencyRaw, ok := raw["currency"].(string)
	if !ok || currencyRaw == "" {
		return nil, fmt.Errorf("currency is required")
	}
	validCurrencies := map[string]bool{"USD": true, "EUR": true, "RUB": true}
	if !validCurrencies[currencyRaw] {
		return nil, fmt.Errorf("invalid currency code: %s", currencyRaw)
	}

	timestampRaw, ok := raw["timestamp"].(string)
	if !ok || timestampRaw == "" {
		return nil, fmt.Errorf("timestamp is required")
	}
	timestamp, err := time.Parse(time.RFC3339, timestampRaw)
	if err != nil {
		return nil, fmt.Errorf("invalid timestamp format")
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
			return nil, fmt.Errorf("location.country is required when location is provided")
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

func (h *RawTransactionHandler) convertDecisionToResponse(decision *model.TransactionDecision) map[string]interface{} {
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

func writeErrorResponse(w http.ResponseWriter, statusCode int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":      code,
		"message":   message,
		"traceId":   uuid.New().String(),
		"timestamp": time.Now().UTC(),
		"path":      "/api/v1/transactions",
	})
}
