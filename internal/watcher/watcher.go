package watcher

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type Reloader interface {
	Reload() error
}

func Start(ctx context.Context, skillsDir string, r Reloader) error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("create watcher: %w", err)
	}
	defer w.Close()

	if err := w.Add(skillsDir); err != nil {
		return fmt.Errorf("watch skills dir: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-w.Errors:
			if err != nil {
				return fmt.Errorf("watch error: %w", err)
			}
		case ev := <-w.Events:
			if !isSkillChange(ev.Name, ev.Op) {
				continue
			}
			if err := r.Reload(); err != nil {
				return fmt.Errorf("reload skills after change: %w", err)
			}
		}
	}
}

func isSkillChange(name string, op fsnotify.Op) bool {
	if filepath.Ext(name) != ".json" {
		return false
	}
	if strings.HasPrefix(filepath.Base(name), ".") {
		return false
	}
	return op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove|fsnotify.Rename) != 0
}
