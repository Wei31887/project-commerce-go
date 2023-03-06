-- name: CreateCart :one
INSERT INTO carts (
  id
) VALUES (
  $1
)
RETURNING *;

-- name: DeleteCart :exec
DELETE FROM carts
WHERE id = $1;