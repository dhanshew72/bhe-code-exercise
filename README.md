# BHE Software Engineer Coding Exercise


## Overview

This module implements a prime number finder using the [Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes). Given a 0-based index `n`, it returns the Nth prime number (e.g., index 0 → 2, index 4 → 11).

The core algorithm is a **segmented sieve** ([pkg/sieve/sieve.go](pkg/sieve/sieve.go)): it sieves a small base segment of size √limit to find base primes, then processes the remaining range in equal-sized chunks, keeping memory usage proportional to √limit rather than the full range. Because finding the Nth prime requires an upper bound, the implementation starts with a small limit and doubles it until enough primes have been found.

A CLI wrapper ([main.go](main.go)) accepts `n` as a command-line argument and prints the result.

## Requirements

- [Go 1.22+](https://go.dev/dl/)

### Setup

```bash
git clone <repo-url>
cd bhe-code-exercise
go mod download
```

### Running the CLI

Find the Nth prime number (0-based index, where index 0 returns 2):

```bash
go run main.go <n>
```

Examples:

```bash
go run main.go 0    # returns 2 (the 1st prime)
go run main.go 99   # returns 541
go run main.go 100  # returns 547
```

### Running Tests

Using Make:

```bash
make test           # run all tests with a 30s timeout
make test-fuzz      # run fuzz tests for 30s (random/edge-case inputs)
```

Using Go directly:

```bash
go test ./...                    # run all tests
go test ./... -fuzz=. -fuzztime=30s  # run fuzz tests
```
