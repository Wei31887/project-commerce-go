// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: orders.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
  id,
  user_id, 
  is_paied
) VALUES (
  $1, $2, $3
)
RETURNING id, user_id, total_amount, total_count, is_paied, created_at
`

type CreateOrderParams struct {
	ID      uuid.UUID `json:"id"`
	UserID  int64     `json:"user_id"`
	IsPaied bool      `json:"is_paied"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder, arg.ID, arg.UserID, arg.IsPaied)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TotalAmount,
		&i.TotalCount,
		&i.IsPaied,
		&i.CreatedAt,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, id)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT id, user_id, total_amount, total_count, is_paied, created_at FROM orders
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, id uuid.UUID) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TotalAmount,
		&i.TotalCount,
		&i.IsPaied,
		&i.CreatedAt,
	)
	return i, err
}

const updateOrderInfo = `-- name: UpdateOrderInfo :one
UPDATE orders SET 
  total_amount = $2, 
  total_count = $3
WHERE id = $1
RETURNING id, user_id, total_amount, total_count, is_paied, created_at
`

type UpdateOrderInfoParams struct {
	ID          uuid.UUID `json:"id"`
	TotalAmount string    `json:"total_amount"`
	TotalCount  int64     `json:"total_count"`
}

func (q *Queries) UpdateOrderInfo(ctx context.Context, arg UpdateOrderInfoParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrderInfo, arg.ID, arg.TotalAmount, arg.TotalCount)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TotalAmount,
		&i.TotalCount,
		&i.IsPaied,
		&i.CreatedAt,
	)
	return i, err
}
