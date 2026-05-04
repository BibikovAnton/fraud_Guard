Anti-Fraud System

> Проект разработан при поддержке **Т-Банк**

📋 Описание проекта

Backend‑сервис, который принимает финансовые транзакции и автоматически проверяет, мошеннические они или нормальные.
* Пользователь отправляет данные операции.
* Система проверяет её по набору заранее заданных правил (например, «слишком большая сумма» или «подозрительный канал платежа»).
* В ответ она выдаёт статус APPROVED (операция ок) или DECLINED (подозрительно, лучше отклонить).
Проект показывает, как строить реальный антифрод‑движок на Go:
* регистрация и логин пользователей,
* админ‑панель и статистика,
* обработка операций по‑одному и пачками,
* хранение и применение правил мошенничества.

🏗 Архитектура

**Монолитный Go-сервис** с чётко разделёнными слоями (Layered Architecture):

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

| Компонент          | Роль                                                   |
|--------------------|--------------------------------------------------------|
| Go service (8080)  | Основной REST API, запущен через go-chi                |
| Postgres 16        | Хранение пользователей, транзакций, фрауд-правил       |
| Redis 7            | Кэширование правил и сессий                            |
| antifraud-checker  | Внешний тестировщик, который автоматически проверяет решение |

Слои приложения (внутри `internal/`):
* **api** — HTTP handlers, маршрутизация, сериализация
* **service** — бизнес-логика
* **repository** — работа с БД (Data Access Layer)
* **model** — доменные модели
* **dsl** — валидация и исполнение правил
* **config** — загрузка конфигурации
* **migrator** — управление миграциями
* **errors** — централизованная обработка ошибок
* **app** — сборка и запуск приложения

🌟 Ключевые особенности

* 🛡 **Rate Limiting** — защита от злоупотреблений API через middleware
* 🔥 **Кэширование правил в Redis** — быстрая оценка транзакций без запросов к БД
* 🔍 **Request ID** — уникальный ID для каждого запроса для трассировки
* 📝 **Структурированные логи** — JSON-логи с zap для удобного анализа
* 📊 **OpenTelemetry** — распределённая трассировка и метрики
* 🧪 **antifraud-checker** — автоматическое тестирование всех эндпоинтов

Как это работает

Пользователь регистрируется → Создаёт транзакцию → 
Система проверяет по fraud-правилам → APPROVED / DECLINED

Пример правила: если сумма транзакции > 100 000 ₽ — пометить как подозрительную. Или если транзакция совершена из необычной геолокации.

📁 Структура проекта

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

API Endpoints

| Метод  | Путь                              | Описание                                      |
|--------|-----------------------------------|-----------------------------------------------|
| POST   | /api/v1/auth/register             | Регистрация                                   |
| POST   | /api/v1/auth/login                | Логин                                         |
| GET    | /api/v1/users                     | Список пользователей (Admin)                  |
| PATCH  | /api/v1/users/{id}                | Обновить пользователя                         |
| PATCH  | /api/v1/users/{id}/block          | Заблокировать                                 |
| PATCH  | /api/v1/users/{id}/unblock        | Разблокировать                                |
| POST   | /api/v1/fraud-rules               | Создать фрауд-правило                         |
| GET    | /api/v1/fraud-rules               | Список правил                                 |
| POST   | /api/v1/fraud-rules/validate      | Валидировать DSL                              |
| POST   | /api/v1/transactions              | Создать транзакцию                            |
| GET    | /api/v1/transactions/{id}         | Получить транзакцию                           |
| POST   | /api/v1/transactions/batch        | Batch-создание (207)                          |
| GET    | /api/v1/stats/overview            | Обзор статистики (Admin)                      |
| GET    | /api/v1/stats/transactions/timeseries | Временные ряды (Admin)                    |
| GET    | /api/v1/ping                      | Health-check                                  |

🚀 Как запустить локально

### Через Docker (рекомендуется)

```bash
git clone https://github.com/ArtemBibikov/antiford-system.git
cd antiford-system
docker compose up -d
```

API откроется на http://localhost:8080  
OpenAPI-документация — http://localhost:8080/docs

### Локальная разработка (без Docker)

1. Убедитесь, что установлены **Go 1.25+**, **Postgres** и **Redis**

2. Установите зависимости:
```bash
cd solution
go mod download
```

3. Запустите миграции:
```bash
go run github.com/pressly/goose/v3/cmd/goose -dir=migrations postgres "postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable" up
```

4. Запустите сервис:
```bash
cd solution
go run cmd/main.go
```

5. Переменные окружения (если нужны отличные от дефолтных):
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=testdb
export DB_USER=postgres
export DB_PASSWORD=postgres
export REDIS_HOST=localhost
export REDIS_PORT=6379
export SERVER_PORT=8080
export ADMIN_EMAIL=admin@mail.ru
export ADMIN_FULLNAME=Test Test
export ADMIN_PASSWORD=<your-secure-password>
export RANDOM_SECRET=<generate-with-openssl>
```

> ⚠️ **Важно:** Замените `ADMIN_PASSWORD` и `RANDOM_SECRET` на безопасные значения в production!

👤 Default-админ (для разработки)

При первом запуске автоматически создаётся администратор:

| Поле         | Значение          |
|--------------|-------------------|
| Email        | admin@mail.ru     |
| Пароль       | 123123123aA!      |

🧪 Тестирование

Проект проверяется внешним **antifraud-checker**, который запускается через docker-compose и генерирует отчёт в `./reports/junit.xml`.

```bash
docker compose up -d
# Checker запустится автоматически и проверит все эндпоинты
```

📝 CI/CD

Проект настроен для GitLab CI/CD (.gitlab-ci.yml):
1. **build** — сборка Docker-образа через Kaniko
2. **test** — запуск docker-compose с checker'ом
3. **export** — отправка результатов в judging API

🗺 Roadmap / Доработки

| Приоритет | Доработка                          | Статус |
|-----------|------------------------------------|--------|
| 🔴 Срочно | Убрать реальные секреты из README  | ✅ Выполнено |
| 🟡 Важно  | Реализовать кэширование правил в Redis | 🔄 В работе |
| 🟡 Важно  | Добавить Request ID и структурированные логи | 🔄 В работе |
| 🟢 Желательно | Rate limiting для API              | 🔄 В работе |

📝 CI/CD

Проект настроен для GitLab CI/CD (.gitlab-ci.yml):
1. **build** — сборка Docker-образа через Kaniko
2. **test** — запуск docker-compose с checker'ом
3. **export** — отправка результатов в judging API

