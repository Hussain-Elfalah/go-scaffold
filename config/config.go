// Package config loads application settings from the environment.
//
// Implement this file first (or second, after models). Every other layer
// depends on knowing APP_NAME, APP_VERSION, ENV, and PORT.
package config

// STEP 1 — Import what you need
//
//   - "os" to read environment variables with os.Getenv
//   - optionally "strconv" if you parse PORT as an integer later
//     (for this project, keeping PORT as a string like "8080" is fine)
//
// Do NOT import a third-party .env loader yet. The standard library does not
// read .env files automatically. For local dev you will either:
//   a) export variables in your shell before `go run .`, or
//   b) load .env yourself later (optional stretch goal), or
//   c) use a tool that injects .env when you run the app
//
// Beginner tip: if APP_NAME is empty at runtime, you probably forgot to
// set env vars — editing .env alone does not load them into Go.

// STEP 2 — Define a Config struct
//
// Create a struct that holds all settings the app needs in one place:
//
//   type Config struct {
//       AppName    string // maps to APP_NAME  — used in health JSON as "service"
//       AppVersion string // maps to APP_VERSION — used in health JSON as "version"
//       Env        string // maps to ENV        — e.g. development, production
//       Port       string // maps to PORT       — e.g. "8080" (include ":" when listening)
//   }
//
// Use exported field names (capitalized) so other packages can read them.

// STEP 3 — Write a Load function
//
//   func Load() (*Config, error) {
//       // read each variable, apply defaults, return &Config{...}, nil
//   }
//
// For each setting, call os.Getenv("KEY").
// If the value is empty (""), substitute a sensible default:
//
//   APP_NAME     default: "backend-api"
//   APP_VERSION  default: "1.0.0"
//   ENV          default: "development"
//   PORT         default: "8080"
//
// Return a pointer to Config so callers can pass it around without copying.
// Returning error is good practice even if you always return nil today —
// you might add validation later (e.g. invalid PORT).

// STEP 4 — (Optional) helper for listen address
//
// When you call http.ListenAndServe in main.go, you need a host:port string.
// Convention: ":" + port listens on all interfaces, e.g. ":8080".
// You can add a small method on Config:
//
//   func (c *Config) Addr() string {
//       return ":" + c.Port
//   }
//
// Or build the string in main.go — either is fine.

// STEP 5 — Verify mentally
//
// After implementing, a quick test in main (temporary) could log cfg.AppName.
// Remove debug prints before committing.
