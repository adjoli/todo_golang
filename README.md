# Task Manager

A Task Manager written in Go, available as both a CLI and an HTTP API.

This project was created as a practical study of the Go language and its ecosystem. Rather than focusing only on implementing a CRUD application, the project explores idiomatic Go, software architecture, testing, and engineering best practices through incremental development.

Every commit represents a complete learning milestone.

---

# Features

* Create, list, update, and remove tasks
* Mark tasks as completed
* SQLite persistence (pure Go, no CGO required)
* Command-line interface (Cobra)
* HTTP API (stdlib `net/http`)
* Structured logging (`log/slog`)
* Configuration via environment variables
* Automated tests
* Layered architecture

---

# Technology Stack

* Go
* SQLite (`modernc.org/sqlite`)
* `database/sql`
* `log/slog`
* `net/http`
* Cobra

---

# Project Structure

```text
task-manager/
│
├── cmd/
│   ├── taskmanager/
│   │   └── main.go              # CLI entrypoint
│   └── taskmanager-api/
│       └── main.go              # HTTP API entrypoint
│
├── data/
│   └── tasks.db                 # SQLite database (gitignored)
│
├── docs/
│   ├── architecture.md
│   ├── decisions.md
│   ├── diary.md
│   ├── glossary.md
│   ├── principles.md
│   └── roadmap.md
│
├── internal/
│   ├── api/                     # HTTP handlers and routes
│   │   ├── dto.go
│   │   ├── response.go
│   │   ├── routes.go
│   │   ├── server.go
│   │   └── tasks.go
│   │
│   ├── app/                     # Application bootstrap (composition root)
│   │   └── app.go
│   │
│   ├── cli/                     # Cobra command definitions
│   │   ├── add.go
│   │   ├── done.go
│   │   ├── errors.go
│   │   ├── list.go
│   │   ├── remove.go
│   │   ├── root.go
│   │   └── update.go
│   │
│   ├── config/                  # Config loading with env var support
│   │   ├── config.go
│   │   ├── config_test.go
│   │   ├── database.go
│   │   └── environment.go
│   │
│   ├── database/                # DB connection and schema creation
│   │   └── db.go
│   │
│   ├── logger/                  # Structured logger setup
│   │   └── logger.go
│   │
│   ├── models/                  # Domain models
│   │   ├── task.go
│   │   └── task_filter.go
│   │
│   ├── repository/              # Persistence layer
│   │   ├── sql.go
│   │   ├── task_queries.go
│   │   ├── task_repository.go
│   │   ├── task_repository_test.go
│   │   └── test_helpers.go
│   │
│   └── service/                 # Business logic layer
│       ├── errors.go
│       ├── inputs.go
│       ├── task_service.go
│       ├── task_service_test.go
│       └── test_helpers.go
│
├── AGENTS.md
├── LICENSE
├── README.md
├── go.mod
└── go.sum
```

---

# Architecture

```text
         CLI / API
              │
              ▼
         TaskService
              │
              ▼
      TaskRepository
              │
              ▼
        database/sql
              │
              ▼
            SQLite
```

Each layer has a single responsibility.

* **CLI / API** receive user input.
* **Service** implements business rules.
* **Repository** handles persistence.
* **SQLite** stores application data.

---

# Design Principles

This project follows several principles commonly adopted by the Go community.

* Simplicity over unnecessary abstraction.
* Explicit code over magic.
* Small packages with well-defined responsibilities.
* Business rules isolated from infrastructure.
* Test behavior instead of implementation.
* Introduce abstractions only when there is a real need (YAGNI).

---

# CLI Commands

## Add a task

```bash
taskmanager add "Study Go"
```

## List tasks

```bash
taskmanager list
```

## List all tasks (including completed)

```bash
taskmanager list --all
```

## Update a task

```bash
taskmanager update 1 "Study Go deeply"
```

## Complete a task

```bash
taskmanager done 1
```

## Remove a task

```bash
taskmanager remove 1
```

Alias:

```bash
taskmanager rm 1
```

---

# Running the Project

Clone the repository:

```bash
git clone <repository-url>
```

Enter the project directory:

```bash
cd task-manager
```

## Run the CLI

```bash
go run ./cmd/taskmanager
```

## Run the HTTP API

```bash
go run ./cmd/taskmanager-api
```

The API listens on `:8080`. Available endpoints:

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Health check |
| GET | `/tasks` | List all tasks |
| POST | `/tasks` | Create a task |
| GET | `/tasks/{id}` | Get a task by ID |
| PUT | `/tasks/{id}` | Update a task |

---

# Running Tests

Run all tests:

```bash
go test ./...
```

Coverage:

```bash
go test -cover ./...
```

Static analysis:

```bash
go vet ./...
```

Format the project:

```bash
go fmt ./...
```

---

# Configuration

The database path can be overridden via environment variable:

```bash
export TASKMANAGER_DB="custom/path/tasks.db"
```

Default: `data/tasks.db`

---

# Learning Roadmap

## Completed

* Project structure
* SQLite integration
* Repository Pattern
* Automated tests
* Service Layer
* Command-Line Interface
* Structured logging (`log/slog`)
* Configuration management (env vars)
* HTTP API (`net/http`)

## Planned

* Docker
* GitHub Actions
* Release automation
* Project review and final refactoring

---

# Documentation

Additional documentation is available in the `docs` directory.

* `architecture.md`
* `decisions.md`
* `diary.md`
* `glossary.md`
* `principles.md`
* `roadmap.md`

---

# Project Goals

This repository has two primary goals.

1. Build a complete application in Go.
2. Serve as a long-term reference for idiomatic Go development.

The focus is not only on writing code, but also on understanding the reasoning behind architectural decisions, testing strategies, and software engineering practices.

---

# License

This project is licensed under the MIT License.
