# Learn Go with Tests

A personal repository for working through [Chris James' "Learn Go with Tests"](https://quii.gitbook.io/learn-go-with-tests) course.

## About

This repository contains my solutions and exercises from the "Learn Go with Tests" course, which teaches Go fundamentals using a test-driven development (TDD) approach. Each chapter introduces new Go concepts while reinforcing the TDD cycle: write a failing test, make it pass, then refactor.

## Chapters

### Go Fundamentals

| Chapter | Topics Covered |
|---------|----------------|
| [Hello, World](./hello_world) | Functions, constants, if/else, TDD basics |
| [Integers](./integers) | Integer operations, example tests |
| [Iteration](./iteration) | For loops, benchmarking |
| [Arrays and Slices](./arrays_and_slices) | Arrays, slices, range, variadic functions |
| [Structs, Methods & Interfaces](./structs) | Structs, methods, receivers, interfaces |
| [Pointers & Errors](./pointers_errors) | Pointers, errors, nil |
| [Maps](./maps) | Maps, custom types, error handling |
| [Dependency Injection](./di) | Dependency injection, interfaces |
| [Mocking](./mocking) | Mocking, dependency injection in tests |
| [Concurrency](./concurrency) | Goroutines, channels |
| [Select](./select) | Select statement, timeouts |
| [Reflection](./reflection) | Reflection API, type inspection |
| [Sync](./sync) | Mutexes, WaitGroups, synchronization |
| [Context](./context) | Context package, cancellation, deadlines |
| [Intro to Property Based Tests](./property_based_tests) | Property-based testing |
| [Maths](./maths) | Math operations, algorithms |
| [Reading Files](./reading_files) | File I/O, reading files |
| [Templating](./templating) | HTML/text templates |
| [Generics](./generics) | Type parameters, generic functions |
| [Revisiting Arrays and Slices with Generics](./generics) | Applying generics to collections |

## Getting Started

### Prerequisites

- [Go 1.18+](https://go.dev/dl/) (for generics support)

### Setup

Initialize the Go module (if not already done):

```bash
go mod init github.com/yourusername/tdd-learning-go
```

Download dependencies:

```bash
go mod download
go mod tidy
```

## Running Tests

Run all tests from the root directory:

```bash
go test ./...
```

Run tests for a specific chapter:

```bash
go test ./hello_world
go test ./maps/...
```

Run tests with verbose output:

```bash
go test -v ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## Resources

- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) - The course this repo follows
- [Go Documentation](https://go.dev/doc/) - Official Go docs
- [Effective Go](https://go.dev/doc/effective_go) - Best practices
