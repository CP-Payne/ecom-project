package product

import (
	"encoding/json"
	"github.com/CP-Payne/ecommerce-server/internal/config"
	"github.com/CP-Payne/ecommerce-server/internal/database"
	"github.com/CP-Payne/ecommerce-server/internal/models"
	"github.com/CP-Payne/ecommerce-server/pkg/encoding"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	apiCfg *config.ApiConfig
}

func NewProductHandler(apiCfg *config.ApiConfig) *Handler {
	return &Handler{apiCfg: apiCfg}
}

type PaginatedResponse struct {
	Products   []models.Product `json:"products"`
	NextCursor string           `json:"next_cursor"`
}

func (h *Handler) getProductsPagination(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var lastCreatedAt time.Time
	var lastID uuid.UUID
	var limit int
	var err error

	cursor := r.URL.Query().Get("cursor")

	if cursor != "" {
		lastCreatedAt, lastID, err = encoding.DecodeCursor(cursor)
		if err != nil {
			http.Error(w, "Invalid cursor", http.StatusBadRequest)
			return
		}
	} else {
		lastCreatedAt = time.Time{} // Default value for initial call
		lastID = uuid.Nil           // Default value for initial call
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}
	} else {
		limit = 10 // Default value
	}

	// Call the ListProducts function to get paginated products
	products, err := h.apiCfg.DB.ListProducts(ctx, database.ListProductsParams{
		CreatedAt: lastCreatedAt,
		ID:        lastID,
		Limit:     int32(limit),
	})
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
	}

	var nextCursor string
	if len(products) > 0 {
		lastProduct := products[len(products)-1]
		nextCursor = encoding.EncodeCursor(lastProduct.CreatedAt, lastProduct.ID)
	}

	// Paginated response
	response := PaginatedResponse{
		Products:   models.DatabaseProductsToProducts(products),
		NextCursor: nextCursor,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
