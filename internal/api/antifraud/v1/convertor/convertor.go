package convertor

import (
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"net/netip"
	"solution/internal/model"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
)

func ConvertUserToAPI(user *model.User) antifraud_v1.User {
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

func ConvertFraudRuleToAPI(rule model.FraudRule) antifraud_v1.FraudRule {
	id, err := uuid.Parse(rule.ID)
	if err != nil {
		id = uuid.New()
	}

	return antifraud_v1.FraudRule{
		ID:            id,
		Name:          rule.Name,
		Description:   antifraud_v1.OptString{Value: rule.Description, Set: rule.Description != ""},
		DslExpression: rule.DslExpression,
		Enabled:       rule.Enabled,
		Priority:      rule.Priority,
		CreatedAt:     rule.CreatedAt,
		UpdatedAt:     rule.UpdatedAt,
	}
}

func ConvertTransactionCreateRequest(req *antifraud_v1.TransactionCreateRequest, userID string) model.TransactionCreateRequest {
	createReq := model.TransactionCreateRequest{
		UserID:    &req.UserId,
		Amount:    req.Amount,
		Currency:  model.CurrencyCode(req.Currency),
		Timestamp: req.Timestamp,
	}

	if req.MerchantId.Set {
		createReq.MerchantID = &req.MerchantId.Value
	}

	if req.MerchantCategoryCode.Set {
		mcc := model.MCCCode(req.MerchantCategoryCode.Value)
		createReq.MerchantCategoryCode = &mcc
	}

	if req.IpAddress.Set {
		if ip, err := netip.ParseAddr(req.IpAddress.Value); err == nil {
			createReq.IPAddress = &ip
		}
	}

	if req.DeviceId.Set {
		createReq.DeviceID = &req.DeviceId.Value
	}

	if req.Channel.Set {
		channel := model.TransactionChannel(req.Channel.Value)
		createReq.Channel = &channel
	}

	if req.Location.Set {
		if req.Location.Value.Latitude.Set && req.Location.Value.Longitude.Set {
			createReq.Location = &model.TransactionLocation{
				Latitude:  &req.Location.Value.Latitude.Value,
				Longitude: &req.Location.Value.Longitude.Value,
			}
			if req.Location.Value.Country.Set {
				createReq.Location.Country = req.Location.Value.Country.Value
			}
		}
	}

	if req.Metadata.Set {
		metadata := make(model.TransactionMetadata)
		for k, v := range req.Metadata.Value {
			metadata[k] = v
		}
		createReq.Metadata = &metadata
	}

	return createReq
}

func ConvertTransactionToAPI(t *model.Transaction) antifraud_v1.Transaction {
	transaction := antifraud_v1.Transaction{
		ID:        t.ID,
		UserId:    *t.UserID,
		Amount:    t.Amount,
		Currency:  antifraud_v1.CurrencyCode(t.Currency),
		Status:    antifraud_v1.TransactionStatus(t.Status),
		Timestamp: t.Timestamp,
		Channel:   antifraud_v1.OptTransactionChannel{},
		IsFraud:   t.IsFraud,
		CreatedAt: t.CreatedAt,
	}

	if t.MerchantID != nil {
		transaction.MerchantId = antifraud_v1.OptString{Set: true, Value: *t.MerchantID}
	}

	if t.MerchantCategoryCode != nil {
		transaction.MerchantCategoryCode = antifraud_v1.OptMccCode{Set: true, Value: antifraud_v1.MccCode(*t.MerchantCategoryCode)}
	}

	if t.IPAddress != nil {
		transaction.IpAddress = antifraud_v1.OptString{Set: true, Value: t.IPAddress.String()}
	}

	if t.DeviceID != nil {
		transaction.DeviceId = antifraud_v1.OptString{Set: true, Value: *t.DeviceID}
	}

	if t.Channel != nil {
		transaction.Channel = antifraud_v1.OptTransactionChannel{Set: true, Value: antifraud_v1.TransactionChannel(*t.Channel)}
	}

	if t.Location != nil {
		transaction.Location = antifraud_v1.OptTransactionLocation{
			Set: true,
			Value: antifraud_v1.TransactionLocation{
				Latitude:  antifraud_v1.OptFloat64{Set: true, Value: *t.Location.Latitude},
				Longitude: antifraud_v1.OptFloat64{Set: true, Value: *t.Location.Longitude},
			},
		}
		if t.Location.Country != "" {
			transaction.Location.Value.Country = antifraud_v1.OptString{Set: true, Value: t.Location.Country}
		}
	}

	if t.Metadata != nil {
		metadata := make(antifraud_v1.TransactionMetadata)
		for k, v := range *t.Metadata {
			if str, ok := v.(string); ok {
				metadata[k] = jx.Raw(str)
			}
		}
		transaction.Metadata = antifraud_v1.OptTransactionMetadata{Set: true, Value: metadata}
	}

	return transaction
}
