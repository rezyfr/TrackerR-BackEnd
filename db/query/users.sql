-- name: CreateUser :one
INSERT INTO users (
  email, 
  full_name,
  hashed_password
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users SET
  full_name = $1
WHERE id = $2
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;