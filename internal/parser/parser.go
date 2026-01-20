package parser

import (
	"encoding/json"
	"os"
	"parsertry/internal/tools/common/kitty"
)

type Theme struct {
	Kitty kitty.KittyConfig `json:"kitty"`
}

func Parse(path string) (*Theme, error) {
	cfg, err := os.ReadFile(path);
	if err != nil {
		return nil, err;
	}

	var t Theme;
	if err := json.Unmarshal(cfg, &t); err != nil {
		return nil, err;
	}

	return &t, nil;
}
