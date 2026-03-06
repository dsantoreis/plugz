---
title: API Reference
description: HTTP endpoints for catalog and skill runtime.
---

## Auth

All protected endpoints require:

```http
Authorization: Bearer <token>
```

## Endpoints

### GET `/api/v1/catalog`
Returns available skills and versions.

### POST `/api/v1/skills/{id}/run`
Executes a skill by id with input payload.

### GET `/healthz`
Readiness and liveness probe.

## Example request

```bash
curl -X POST \
  -H "Authorization: Bearer dev-token" \
  -H "Content-Type: application/json" \
  -d '{"input":{"text":"summarize this"}}' \
  http://localhost:8080/api/v1/skills/summarizer/run
```
