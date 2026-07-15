# Principles

This document collects the principles that guided the development of this project.

These principles are more important than any specific implementation because they can be applied to future projects.

---

# Simplicity over unnecessary abstraction

★★★★★ Fundamental

Always choose the simplest solution that correctly solves the current problem.

Avoid introducing abstractions before they become necessary.

---

# Clarity over cleverness

★★★★★ Fundamental

Code is read far more often than it is written.

Prefer code that is immediately understandable.

---

# Explicit is better than implicit

★★★★★ Fundamental

Dependencies, data flow and responsibilities should be visible.

Avoid hidden behavior.

---

# Business rules belong to the Service

★★★★★ Fundamental

The Service implements use cases.

It should not know SQL.

---

# Persistence belongs to the Repository

★★★★★ Fundamental

The Repository stores and retrieves data.

It should not implement business rules.

---

# The main package composes the application

★★★★★ Fundamental

The entry point should only assemble dependencies and start the application.

---

# Test behavior, not implementation

★★★★★ Fundamental

Tests should validate observable behavior.

Implementation details are free to change.

---

# Every package should have one responsibility

★★★★★ Fundamental

Small, focused packages are easier to understand and maintain.

---

# Introduce abstractions only when necessary

★★★★★ Fundamental

Abstractions should emerge from repeated patterns.

Not from anticipation.

---

# Accept interfaces, return structs

★★★★☆ Community Practice

Interfaces represent behavior.

They should usually be owned by the consumer.

---

# The interface belongs to the consumer

★★★★☆ Community Practice

The package that depends on a behavior defines the interface.

Not the package that implements it.

---

# Keep commands thin

★★★★☆ Community Practice

CLI commands should:

* validate arguments;
* call the Service;
* display results.

Nothing more.

---

# Every commit should represent a complete milestone

★★★★★ Fundamental

Each commit should leave the project in a working state.

---

# Documentation is part of the project

★★★★★ Fundamental

Documentation evolves together with the code.

Outdated documentation is considered a bug.

---

# Continuous improvement

★★★★★ Fundamental

Architecture evolves.

Refactoring is part of software development.

The goal is not perfection on day one, but continuous improvement.
