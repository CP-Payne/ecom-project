package fileserver

import (
	"github.com/CP-Payne/ecommerce-server/internal/config"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes(apiCfg *config.ApiConfig) *chi.Mux {
	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./internal/api/fileserver/images"))
	apiCfg.Logger.Info("Serving files from ./internal/api/fileserver/images")
	r.Handle("/*", http.StripPrefix("/images/", fileServer))
	return r
}
