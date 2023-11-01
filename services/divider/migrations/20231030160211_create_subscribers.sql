-- 20231030160211_create_subscribers.sql
-- +goose Up
CREATE TABLE subscribers (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    packet_id UUID DEFAULT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE subscribers;