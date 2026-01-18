-- +goose Up
CREATE SCHEMA IF NOT EXISTS antifraud;

-- +goose Down
DROP SCHEMA IF EXISTS antifraud;
