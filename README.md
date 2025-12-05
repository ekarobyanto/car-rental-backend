# Car Rental Backend

A Go backend application using clean architecture with Fiber framework.

## Tech Stack

- **Framework**: Fiber v2
- **ORM**: GORM
- **Database**: PostgreSQL
- **Dependency Injection**: Wire
- **Migrations**: go-migrate
- **Storage**: MinIO
- **Configuration**: Environment variables

## Project Structure

```
.
├── cmd/
│   └── server/          # Application entry point
│       ├── main.go      # Main application
│       ├── wire.go      # Wire dependency injection
│       └── wire_gen.go  # Generated wire code
├── config/              # Configuration management
├── internal/            # Private application code
│   ├── dto/            # Data Transfer Objects
│   ├── handler/        # HTTP handlers (controllers)
│   ├── model/          # Domain models
│   ├── repository/     # Data access layer
│   └── service/        # Business logic layer
├── pkg/                # Public packages
│   ├── database/       # Database connection
│   ├── minio/          # MinIO client
│   └── response/       # HTTP response helpers
├── migrations/         # Database migrations
└── go.mod
```

## Prerequisites

- Go 1.23+
- PostgreSQL
- MinIO (or S3-compatible storage)
- migrate CLI tool
- Docker & Docker Compose (for containerized deployment)

## Quick Start with Docker

The fastest way to get started:

```bash
# Start all services (PostgreSQL, MinIO, Backend)
make docker-up

# View logs
make docker-logs

# Stop services
make docker-down
```

The application will run at http://localhost:8080 with pre-seeded data.

See [DEPLOYMENT.md](DEPLOYMENT.md) for detailed deployment instructions.

## Installation (Local Development)

1. Clone the repository
2. Copy environment variables:
```bash
cp .env.example .env
```

3. Install dependencies:
```bash
go mod download
```

4. Install go-migrate:
```bash
# macOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# Windows
scoop install migrate
```

5. Install Wire:
```bash
go install github.com/google/wire/cmd/wire@latest
```

## Configuration

Edit the `.env` file with your settings:

```env
# Server
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=car_rental
DB_SSLMODE=disable

# MinIO
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin
MINIO_USE_SSL=false
MINIO_BUCKET_NAME=car-rental
```

## Running Migrations

Make the migration script executable:
```bash
chmod +x migrate.sh
```

Run migrations:
```bash
# Run all migrations
./migrate.sh up

# Rollback last migration
./migrate.sh down

# Create new migration
./migrate.sh create add_cars_table

# Check current version
./migrate.sh version

# Force to specific version
./migrate.sh force 1
```

## Generate Wire Dependencies

When you modify the dependency injection setup:

```bash
cd cmd/server
wire
```

## Running the Application

### With Docker (Recommended)

```bash
# Production mode (all services)
make docker-up

# Development mode (only database services)
make docker-dev
go run cmd/server/main.go
```

### Without Docker

```bash
go run cmd/server/main.go
```

Or build and run:

```bash
go build -o bin/server cmd/server/main.go
./bin/server
```

## Test Credentials

After running migrations with seed data:

```
Admin User:
  Email: admin@carental.com
  Password: password123

Staff User:
  Email: staff@carental.com
  Password: password123
```

## API Endpoints

### Health Check
- `GET /health` - Check server health

### Public Endpoints (No Authentication)
- `POST /api/v1/public/register` - Register new user
- `POST /api/v1/public/login` - User login (returns JWT token)

### Protected Endpoints (Requires JWT Token)
- `GET /api/v1/users/me` - Get current user info
- `GET /api/v1/users` - List users (with pagination)
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

**Authentication**: Protected endpoints require `Authorization: Bearer <token>` header.

See [JWT Authentication Documentation](docs/JWT_AUTHENTICATION.md) for detailed usage.

## Example API Requests

### Register User
```bash
curl -X POST http://localhost:8080/api/v1/public/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "name": "John Doe",
    "password": "password123",
    "phone": "+1234567890"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/api/v1/public/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

### Access Protected Endpoint
```bash
# Get token from login response, then:
curl -X GET http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Development

### Adding New Features

1. Create model in `internal/model/`
2. Create DTOs in `internal/dto/`
3. Create repository interface and implementation in `internal/repository/`
4. Create service interface and implementation in `internal/service/`
5. Create handler in `internal/handler/`
6. Register handler in Wire (`cmd/server/wire.go`)
7. Run `wire` to regenerate dependencies
8. Register routes in `cmd/server/main.go`

### Creating Migrations

```bash
./migrate.sh create your_migration_name
```

This creates two files:
- `migrations/XXXXXX_your_migration_name.up.sql`
- `migrations/XXXXXX_your_migration_name.down.sql`

## Docker Setup (Optional)

### PostgreSQL
```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=car_rental -p 5432:5432 -d postgres:15-alpine
```

### MinIO
```bash
docker run --name minio -p 9000:9000 -p 9001:9001 -e MINIO_ROOT_USER=minioadmin -e MINIO_ROOT_PASSWORD=minioadmin -d quay.io/minio/minio server /data --console-address ":9001"
```

## License

MIT
