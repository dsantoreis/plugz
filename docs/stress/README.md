# Stress Suite

This suite defines three resilience phases.

## 1) Load test

```bash
go run ./scripts/loadtest.go -url http://localhost:8080/api/v1/catalog -token dev-token -requests 2000 -concurrency 50
```

## 2) Chaos test

Inject instability by restarting containers while load test runs.

```bash
docker compose restart api
```

Expected: transient 5xx, system recovers without manual intervention.

## 3) Soak test

Sustain moderate traffic for 60+ minutes.

```bash
go run ./scripts/loadtest.go -url http://localhost:8080/api/v1/catalog -token dev-token -requests 50000 -concurrency 20
```

Capture p95 latency, error rate, and memory growth.
