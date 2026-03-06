# Plugz

[![CI](https://github.com/dsantoreis/plugz/actions/workflows/ci.yml/badge.svg)](https://github.com/dsantoreis/plugz/actions/workflows/ci.yml) [![Docs](https://github.com/dsantoreis/plugz/actions/workflows/docs.yml/badge.svg)](https://github.com/dsantoreis/plugz/actions/workflows/docs.yml) [![Coverage >= 80%](https://img.shields.io/badge/coverage-80%25%2B-success)](#ci--coverage) [![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](./LICENSE)

## Hero

Production-ready Go skill marketplace and runtime for internal AI platform teams.

Plugz gives teams one path to discover, install, and execute skills with governance, observability, and deterministic behavior.

![Plugz hero](docs/screenshots/plugz-hero.png)
![Plugz demo GIF](docs/screenshots/plugz-demo.gif)

## Problem

Most internal AI tools break at scale because skills are scattered across scripts with no install contract, no audit trail, and no runtime policy.

Plugz solves this with a versioned catalog API, controlled execution, and a frontend that operators can actually use.

## Quickstart (3 comandos)

```bash
go run ./cmd/skillsd serve --skills-dir ./examples/skills --addr :8080
cd web && npm ci && VITE_API_URL=http://localhost:8080 VITE_API_TOKEN=dev-token npm run dev
curl -H "Authorization: Bearer dev-token" http://localhost:8080/api/v1/catalog
```

## Docs Site (Astro Starlight)

```bash
cd docs-site
npm install
npm run dev
```

Build static docs:

```bash
npm run build
```

Published docs: https://dsantoreis.github.io/plugz/ (GitHub Pages via `.github/workflows/docs.yml`)

## Docker

```bash
docker compose up --build
```

## Kubernetes

```bash
kubectl apply -f k8s/
```

Manifests include deployment, service, and ingress resources for production-like rollouts.

## CI + Coverage

CI validates lint, tests, build, frontend bundle, race checks, and coverage gate:

- `golangci-lint`
- `go test ./...`
- `go build ./...`
- `go test ./internal/orchestrator -coverprofile=coverage.out`
- `go tool cover -func=coverage.out` with `>=80%` threshold

## Architecture

- API server: `cmd/skillsd`
- Core runtime: `internal/orchestrator`, `internal/executor`, `internal/registry`
- Frontend catalog: `web/`
- Docs UI: `docs-site/`
- Published: https://dsantoreis.github.io/plugz/ (GitHub Pages via `.github/workflows/docs.yml`)

## CTA

If this helps your team ship governed AI skills faster, star the repo and open an issue with your use case.
PRs with reproducible load-test notes are welcome.
