package renderer

import (
	"encoding/json"
	"os"

	"parsertry/internal/state"
	"parsertry/internal/theme"
	"parsertry/internal/waybar"
)

func LoadTheme(path string) (*theme.Theme, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var t theme.Theme
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func LoadWaybar(path string) (*waybar.Waybar, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var t waybar.Waybar
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func LoadState(path string) (*state.State, error){
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var s state.State
	return &s, json.Unmarshal(data, &s)
}
