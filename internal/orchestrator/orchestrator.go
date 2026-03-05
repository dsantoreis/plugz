package orchestrator

import "strings"

type Output struct {
	Plan    string
	Context string
	Draft   string
	Review  string
}

func Run(query string) Output {
	ctx := "retrieved: onboarding, activation, retention"
	return Output{
		Plan:    "planner -> researcher -> writer -> reviewer",
		Context: ctx,
		Draft:   "Recommended approach: " + strings.TrimSpace(query) + " with metrics and guardrails.",
		Review:  "Review notes: concise, actionable, measurable.",
	}
}
