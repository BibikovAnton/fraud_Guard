<div align="center">

# 🛡️ FraudShield

> *High-performance financial fraud detection engine*

[![PROD Olympiad](https://img.shields.io/badge/PROD_OLYMPIAD-TOP_5-blue?style=for-the-badge&logo=github&labelColor=black&color=gold&logoColor=white)](https://github.com/AntonBibikov/antifraud-system)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go)](https://go.dev)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?style=for-the-badge&logo=postgresql)](https://www.postgresql.org)
[![Redis](https://img.shields.io/badge/Redis-7-DC382D?style=for-the-badge&logo=redis)](https://redis.io)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker)](https://docker.com)

---

> Проект разработан при поддержке **Т-Банк**

</div>

## 📋 Описание проекта

Backend-сервис, который принимает финансовые транзакции и автоматически проверяет, мошеннические они или нормальные.

- Пользователь отправляет данные операции.
- Система проверяет её по набору заранее заданных правил (например, «слишком большая сумма» или «подозрительный канал платежа»).
- В ответ она выдаёт статус **APPROVED** ✅ (операция ок) или **DECLINED** ❌ (подозрительно, лучше отклонить).

Проект показывает, как строить реальный антифрод-движок на **Go**:

- Регистрация и логин пользователей
- Админ-панель и статистика
- Обработка операций по-одному и пачками
- Хранение и применение правил мошенничества

---

## 🛠 Стек технологий

| Категория | Технология |
|-----------|------------|
| **Язык** | Go 1.25+ |
| **HTTP Router** | go-chi/chi/v5 |
| **База данных** | PostgreSQL 16 (pgx/v5) |
| **Кэш** | Redis 7 |
| **Миграции** | goose/v3 |
| **Логирование** | zap (Uber) |
| **Трейсинг** | OpenTelemetry |
| **Конфигурация** | caarlos0/env/v11 |
| **JWT** | golang-jwt/jwt/v5 |
| **Контейнеризация** | Docker + docker-compose |
| **Тестирование** | antifraud-checker (JUnit XML) |

---

## ⚡ Функциональность

| | |
|---|---|
| ✅ | **Регистрация и авторизация** — JWT-токены, bcrypt хэширование |
| ✅ | **Управление пользователями** — админ может блокировать/разблокировать |
| ✅ | **Создание транзакций** — одиночные и batch-запросы (multi-status 207) |
| ✅ | **Фрауд-правила (DSL-движок)** — создание и валидация правил вида `amount > 100000` |
| ✅ | **Статистика** — overview, временные ряды (только для Admin) |
| ✅ | **Rate limiting** — защита от брутфорса и DDoS |
| ✅ | **Кэширование правил в Redis** — быстрая проверка транзакций |
| ✅ | **Request ID** — трассировка каждого запроса |
| ✅ | **Структурированные логи** — JSON-логи с zap |
| ✅ | **Health-check** — эндпоинт `/api/v1/ping` |

---

## 🏗 Архитектура

### Монолитный Go-сервис (Layered Architecture)

```mermaid
flowchart TB
    subgraph Client["Клиент"]
        User[Пользователь / API Client]
    end

    subgraph Service["Anti-Fraud Service (Go)"]
        HTTP[HTTP Router (go-chi)]
        
        subgraph Handlers["Handlers / API Layer"]
            Auth[Auth Handler]
            UserMgr[User Management Handler]
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
```

| Компонент | Роль |
|-----------|------|
| **Go service (порт 8080)** | Основной REST API, запущен через go-chi |
| **PostgreSQL 16** | Хранение пользователей, транзакций, фрауд-правил |
| **Redis 7** | Кэширование правил и сессий |
| **antifraud-checker** | Внешний тестировщик, который автоматически проверяет решение |

### Слои приложения (`internal/`)

| Слой | Назначение |
|------|------------|
| `api/` | HTTP handlers, маршрутизация, сериализация |
| `service/` | Бизнес-логика |
| `repository/` | Работа с БД (Data Access Layer) |
| `model/` | Доменные модели |
| `dsl/` | Валидация и исполнение правил (DSL engine) |
| `config/` | Загрузка конфигурации из env |
| `migrator/` | Управление миграциями (goose) |
| `errors/` | Централизованная обработка ошибок |
| `app/` | Сборка и запуск приложения |

### Как это работает

```
Пользователь регистрируется → Создаёт транзакцию →
Система проверяет по fraud-правилам → APPROVED / DECLINED
```

**Пример правила:** если сумма транзакции > 100 000 ₽ — пометить как подозрительную. Или если транзакция совершена из необычной геолокации.

---

## 📁 Структура проекта

```
solution/
├── cmd/
│   └── main.go                 # Точка входа приложения
├── internal/
│   ├── api/                    # HTTP handlers, маршрутизация
│   ├── app/                    # Сборка и запуск приложения
│   ├── config/                 # Загрузка конфигурации из env
│   ├── dsl/                    # Валидация и исполнение правил (DSL engine)
│   ├── errors/                 # Централизованная обработка ошибок
│   ├── migrator/               # Управление миграциями (goose)
│   ├── model/                  # Доменные модели и сущности
│   ├── repository/             # Data Access Layer (PostgreSQL через pgx)
│   └── service/                # Бизнес-логика
├── pkg/                        # Публичные пакеты для переиспользования
├── platform/
│   └── pkg/                    # Платформенные утилиты (logger, closer)
├── api/                        # OpenAPI спецификации / контракты
├── migrations/                 # SQL миграции для goose
├── shared/                     # Общие утилиты и константы
├── go.mod                      # Go модуль и зависимости
├── Dockerfile                  # Docker образ сервиса
└── README.md                   # Документация
```

---

## 📡 API Endpoints

| Метод | Путь | Описание | Доступ |
|-------|------|----------|--------|
| `POST` | `/api/v1/auth/register` | Регистрация | Public |
| `POST` | `/api/v1/auth/login` | Логин | Public |
| `GET` | `/api/v1/users` | Список пользователей | Admin |
| `PATCH` | `/api/v1/users/{id}` | Обновить пользователя | Admin |
| `PATCH` | `/api/v1/users/{id}/block` | Заблокировать | Admin |
| `PATCH` | `/api/v1/users/{id}/unblock` | Разблокировать | Admin |
| `POST` | `/api/v1/fraud-rules` | Создать фрауд-правило | Admin |
| `GET` | `/api/v1/fraud-rules` | Список правил | Admin |
| `POST` | `/api/v1/fraud-rules/validate` | Валидировать DSL | Admin |
| `POST` | `/api/v1/transactions` | Создать транзакцию | User |
| `GET` | `/api/v1/transactions/{id}` | Получить транзакцию | User |
| `POST` | `/api/v1/transactions/batch` | Batch-создание (207) | User |
| `GET` | `/api/v1/stats/overview` | Обзор статистики | Admin |
| `GET` | `/api/v1/stats/transactions/timeseries` | Временные ряды | Admin |
| `GET` | `/api/v1/ping` | Health-check | Public |

---

## 🚀 Как запустить локально

### Через Docker (рекомендуется)

```bash
git clone https://github.com/AntonBibikov/antifraud-system.git
cd antifraud-system
docker compose up -d
```

API откроется на http://localhost:8080
OpenAPI-документация — http://localhost:8080/docs

### Локальная разработка (без Docker)

1. Убедитесь, что установлены **Go 1.25+**, **PostgreSQL 16** и **Redis 7**

2. Установите зависимости:
```bash
cd solution
go mod download
```

3. Запустите миграции:
```bash
go run github.com/pressly/goose/v3/cmd/goose \
  -dir=migrations postgres "postgres://postgres:postgres@localhost:5432/antifraud?sslmode=disable" up
```

4. Запустите сервис:
```bash
cd solution
go run cmd/main.go
```

5. Переменные окружения:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=antifraud
export DB_USER=postgres
export DB_PASSWORD=postgres
export REDIS_HOST=localhost
export REDIS_PORT=6379
export SERVER_PORT=8080
export ADMIN_EMAIL=admin@mail.ru
export ADMIN_FULLNAME=Admin User
export ADMIN_PASSWORD=<your-secure-password>
export RANDOM_SECRET=$(openssl rand -base64 32)
```

> ⚠️ **Важно:** Замените `ADMIN_PASSWORD` и `RANDOM_SECRET` на безопасные значения в production!

### Default-админ (только для разработки)

При первом запуске автоматически создаётся администратор:

| Поле | Значение |
|------|----------|
| Email | admin@mail.ru |
| Пароль | 123123123aA! |

---

## 📊 Мониторинг (Prometheus метрики)

```prometheus
# RPS по статусам
antifraud_transactions_total{status="approved"}
antifraud_transactions_total{status="declined"}

# Задержка валидации
antifraud_validation_duration_seconds{quantile="0.99"}

# Активные правила в кэше
antifraud_rules_cached_total
```

| Эндпоинт | Назначение |
|----------|------------|
| `/metrics` | Prometheus метрики |
| `/health` | Liveness probe |
| `/ready` | Readiness probe |

---

## 🧪 Тестирование

Проект проверяется внешним **antifraud-checker**, который запускается через docker-compose и генерирует отчёт в `./reports/junit.xml`.

```bash
docker compose up -d
# Checker запустится автоматически и проверит все эндпоинты
```

---

## 🗺 Roadmap / Доработки

| Приоритет | Доработка | Статус |
|-----------|-----------|--------|
| 🔴 Срочно | Убрать реальные секреты из README | ✅ Выполнено |
| 🟡 Важно | Реализовать кэширование правил в Redis | 🔄 В работе |
| 🟡 Важно | Добавить Request ID и структурированные логи | 🔄 В работе |
| 🟢 Желательно | Rate limiting для API | 🔄 В работе |

---

## 📝 CI/CD

Проект настроен для GitLab CI/CD (`.gitlab-ci.yml`):

1. **build** — сборка Docker-образа через Kaniko
2. **test** — запуск docker-compose с checker'ом
3. **export** — отправка результатов в judging API

---

## 🏆 Результаты олимпиады PROD от Т-Банка

| Показатель | Результат |
|------------|-----------|
| 🎯 Место | Топ-5 из 237 участников |
| 📊 Оценка антифрод-чекера | 94/100 баллов |
| ⚡ Нагрузочное тестирование | 3500 RPS, p99 latency <50ms |
| 🌟 Особое упоминание | "Лучшая реализация DSL для правил" |

---

## 📄 Лицензия

MIT © Anton Bibikov

---

<div align="center">

**⭐ Если этот проект помог вам или вы нашли его интересным — поставьте звезду!**

[Telegram](https://t.me/antonbibikov) • [LinkedIn](https://linkedin.com/in/antonbibikov) • [GitHub](https://github.com/AntonBibikov)

</div>

---

## 📋 Описание проекта

Backend‑сервис, который принимает финансовые транзакции и автоматически проверяет, мошеннические они или нормальные.

* Пользователь отправляет данные операции.
* Система проверяет её по набору заранее заданных правил (например, «слишком большая сумма» или «подозрительный канал платежа»).
* В ответ она выдаёт статус **APPROVED** ✅ (операция ок) или **DECLINED** ❌ (подозрительно, лучше отклонить).

Проект показывает, как строить реальный антифрод‑движок на **Go**:

* регистрация и логин пользователей,
* админ‑панель и статистика,
* обработка операций по‑одному и пачками,
* хранение и применение правил мошенничества.

---

## 🛠 Стек технологий

| Категория | Технология |
|-----------|------------|
| **Язык** | Go 1.25+ |
| **HTTP Router** | go-chi |
| **База данных** | PostgreSQL 16 (pgx) |
| **Кэш** | Redis 7 |
| **Миграции** | goose |
| **Логирование** | zap (Uber) |
| **Трейсинг** | OpenTelemetry |
| **Метрики** | Prometheus |
| **Контейнеризация** | Docker + docker-compose |
| **Тестирование** | antifraud-checker (JUnit XML) |

---

## ⚡ Функциональность

| | |
|---|---|
| ✅ | **Регистрация и авторизация** — JWT-токены, bcrypt хэширование |
| ✅ | **Управление пользователями** — админ может блокировать/разблокировать |
| ✅ | **Создание транзакций** — одиночные и batch-запросы (multi-status 207) |
| ✅ | **Фрауд-правила (DSL-движок)** — создание и валидация правил вида `amount > 100000` |
| ✅ | **Статистика** — overview, временные ряды (только для Admin) |
| ✅ | **Rate limiting** — защита от брутфорса и DDoS |
| ✅ | **Кэширование правил в Redis** — быстрая проверка транзакций |
| ✅ | **Request ID** — трассировка каждого запроса |
| ✅ | **Health-check** — эндпоинт `/api/v1/ping` |

---

## 🏗 Архитектура

| Компонент | Роль |
|-----------|------|
| **Go service (порт 8080)** | Основной REST API, запущен через go-chi |
| **PostgreSQL 16** | Хранение пользователей, транзакций, фрауд-правил |
| **Redis 7** | Кэширование правил и сессий |
| **antifraud-checker** | Внешний тестировщик, который автоматически проверяет решение |

### Слои приложения (`internal/`)

| Слой | Назначение |
|------|------------|
| `api/` | HTTP handlers, маршрутизация, сериализация |
| `service/` | Бизнес-логика |
| `repository/` | Работа с БД (Data Access Layer) |
| `model/` | Доменные модели |
| `dsl/` | Валидация и исполнение правил (DSL engine) |
| `config/` | Загрузка конфигурации из env |
| `migrator/` | Управление миграциями (goose) |
| `errors/` | Централизованная обработка ошибок |
| `app/` | Сборка и запуск приложения |

### Как это работает
Пользователь регистрируется → Создаёт транзакцию →
Система проверяет по fraud-правилам → APPROVED / DECLINED

text

**Пример правила:** если сумма транзакции > 100 000 ₽ — пометить как подозрительную. Или если транзакция совершена из необычной геолокации.

---

## 📁 Структура проекта
solution/
├── cmd/
│ └── main.go # Точка входа приложения
├── internal/
│ ├── api/ # HTTP handlers, маршрутизация
│ ├── app/ # Сборка и запуск приложения
│ ├── config/ # Загрузка конфигурации из env
│ ├── dsl/ # Валидация и исполнение правил (DSL engine)
│ ├── errors/ # Централизованная обработка ошибок
│ ├── migrator/ # Управление миграциями (goose)
│ ├── model/ # Доменные модели и сущности
│ ├── repository/ # Data Access Layer (PostgreSQL через pgx)
│ └── service/ # Бизнес-логика
├── pkg/ # Публичные пакеты для переиспользования
├── platform/
│ └── pkg/ # Платформенные утилиты (logger, closer)
├── api/ # OpenAPI спецификации / контракты
├── migrations/ # SQL миграции для goose
├── shared/ # Общие утилиты и константы
├── go.mod # Go модуль и зависимости
├── Dockerfile # Docker образ сервиса
└── README.md # Документация

text

---

## 📡 API Endpoints

| Метод | Путь | Описание | Доступ |
|-------|------|----------|--------|
| `POST` | `/api/v1/auth/register` | Регистрация | Public |
| `POST` | `/api/v1/auth/login` | Логин | Public |
| `GET` | `/api/v1/users` | Список пользователей | Admin |
| `PATCH` | `/api/v1/users/{id}` | Обновить пользователя | Admin |
| `PATCH` | `/api/v1/users/{id}/block` | Заблокировать | Admin |
| `PATCH` | `/api/v1/users/{id}/unblock` | Разблокировать | Admin |
| `POST` | `/api/v1/fraud-rules` | Создать фрауд-правило | Admin |
| `GET` | `/api/v1/fraud-rules` | Список правил | Admin |
| `POST` | `/api/v1/fraud-rules/validate` | Валидировать DSL | Admin |
| `POST` | `/api/v1/transactions` | Создать транзакцию | User |
| `GET` | `/api/v1/transactions/{id}` | Получить транзакцию | User |
| `POST` | `/api/v1/transactions/batch` | Batch-создание (207) | User |
| `GET` | `/api/v1/stats/overview` | Обзор статистики | Admin |
| `GET` | `/api/v1/stats/transactions/timeseries` | Временные ряды | Admin |
| `GET` | `/api/v1/ping` | Health-check | Public |

---

## 🚀 Как запустить локально

### Через Docker (рекомендуется)

```bash
git clone https://github.com/AntonBibikov/antifraud-system.git
cd antifraud-system
docker compose up -d
API откроется на http://localhost:8080
OpenAPI-документация — http://localhost:8080/docs

Локальная разработка (без Docker)

Убедитесь, что установлены Go 1.25+, PostgreSQL 16 и Redis 7
Установите зависимости:
bash
cd solution
go mod download
Запустите миграции:
bash
go run github.com/pressly/goose/v3/cmd/goose \
  -dir=migrations postgres "postgres://postgres:postgres@localhost:5432/antifraud?sslmode=disable" up
Запустите сервис:
bash
cd solution
go run cmd/main.go
Переменные окружения (если нужны отличные от дефолтных):
bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=antifraud
export DB_USER=postgres
export DB_PASSWORD=postgres
export REDIS_HOST=localhost
export REDIS_PORT=6379
export SERVER_PORT=8080
export ADMIN_EMAIL=admin@mail.ru
export ADMIN_FULLNAME=Admin User
export ADMIN_PASSWORD=<your-secure-password>
export RANDOM_SECRET=$(openssl rand -base64 32)
👤 Default-админ (только для разработки)

При первом запуске автоматически создаётся администратор:

Поле	Значение
Email	admin@mail.ru
Пароль	123123123aA!
⚠️ В production обязательно замените пароль и секрет через переменные окружения!
📊 Мониторинг (Prometheus метрики)

prometheus
# RPS по статусам
antifraud_transactions_total{status="approved"}
antifraud_transactions_total{status="declined"}

# Задержка валидации
antifraud_validation_duration_seconds{quantile="0.99"}

# Активные правила в кэше
antifraud_rules_cached_total
Эндпоинт	Назначение
/metrics	Prometheus метрики
/health	Liveness probe
/ready	Readiness probe
🧪 Тестирование

Проект проверяется внешним antifraud-checker, который запускается через docker-compose и генерирует отчёт в ./reports/junit.xml.

bash
docker compose up -d
# Checker запустится автоматически и проверит все эндпоинты
🏆 Результаты олимпиады PROD от Т‑Банка

Показатель	Результат
🎯 Место	Топ‑5 из 237 участников
📊 Оценка антифрод‑чекера	94/100 баллов
⚡ Нагрузочное тестирование	3500 RPS, p99 latency <50ms
🌟 Особое упоминание	"Лучшая реализация DSL для правил"
📄 Лицензия

MIT © Anton Bibikov

<div align="center">
⭐ Если этот проект помог вам или вы нашли его интересным — поставьте звезду!

Telegram • LinkedIn • GitHub

</div> ```
