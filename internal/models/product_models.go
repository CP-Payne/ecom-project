package models

import (
	"database/sql"
	"github.com/CP-Payne/ecommerce-server/internal/database"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Sku         string    `json:"sku"`
	Price       float32   `json:"price"`
	Stock       int32     `json:"stock"`
	Description string    `json:"description"`
	CategoryID  uuid.UUID `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
