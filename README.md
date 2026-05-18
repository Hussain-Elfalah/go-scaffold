# Go Backend Scaffold (Learning Guide)

This repository is a **comment-driven scaffold**. Each Go file tells you what to implement step by step. The goal is to build a standard-library HTTP server with:

```http
GET /api/v1/health
```

Expected JSON:

```json
{
  "status": "ok",
  "service": "backend-api",
  "version": "1.0.0"
}
```

No Gin, Echo, Fiber, Chi, or Gorilla — only `net/http` and the Go standard library.

---

## Recommended implementation order

Work through the files in this order so each layer only depends on what you already built:

| Order | File | Why |
| ----- | ---- | --- |
| 1 | `models/health.go` | Plain data types; no internal imports |
| 2 | `config/config.go` | Loads `APP_NAME`, `APP_VERSION`, `ENV`, `PORT` |
| 3 | `middleware/json.go` | Reusable JSON writer for handlers |
| 4 | `controllers/health_controller.go` | Builds `HealthResponse`, calls middleware |
| 5 | `routes/routes.go` | Registers `GET /api/v1/health` |
| 6 | `main.go` | Loads config, starts server on configured port |

Migrations (`migrations/*.sql`) are placeholders until you add a database. Skip them for the health endpoint exercise.

---

## How to run the server

1. Implement the files above until `go build` succeeds.
2. Set environment variables (Go does **not** read `.env` automatically):

   **PowerShell (manual):**

   ```powershell
   $env:APP_NAME="backend-api"
   $env:APP_VERSION="1.0.0"
   $env:ENV="development"
   $env:PORT="8080"
   go run .
   ```

   Or copy values from `.env` into your shell session before `go run .`.

3. You should see a log line that the server is listening on `:8080` (or your `PORT`).

---

## How to test with curl

With the server running:

```bash
curl -i http://localhost:8080/api/v1/health
```

You want:

- Status: `200 OK`
- Header: `Content-Type: application/json`
- Body: `status`, `service`, and `version` matching your config

Pretty-printed (optional):

```bash
curl -s http://localhost:8080/api/v1/health | python -m json.tool
```

---

## Common beginner mistakes

1. **Editing `.env` but not exporting variables** — `os.Getenv` only sees the process environment. Loading `.env` requires extra tooling or shell export.
2. **Wrong import path** — Use the module name from `go.mod` (`go-scaffold/...`), not a random GitHub path.
3. **Writing the response body before `WriteHeader`** — Set headers and status before encoding JSON.
4. **Forgotten struct tags** — Without `` `json:"service"` ``, JSON keys will be capitalized (`Service` instead of `service`).
5. **Hardcoding service name in the controller** — Use `config.AppName` so the response reflects `APP_NAME`.
6. **Route path mismatch** — Handler registered as `/health` but curl hits `/api/v1/health` → 404.
7. **Listen address format** — Use `":" + port` (e.g. `:8080`), not `port` alone.
8. **Building `main.go` first** — You will get import errors until other packages compile; follow the order table above.

---

## Project layout

```
.
├── main.go
├── go.mod
├── config/config.go
├── models/health.go
├── middleware/json.go
├── controllers/health_controller.go
├── routes/routes.go
└── migrations/          # SQL placeholders for later
```

Open each `.go` file and follow the `STEP` comments inside.
