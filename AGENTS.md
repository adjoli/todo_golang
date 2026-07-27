# AGENTS.md

## Build & Run

```bash
# Run the CLI
go run ./cmd/taskmanager

# Run the HTTP API server (listens on :8080)
go run ./cmd/taskmanager-api
```

## Verify

```bash
go vet ./...
go test ./...
go fmt ./...
```

## Project Structure

- `cmd/taskmanager/` — CLI entrypoint (Cobra)
- `cmd/taskmanager-api/` — HTTP API entrypoint (stdlib net/http)
- `internal/app/` — Application bootstrap, wires config, db, repo, service
- `internal/config/` — Config loading with env var support
- `internal/cli/` — Cobra command definitions
- `internal/api/` — HTTP handlers and routes
- `internal/service/` — Business logic layer
- `internal/repository/` — Persistence layer (SQLite or Postgres via `database/sql`)
- `internal/database/` — DB connection and schema creation
- `internal/models/` — Domain models
- `internal/logger/` — Structured logger setup (`slog`)
- `data/tasks.db` — SQLite database file (gitignored)

## Key Facts

- **Module**: `github.com/adjoli/todo_golang`
- **Go version**: 1.26.4
- **SQLite driver**: `modernc.org/sqlite` (pure Go, no CGO required)
- **Postgres driver**: `github.com/jackc/pgx/v5` (registered as `"postgres"` via `sql.Register`)
- **Database driver**: select via `TASKMANAGER_DB_DRIVER` env var (`"sqlite"` default, or `"postgres"`)
- **Database DSN**: select via `TASKMANAGER_DB_DSN` env var (`data/tasks.db` default for SQLite)
- **Legacy env var**: `TASKMANAGER_DB` works as alias for `TASKMANAGER_DB_DSN`
- **Tests use in-memory SQLite** (`:memory:`) — no external services needed
- **`.env` support**: application loads `.env` from working directory via `godotenv`; real env vars always override
- **`.env.example`** is tracked — copy to `.env` and edit; `.env` itself is gitignored
- **`requests.http`** is gitignored but present — likely HTTP client test file, ignore it

## Architecture

Layered: CLI/API → Service → Repository → database/sql → SQLite or Postgres. The `app.App` struct is the composition root that holds config, db, logger, and service. Both CLI and API entrypoints create an `App` first, then hand off to their respective layer.

## Conventions

- All internal packages live under `internal/` (enforced by Go toolchain)
- Table-driven tests with `t.Helper()` and `t.Cleanup()`
- Error wrapping with `fmt.Errorf("context: %w", err)`
- Config via environment variables only (no config files)
- Repository uses `dialect` interface for multi-driver SQL adaptation
