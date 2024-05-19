package api

import (
	"github.com/CP-Payne/ecommerce-server/internal/api/fileserver"
	"github.com/CP-Payne/ecommerce-server/internal/api/product"
	"github.com/CP-Payne/ecommerce-server/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(apiCfg *config.ApiConfig) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Mount("/images", fileserver.Routes(apiCfg))
	r.Mount("/products", product.Routes(apiCfg))

	return r
}
