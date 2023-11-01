-- 20231101200957_create_packets.sql

-- +goose Up
CREATE TABLE packets (
    id UUID PRIMARY KEY,
    message VARCHAR(255) DEFAULT '',
    created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE packets;
