-- name: CreateSeller :one
INSERT INTO sellers (
  seller_name, bank_account
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetSeller :one
SELECT * FROM sellers
WHERE id = $1 LIMIT 1;

-- name: DeleteSeller :exec
DELETE FROM sellers
WHERE id = $1;