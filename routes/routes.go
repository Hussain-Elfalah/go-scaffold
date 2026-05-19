package routes

import (
	"net/http"
	"github.com/go-chi/chi/v5"

	"go-scaffold/config"
	"go-scaffold/controllers"
)

func NewMux(cfg *config.Config) http.Handler {
	r := chi.NewRouter()
	health := &controllers.HealthController{Config: cfg}
	r.Route("/api/v1/", func(c chi.Router) {
		c.Get("/health", health.ServeHTTP)
	})
	return r
}
