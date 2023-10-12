#!/bin/sh

set -e

echo "Migrating database"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Starting test application"
exec "$@"