#!/bin/sh
set -e

# Load .env if present (useful for local docker run without compose)
if [ -f /app/.env ]; then
  export $(grep -v '^#' /app/.env | xargs)
fi

echo "Starting money-manager..."
exec ./money-manager
