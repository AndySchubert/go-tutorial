.PHONY: run build test test-v race test-no-cache stress fmt vet clean

# Run server
run:
	go run ./cmd/server

# Build binary
build:
	mkdir -p bin
	go build -o bin/server ./cmd/server

# Run all tests
test:
	go test ./...

# Run tests with verbose output
test-v:
	go test -v ./...

# Run with race detector
race:
	go test -race ./...

# Disable test cache (force re-run)
test-no-cache:
	go test -count=1 ./...

# Stress tests (run multiple times to catch flaky concurrency bugs)
stress:
	go test -race -count=100 ./...

# Format code
fmt:
	go fmt ./...

# Static analysis
vet:
	go vet ./...

check: 
	fmt vet test race

# Clean build artifacts
clean:
	rm -rf bin