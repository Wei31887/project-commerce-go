-- name: CreateCartItem :one
INSERT INTO cart_items (
  cart_id,
  count,
  product_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListCartItem :many
SELECT
	cart_items.cart_id,
	cart_items.product_id,
	cart_items.count,
	products.name,
	products.stock,
	products.sell,
	products.price,
	products.on_sell
FROM
	cart_items
	INNER JOIN products ON cart_items.product_id = products.id
WHERE
	cart_items.cart_id = $1
ORDER BY
	product_id ASC;

-- name: DeleteCartItem :exec
DELETE FROM cart_items
WHERE cart_id = $1 and product_id = $2;
