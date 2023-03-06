-- name: CreateOrder :one
INSERT INTO orders (
  id,
  user_id, 
  is_paied
) VALUES (
  $1, $2, $3
)
RETURNING *;


-- name: UpdateOrderInfo :one
UPDATE orders SET 
  total_amount = $2, 
  total_count = $3
WHERE id = $1
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;