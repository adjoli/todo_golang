# Glossary

This glossary contains the main concepts introduced during the project.

---

# App

The object responsible for composing the application and managing shared resources.

---

# CLI

Command-Line Interface.

The layer responsible for interacting with the user.

---

# Cobra

The de facto standard Go library for building command-line applications.

---

# Composition Root

The place where all dependencies are created and connected.

---

# Context

An object used to propagate cancellation, deadlines and request-scoped values.

---

# CRUD

Create, Read, Update and Delete.

Basic persistence operations.

---

# Domain Error

An error that represents a business concept instead of an infrastructure detail.

Example:

```text id="2uhov4"
ErrTaskNotFound
```

---

# Repository

The layer responsible for persistence.

It communicates with the database.

---

# Service

The layer responsible for business rules and use cases.

---

# SQLite

An embedded relational database used by the project.

---

# Table-Driven Test

A Go testing pattern where multiple scenarios are executed through a table of test cases.

---

# Use Case

A business operation offered by the application.

Examples:

* CreateTask
* ListTasks
* CompleteTask
* DeleteTask

---

# YAGNI

"You Aren't Gonna Need It"

A principle that recommends delaying abstractions until they become necessary.

---

# KISS

"Keep It Simple, Stupid"

A principle that favors the simplest solution that solves the problem correctly.

---

# Single Responsibility Principle

Each package or type should have one clear responsibility.

---

# Zero Value

The default value automatically assigned to variables in Go.

One of the language's defining characteristics.
