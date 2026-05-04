-- +goose Up
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       full_name VARCHAR(200) NOT NULL,
                       age INTEGER CHECK (age >= 18 AND age <= 120),
                       region VARCHAR(32),
                       gender VARCHAR(10) CHECK (gender IN ('MALE', 'FEMALE')),
                       marital_status VARCHAR(10) CHECK (marital_status IN ('SINGLE', 'MARRIED', 'DIVORCED', 'WIDOWED')),
                       role VARCHAR(10) NOT NULL DEFAULT 'USER' CHECK (role IN ('USER', 'ADMIN')),
                       is_active BOOLEAN NOT NULL DEFAULT true,
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

                       CONSTRAINT users_email_key UNIQUE (email)
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_is_active ON users(is_active);
CREATE INDEX idx_users_role ON users(role);

-- +goose Down
drop table users
