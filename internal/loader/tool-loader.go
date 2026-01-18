package loader

import (
	"encoding/json"
	"os"
	"parsertry/internal/tools"
)

func LoadToolFromTheme(path string) (*tools.Tool, error) {
	raw, err := os.ReadFile(path);
	if err != nil {
		return nil, err;
	}

	var t tools.Tool;
	if err := json.Unmarshal(raw, &t); err != nil {
		return nil, err;
	}

	return &t, nil;
}
