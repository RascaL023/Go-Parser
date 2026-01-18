package common

import (
	"parsertry/internal/renderer"
	"parsertry/internal/state"
	"parsertry/internal/tools/waybar"
)

type Tool struct {
	Wlogout Wlogout  `json:"wlogout"`
	Cava 		Cava     `json:"cava"`
	Kitty 	Kitty    `json:"kitty"`
	Foot  	Foot     `json:"foot"`
	Hypr  	Hyprland `json:"hyprland"`
}

func RegisterTool(t Tool, s state.State, wb *waybar.Waybar) map[string]renderer.Renderer {
	return map[string]renderer.Renderer{
		"waybar"		: renderer.WaybarRender{
			Waybar: *wb,
			Preset: s.Waybar,
		},
		"wlogout"		: t.Wlogout,
		"cava"	 		: t.Cava,
		"foot"	 		: t.Foot,
		"kitty"	 		: t.Kitty,
		"hypr"			: t.Hypr,
	}
}

