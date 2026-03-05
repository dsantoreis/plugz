# AI Agent Skills Marketplace & Registry

Enterprise-ready Go + React reference implementation for a **skill marketplace**:
- Web catalog (React) for discovery, install, and test flows
- Registry/API in Go with auth, rate limiting, JSON logs, Prometheus + OpenTelemetry
- Docker + Kubernetes deployment assets
- CI with lint/test/build, integration/e2e, and coverage artifacts
- Stress testing playbook (load/chaos/soak)

## Product Pitch

This project demonstrates how to package AI skills as versioned contracts and expose them through a marketplace user experience. It is aimed at teams building internal AI platforms and partner ecosystems where governance, observability, and reliability matter as much as speed.

## Architecture

- **Backend**: `cmd/skillsd serve` + `internal/api`
- **Skill Runtime**: `internal/registry`, `internal/executor`
- **Frontend**: `web/` (React + Vite)
- **Telemetry**:
  - JSON request logs
  - Prometheus metrics (`/metrics`)
  - OpenTelemetry tracing (stdout exporter)

See full technical layout in `docs/architecture.md`.

## Quickstart

### 1) Backend API
```bash
go run ./cmd/skillsd serve --skills-dir ./examples/skills --addr :8080
```

### 2) Frontend catalog
```bash
cd web
npm ci
VITE_API_URL=http://localhost:8080 VITE_API_TOKEN=dev-token npm run dev
```

### 3) Use the API directly
```bash
curl -H "Authorization: Bearer dev-token" http://localhost:8080/api/v1/catalog
curl -X POST -H "Authorization: Bearer dev-token" -H "Content-Type: application/json" \
  -d '{"name":"echo"}' http://localhost:8080/api/v1/install
curl -X POST -H "Authorization: Bearer dev-token" -H "Content-Type: application/json" \
  -d '{"input":"hello"}' http://localhost:8080/api/v1/test/echo
```

## Security Defaults

- Bearer token enforced for marketplace routes
- Per-IP rate limiting middleware
- Structured logs for auditability
- Security policy in `SECURITY.md`

## CI Quality Gate

GitHub Actions (`.github/workflows/ci.yml`) runs:
- `golangci-lint`
- `go test ./... -coverprofile=coverage.out`
- `go build ./...`
- coverage text + HTML report artifact
- frontend build (`npm run build`)

## Docker & Kubernetes

```bash
docker compose up --build
kubectl apply -f k8s/
```

## Stress Suite

Runbook and scripts in `docs/stress/` + `scripts/`.

## Governance & Community

- `SECURITY.md`
- `CONTRIBUTING.md`
- `CODE_OF_CONDUCT.md`
- `CHANGELOG.md`

## License

MIT
