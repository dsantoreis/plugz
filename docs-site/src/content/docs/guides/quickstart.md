---
title: Quickstart
---

```bash
go run ./cmd/skillsd serve --skills-dir ./examples/skills --addr :8080
```

```bash
cd web && npm ci && VITE_API_URL=http://localhost:8080 VITE_API_TOKEN=dev-token npm run dev
```

```bash
curl -H "Authorization: Bearer dev-token" http://localhost:8080/api/v1/catalog
```
