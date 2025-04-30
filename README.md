# Go Design Patterns

This repository contains implementations of various design patterns in Go programming language. The examples are organized by pattern type and demonstrate how to apply these patterns in Go projects.

## Patterns Included

### Behavioral Patterns

#### Memento
- **Simple Memento**: Basic implementation of the Memento pattern that allows capturing and storing an object's internal state
- **Memento with Undo/Redo**: Advanced implementation with undo and redo functionality

#### State
- **State Classic**: Implementation of the State pattern that allows an object to alter its behavior when its internal state changes. The example demonstrates a light switch that can be in On or Off states.
- **State Rules**: Alternative implementation using a rules-based approach with a finite state machine. Demonstrates a phone system with states like OffHook, Connecting, Connected, OnHold, and OnHook.

### Creational Patterns

#### Builder
- **Functional Builder**: Implementation of the Builder pattern using functional programming concepts in Go

## Getting Started

Each pattern implementation is contained in its own directory with a dedicated `main.go` file which demonstrates how to use the pattern.

To run any example:

```bash
cd <pattern-directory>
go run .
```

## Project Structure

```
.
├── memento/
│   ├── simple-memento/           # Basic implementation of Memento pattern
│   └── memento-undo-redo/        # Memento with undo/redo functionality
├── state/
│   ├── state-classic/            # Classic implementation of State pattern
│   └── state-rules/              # Rules-based implementation of State pattern
└── functional-builder/           # Functional approach to the Builder pattern
```

## Requirements

- Go 1.16+
