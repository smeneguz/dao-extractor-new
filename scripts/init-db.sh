#!/usr/bin/env bash
# init-db.sh — Initialize the PostgreSQL schema for dao-portal extractor.
#
# Usage:
#   ./scripts/init-db.sh [DATABASE_URL]
#
# If DATABASE_URL is not provided, it defaults to:
#   postgres://dao_portal:password@localhost:5432/dao_portal?sslmode=disable
#
set -euo pipefail

DB_URL="${1:-postgres://dao_portal:password@localhost:5432/dao_portal?sslmode=disable}"
SCHEMA_DIR="$(cd "$(dirname "$0")/../database/postgresql/schema" && pwd)"

echo "==> Initializing database schema"
echo "    URL: ${DB_URL%%\?*}?..."
echo "    Schema dir: ${SCHEMA_DIR}"
echo ""

for sql_file in "$SCHEMA_DIR"/*.sql; do
    filename="$(basename "$sql_file")"
    echo "    Applying ${filename}..."
    psql "$DB_URL" -f "$sql_file" --quiet 2>&1 | grep -v "already exists" || true
done

echo ""
echo "==> Schema initialized successfully"
