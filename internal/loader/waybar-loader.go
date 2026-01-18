package loader

import (
	"encoding/json"
	"os"
	"parsertry/internal/tools/waybar"
)

func LoadWaybar(path string) (*waybar.Waybar, error) {
	data, err := os.ReadFile(path);
	if err != nil {
		return nil, err;
	}

	var w waybar.Waybar;
	if err := json.Unmarshal(data, &w); err != nil {
		return nil, err;
	}

	return &w, nil;
}
