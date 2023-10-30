-- 20231030160240_create_files.sql

-- +goose Up
CREATE TABLE files (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    success_accounts INT NOT NULL,
    fail_accounts INT NOT NULL,
    loading_accounts INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    status VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE files;