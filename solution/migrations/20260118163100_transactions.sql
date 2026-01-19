-- +goose Up
-- Создание таблицы транзакций
-- Из прошлого проекта: индексы критически важны для производительности на 10k RPS

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    amount DECIMAL(15,2) NOT NULL CHECK (amount > 0),
    currency VARCHAR(3) NOT NULL CHECK (currency IN ('USD', 'EUR', 'RUB')),
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'APPROVED', 'REJECTED', 'PROCESSED')),
    merchant_id VARCHAR(100),
    merchant_category_code VARCHAR(4),
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    ip_address INET,
    device_id VARCHAR(100),
    channel VARCHAR(20) CHECK (channel IN ('ONLINE', 'POS', 'MOBILE', 'ATM')),
    location JSONB,
    is_fraud BOOLEAN NOT NULL DEFAULT false,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Индексы для производительности
CREATE INDEX IF NOT EXISTS idx_transactions_user_id ON transactions (user_id);
CREATE INDEX IF NOT EXISTS idx_transactions_timestamp ON transactions (timestamp DESC);
CREATE INDEX IF NOT EXISTS idx_transactions_status ON transactions (status);
CREATE INDEX IF NOT EXISTS idx_transactions_is_fraud ON transactions (is_fraud);
CREATE INDEX IF NOT EXISTS idx_transactions_merchant_id ON transactions (merchant_id) WHERE merchant_id IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_transactions_amount ON transactions (amount) WHERE amount > 1000;
CREATE INDEX IF NOT EXISTS idx_transactions_currency ON transactions (currency);

-- Составные индексы для частых запросов
CREATE INDEX IF NOT EXISTS idx_transactions_user_timestamp ON transactions (user_id, timestamp DESC);
CREATE INDEX IF NOT EXISTS idx_transactions_fraud_timestamp ON transactions (is_fraud, timestamp DESC) WHERE is_fraud = true;

-- Триггер для updated_at
CREATE OR REPLACE FUNCTION update_transactions_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_transactions_updated_at
    BEFORE UPDATE ON transactions
    FOR EACH ROW
    EXECUTE FUNCTION update_transactions_updated_at();

-- +goose Down
DROP TRIGGER IF EXISTS trigger_transactions_updated_at ON transactions;
DROP FUNCTION IF EXISTS update_transactions_updated_at();
DROP INDEX IF EXISTS idx_transactions_fraud_timestamp;
DROP INDEX IF EXISTS idx_transactions_user_timestamp;
DROP INDEX IF EXISTS idx_transactions_currency;
DROP INDEX IF EXISTS idx_transactions_amount;
DROP INDEX IF EXISTS idx_transactions_merchant_id;
DROP INDEX IF EXISTS idx_transactions_is_fraud;
DROP INDEX IF EXISTS idx_transactions_status;
DROP INDEX IF EXISTS idx_transactions_timestamp;
DROP INDEX IF EXISTS idx_transactions_user_id;
DROP TABLE IF EXISTS transactions;
