-- +goose Up
CREATE TABLE transaction_rule_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transaction_id UUID NOT NULL,
    rule_id UUID NOT NULL,
    rule_name VARCHAR(120) NOT NULL,
    priority INTEGER NOT NULL,
    enabled BOOLEAN NOT NULL,
    matched BOOLEAN NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    
    CONSTRAINT fk_transaction_rule_results_transaction 
        FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE
);

CREATE INDEX idx_transaction_rule_results_transaction_id ON transaction_rule_results(transaction_id);
CREATE INDEX idx_transaction_rule_results_rule_id ON transaction_rule_results(rule_id);
CREATE INDEX idx_transaction_rule_results_matched ON transaction_rule_results(matched);

-- +goose Down
DROP TABLE transaction_rule_results;
