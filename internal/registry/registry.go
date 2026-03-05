package registry

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/dsantoreis/ai-agent-skills-demo/internal/skill"
)

type Registry struct {
	mu     sync.RWMutex
	dir    string
	skills map[string]skill.Definition
}

func New(dir string) (*Registry, error) {
	r := &Registry{dir: dir, skills: map[string]skill.Definition{}}
	if err := r.Reload(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Registry) Reload() error {
	entries, err := os.ReadDir(r.dir)
	if err != nil {
		return fmt.Errorf("read skills dir: %w", err)
	}
	loaded := make(map[string]skill.Definition)
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		path := filepath.Join(r.dir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}
		d, err := skill.Parse(data)
		if err != nil {
			return fmt.Errorf("invalid %s: %w", path, err)
		}
		loaded[d.Name] = d
	}

	r.mu.Lock()
	r.skills = loaded
	r.mu.Unlock()
	return nil
}

func (r *Registry) Get(name string) (skill.Definition, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	d, ok := r.skills[name]
	return d, ok
}

func (r *Registry) List() []skill.Definition {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]skill.Definition, 0, len(r.skills))
	for _, d := range r.skills {
		out = append(out, d)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}
