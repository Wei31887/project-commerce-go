-- name: CreateProductCategory :one
INSERT INTO product_categories (
  name, sort
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetProductCategory :one
SELECT * FROM product_categories
WHERE id = $1 LIMIT 1;

-- name: ListProductCategory :many
SELECT * FROM product_categories
WHERE id = sqlc.arg(category_id)
ORDER BY id;

-- name: ListAllProductCategory :many
SELECT * FROM product_categories
ORDER BY id;

-- name: UpdateProductCategory :one
UPDATE product_categories SET 
    name = $2, sort = $3
WHERE id = $1
RETURNING *;

-- name: DeleteProductCategory :exec
DELETE FROM product_categories
WHERE id = $1;