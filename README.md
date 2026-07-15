# Task Manager

A command-line Task Manager written in Go.

This project was created as a practical study of the Go language and its ecosystem. Rather than focusing only on implementing a CRUD application, the project explores idiomatic Go, software architecture, testing, and engineering best practices through incremental development.

Every commit represents a complete learning milestone.

---

# Features

Current features:

* Create tasks
* List tasks
* Mark tasks as completed
* Remove tasks
* SQLite persistence
* Command-line interface (Cobra)
* Automated tests
* Layered architecture

---

# Technology Stack

* Go
* SQLite
* Standard Library
* `database/sql`
* Cobra

---

# Project Structure

```text
task-manager/
│
├── cmd/
│   └── taskmanager/
│       └── main.go
│
├── data/
│   └── taskmanager.db
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
│   ├── app/
│   │   └── app.go
│   │
│   ├── cli/
│   │   ├── add.go
│   │   ├── done.go
│   │   ├── errors.go
│   │   ├── helpers.go
│   │   ├── list.go
│   │   ├── remove.go
│   │   └── root.go
│   │
│   ├── database/
│   │   ├── config.go
│   │   └── database.go
│   │
│   ├── models/
│   │   └── task.go
│   │
│   ├── repository/
│   │   ├── sql.go
│   │   ├── task_repository.go
│   │   ├── task_repository_test.go
│   │   └── test_helpers.go
│   │
│   └── service/
│       ├── errors.go
│       ├── task_service.go
│       ├── task_service_test.go
│       └── test_helpers.go
│
├── LICENSE
├── README.md
├── go.mod
└── go.sum
```

---

# Architecture

```text
                 CLI
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

* **CLI** receives user input.
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

# Current Commands

## Add a task

```bash
taskmanager add "Study Go"
```

---

## List tasks

```bash
taskmanager list
```

---

## Complete a task

```bash
taskmanager done 1
```

---

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

Run the application:

```bash
go run ./cmd/taskmanager
```

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

# Learning Roadmap

## Completed

* Project structure
* SQLite integration
* Repository Pattern
* Automated tests
* Service Layer
* Command-Line Interface

## Planned

* Structured logging
* Configuration management
* Environment variables
* Docker
* GitHub Actions
* Release automation

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

1. Build a complete command-line application in Go.
2. Serve as a long-term reference for idiomatic Go development.

The focus is not only on writing code, but also on understanding the reasoning behind architectural decisions, testing strategies, and software engineering practices.

---

# License

This project is licensed under the MIT License.
