-- name: CreateUser :one
INSERT INTO users (
  user_name, hashed_password, email, full_name
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUserInfo :one
UPDATE users SET 
  full_name = $2, 
  email = $3,
  gender = $4 
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users SET 
  hashed_password = $2
WHERE id = $1
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;