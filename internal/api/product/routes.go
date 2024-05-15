package product

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	return r
}
