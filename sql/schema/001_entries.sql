-- +goose Up
CREATE TABLE IF NOT EXISTS entries (
    id INTEGER PRIMARY KEY,
    category INTEGER NOT NULL,
    friendly_name TEXT NOT NULL,
    host_name TEXT NOT NULL,
    ip_address TEXT NOT NULL,
    description TEXT NOT NULL
);
-- +goose Down
DROP TABLE entries;
