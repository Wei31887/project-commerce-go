-- name: CreateProduct :one
INSERT INTO products (
  name,
  category_id,
  stock,
  sell,
  price,
  on_sell,
  description
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProduct :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListProductByType :many
SELECT * FROM products
WHERE category_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateProduct :one
UPDATE products SET 
    name = $2,
    stock = $3,
    sell = $4,
    category_id = $5,
    price = $6,
    on_sell = $7,
    description = $8,
    updated_at = $9
WHERE id = $1
RETURNING *;

-- name: UpdateProductStockSell :one
UPDATE products SET 
    stock = stock + sqlc.arg(count),
    sell = sell + sqlc.arg(sell_count),
    updated_at = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;