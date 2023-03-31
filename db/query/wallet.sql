-- name: CreateWallet :one
INSERT INTO wallet (
  user_id, 
  name,
  balance,
  icon
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetWalletForUpdate :one
SELECT * FROM wallet
WHERE id = $1 FOR NO KEY UPDATE;

-- name: GetWallet :one
SELECT * FROM wallet
WHERE id = $1 LIMIT 1;

-- name: ListWallets :many
SELECT * FROM wallet 
WHERE user_id = $1
ORDER BY name;

-- name: UpdateWallet :one
UPDATE wallet SET
  name = $1,
  balance = $2,
  icon = $3
WHERE id = $4
RETURNING *;

-- name: AddWalletBalance :one
UPDATE wallet SET
  balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: UpdateWalletBalance :one
UPDATE wallet SET
  balance = sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteWallet :exec
DELETE FROM wallet
WHERE id = $1;