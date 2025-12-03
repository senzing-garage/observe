# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`observe` is a Go command-line tool that receives Observer messages over gRPC. It's part of the [senzing-tools](https://github.com/senzing-garage/senzing-tools) suite.

## Common Commands

### Build

```bash
make clean build        # Build binary for current platform
make build-all          # Build for all platforms (darwin/linux/windows, amd64/arm64)
```

Binaries output to `target/<os>-<arch>/observe`

### Test

```bash
make clean setup test   # Run all tests
go test -v ./...        # Run tests directly
go test -v -run TestName ./path/to/package  # Run single test
```

### Lint

```bash
make lint               # Run golangci-lint, govulncheck, and cspell
make golangci-lint      # Run only golangci-lint
make fix                # Auto-fix lint issues where possible
```

### Other

```bash
make dependencies              # Update Go dependencies
make dependencies-for-development  # Install dev tools (golangci-lint, godoc, etc.)
make coverage                  # Run tests with coverage report
make docker-build             # Build Docker image
go run main.go                # Run directly
```

## Architecture

- `main.go` - Entry point, calls `cmd.Execute()`
- `cmd/` - Cobra command setup
  - `root.go` - Main command definition, starts gRPC server
  - `context_*.go` - Platform-specific context variables (linux/darwin/windows)
  - `github.go` - Version info from GitHub
- `observer/` - Core observer implementation
  - `observer_simple.go` - `SimpleObserver` struct that runs gRPC server using `go-observing` library

The tool creates a gRPC server that:

1. Creates a Subject (`go-observing/subject.SimpleSubject`)
2. Registers a `RawObserver` with the subject
3. Serves via `grpcserver.SimpleGrpcServer`

## Key Dependencies

- `github.com/senzing-garage/go-cmdhelping` - Command-line helpers for Cobra/Viper
- `github.com/senzing-garage/go-observing` - Observer pattern implementation with gRPC
- `github.com/spf13/cobra` and `github.com/spf13/viper` - CLI framework

## Linting Configuration

Uses golangci-lint with extensive linters enabled. Config at `.github/linters/.golangci.yaml`. Line length limit is 120 characters.
