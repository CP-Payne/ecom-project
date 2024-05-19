package models

import (
	"database/sql"
	"github.com/CP-Payne/ecommerce-server/internal/database"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Sku         string
	Price       float32
	Stock       int32
	Description string
	CategoryID  uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func DatabaseProductToProduct(product database.Product) Product {
	return Product{
		ID:          product.ID,
		Name:        product.Name,
		Sku:         product.Sku,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: SqlNulStringToString(product.Description),
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func DatabaseProductsToProducts(products []database.Product) []Product {
	result := make([]Product, len(products))
	for i, product := range products {
		result[i] = DatabaseProductToProduct(product)
	}
	return result
}

func SqlNulStringToString(sqlString sql.NullString) string {
	if sqlString.Valid {
		return sqlString.String
	}
	return ""
}
