package convertor

import (
	"github.com/google/uuid"
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
		DslExpression: rule.DSL,
		Enabled:       rule.IsActive,
		Priority:      rule.Priority,
		CreatedAt:     rule.CreatedAt,
		UpdatedAt:     rule.UpdatedAt,
	}
}
