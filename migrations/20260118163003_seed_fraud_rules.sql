-- +goose Up
INSERT INTO fraud_rules (id, name, description, dsl, priority, is_active) VALUES
('76a681b7-d876-4ea3-8514-35deba891bcc', 'High amount check', 'AMOUNT > 10000', 'amount > 10000', 10, true),
('6e9df41b-3b83-4968-8851-b047845c1d64', 'Old humans risk', 'USER.AGE > 90', 'user.age > 90', 11, true);

-- +goose Down
DELETE FROM fraud_rules WHERE id IN ('76a681b7-d876-4ea3-8514-35deba891bcc', '6e9df41b-3b83-4968-8851-b047845c1d64');
