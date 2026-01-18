package loader

import (
	"encoding/json"
	"os"
	"parsertry/internal/tools/common"
)

func LoadToolFromTheme(path string) (*common.Tool, error) {
	raw, err := os.ReadFile(path);
	if err != nil {
		return nil, err;
	}

	var t common.Tool;
	if err := json.Unmarshal(raw, &t); err != nil {
		return nil, err;
	}

	return &t, nil;
}
