.PHONY: build test coverage bench clean test-bats

build:
	@echo "Building the project..."
	go build -o bin/coolsnippets .
	@echo "Build completed."

test:
	@echo "Running tests..."
	go test ./...
	@echo "Tests completed."

test-bats:
	@echo "Running BATS tests..."
	bats tests/
	@echo "BATS tests completed."

coverage:
	@echo "Running tests with coverage..."
	mkdir -p coverage
	go test -coverprofile=coverage/coverage.out ./...
	go tool cover -html=coverage/coverage.out -o coverage.html
	open coverage.html
	@echo "Coverage report generated: coverage.html"

bench:
	@echo "Running benchmarks..."
	go test ./internal/snippets -bench=. -benchmem
	@echo "Benchmarks completed."

clean:
	@echo "Cleaning..."
	rm -rf bin/ coverage/
	@echo "Clean completed."
