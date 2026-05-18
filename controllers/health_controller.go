// Package controllers contains HTTP handlers (business-facing entry points per route).
//
// Implement after: models, config, middleware.
package controllers

// STEP 1 — Imports you will need
//
//   import (
//       "net/http"
//
//       "go-scaffold/config"      // adjust if your module path differs
//       "go-scaffold/middleware"
//       "go-scaffold/models"
//   )
//
// Check go.mod for the exact module path at the top of the file.

// STEP 2 — Decide how the handler gets config
//
// Option A (recommended for learning): struct with config dependency
//
//   type HealthController struct {
//       Config *config.Config
//   }
//
//   func (c *HealthController) ServeHTTP(w http.ResponseWriter, r *http.Request) { ... }
//
// Option B: package-level function that closes over config
//
//   func HealthHandler(cfg *config.Config) http.HandlerFunc {
//       return func(w http.ResponseWriter, r *http.Request) { ... }
//   }
//
// Both work with net/http. Option B is common for a single health route.

// STEP 3 — Implement the handler logic
//
// Inside the handler:
//   1. (Optional) check r.Method == http.MethodGet — return 405 if wrong method
//   2. Build a models.HealthResponse:
//        Status:  "ok"
//        Service: cfg.AppName    (from config)
//        Version: cfg.AppVersion (from config)
//   3. Call middleware.WriteJSON(w, http.StatusOK, response)
//   4. If WriteJSON returns an error, log it — the client may already have partial data
//
// Do not hardcode "backend-api" in the controller — read from config so
// changing .env changes the response.

// STEP 4 — Keep handlers thin
//
// Health check has no database call yet. Later handlers might call services/repos;
// for now, assemble the struct and return JSON. That is the whole handler.
