-- 20231027235750_create_users.sql

-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP
);

-- +goose Down
DROP TABLE users;