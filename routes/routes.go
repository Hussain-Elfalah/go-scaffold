// Package routes registers all HTTP paths and wires them to controllers.
//
// Implement after controllers exist. main.go will call into this package.
package routes

// STEP 1 — Imports
//
//   import (
//       "net/http"
//
//       "go-scaffold/config"
//       "go-scaffold/controllers"
//   )

// STEP 2 — Create a function that returns an http.Handler
//
//   func NewMux(cfg *config.Config) http.Handler
//
// Using http.ServeMux (Go 1.22+): patterns can include method and path:
//
//   mux := http.NewServeMux()
//   health := &controllers.HealthController{Config: cfg}
//   mux.Handle("GET /api/v1/health", health)
//
// Or with a handler func:
//
//   mux.HandleFunc("GET /api/v1/health", controllers.HealthHandler(cfg))
//
// Go 1.21 and earlier: ServeMux did not support "GET /path" patterns.
// You would use mux.HandleFunc("/api/v1/health", ...) and check method inside the handler.
// This project targets modern Go — use the method-aware pattern if your version supports it.

// STEP 3 — Connect route to controller
//
// The path GET /api/v1/health must match exactly what you test with curl.
// The handler is whatever you defined in health_controller.go.
//
// Flow:
//   Client → GET /api/v1/health → ServeMux → HealthController → JSON response

// STEP 4 — Return the mux
//
//   return mux
//
// main.go passes this handler to http.ListenAndServe(cfg.Addr(), handler).

// STEP 5 — (Later) grouping
//
// When you add more routes, keep them in this file or split by version:
//   /api/v1/...
// Do not register routes in main.go — keeps main small and testable.
