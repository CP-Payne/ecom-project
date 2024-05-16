-- name: ListProducts :many
SELECT * FROM products WHERE (created_at, id) > ($1, $2)
ORDER BY created_at, id
LIMIT $3;