# Architecture

## Overview

The application follows a layered architecture with clear separation of responsibilities.

```text
                  CLI
                   │
                   ▼
             Task Service
                   │
                   ▼
           Task Repository
                   │
                   ▼
             database/sql
                   │
                   ▼
                SQLite
```

Each layer communicates only with the layer immediately below it.

---

# Directory Structure

```text
task-manager/
│
├── cmd/
│   └── taskmanager/
│       └── main.go
│
├── docs/
│
├── internal/
│   ├── app/
│   ├── cli/
│   ├── database/
│   ├── models/
│   ├── repository/
│   └── service/
│
├── go.mod
└── go.sum
```

The project uses the `internal` directory to prevent packages from being imported by external applications.

---

# Application Entry Point

```
cmd/taskmanager/main.go
```

Responsibilities:

* create the application;
* handle startup errors;
* defer resource cleanup;
* execute the CLI.

The `main` package contains no business rules.

---

# Composition Root

```
internal/app
```

The `App` type is responsible for composing the application.

Responsibilities:

* create the database connection;
* create repositories;
* create services;
* expose application services;
* manage application resources.

Current structure:

```text
App
 ├── db
 └── TaskService
```

Future versions may also include:

* logger;
* configuration;
* metrics;
* cache.

---

# CLI Layer

```
internal/cli
```

Responsibilities:

* parse command-line arguments;
* validate user input;
* call the appropriate service;
* display results to the user.

The CLI never communicates directly with the database.

Current commands:

* add
* list
* done
* remove

---

# Service Layer

```
internal/service
```

The Service Layer contains the application's business rules.

Responsibilities:

* implement use cases;
* validate input;
* enforce business rules;
* translate infrastructure errors into domain errors.

Examples:

```
CreateTask()

ListTasks()

CompleteTask()

DeleteTask()
```

The Service does not know SQL statements or database details.

---

# Repository Layer

```
internal/repository
```

Responsibilities:

* execute SQL statements;
* map rows to domain models;
* persist data.

The Repository contains no business rules.

Typical operations:

* INSERT
* SELECT
* UPDATE
* DELETE

---

# Database Layer

```
internal/database
```

Responsibilities:

* open the SQLite database;
* configure the connection;
* create required tables;
* expose the configured `*sql.DB`.

The rest of the application never worries about database initialization.

---

# Domain Model

```
internal/models
```

The project currently contains a single domain entity.

```
Task
```

The model represents application data only.

Business rules belong to the Service.

Persistence belongs to the Repository.

---

# Error Flow

Infrastructure errors never leave the Service layer.

Example:

```text
SQLite

↓

sql.ErrNoRows

↓

Repository

↓

Service

↓

ErrTaskNotFound

↓

CLI

↓

"Task not found."
```

Each layer translates errors into the language of its own responsibility.

---

# Testing Strategy

The project contains two kinds of tests.

## Repository Tests

Validate:

* SQL;
* persistence;
* CRUD operations.

SQLite in-memory databases are used to keep tests isolated.

---

## Service Tests

Validate:

* business rules;
* use cases;
* domain errors.

The Service is tested through the Repository using an in-memory SQLite database.

---

# Architectural Principles

The project follows a few fundamental principles.

## Single Responsibility

Each package has one primary responsibility.

---

## Layer Isolation

Each layer communicates only with the layer directly below it.

---

## Business First

Business rules belong to the Service.

Infrastructure concerns belong to the Repository.

---

## Explicit Dependencies

Dependencies are created explicitly.

No dependency injection framework is used.

---

## Simplicity

The simplest solution that solves the current problem is preferred.

---

## YAGNI

Abstractions are introduced only when a real problem appears.

---

## Test Behavior

Tests verify observable behavior rather than implementation details.

---

# Dependency Flow

Dependencies always point downward.

```text
main
 │
 ▼
App
 │
 ▼
CLI
 │
 ▼
Service
 │
 ▼
Repository
 │
 ▼
Database
```

Lower layers never import higher layers.

---

# Why Cobra?

Cobra was chosen because it is the de facto standard library for building command-line applications in Go.

It provides:

* command hierarchy;
* automatic help generation;
* argument validation;
* shell completion;
* aliases;
* consistent command structure.

---

# Future Evolution

The current architecture is intentionally small.

Future commits will extend it with:

* structured logging (`log/slog`);
* configuration management;
* environment variables;
* Docker support;
* GitHub Actions;
* release automation.

These features should integrate naturally without changing the overall architecture.

---

# Summary

The architecture intentionally favors:

* clarity over cleverness;
* composition over unnecessary abstraction;
* small packages;
* explicit dependencies;
* idiomatic Go.

The goal is to provide a codebase that is easy to understand, easy to test and easy to evolve.
