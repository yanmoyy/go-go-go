-- name: CreateUser :exec
INSERT INTO
  users (id, created_at, updated_at, name)
VALUES
  (gen_random_uuid (), now(), now(), $1);
