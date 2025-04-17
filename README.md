# Go Design Patterns

This repository contains implementations of various design patterns in Go programming language. The examples are organized by pattern type and demonstrate how to apply these patterns in Go projects.

## Patterns Included

### Behavioral Patterns

#### Memento
- **Simple Memento**: Basic implementation of the Memento pattern that allows capturing and storing an object's internal state
- **Memento with Undo/Redo**: Advanced implementation with undo and redo functionality

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
└── memento/
    ├── simple-memento/           # Basic implementation of Memento pattern
    └── memento-undo-redo/        # Memento with undo/redo functionality
```

## Requirements

- Go 1.16+
