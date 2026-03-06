---
title: Deployment
description: Run Plugz in Docker and Kubernetes.
---

## Docker Compose

```bash
docker compose up --build
```

## Kubernetes

```bash
kubectl apply -f k8s/
```

`k8s/` includes deployment, service, and ingress manifests.

## Docs publish

Docs build and GitHub Pages deployment run via `.github/workflows/docs.yml`.
