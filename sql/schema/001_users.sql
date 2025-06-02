-- +goose Up
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;
