#!/bin/bash

# Function to run golangci-lint
run_linter() {
  golangci-lint run
}

# Run golangci-lint initially
run_linter

# Watch for file changes and rerun golangci-lint
nodemon --ext go --exec run_linter
linters-settings:
  gosimple:
    # Deprecated: use the global `run.go` instead.
    go: latest
    # Sxxxx checks in https://staticcheck.io/docs/configuration/options/#checks
    # Default: ["*"]
    checks: ["all"]