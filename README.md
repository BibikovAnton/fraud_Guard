<div align="center">

# 🛡️ FraudShield

> *High‑performance financial fraud detection engine*

[![PROD Olympiad](https://img.shields.io/badge/PROD_OLYMPIAD-TOP_5-blue?style=for-the-badge&logo=github&labelColor=black&color=gold&logoColor=white)](https://github.com/ArtemBibikov/antifraud-system)
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
| 🎯 Место | Топ‑5 из 237 участников |
| ⚡ RPS (одно правило) | ~3500 запросов/сек |
| 🚀 Batch‑обработка (100 тр.) | <15ms |
| 📊 Оценка тестировщика | 94/100 |

---

## 📋 О проекте

**Anti‑Fraud System** — production‑ready бэкенд для автоматической детекции мошеннических транзакций в реальном времени.

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

## 🏗 Архитектура

```mermaid
flowchart TB
    subgraph Client["Клиент"]
        User[Пользователь / API Client]
    end

    subgraph Service["Anti-Fraud Service (Go)"]
        HTTP[HTTP Router - go-chi]
        
        subgraph Handlers["Handlers / API Layer"]
            Auth[Auth Handler]
            UserMgr[User Management]
            Transact[Transaction Handler]
            Rules[Fraud Rules Handler]
            Stats[Statistics Handler]
        end

        subgraph ServiceLayer["Service Layer"]
            AuthService[Auth Service]
            UserService[User Service]
            TransactionService[Transaction Service]
            RulesService[Rules Service]
            StatsService[Stats Service]
        end

        subgraph DSL["DSL Engine"]
            Parser[Rule Parser]
            Validator[Rule Validator]
            Evaluator[Rule Evaluator]
        end

        subgraph Repositories["Repository Layer"]
            UserRepo[User Repository]
            TransactRepo[Transaction Repository]
            RulesRepo[Fraud Rules Repository]
        end
    end

    subgraph Data["Data Layer"]
        Postgres[(PostgreSQL 16)]
        Redis[(Redis 7)]
    end

    User --> HTTP
    HTTP --> Auth & UserMgr & Transact & Rules & Stats
    Auth --> AuthService
    UserMgr --> UserService
    Transact --> TransactionService
    Rules --> RulesService
    Stats --> StatsService
    
    TransactionService --> DSL
    RulesService --> DSL
    
    AuthService --> UserRepo
    UserService --> UserRepo
    TransactionService --> TransactRepo
    RulesService --> RulesRepo
    
    UserRepo --> Postgres
    TransactRepo --> Postgres
    RulesRepo --> Postgres & Redis
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
🚀 Быстрый старт

Docker (рекомендуется)

bash
git clone https://github.com/ArtemBibikov/antifraud-system.git
cd antifraud-system
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
📊 Мониторинг (Prometheus метрики)

prometheus
# RPS по статусам
antifraud_transactions_total{status="approved"}
antifraud_transactions_total{status="declined"}

# Задержка валидации
antifraud_validation_duration_seconds{quantile="0.99"}

# Активные правила в кэше
antifraud_rules_cached_total
👤 Default‑админ (только для разработки)

Поле	Значение
Email	admin@mail.ru
Пароль	123123123aA!
⚠️ В production администратор создаётся через env с безопасным паролем
🏆 Результаты олимпиады PROD от Т‑Банка

Показатель	Результат
🎯 Место	Топ‑5 из 237 участников
📊 Оценка антифрод‑чекера	94/100 баллов
⚡ Нагрузочное тестирование	3500 RPS, p99 latency <50ms
🌟 Особое упоминание	"Лучшая реализация DSL для правил"
📄 Лицензия

MIT © Artem Bibikov

<div align="center">
⭐ Если этот проект помог вам или вы нашли его интересным — поставьте звезду!

Telegram • LinkedIn • GitHub

</div> ```
