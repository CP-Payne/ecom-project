// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID          uuid.UUID
	Name        string
	Sku         string
	Price       float32
	Stock       int32
	Description sql.NullString
	CategoryID  uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductImage struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	ImageType sql.NullString
	ImageUrl  sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}
