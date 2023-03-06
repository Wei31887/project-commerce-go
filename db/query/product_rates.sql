-- name: CreateProductRate :one
INSERT INTO product_rates (
  user_id, product_id, rate
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateProductRateByUserId :one
UPDATE product_rates SET 
    rate = $2
WHERE user_id = $1
RETURNING *;


-- name: GetProductRateByProductId :one
SELECT * FROM product_rates
WHERE product_id = $1 LIMIT 1;

-- name: GetProductRateByUserId :one
SELECT * FROM product_rates
WHERE user_id = $1 LIMIT 1;


-- name: DeleteProductRate :exec
DELETE FROM product_rates
WHERE id = $1;