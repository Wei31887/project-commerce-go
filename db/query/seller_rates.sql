-- name: CreateSellerRate :one
INSERT INTO seller_rates (
  user_id, rate
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateSellerRate :one
UPDATE seller_rates SET rate = $2
WHERE user_id = $1
RETURNING *;


-- name: GetSellerRateByUserId :one
SELECT * FROM seller_rates
WHERE user_id = $1 LIMIT 1;


-- name: DeleteSellerRate :exec
DELETE FROM seller_rates
WHERE id = $1;