# Architecture

`skillsd` is split into small internal packages:

- `internal/skill`: schema + strict validation
- `internal/registry`: in-memory registry loaded from JSON files
- `internal/executor`: command execution with timeout and structured result
- `internal/watcher`: directory watcher (`fsnotify`) with hot-reload

Skills are plain JSON files. Any create/write/remove event on `*.json` triggers registry reload.
