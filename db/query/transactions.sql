-- name: CreateTransaction :one
INSERT INTO transactions (
  user_id, 
  amount,
  type,
  category_id,
  wallet_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListTransactions :many
SELECT * FROM transactions 
WHERE user_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3;

-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: UpdateTransaction :one
UPDATE transactions SET
  amount = $1,
  type = $2,
  category_id = $3,
  wallet_id = $4
WHERE id = $5
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;