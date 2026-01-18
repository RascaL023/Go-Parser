package loader

import (
	"encoding/json"
	"os"
	"parsertry/internal/state"
)

func LoadState(path string) (*state.State, error) {
	data, err := os.ReadFile(path);
	if err != nil {
		return nil, err;
	}

	var s state.State
	return &s, json.Unmarshal(data, &s);
}
