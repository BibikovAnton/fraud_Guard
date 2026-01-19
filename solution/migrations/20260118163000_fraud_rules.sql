CREATE TABLE IF NOT EXISTS fraud_rules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(500) NOT NULL,
    dsl TEXT NOT NULL,
    priority INTEGER NOT NULL DEFAULT 100,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_fraud_rules_active_priority ON fraud_rules (is_active, priority) WHERE is_active = true;
CREATE INDEX IF NOT EXISTS idx_fraud_rules_name ON fraud_rules (name);

CREATE OR REPLACE FUNCTION update_fraud_rules_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER trigger_fraud_rules_updated_at
    BEFORE UPDATE ON fraud_rules
    FOR EACH ROW
    EXECUTE FUNCTION update_fraud_rules_updated_at();
