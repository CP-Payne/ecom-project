package product

import (
	"github.com/CP-Payne/ecommerce-server/internal/config"
	"github.com/go-chi/chi/v5"
)

func Routes(apiCfg *config.ApiConfig) *chi.Mux {
	r := chi.NewRouter()
	ph := NewProductHandler(apiCfg)
	r.Get("/", ph.getProductsPagination)

	return r
}
