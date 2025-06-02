-- name: CreateUser :one
INSERT INTO
  users (
    id,
    created_at,
    updated_at,
    username,
    hashed_password
  )
VALUES
  (gen_random_uuid (), now(), now(), $1, $2)
RETURNING
  *;

-- name: GetUserByUsername :one
SELECT
  *
FROM
  users
WHERE
  username = $1;
