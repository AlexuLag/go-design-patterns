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
- **State Switch**: Implementation using a switch statement to handle state transitions. Demonstrates a code lock system with states Locked, Failed, and Unlocked, showing how to handle state changes in a more procedural way.

### Creational Patterns

#### Builder
- **Functional Builder**: Implementation of the Builder pattern using functional programming concepts in Go

### Structural Patterns

#### Flyweight
- **Basic Flyweight**: Implementation of the Flyweight pattern that demonstrates how to share common parts of object state among multiple objects. The example shows how to efficiently store and reuse names in a user system, reducing memory usage by storing common strings only once.

## SOLID Principles

### Single Responsibility Principle (SRP)
- **Journal Example**: Implementation demonstrating the Single Responsibility Principle through a journal system. Shows how to separate concerns between journal management and persistence, with examples of both violating and following the principle.

### Open/Closed Principle (OCP)
- **Product Filter Example**: Implementation demonstrating the Open/Closed Principle through a product filtering system. Shows how to extend functionality without modifying existing code by using interfaces and specifications. Includes examples of color and size filtering, and how to combine specifications.

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
│   ├── state-rules/              # Rules-based implementation of State pattern
│   └── state-switch/             # Switch-based implementation of State pattern
├── flyweight/                    # Implementation of the Flyweight pattern
├── single-responsability/        # Implementation of Single Responsibility Principle
├── open-close/                   # Implementation of Open/Closed Principle
└── functional-builder/           # Functional approach to the Builder pattern
```

## Requirements

- Go 1.16+
