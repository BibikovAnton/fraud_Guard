package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTData struct {
	UserID string
	Role   string
}

type JWT struct {
	secret      string
	tokenExpiry time.Duration
}

func NewJWT(secret string) *JWT {
	return &JWT{
		secret:      secret,
		tokenExpiry: time.Hour,
	}
}

type CustomClaims struct {
	UserID string `json:"sub"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func (j *JWT) Create(data JWTData) (string, error) {
	if data.Role != "USER" && data.Role != "ADMIN" {
		return "", errors.New("invalid role, must be USER or ADMIN")
	}

	if data.UserID == "" {
		return "", errors.New("userID cannot be empty")
	}

	now := time.Now()
	claims := CustomClaims{
		UserID: data.UserID,
		Role:   data.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.tokenExpiry)),
			Subject:   data.UserID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) Parse(tokenString string) (bool, *JWTData, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return false, nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwtData := &JWTData{
			UserID: claims.UserID,
			Role:   claims.Role,
		}
		return true, jwtData, nil
	}

	return false, nil, errors.New("invalid token")
}

func ValidateSecret(secret string) error {
	if len(secret) != 128 {
		return errors.New("secret must be exactly 128 characters long")
	}
	return nil
}

func (j *JWT) GetTokenExpiry() time.Duration {
	return j.tokenExpiry
}

// GenerateToken - удобная функция для генерации токена
// Добавил для совместимости с UserService, из прошлого проекта так было удобнее
func GenerateToken(userID, role, secret string, expiresIn time.Duration) (string, error) {
	if role != "USER" && role != "ADMIN" {
		return "", errors.New("invalid role, must be USER or ADMIN")
	}

	if userID == "" {
		return "", errors.New("userID cannot be empty")
	}

	now := time.Now().UTC() // всегда UTC для воспроизводимости
	claims := CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiresIn)),
			Subject:   userID,
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
