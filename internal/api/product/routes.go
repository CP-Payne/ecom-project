package product

import "github.com/go-chi/chi/v5"

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", GetProducts)

	return r
}
