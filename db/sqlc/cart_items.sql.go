// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: cart_items.sql

package sqlc

import (
	"context"
)

const createCartItem = `-- name: CreateCartItem :one
INSERT INTO cart_items (
  cart_id,
  count,
  product_id
) VALUES (
  $1, $2, $3
)
RETURNING id, cart_id, product_id, count
`

type CreateCartItemParams struct {
	CartID    int64 `json:"cart_id"`
	Count     int64 `json:"count"`
	ProductID int64 `json:"product_id"`
}

func (q *Queries) CreateCartItem(ctx context.Context, arg CreateCartItemParams) (CartItem, error) {
	row := q.db.QueryRowContext(ctx, createCartItem, arg.CartID, arg.Count, arg.ProductID)
	var i CartItem
	err := row.Scan(
		&i.ID,
		&i.CartID,
		&i.ProductID,
		&i.Count,
	)
	return i, err
}

const deleteCartItem = `-- name: DeleteCartItem :exec
DELETE FROM cart_items
WHERE cart_id = $1 and product_id = $2
`

type DeleteCartItemParams struct {
	CartID    int64 `json:"cart_id"`
	ProductID int64 `json:"product_id"`
}

func (q *Queries) DeleteCartItem(ctx context.Context, arg DeleteCartItemParams) error {
	_, err := q.db.ExecContext(ctx, deleteCartItem, arg.CartID, arg.ProductID)
	return err
}

const listCartItem = `-- name: ListCartItem :many
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
	product_id ASC
`

type ListCartItemRow struct {
	CartID    int64  `json:"cart_id"`
	ProductID int64  `json:"product_id"`
	Count     int64  `json:"count"`
	Name      string `json:"name"`
	Stock     int64  `json:"stock"`
	Sell      int64  `json:"sell"`
	Price     string `json:"price"`
	OnSell    string `json:"on_sell"`
}

func (q *Queries) ListCartItem(ctx context.Context, cartID int64) ([]ListCartItemRow, error) {
	rows, err := q.db.QueryContext(ctx, listCartItem, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListCartItemRow{}
	for rows.Next() {
		var i ListCartItemRow
		if err := rows.Scan(
			&i.CartID,
			&i.ProductID,
			&i.Count,
			&i.Name,
			&i.Stock,
			&i.Sell,
			&i.Price,
			&i.OnSell,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
