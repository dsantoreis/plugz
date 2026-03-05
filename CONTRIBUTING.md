# Contributing

1. Fork and create a topic branch
2. Use Conventional Commits
3. Add tests for behavior changes
4. Run locally:
   - `go test ./...`
   - `go test ./... -coverprofile=coverage.out`
   - `cd web && npm ci && npm run build`
5. Open PR with problem/root-cause/fix/testing sections
