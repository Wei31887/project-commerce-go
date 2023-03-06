-- name: CreateComment :one
INSERT INTO comments (
  user_id, product_id, content
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateComment :one
UPDATE comments SET 
    content = $2, updated_at = $3
WHERE id = $1
RETURNING *;

-- name: UpdateCommentStatus :exec
UPDATE comments SET 
    status = $2
WHERE id = $1;


-- name: GetCommentByProductId :one
SELECT * FROM comments
WHERE product_id = $1 LIMIT 1;

-- name: GetCommentByUserId :one
SELECT * FROM comments
WHERE user_id = $1 LIMIT 1;


-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;