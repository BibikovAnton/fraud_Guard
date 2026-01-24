package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            string         `json:"id"`
	Email         string         `json:"email"`
	PasswordHash  string         `json:"-"`
	FullName      string         `json:"fullName"`
	Age           *int           `json:"age,omitempty"`
	Region        *string        `json:"region,omitempty"`
	Gender        *Gender        `json:"gender,omitempty"`
	MaritalStatus *MaritalStatus `json:"maritalStatus,omitempty"`
	Role          UserRole       `json:"role"`
	IsActive      bool           `json:"isActive"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}

type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
)

type MaritalStatus string

const (
	MaritalStatusSingle   MaritalStatus = "SINGLE"
	MaritalStatusMarried  MaritalStatus = "MARRIED"
	MaritalStatusDivorced MaritalStatus = "DIVORCED"
	MaritalStatusWidowed  MaritalStatus = "WIDOWED"
)

type UserRole string

const (
	UserRoleConst UserRole = "USER"
	AdminRole     UserRole = "ADMIN"
)

type RegisterRequest struct {
	Email         string         `json:"email" validate:"required,email,max=254"`
	Password      string         `json:"password" validate:"required,min=8,max=72"`
	FullName      string         `json:"fullName" validate:"required,min=2,max=200"`
	Age           *int           `json:"age,omitempty" validate:"omitempty,ge=18,le=120"`
	Region        *string        `json:"region,omitempty" validate:"omitempty,max=32"`
	Gender        *Gender        `json:"gender,omitempty"`
	MaritalStatus *MaritalStatus `json:"maritalStatus,omitempty"`
}

type RegisterResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
	User        User   `json:"user"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=254"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
	User        User   `json:"user"`
}

type UserCreateRequest struct {
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	FullName      string         `json:"fullName"`
	Region        *string        `json:"region,omitempty"`
	Gender        *Gender        `json:"gender,omitempty"`
	Age           *int           `json:"age,omitempty"`
	MaritalStatus *MaritalStatus `json:"maritalStatus,omitempty"`
	Role          UserRole       `json:"role"`
}

type UserUpdateRequest struct {
	FullName      *string        `json:"fullName,omitempty"`
	Region        *string        `json:"region,omitempty"`
	Gender        *Gender        `json:"gender,omitempty"`
	Age           *int           `json:"age,omitempty"`
	MaritalStatus *MaritalStatus `json:"maritalStatus,omitempty"`
	Role          *UserRole      `json:"role,omitempty"`
	IsActive      *bool          `json:"isActive,omitempty"`
}

func NewUser(email, passwordHash, fullName string, role UserRole, age *int, region *string, gender *Gender, maritalStatus *MaritalStatus) User {
	now := time.Now().UTC() // всегда храним время в UTC для консистентности
	return User{
		ID:            uuid.New().String(),
		Email:         email,
		PasswordHash:  passwordHash,
		FullName:      fullName,
		Age:           age,
		Region:        region,
		Gender:        gender,
		MaritalStatus: maritalStatus,
		Role:          role,
		IsActive:      true,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}
