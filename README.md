# Go Design Patterns

This repository contains implementations of various design patterns in Go programming language. The examples are organized by pattern type and demonstrate how to apply these patterns in Go projects.

## Patterns Included

### Behavioral Patterns

#### Chain of Responsibility
- **Basic Chain**: Implementation of the Chain of Responsibility pattern that allows passing requests along a chain of handlers.
- **Chain with Steroids**: Advanced implementation with enhanced functionality and flexibility.

#### Observer
- **Basic Observer**: Implementation of the Observer pattern for event handling and notification.
- **Observer with Properties**: Enhanced implementation with property-based observation.

#### Memento
- **Simple Memento**: Basic implementation of the Memento pattern that allows capturing and storing an object's internal state
- **Memento with Undo/Redo**: Advanced implementation with undo and redo functionality

#### State
- **State Classic**: Implementation of the State pattern that allows an object to alter its behavior when its internal state changes. The example demonstrates a light switch that can be in On or Off states.
- **State Rules**: Alternative implementation using a rules-based approach with a finite state machine. Demonstrates a phone system with states like OffHook, Connecting, Connected, OnHold, and OnHook.
- **State Switch**: Implementation using a switch statement to handle state transitions. Demonstrates a code lock system with states Locked, Failed, and Unlocked, showing how to handle state changes in a more procedural way.

#### Visitor
- **Classic Visitor**: Implementation of the Visitor pattern that allows adding new operations to existing object structures without modifying them. The example demonstrates a mathematical expression evaluator and printer, showing how to separate algorithms from the objects they operate on.
- **Functional Visitor**: Alternative implementation of the Visitor pattern using functional programming concepts in Go, providing a more flexible and composable approach to visiting object structures.

### Creational Patterns

#### Builder
- **Functional Builder**: Implementation of the Builder pattern using functional programming concepts in Go
- **Parameter Builder**: Implementation focusing on parameter handling in the Builder pattern
- **Builder Pattern**: Classic implementation of the Builder pattern

#### Factory
- **Basic Factory**: Implementation of the Factory pattern for object creation

#### Prototype
- **Basic Prototype**: Implementation of the Prototype pattern for object cloning

#### Singleton
- **Basic Singleton**: Implementation of the Singleton pattern for single instance management

### Structural Patterns

#### Composite
- **Graphics Composite**: Implementation of the Composite pattern for graphics manipulation
- **Neural Network Composite**: Implementation of the Composite pattern for neural network structures

#### Flyweight
- **Basic Flyweight**: Implementation of the Flyweight pattern that demonstrates how to share common parts of object state among multiple objects. The example shows how to efficiently store and reuse names in a user system, reducing memory usage by storing common strings only once.

## SOLID Principles

### Single Responsibility Principle (SRP)
- **Journal Example**: Implementation demonstrating the Single Responsibility Principle through a journal system. Shows how to separate concerns between journal management and persistence, with examples of both violating and following the principle.

### Open/Closed Principle (OCP)
- **Product Filter Example**: Implementation demonstrating the Open/Closed Principle through a product filtering system. Shows how to extend functionality without modifying existing code by using interfaces and specifications. Includes examples of color and size filtering, and how to combine specifications.

### Dependency Inversion Principle (DIP)
- **Basic DIP**: Implementation demonstrating the Dependency Inversion Principle through dependency injection and interface-based design.

## Additional Concepts

### Lambda in Go
- **Basic Lambda**: Examples of lambda-like functionality in Go

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
├── ChainOfResponsability/         # Basic Chain of Responsibility implementation
├── ChainOfResponsabilityWithSteroids/  # Enhanced Chain of Responsibility
├── Observer/                      # Basic Observer pattern
├── ObserverWithProperties/        # Enhanced Observer with properties
├── memento/
│   ├── simple-memento/           # Basic implementation of Memento pattern
│   └── memento-undo-redo/        # Memento with undo/redo functionality
├── state/
│   ├── state-classic/            # Classic implementation of State pattern
│   ├── state-rules/              # Rules-based implementation of State pattern
│   └── state-switch/             # Switch-based implementation of State pattern
├── classic-visitor/              # Classic implementation of the Visitor pattern
├── functional-visitor/           # Functional implementation of the Visitor pattern
├── Factory/                      # Factory pattern implementation
├── Prototype/                    # Prototype pattern implementation
├── Singleton/                    # Singleton pattern implementation
├── StructuralCompositeGraphics/  # Composite pattern for graphics
├── structuralCompositeNeuralNetwork/  # Composite pattern for neural networks
├── BuilderPattern/               # Classic Builder pattern
├── BulderParameter/              # Parameter-based Builder pattern
├── flyweight/                    # Implementation of the Flyweight pattern
├── single-responsability/        # Implementation of Single Responsibility Principle
├── open-close/                   # Implementation of Open/Closed Principle
├── dependency-inversion/         # Implementation of Dependency Inversion Principle
├── LambdaInGo/                   # Lambda functionality examples
└── functional-builder/           # Functional approach to the Builder pattern
```

## Requirements

- Go 1.16+
