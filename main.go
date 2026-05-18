// Application entry point.
//
// Implement this file LAST, after config, models, middleware, controllers, and routes.
// Until then, `go run .` will not build — that is expected.
package main

// STEP 1 — Imports
//
//   import (
//       "log"
//       "net/http"
//
//       "go-scaffold/config"
//       "go-scaffold/routes"
//   )

// STEP 2 — func main()
//
// This is where the process starts. Typical startup order:
//
//   1. Load configuration
//        cfg, err := config.Load()
//        if err != nil { log.Fatal(err) }
//
//   2. Build the router
//        handler := routes.NewMux(cfg)
//
//   3. Start the HTTP server
//        addr := ":" + cfg.Port   // or cfg.Addr() if you added that helper
//        log.Printf("listening on %s", addr)
//        if err := http.ListenAndServe(addr, handler); err != nil {
//            log.Fatal(err)
//        }
//
// ListenAndServe blocks forever until the process exits. It is a simple
// production-style server for learning; later you might use http.Server
// with graceful shutdown.

// STEP 3 — Environment variables
//
// config.Load() reads os.Getenv. Before running locally, set variables from .env:
//
//   PowerShell:
//     Get-Content .env | ForEach-Object {
//       if ($_ -match '^\s*([^#][^=]+)=(.*)$') {
//         [Environment]::SetEnvironmentVariable($matches[1].Trim(), $matches[2].Trim(), 'Process')
//       }
//     }
//     go run .
//
//   Or set manually:
//     $env:APP_NAME="backend-api"; $env:PORT="8080"; go run .
//
// STEP 4 — What success looks like
//
//   - Terminal shows listening on :8080 (or your PORT)
//   - curl http://localhost:8080/api/v1/health returns the JSON health payload
//
// No frameworks: only net/http from the standard library wires everything together.
