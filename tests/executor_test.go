package tests

import (
	"context"
	"testing"
	"time"

	"github.com/dsantoreis/ai-agent-skills-demo/internal/executor"
	"github.com/dsantoreis/ai-agent-skills-demo/internal/skill"
)

func TestExecutorRunSuccess(t *testing.T) {
	s := skill.Definition{Name: "echo", Command: "/bin/sh", Args: []string{"-c", "cat"}}
	res := executor.Run(context.Background(), s, "hello", time.Second)
	if res.Status != "ok" {
		t.Fatalf("expected ok status, got %+v", res)
	}
	if res.Stdout != "hello" {
		t.Fatalf("expected stdout=hello, got %q", res.Stdout)
	}
}

func TestExecutorTimeout(t *testing.T) {
	s := skill.Definition{Name: "sleep", Command: "/bin/sh", Args: []string{"-c", "sleep 2"}}
	res := executor.Run(context.Background(), s, "", 100*time.Millisecond)
	if res.Status != "timeout" {
		t.Fatalf("expected timeout, got %+v", res)
	}
}
