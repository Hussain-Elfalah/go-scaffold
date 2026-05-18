// Package middleware provides small HTTP helpers used across handlers.
//
// Implement this file after models and config, before controllers.
package middleware

// STEP 1 — Imports
//
//   import (
//       "encoding/json"
//       "net/http"
//   )
//
// Stick to the standard library — no framework response helpers.

// STEP 2 — Create a JSON response helper
//
// Define a function that any handler can call to send JSON consistently:
//
//   func WriteJSON(w http.ResponseWriter, statusCode int, payload any) error
//
// Parameters:
//   w           — the ResponseWriter from the handler
//   statusCode  — e.g. http.StatusOK (200)
//   payload     — any value json.Marshal can encode (struct, map, etc.)
//
// STEP 3 — Set headers before writing the body
//
// Order matters in net/http:
//   1. w.Header().Set("Content-Type", "application/json")
//   2. w.WriteHeader(statusCode)   // must happen before body writes
//   3. json.NewEncoder(w).Encode(payload)
//
// Common mistake: calling w.Write() before WriteHeader — then status defaults to 200
// and you cannot change it.
//
// STEP 4 — Handle encode errors
//
// If json.Encoder.Encode returns an error:
//   - log it (fmt.Println or log.Println for now), and/or
//   - send http.StatusInternalServerError with a plain text fallback
//
// For learning, returning the error to the caller lets the handler decide.
//
// STEP 5 — Usage pattern (you will write this in the controller)
//
//   err := middleware.WriteJSON(w, http.StatusOK, response)
//   if err != nil { ... }
//
// Why a shared helper?
//   Every endpoint would otherwise duplicate Content-Type and encoding logic.
