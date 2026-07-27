# Decisions

This document records the most important architectural decisions made during the project.

---

# ADR-001

## Repository Pattern

### Context

The application needed to isolate persistence logic.

### Decision

Introduce a Repository layer responsible for SQL operations.

### Consequences

* SQL is isolated.
* Business rules remain independent from persistence.
* Testing becomes easier.

---

# ADR-002

## Service Layer

### Context

Business rules started to grow beyond simple CRUD operations.

### Decision

Create a Service layer responsible for implementing use cases.

### Consequences

* Business rules became centralized.
* The Repository remained focused on persistence.

---

# ADR-003

## Domain Errors

### Context

Infrastructure errors should not be exposed outside the Service.

### Decision

Translate infrastructure errors into domain errors.

Example:

```text id="0j25tt"
sql.ErrNoRows

↓

ErrTaskNotFound
```

### Consequences

The CLI works only with domain concepts.

---

# ADR-004

## Composition Root

### Context

The application needed a single place to assemble dependencies.

### Decision

Introduce the `App` type.

### Consequences

* Centralized initialization.
* Cleaner `main.go`.
* Easier future expansion.

---

# ADR-005

## Cobra for CLI

### Context

The application required a maintainable command-line interface.

### Decision

Use Cobra.

### Consequences

* Standard command structure.
* Automatic help generation.
* Argument validation.
* Aliases.
* Shell completion support.

---

# ADR-006

## Delayed Abstractions

### Context

Several opportunities appeared to introduce interfaces and helper types.

### Decision

Only introduce abstractions after a concrete need arises.

### Consequences

* Simpler code.
* Less accidental complexity.
* Easier maintenance.

---

# ADR-007

## Multi-Driver Database Support

### Context

The application initially supported only SQLite. As the project evolved, the need to support Postgres arose — both for production use cases and to learn how Go applications handle multiple database drivers through `database/sql`.

### Decision

Introduce a `dialect` interface in the repository layer that abstracts SQL differences between SQLite and Postgres. Each driver provides its own implementation (`sqliteDialect`, `postgresDialect`), and the repository receives the appropriate dialect at construction time.

Driver selection is done via environment variables (`TASKMANAGER_DB_DRIVER`, `TASKMANAGER_DB_DSN`), with SQLite remaining the default for backwards compatibility.

Postgres uses the `pgx` driver, registered as `"postgres"` via `sql.Register` in an `init()` function.

### Consequences

* Repository code remains driver-agnostic — SQL differences are encapsulated in the dialect.
* Adding a new driver requires only a new dialect implementation and schema branch.
* `Create` operations branch internally: SQLite uses `LastInsertId()`, Postgres uses `RETURNING id`.
* Tests continue to use in-memory SQLite — no external services required.

---

# ADR-008

## .env File Support

### Context

Configuration is done via environment variables, but requiring users to `export` variables manually is inconvenient for local development. A `.env` file provides a zero-friction way to configure the application without polluting the shell environment.

### Decision

Load `.env` files automatically using `github.com/joho/godotenv` at the start of `config.New()`. If no `.env` file exists, the application continues normally.

### Consequences

* `loadEnvironment()` remains unchanged — it still reads from `os.LookupEnv`.
* Real environment variables always override `.env` values.
* `.env` is gitignored — secrets are never committed.
* `.env.example` is tracked — documents expected variables for new developers.
