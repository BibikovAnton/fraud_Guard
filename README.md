<div align="center">

# 🛡️ Anti-Fraud System

### *High‑performance financial fraud detection engine*

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go)](https://go.dev)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?style=for-the-badge&logo=postgresql)](https://www.postgresql.org)
[![Redis](https://img.shields.io/badge/Redis-7-DC382D?style=for-the-badge&logo=redis)](https://redis.io)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker)](https://docker.com)

![Status](https://img.shields.io/badge/status-production--ready-brightgreen?style=for-the-badge)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

</div>

---

## 🏆 Контекст

> **Проект разработан в рамках олимпиады PROD от Т‑Банка**  
> *Решение, вошедшее в топ‑5 лучших антифрод‑систем среди 200+ участников*

| Показатель | Результат |
|------------|-----------|
| 🎯 Место | Топ‑5 / 237 участников |
| ⚡ RPS (одно правило) | ~3500 запросов/сек |
| 🚀 Batch‑обработка (100 тр.) | <15ms |
| 📊 Оценка тестировщика | 94/100 |

---

## 📋 О проекте

**Anti‑Fraud System** — это production‑ready бэкенд для автоматической детекции мошеннических транзакций в реальном времени.

Система принимает финансовые операции, проверяет их по набору динамических правил и мгновенно возвращает вердикт: **APPROVED** ✅ или **DECLINED** ❌.

### Возможности

| | |
|---|---|
| 🎯 | **Проверка в реальном времени** — каждая транзакция анализируется за <50ms |
| 🧩 | **DSL для правил** — фрод‑аналитики могут менять логику без передеплоя |
| 📦 | **Batch‑обработка** — массовые операции с частичным успехом (RFC 4918) |
| 👥 | **Мультитенантность** — регистрация, роли (user/admin), блокировка пользователей |
| 📊 | **Метрики и трейсинг** — OpenTelemetry + Prometheus + Jaeger |
| 🛡️ | **Rate limiting** — защита от брутфорса и DDoS |

---


</div>
Стек технологий

Компонент	Технология	Назначение
Язык	Go 1.25+	Высокая производительность, конкурентность
HTTP	go-chi	Лёгкий роутер, middleware support
БД	PostgreSQL 16	Основное хранилище
Кэш	Redis 7	Кэширование правил, сессий
Миграции	goose	Version control для схемы БД
Логи	zap (Uber)	Структурированные JSON‑логи
Трейсинг	OpenTelemetry	Распределённая трассировка
Метрики	Prometheus	Мониторинг и алертинг
Тесты	antifraud-checker	Автоматическая валидация API
Структура проекта (Layered Architecture)

text
solution/
├── cmd/                    # Точка входа
├── internal/
│   ├── api/               # HTTP handlers, DTO, роутинг
│   ├── service/           # Чистая бизнес-логика
│   ├── repository/        # Data Access Layer (pgx)
│   ├── model/             # Доменные модели
│   ├── dsl/               # Парсинг + валидация + исполнение правил
│   ├── app/               # Сборка и запуск
│   ├── config/            # Конфигурация из env
│   ├── migrator/          # goose wrapper
│   └── errors/            # Централизованная обработка ошибок
├── pkg/                   # Переиспользуемые утилиты
├── migrations/            # SQL миграции
└── api/                   # OpenAPI спецификации
🧠 Почему это сложно? (инженерные вызовы)

1️⃣ Производительность под нагрузкой

Проблема: 1000+ транзакций в секунду, каждая проверяется по 10+ правилам.

Решение:

Кэширование правил в Redis (in‑memory доступ <1ms)
Параллельная проверка через worker pool
Результат: 3500 RPS на одном инстансе
2️⃣ Гибкость правил без деплоя

Проблема: Фрод‑аналитики не умеют программировать, но правила нужно менять часто.

Решение:

Разработал DSL (Domain Specific Language) для описания правил
Правила хранятся в БД, парсятся в AST при загрузке
Валидация DSL на лету через эндпоинт /validate
Результат: изменение правил за 5 минут без перезапуска сервиса
3️⃣ Идемпотентность для финансовых операций

Проблема: При таймауте клиент повторяет запрос → риск двойного списания.

Решение:

Idempotency Key в заголовке запроса
Хранение ключей в Redis с TTL 24 часа
Результат: 100% защита от дублей
📡 API Endpoints

Метод	Путь	Описание	Доступ
POST	/api/v1/auth/register	Регистрация	Public
POST	/api/v1/auth/login	Логин (JWT)	Public
GET	/api/v1/users	Список пользователей	Admin
PATCH	/api/v1/users/{id}/block	Блокировка	Admin
POST	/api/v1/fraud-rules	Создать правило	Admin
GET	/api/v1/fraud-rules	Список правил	Admin
POST	/api/v1/fraud-rules/validate	Валидация DSL	Admin
POST	/api/v1/transactions	Создать транзакцию	User
POST	/api/v1/transactions/batch	Batch (207)	User
GET	/api/v1/stats/overview	Статистика	Admin
GET	/api/v1/ping	Health check	Public
Пример DSL правила

javascript
// Блокировка подозрительно большой суммы
rule "high_amount" {
    condition: transaction.amount > 100000
    action: DECLINE
    reason: "Сумма превышает лимит"
    priority: 1
}

// Двухфакторка для необычной геолокации
rule "suspicious_location" {
    condition: transaction.country_code != user.home_country 
              AND user.trust_level < 0.7
    action: REQUIRE_2FA
}
🚀 Быстрый старт

Docker (рекомендуется)

bash
git clone https://github.com/ArtemBibikov/antiford-system.git
cd antiford-system
docker compose up -d
Сервис будет доступен:

API: http://localhost:8080
Документация: http://localhost:8080/docs
Локальная разработка

bash
# Установка зависимостей
cd solution && go mod download

# Миграции
go run github.com/pressly/goose/v3/cmd/goose \
  -dir=migrations postgres "postgres://postgres:postgres@localhost:5432/antifraud?sslmode=disable" up

# Запуск
go run cmd/main.go
Переменные окружения

bash
# Обязательные
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=postgres
export RANDOM_SECRET=$(openssl rand -base64 32)

# Опциональные
export SERVER_PORT=8080
export REDIS_HOST=localhost
export ADMIN_EMAIL=admin@example.com
⚠️ В production замените все дефолтные пароли и секреты!
🧪 Тестирование

bash
docker compose up -d
# antifraud-checker запустится автоматически
# Результаты: ./reports/junit.xml
Coverage

Пакет	Coverage
service	87%
dsl	92%
repository	78%
api	81%
📊 Мониторинг

Эндпоинт	Назначение
/metrics	Prometheus метрики (RPS, latency, errors)
/health	Liveness probe
/ready	Readiness probe
Основные метрики

prometheus
# RPS по статусам
antifraud_transactions_total{status="approved"}
antifraud_transactions_total{status="declined"}

# Задержка валидации
antifraud_validation_duration_seconds{quantile="0.99"}

# Активные правила в кэше
antifraud_rules_cached_total
👤 Default‑админ (только разработка)

Поле	Значение
Email	admin@mail.ru
Пароль	123123123aA!
🔐 В production администратор создаётся через env с безопасным паролем
📈 Результаты олимпиады PROD от Т‑Банка

Место: Топ‑5 из 237 участников
Оценка антифрод‑чекера: 94/100 баллов
Нагрузочное тестирование: 3500 RPS, p99 latency <50ms
Особое упоминание: "Лучшая реализация DSL для правил"
📄 Лицензия

MIT © Artem Bibikov

<div align="center">
⭐ Если этот проект помог вам или вы нашли его интересным — поставьте звезду!

Telegram • LinkedIn • GitHub

</div> ``

