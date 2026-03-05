package skill

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Definition struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Command     string            `json:"command"`
	Args        []string          `json:"args,omitempty"`
	Env         map[string]string `json:"env,omitempty"`
	TimeoutMs   int               `json:"timeout_ms,omitempty"`
}

func Parse(data []byte) (Definition, error) {
	var d Definition
	if err := json.Unmarshal(data, &d); err != nil {
		return Definition{}, fmt.Errorf("decode skill: %w", err)
	}
	if err := Validate(d); err != nil {
		return Definition{}, err
	}
	return d, nil
}

func Validate(d Definition) error {
	switch {
	case strings.TrimSpace(d.Name) == "":
		return errors.New("skill.name is required")
	case strings.TrimSpace(d.Command) == "":
		return errors.New("skill.command is required")
	case d.TimeoutMs < 0:
		return errors.New("skill.timeout_ms must be >= 0")
	}
	return nil
}
