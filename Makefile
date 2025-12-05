# Makefile for Car Rental Backend

.PHONY: help run build test clean migrate-up migrate-down migrate-create wire deps

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

wire: ## Generate wire dependencies
	@echo "Generating wire dependencies..."
	cd cmd/server && wire

run: ## Run the application
	@echo "Rewiring dependencies..."
	cd cmd/server && wire
	@echo "Running application..."
	go run ./cmd/server

build: ## Build the application
	@echo "Building application..."
	go build -o bin/server cmd/server/*.go

test: ## Run tests
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...

coverage: test ## Show test coverage
	@echo "Showing coverage..."
	go tool cover -html=coverage.out

migrate-up: ## Run database migrations
	@echo "Running migrations..."
	./migrate.sh up

migrate-down: ## Rollback last migration
	@echo "Rolling back migration..."
	./migrate.sh down

migrate-create: ## Create new migration (usage: make migrate-create name=create_users)
	@echo "Creating migration: $(name)"
	./migrate.sh create $(name)

migrate-version: ## Show current migration version
	@echo "Current migration version:"
	./migrate.sh version

migrate-fresh: ## Drop all tables and re-run all migrations
	@echo "Dropping all tables and re-running migrations..."
	./migrate.sh drop
	./migrate.sh up

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out

docker-db: ## Start PostgreSQL in Docker
	docker run --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=car_rental -p 5432:5432 -d postgres:15-alpine

docker-minio: ## Start MinIO in Docker
	docker run --name minio -p 9000:9000 -p 9001:9001 -e MINIO_ROOT_USER=minioadmin -e MINIO_ROOT_PASSWORD=minioadmin -d quay.io/minio/minio server /data --console-address ":9001"

docker-stop: ## Stop Docker containers
	docker stop postgres minio || true
	docker rm postgres minio || true

install-tools: ## Install development tools
	@echo "Installing tools..."
	go install github.com/google/wire/cmd/wire@latest
	@echo "Please install golang-migrate manually:"
	@echo "  macOS: brew install golang-migrate"
	@echo "  Linux: See https://github.com/golang-migrate/migrate/releases"

docker-up: ## Start all services with Docker Compose
	@echo "Starting all services..."
	docker-compose up -d

docker-down: ## Stop all services
	@echo "Stopping all services..."
	docker-compose down

docker-logs: ## View logs from all services
	docker-compose logs -f

docker-rebuild: ## Rebuild and restart all services
	@echo "Rebuilding services..."
	docker-compose down
	docker-compose build --no-cache
	docker-compose up -d

docker-dev: ## Start only database and minio for local development
	@echo "Starting development services..."
	docker-compose -f docker-compose.dev.yml up -d

docker-dev-down: ## Stop development services
	@echo "Stopping development services..."
	docker-compose -f docker-compose.dev.yml down

docker-clean: ## Clean all Docker resources
	@echo "Cleaning Docker resources..."
	docker-compose down -v
	docker system prune -f

docker-migrate-up: ## Run migrations in Docker
	@echo "Running migrations in Docker..."
	docker-compose exec backend sh -c "cd /root && ./migrate.sh up"

docker-migrate-down: ## Rollback migrations in Docker
	@echo "Rolling back migrations in Docker..."
	docker-compose exec backend sh -c "cd /root && ./migrate.sh down"
