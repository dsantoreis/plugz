---
title: Architecture
description: Core components of the Plugz runtime and marketplace.
---

## System layout

- `cmd/skillsd`: API entrypoint
- `internal/orchestrator`: runtime coordination and execution lifecycle
- `internal/registry`: versioned skill catalog and metadata
- `internal/executor`: invocation path and runtime policy enforcement
- `web/`: React operator catalog

## Request flow

1. Client authenticates with API token
2. Catalog endpoint returns available skills and versions
3. Runtime validates request and policy constraints
4. Skill executes with structured output and trace metadata

## Runtime surfaces

- Local developer run
- Docker compose stack
- Kubernetes deployment with ingress
