#!/usr/bin/env bash
# fill-gaps.sh — Find and fill missing blocks for an indexer.
#
# Usage:
#   ./scripts/fill-gaps.sh <indexer-name> [db-url]
#
# Example:
#   ./scripts/fill-gaps.sh ethereum-mainnet
#   ./scripts/fill-gaps.sh arbitrum "postgres://user:pass@host:5432/db?sslmode=disable"
#
# This script:
#   1. Queries the DB for the min/max block height for the given indexer
#   2. Finds all gaps (missing block heights) in that range
#   3. Uses `parse range` to re-index each missing block
#
set -euo pipefail

INDEXER="${1:?Usage: fill-gaps.sh <indexer-name> [db-url]}"
DB_URL="${2:-postgres://dao_portal:password@localhost:5432/dao_portal?sslmode=disable}"
EXTRACTOR="./build/extractor"
HOME_DIR="${PWD}"

echo "==> Finding gaps for indexer: ${INDEXER}"

# Get missing block heights
MISSING=$(psql "${DB_URL}" -t -A -c "
WITH bounds AS (
  SELECT min(height) as lo, max(height) as hi
  FROM blocks WHERE indexer = '${INDEXER}'
),
seq AS (
  SELECT generate_series(lo, hi) as h FROM bounds
),
indexed AS (
  SELECT height FROM blocks WHERE indexer = '${INDEXER}'
)
SELECT s.h
FROM seq s
LEFT JOIN indexed i ON s.h = i.height
WHERE i.height IS NULL
ORDER BY s.h;
")

if [ -z "${MISSING}" ]; then
  echo "==> No gaps found. All blocks are indexed."
  exit 0
fi

COUNT=$(echo "${MISSING}" | wc -l | tr -d ' ')
echo "==> Found ${COUNT} missing blocks"

FILLED=0
FAILED=0
for HEIGHT in ${MISSING}; do
  echo -n "    Filling block ${HEIGHT}... "
  if ${EXTRACTOR} parse range "${INDEXER}" "${HEIGHT}" --home "${HOME_DIR}" > /dev/null 2>&1; then
    echo "OK"
    FILLED=$((FILLED + 1))
  else
    echo "FAILED"
    FAILED=$((FAILED + 1))
  fi
done

echo ""
echo "==> Done. Filled: ${FILLED}, Failed: ${FAILED}, Total: ${COUNT}"
