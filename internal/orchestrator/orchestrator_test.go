package orchestrator

import "testing"

func TestRun(t *testing.T) {
	out := Run("improve onboarding")
	if out.Plan == "" || out.Context == "" || out.Draft == "" || out.Review == "" {
		t.Fatal("expected full workflow output")
	}
}
