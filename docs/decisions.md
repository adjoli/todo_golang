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
