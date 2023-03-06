-- name: CreateOrderItem :one
INSERT INTO order_items (
  order_id,
  count,
  product_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetOrderItemByOrderId :one
SELECT * FROM order_items
WHERE order_id = $1 LIMIT 1;

-- name: GetOrderItemByProductId :many
SELECT * FROM order_items
WHERE product_id = $1;

-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE id = $1;