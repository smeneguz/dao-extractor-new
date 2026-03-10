#!/usr/bin/env bash
# backfill.sh — Historical backfill using `parse range`.
#
# Unlike `start` (which runs a massive generate_series query at startup),
# `parse range` feeds blocks sequentially into the worker queue with no
# heavy DB query. This is the correct way to index historical data.
#
# Usage:
#   ./scripts/backfill.sh <indexer-name> <start-height> <end-height> [chunk-size]
#
# Examples:
#   # Backfill ethereum from deployment to block 24M
#   ./scripts/backfill.sh ethereum-mainnet 4620855 24000000
#
#   # Backfill with 200K block chunks (default: 100000)
#   ./scripts/backfill.sh ethereum-mainnet 4620855 24000000 200000
#
#   # Backfill all chains (run in separate terminals):
#   ./scripts/backfill.sh ethereum-mainnet 4620855  24620000
#   ./scripts/backfill.sh arbitrum        70397646 300000000
#   ./scripts/backfill.sh optimism        116055145 130000000
#   ./scripts/backfill.sh base            5842016  25000000
#   ./scripts/backfill.sh gnosis          35476313 38000000
#   ./scripts/backfill.sh lisk            1        10000000
#   ./scripts/backfill.sh polygon         25818525 70000000
#
# The script splits the range into chunks and runs `parse range` for each.
# Progress is tracked in a file so it can be resumed after interruption.
#
set -euo pipefail

INDEXER="${1:?Usage: backfill.sh <indexer-name> <start-height> <end-height> [chunk-size]}"
START="${2:?Usage: backfill.sh <indexer-name> <start-height> <end-height> [chunk-size]}"
END="${3:?Usage: backfill.sh <indexer-name> <start-height> <end-height> [chunk-size]}"
CHUNK="${4:-100000}"
EXTRACTOR="./build/extractor"
HOME_DIR="${PWD}"
PROGRESS_FILE=".backfill-progress-${INDEXER}"

if [ ! -f "${EXTRACTOR}" ]; then
  echo "Error: ${EXTRACTOR} not found. Run: go build -o build/extractor ./cmd/extractor"
  exit 1
fi

TOTAL=$(( END - START + 1 ))
CHUNKS=$(( (TOTAL + CHUNK - 1) / CHUNK ))

echo "==> Backfill: ${INDEXER}"
echo "    Range: ${START} -> ${END} (${TOTAL} blocks)"
echo "    Chunk size: ${CHUNK} (${CHUNKS} chunks)"
echo "    Progress file: ${PROGRESS_FILE}"
echo ""

# Resume from last completed chunk if progress file exists.
COMPLETED=0
if [ -f "${PROGRESS_FILE}" ]; then
  COMPLETED=$(cat "${PROGRESS_FILE}")
  echo "    Resuming from chunk $((COMPLETED + 1))/${CHUNKS}"
  echo ""
fi

CHUNK_IDX=0
CURRENT="${START}"

while [ "${CURRENT}" -le "${END}" ]; do
  CHUNK_END=$(( CURRENT + CHUNK - 1 ))
  if [ "${CHUNK_END}" -gt "${END}" ]; then
    CHUNK_END="${END}"
  fi

  CHUNK_IDX=$(( CHUNK_IDX + 1 ))

  # Skip already-completed chunks.
  if [ "${CHUNK_IDX}" -le "${COMPLETED}" ]; then
    CURRENT=$(( CHUNK_END + 1 ))
    continue
  fi

  echo "==> Chunk ${CHUNK_IDX}/${CHUNKS}: blocks ${CURRENT} -> ${CHUNK_END}"

  if ${EXTRACTOR} parse range "${INDEXER}" "${CURRENT}" "${CHUNK_END}" --home "${HOME_DIR}"; then
    echo "${CHUNK_IDX}" > "${PROGRESS_FILE}"
    echo "    Chunk ${CHUNK_IDX} complete."
  else
    echo "    Chunk ${CHUNK_IDX} FAILED (exit code $?). Re-run to retry from this chunk."
    exit 1
  fi

  CURRENT=$(( CHUNK_END + 1 ))
done

echo ""
echo "==> Backfill complete: ${INDEXER} (${START} -> ${END})"
echo "    Run ./scripts/fill-gaps.sh ${INDEXER} to verify zero gaps."

# Clean up progress file.
rm -f "${PROGRESS_FILE}"
