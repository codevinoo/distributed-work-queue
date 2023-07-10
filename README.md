# Distributed Work Queue

Redis-backed distributed task queue with worker pools, retries, and dead-letter handling.

## Components

- **Broker** — task enqueue/dequeue via Redis lists
- **Worker** — concurrent task execution with retry logic
- **Scheduler** — delayed and periodic task support

## Run

```bash
go run ./cmd/broker &
go run ./cmd/worker --concurrency 4
```

## License

MIT
