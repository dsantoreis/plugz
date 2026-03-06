---
title: Getting Started
description: Start Plugz API and web catalog locally.
---

## Prerequisites

- Go 1.22+
- Node.js 20+

## Run in 3 commands

```bash
go run ./cmd/skillsd serve --skills-dir ./examples/skills --addr :8080
cd web && npm ci && VITE_API_URL=http://localhost:8080 VITE_API_TOKEN=dev-token npm run dev
curl -H "Authorization: Bearer dev-token" http://localhost:8080/api/v1/catalog
```

## Local quality gate

```bash
go test ./...
go build ./...
cd web && npm run build
```
