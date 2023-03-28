-- name: CreateCategory :one
INSERT INTO category (
  user_id, 
  name,
  type,
  icon
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetCategory :one
SELECT * FROM category
WHERE id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM category 
WHERE user_id = $1
ORDER BY name;

-- name: UpdateCategory :one
UPDATE category SET
  name = $1,
  type = $2,
  icon = $3
WHERE id = $4
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category
WHERE id = $1;