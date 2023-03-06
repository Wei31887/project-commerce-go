-- name: CreateUserDeliver :one
INSERT INTO user_delivers (
  user_id, address
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserDeliver :one
SELECT * FROM user_delivers
WHERE id = $1 LIMIT 1;

-- name: ListUserDeliverByUserId :many
SELECT * FROM user_delivers
WHERE user_id = $1
ORDER BY user_id;

-- name: UpdateUserDeliver :one
UPDATE user_delivers SET 
  address = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUserDeliver :exec
DELETE FROM user_delivers
WHERE id = $1;