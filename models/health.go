// Package models holds data shapes shared across the application.
//
// Implement this file early — it has no dependencies on other project packages.
package models

// STEP 1 — Define HealthResponse
//
// This struct is the JSON body returned by GET /api/v1/health.
// Field names in Go use PascalCase; JSON keys use snake_case via struct tags.
//
//   type HealthResponse struct {
//       Status  string `json:"status"`  // always "ok" when the service is up
//       Service string `json:"service"` // your app name, e.g. from config AppName
//       Version string `json:"version"` // semver or build label from config AppVersion
//   }
//
// Why struct tags?
//   Without `json:"service"`, encoding/json would emit "Service" (capital S).
//   Tags tell the encoder the exact key names your API contract expects.
//
// Expected JSON after you wire everything up:
//
//   {
//     "status": "ok",
//     "service": "backend-api",
//     "version": "1.0.0"
//   }
//
// STEP 2 — Keep models dumb
//
// Do not put HTTP logic here. No http.ResponseWriter, no status codes.
// Models are plain data — controllers fill them, middleware encodes them.
