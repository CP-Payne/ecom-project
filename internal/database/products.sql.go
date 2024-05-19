// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: products.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const listProducts = `-- name: ListProducts :many
SELECT id, name, sku, price, stock, description, category_id, created_at, updated_at FROM products
WHERE (created_at > $1 OR (created_at = $1 AND id > $2))
ORDER BY created_at, id
LIMIT $3
`

type ListProductsParams struct {
	CreatedAt time.Time
	ID        uuid.UUID
	Limit     int32
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.CreatedAt, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Sku,
			&i.Price,
			&i.Stock,
			&i.Description,
			&i.CategoryID,
			&i.CreatedAt,
			&i.UpdatedAt,
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
