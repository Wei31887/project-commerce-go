// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: users.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  user_name, hashed_password, email, full_name
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, user_name, hashed_password, email, full_name, gender, user_rank, created_at, changed_password_at
`

type CreateUserParams struct {
	UserName       string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
	FullName       string `json:"full_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserName,
		arg.HashedPassword,
		arg.Email,
		arg.FullName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.HashedPassword,
		&i.Email,
		&i.FullName,
		&i.Gender,
		&i.UserRank,
		&i.CreatedAt,
		&i.ChangedPasswordAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, user_name, hashed_password, email, full_name, gender, user_rank, created_at, changed_password_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.HashedPassword,
		&i.Email,
		&i.FullName,
		&i.Gender,
		&i.UserRank,
		&i.CreatedAt,
		&i.ChangedPasswordAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, user_name, hashed_password, email, full_name, gender, user_rank, created_at, changed_password_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.HashedPassword,
		&i.Email,
		&i.FullName,
		&i.Gender,
		&i.UserRank,
		&i.CreatedAt,
		&i.ChangedPasswordAt,
	)
	return i, err
}

const updateUserInfo = `-- name: UpdateUserInfo :one
UPDATE users SET 
  full_name = $2, 
  email = $3,
  gender = $4 
WHERE id = $1
RETURNING id, user_name, hashed_password, email, full_name, gender, user_rank, created_at, changed_password_at
`

type UpdateUserInfoParams struct {
	ID       int64         `json:"id"`
	FullName string        `json:"full_name"`
	Email    string        `json:"email"`
	Gender   sql.NullInt16 `json:"gender"`
}

func (q *Queries) UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserInfo,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.Gender,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.HashedPassword,
		&i.Email,
		&i.FullName,
		&i.Gender,
		&i.UserRank,
		&i.CreatedAt,
		&i.ChangedPasswordAt,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :one
UPDATE users SET 
  hashed_password = $2
WHERE id = $1
RETURNING id, user_name, hashed_password, email, full_name, gender, user_rank, created_at, changed_password_at
`

type UpdateUserPasswordParams struct {
	ID             int64  `json:"id"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserPassword, arg.ID, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.HashedPassword,
		&i.Email,
		&i.FullName,
		&i.Gender,
		&i.UserRank,
		&i.CreatedAt,
		&i.ChangedPasswordAt,
	)
	return i, err
}
