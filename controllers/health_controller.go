package controllers

import (
	"log"
	"net/http"

	"go-scaffold/config"
	"go-scaffold/middleware"
	"go-scaffold/models"
)

type HealthController struct {
	Config *config.Config
}

func (c *HealthController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	response := models.HealthResponse{
		Status:  "ok",
		Service: c.Config.AppName,
		Version: c.Config.AppVersion,
	}

	if err := middleware.WriteJSON(w, http.StatusOK, response); err != nil {
		log.Printf("write health response: %v", err)
	}
}
