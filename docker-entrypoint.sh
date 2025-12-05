#!/bin/bash

# Entrypoint script for Docker container
# This script runs migrations before starting the application

set -e

echo "Starting Car Rental Backend..."

# Wait for database to be ready
echo "Waiting for database..."
until nc -z -v -w30 ${DB_HOST:-postgres} ${DB_PORT:-5432}
do
  echo "Waiting for database connection..."
  sleep 2
done

echo "Database is ready!"

# Run migrations
echo "Running database migrations..."
./migrate.sh up

# Start the application
echo "Starting application..."
exec ./main
