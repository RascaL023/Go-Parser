package terminal

import (
	"encoding/json"
	// "parsertry/internal/context"
)

type TerminalProcessor struct{}

func (TerminalProcessor) Name() string { return "terminal" }

func (TerminalProcessor) Parse(raw json.RawMessage) (TerminalInput, error) {
	var cfg TerminalInput;
	err := json.Unmarshal(raw, &cfg);
	return cfg, err;
}

func (TerminalProcessor) Resolve(in any) (*Terminal, error) {
	return Resolve(in.(TerminalInput));
}
