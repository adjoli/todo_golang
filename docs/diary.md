# Diary

This document records the evolution of the project.

Each commit represents a complete learning milestone.

---

# Commit 01 — Project Setup

## Goal

Create the initial project structure.

## Implemented

* Go module
* Initial directory structure
* First executable
* Git repository

## Concepts Learned

* Modules
* Packages
* Project organization
* `go mod init`

## Takeaways

* Every project starts simple.
* Organize the project before adding complexity.

---

# Commit 02 — SQLite Integration

## Goal

Persist application data.

## Implemented

* SQLite database
* Connection management
* Table creation
* Database package

## Concepts Learned

* `database/sql`
* Drivers
* Connection Pool
* Context
* Resource management

## Takeaways

* Prefer the Standard Library.
* Always close resources.
* Keep database initialization isolated.

---

# Commit 03 — Repository Layer

## Goal

Separate persistence from business logic.

## Implemented

* Repository Pattern
* CRUD operations
* SQL queries
* Data mapping

## Concepts Learned

* Repository Pattern
* `QueryContext`
* `QueryRowContext`
* `ExecContext`
* `Scan`
* Parameterized SQL

## Takeaways

* The Repository speaks the language of the database.
* Keep SQL isolated.

---

# Commit 04 — Automated Tests

## Goal

Introduce automated testing.

## Implemented

* Repository tests
* Test helpers
* In-memory SQLite
* Table-Driven Tests

## Concepts Learned

* package `testing`
* `t.Helper`
* `t.Cleanup`
* Test isolation

## Takeaways

* Test behavior.
* Small tests are easier to maintain.
* Fast tests encourage frequent execution.

---

# Commit 05 — Service Layer

## Goal

Move business rules out of the Repository.

## Implemented

* Service Layer
* Domain errors
* Use cases
* Service tests

## Concepts Learned

* Use Cases
* Domain Errors
* Error translation
* Layer separation

## Takeaways

* The Service speaks the language of the domain.
* Infrastructure errors should not leak.
* Business rules belong in the Service.

---

# Commit 06 — Command-Line Interface

## Goal

Build a complete command-line interface.

## Implemented

* Cobra integration
* Root command
* add
* list
* done
* remove
* Error handling
* Application composition

## Concepts Learned

* Cobra
* Commands
* Aliases
* Argument validation
* Composition Root

## Takeaways

* The CLI is only an interface.
* Business rules remain in the Service.
* Commands should remain thin.

---

# Project Evolution

```text id="cwk0b9"
Commit 01
        │
        ▼
Project Structure
        │
        ▼
Commit 02
        │
        ▼
SQLite
        │
        ▼
Commit 03
        │
        ▼
Repository
        │
        ▼
Commit 04
        │
        ▼
Tests
        │
        ▼
Commit 05
        │
        ▼
Service
        │
        ▼
Commit 06
        │
        ▼
CLI
```

---

# Current Status

The application is now fully functional from the command line.

Future commits will focus on improving infrastructure and developer experience rather than adding core functionality.
