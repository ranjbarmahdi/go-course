-- +goose Up
CREATE TABLE IF NOT EXISTS sample_table (
    id         SERIAL PRIMARY KEY,
    uuid       UUID NOT NULL UNIQUE,
    name       TEXT,
    number     INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS sample_table;