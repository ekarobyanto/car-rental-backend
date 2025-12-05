#!/bin/bash
MIGRATION_DIR="./migrations"

# Use environment variables with defaults
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_NAME=${DB_NAME:-car_rental}
DB_SSLMODE=${DB_SSLMODE:-disable}

DB_URL="postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}"

echo "Connecting to database at ${DB_HOST}:${DB_PORT}..."

case "$1" in
    up)
        echo "Running migrations up..."
        migrate -path $MIGRATION_DIR -database "$DB_URL" up
        ;;
    down)
        echo "Running migrations down..."
        migrate -path $MIGRATION_DIR -database "$DB_URL" down
        ;;
    force)
        if [ -z "$2" ]; then
            echo "Please provide version number: ./migrate.sh force <version>"
            exit 1
        fi
        echo "Forcing migration version $2..."
        migrate -path $MIGRATION_DIR -database "$DB_URL" force $2
        ;;
    create)
        if [ -z "$2" ]; then
            echo "Please provide migration name: ./migrate.sh create <migration_name>"
            exit 1
        fi
        echo "Creating migration: $2"
        migrate create -ext sql -dir $MIGRATION_DIR -seq $2
        ;;
    version)
        echo "Current migration version:"
        migrate -path $MIGRATION_DIR -database "$DB_URL" version
        ;;
    drop)
        echo "Dropping all tables..."
        migrate -path $MIGRATION_DIR -database "$DB_URL" drop
        ;;
    *)
        echo "Usage: $0 {up|down|force <version>|create <name>|version|drop}"
        exit 1
        ;;
esac
