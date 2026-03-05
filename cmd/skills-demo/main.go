package main

import (
	"fmt"
	"os"

	"github.com/dsantoreis/ai-agent-skills-demo/internal/orchestrator"
)

func main() {
	query := "improve onboarding"
	if len(os.Args) > 1 {
		query = os.Args[1]
	}
	out := orchestrator.Run(query)
	fmt.Println("Plan:", out.Plan)
	fmt.Println("Context:", out.Context)
	fmt.Println("Draft:", out.Draft)
	fmt.Println("Review:", out.Review)
}
