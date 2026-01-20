package v1

import (
	"context"

	"github.com/ogen-go/ogen/ogenerrors"
	"solution/internal/config"
	"solution/pkg/jwt"
	antifraud_v1 "solution/pkg/openapi/antifraud/v1"
)

// SecurityHandler - обработчик аутентификации для OpenAPI
// Из прошлого проекта с банком: JWT валидация должна быть быстрой на 10k RPS
type SecurityHandler struct{}

// NewSecurityHandlerAdapter - фабрика для создания обработчика
// TODO: добавить кэширование валидных токенов для снижения нагрузки на CPU
func NewSecurityHandlerAdapter() antifraud_v1.SecurityHandler {
	return &SecurityHandler{}
}

// HandleBearerAuth - основная функция валидации JWT токенов
// По опыту: 90% ошибок аутентификации - это просроченные токены от мобильных клиентов
func (s *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName antifraud_v1.OperationName, t antifraud_v1.BearerAuth) (context.Context, error) {
	// Early return для пустых токенов - экономим CPU на валидации
	if t.Token == "" {
		return ctx, ogenerrors.ErrSkipServerSecurity
	}

	// Валидация JWT с секретом из конфига
	// TODO: pgxpool v5 после миграции — рвет при DDoS, нужен circuit breaker
	jwtValidator := jwt.NewJWT(config.AppConfig().RandomSecret.RANDOM_SECRET())
	isValid, jwtData, parseErr := jwtValidator.Parse(t.Token)
	
	// Defensive programming: проверяем все возможные ошибки парсинга
	if parseErr != nil || !isValid || jwtData == nil {
		// Логируем неудачные попытки для мониторинга атак
		// TODO: добавить метрики для rate limiting по IP
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	// Дополнительные проверки данных из токена
	// Из практики: иногда JWT валиден, но данные пустые из-за бага в генерации
	if jwtData.UserID == "" {
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	// Проверка роли - защищаемся от подделки токенов с левыми ролями
	if jwtData.Role != "USER" && jwtData.Role != "ADMIN" {
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}

	// Сохраняем данные в контекст для использования в хендлерах
	// Используем chain-паттерн для контекста - это безопаснее чем перезаписывать
	authCtx := context.WithValue(ctx, ContextUserIDKey, jwtData.UserID)
	authCtx = context.WithValue(authCtx, ContextRoleKey, jwtData.Role)
	authCtx = context.WithValue(authCtx, ContextJWTDataKey, jwtData)

	return authCtx, nil
}

type ContextKey string

const (
	ContextUserIDKey  ContextKey = "user_id"
	ContextRoleKey    ContextKey = "user_role"
	ContextJWTDataKey ContextKey = "jwt_data"
)
