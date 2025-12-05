# Car Rental Backend API

A RESTful API for car rental management system built with Go, Fiber, PostgreSQL, and MinIO.

## Features

- **Authentication** - JWT-based user authentication
- **Car Management** - CRUD operations for car inventory
- **Renter Management** - Customer information with photo uploads
- **Rental Transactions** - Complete rental workflow (booking, pickup, return, cancellation)
- **File Storage** - MinIO integration for photo uploads
- **Database Migrations** - Automated schema management

## Tech Stack

- **Go 1.24** - Programming language
- **Fiber v2** - Web framework
- **PostgreSQL 15** - Database
- **MinIO** - Object storage
- **GORM** - ORM
- **Wire** - Dependency injection
- **JWT** - Authentication
- **Docker** - Containerization

## Prerequisites

- Go 1.24+
- Docker & Docker Compose
- Make (optional)

## Quick Start

### Using Docker (Recommended)

```bash
# Clone the repository
git clone <repository-url>
cd carcirus_backend_test

# Start all services
docker compose up --build

# API will be available at http://localhost:8080
```

### Local Development

```bash
# Install dependencies
go mod download

# Install Wire
go install github.com/google/wire/cmd/wire@latest

# Copy environment file
cp .env.example .env

# Generate dependency injection code
make wire

# Run migrations
make migrate-up

# Run the application
make run
```

## Environment Variables

```env
# Server
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=car_rental
DB_SSLMODE=disable

# MinIO
MINIO_ENDPOINT=minio:9000
MINIO_PUBLIC_URL=http://localhost:9000
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin
MINIO_USE_SSL=false
MINIO_BUCKET_NAME=car-rental

# JWT
JWT_SECRET_KEY=super-secret
JWT_TOKEN_DURATION=24h
```

## API Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - User login

### Cars (Protected)
- `GET /api/v1/cars` - List cars (paginated)
- `POST /api/v1/cars` - Create car
- `PUT /api/v1/cars/:id` - Update car
- `DELETE /api/v1/cars/:id` - Delete car

### Renters (Protected)
- `GET /api/v1/renters` - List renters (paginated)
- `POST /api/v1/renters` - Create renter
- `PUT /api/v1/renters/:id` - Update renter
- `DELETE /api/v1/renters/:id` - Delete renter

### Rental Transactions (Protected)
- `GET /api/v1/rental-transactions` - List transactions (paginated)
- `GET /api/v1/rental-transactions/:id` - Get transaction detail
- `POST /api/v1/rental-transactions` - Create booking
- `PUT /api/v1/rental-transactions/:id/pickup` - Mark as picked up
- `PUT /api/v1/rental-transactions/:id/return` - Process return
- `PUT /api/v1/rental-transactions/:id/cancel` - Cancel booking
- `DELETE /api/v1/rental-transactions/:id` - Delete transaction

## Makefile Commands

```bash
make run           # Run the application
make build         # Build the application
make wire          # Generate Wire dependency injection
make migrate-up    # Run database migrations
make migrate-down  # Rollback migrations
make migrate-fresh # Drop and recreate database
```

## Database Migrations

Migrations are located in `migrations/` directory and run automatically on Docker startup.

Manual migration:
```bash
# Up
./migrate.sh up

# Down
./migrate.sh down

# Create new migration
./migrate.sh create <migration_name>
```

## Project Structure

```
.
├── cmd/server/          # Application entry point
├── config/              # Configuration management
├── internal/
│   ├── dto/            # Data transfer objects
│   ├── handler/        # HTTP handlers
│   ├── middleware/     # Middlewares
│   ├── model/          # Database models
│   ├── repository/     # Data access layer
│   ├── service/        # Business logic
│   └── providers/      # Wire providers
├── pkg/
│   ├── database/       # Database connection
│   ├── jwt/            # JWT utilities
│   ├── minio/          # MinIO client
│   └── response/       # Response helpers
├── migrations/         # SQL migrations
├── docker-compose.yml  # Docker services
├── Dockerfile          # Application container
└── Makefile           # Build commands
```

## License

MIT
