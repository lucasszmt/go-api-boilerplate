.PHONY: build run test lint clean docker-build docker-up docker-down migrate-up migrate-down sqlc help

# Application
APP_NAME := go-api-boilerplate
BUILD_DIR := bin
MAIN_PATH := cmd/api/main.go

# Go
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOMOD := $(GOCMD) mod
GOFMT := gofmt

# Tools
SQLC := sqlc
MIGRATE := migrate
GOLANGCI_LINT := golangci-lint

# Docker
DOCKER_COMPOSE := docker-compose

# Database
DB_URL := postgres://postgres:postgres@localhost:5432/go_api_boilerplate?sslmode=disable

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'

## build: Build the application
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

## run: Run the application
run:
	@echo "Running..."
	$(GOCMD) run $(MAIN_PATH)

## test: Run tests
test:
	@echo "Testing..."
	$(GOTEST) -v -race -cover ./...

## test-coverage: Run tests with coverage report
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -race -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## lint: Run linter
lint:
	@echo "Linting..."
	$(GOLANGCI_LINT) run ./...

## fmt: Format code
fmt:
	@echo "Formatting..."
	$(GOFMT) -s -w .

## tidy: Tidy go modules
tidy:
	@echo "Tidying modules..."
	$(GOMOD) tidy

## clean: Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html

## docker-build: Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):latest .

## docker-up: Start Docker Compose services
docker-up:
	@echo "Starting services..."
	$(DOCKER_COMPOSE) up -d

## docker-down: Stop Docker Compose services
docker-down:
	@echo "Stopping services..."
	$(DOCKER_COMPOSE) down

## docker-logs: Show Docker Compose logs
docker-logs:
	$(DOCKER_COMPOSE) logs -f

## migrate-up: Run database migrations
migrate-up:
	@echo "Running migrations..."
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" up

## migrate-down: Rollback database migrations
migrate-down:
	@echo "Rolling back migrations..."
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" down 1

## migrate-create: Create a new migration (usage: make migrate-create name=migration_name)
migrate-create:
	@echo "Creating migration..."
	$(MIGRATE) create -ext sql -dir db/migrations -seq $(name)

## sqlc: Generate sqlc code
sqlc:
	@echo "Generating sqlc..."
	$(SQLC) generate

## install-tools: Install development tools
install-tools:
	@echo "Installing tools..."
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

## dev: Run with hot reload (requires air)
dev:
	@echo "Starting development server..."
	air
