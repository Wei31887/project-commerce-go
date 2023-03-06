// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: carts.sql

package sqlc

import (
	"context"
)

const createCart = `-- name: CreateCart :one
INSERT INTO carts (
  id
) VALUES (
  $1
)
RETURNING id
`

func (q *Queries) CreateCart(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, createCart, id)
	err := row.Scan(&id)
	return id, err
}

const deleteCart = `-- name: DeleteCart :exec
DELETE FROM carts
WHERE id = $1
`

func (q *Queries) DeleteCart(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCart, id)
	return err
}