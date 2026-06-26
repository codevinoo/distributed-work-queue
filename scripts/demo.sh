#!/usr/bin/env bash
set -euo pipefail

trap 'kill $(jobs -p) 2>/dev/null' EXIT

lsof -ti:7001 | xargs kill -9 2>/dev/null || true

go run ./cmd/broker &
go run ./cmd/worker --concurrency 2 &
sleep 0.5

echo "→ Enqueue task"
curl -s -X POST http://localhost:7001/enqueue \
  -H 'Content-Type: application/json' \
  -d '{"type":"email.send","payload":{"to":"user@example.com"}}' | python3 -m json.tool

sleep 1
echo "✓ Demo complete (check worker output above)"
