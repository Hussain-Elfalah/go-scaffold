package main

import (
	"log"
	"net/http"

	"go-scaffold/config"
	"go-scaffold/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	handler := routes.NewMux(cfg)
	addr := cfg.Addr()
	log.Printf("listening on http://localhost%s, health check at http://localhost%s/api/v1/health", addr, addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
