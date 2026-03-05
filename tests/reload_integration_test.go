package tests

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/dsantoreis/ai-agent-skills-demo/internal/registry"
	"github.com/dsantoreis/ai-agent-skills-demo/internal/watcher"
)

func TestWatcherHotReloadIntegration(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "a.json"), `{"name":"a","command":"/bin/echo","args":["a"]}`)

	r, err := registry.New(dir)
	if err != nil {
		t.Fatalf("new registry: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() { _ = watcher.Start(ctx, dir, r) }()
	time.Sleep(150 * time.Millisecond)

	writeFile(t, filepath.Join(dir, "b.json"), `{"name":"b","command":"/bin/echo","args":["b"]}`)

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if _, ok := r.Get("b"); ok {
			return
		}
		time.Sleep(50 * time.Millisecond)
	}
	t.Fatal("expected hot reload to load new skill b")
}
