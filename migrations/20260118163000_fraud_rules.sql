-- +goose Up
CREATE TABLE fraud_rules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    dsl TEXT NOT NULL,
    priority INTEGER NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_fraud_rules_name ON fraud_rules(name);
CREATE INDEX idx_fraud_rules_is_active ON fraud_rules(is_active);
CREATE INDEX idx_fraud_rules_priority ON fraud_rules(priority);

-- +goose Down
DROP TABLE fraud_rules;
