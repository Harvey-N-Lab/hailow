.PHONY: build test clean install run help

# Variables
BINARY_NAME=hailow
VERSION?=dev
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(BUILD_TIME)"

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the CLI binary
	@echo "Building $(BINARY_NAME)..."
	@go build $(LDFLAGS) -o bin/$(BINARY_NAME) ./cmd/hailow

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html

install: build ## Install the binary to ~/.local/bin
	@echo "Installing to ~/.local/bin..."
	@mkdir -p ~/.local/bin
	@cp bin/$(BINARY_NAME) ~/.local/bin/
	@echo "Installed! Make sure ~/.local/bin is in your PATH"

run: build ## Build and run the CLI
	@./bin/$(BINARY_NAME)

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

lint: ## Lint code
	@echo "Linting code..."
	@golangci-lint run

release-test: ## Test release process
	@echo "Testing release process..."
	@goreleaser release --snapshot --clean

.DEFAULT_GOAL := help
