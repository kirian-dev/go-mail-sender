-- 20231031235822_create_newsletters.sql

-- +goose Up
CREATE TABLE newsletters (
    id UUID PRIMARY KEY,
    message VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE newsletters;