# Architecture

## Components

1. **Registry Core**
   - Loads JSON skills from disk
   - Validates schema constraints
   - Exposes list/get operations

2. **Execution Engine**
   - Executes skill command with timeout
   - Captures stdout/stderr/exit status

3. **Marketplace API** (`/api/v1`)
   - `GET /catalog`
   - `POST /install`
   - `POST /test/{name}`
   - `GET /installed`

4. **Cross-Cutting Controls**
   - Auth: bearer token (`SKILLS_API_TOKEN`)
   - Rate limit: per IP token bucket
   - JSON logging middleware
   - Metrics: Prometheus endpoint (`/metrics`)
   - Tracing: OpenTelemetry provider initialization

5. **Frontend (React)**
   - Catalog cards
   - Install action
   - Test action with inline result rendering

## Deployment

- Local: Docker Compose (`api` + `web`)
- Kubernetes:
  - Deployment (2 replicas)
  - Service
  - Ingress

## Testing Pyramid

- Unit tests (registry, executor)
- Integration tests (API auth + install/test)
- E2E behavior check (rate-limit response)
- Stress tests (load/chaos/soak runbook)
