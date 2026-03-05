package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dsantoreis/ai-agent-skills-demo/internal/registry"
)

func TestRegistryLoadsSkills(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "echo.json"), `{"name":"echo","command":"/bin/echo","args":["hi"]}`)

	r, err := registry.New(dir)
	if err != nil {
		t.Fatalf("new registry: %v", err)
	}

	items := r.List()
	if len(items) != 1 || items[0].Name != "echo" {
		t.Fatalf("expected single loaded skill, got %+v", items)
	}
}

func TestRegistryRejectsMalformedSkill(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "bad.json"), `{"name":"broken"}`)

	_, err := registry.New(dir)
	if err == nil {
		t.Fatal("expected validation error for malformed skill")
	}
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write file: %v", err)
	}
}
