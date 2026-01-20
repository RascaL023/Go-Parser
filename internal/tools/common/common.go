package common

import (
	"parsertry/internal/renderer"
	"parsertry/internal/state"
	"parsertry/internal/tools/common/kitty"
	"parsertry/internal/tools/waybar"
)

type Tool struct {
	// Wlogout Wlogout  `json:"wlogout"`
	// Cava 		Cava     `json:"cava"`
	Kitty 	kitty.KittyConfig    `json:"kitty"`
	// Foot  	Foot     `json:"foot"`
	// Hypr  	Hyprland `json:"hyprland"`
}

type RenderTool struct {
	Kitty kitty.Kitty
}

func RegisterTool(t RenderTool, s state.State, wb *waybar.Waybar) map[string]renderer.Renderer {
	return map[string]renderer.Renderer{
		"waybar"		: renderer.WaybarRender{
			Waybar: *wb,
			Preset: s.Waybar,
		},
		// "wlogout"		: t.Wlogout,
		// "cava"	 		: t.Cava,
		// "foot"	 		: t.Foot,
		"kitty"	 		: t.Kitty,
		// "hypr"			: t.Hypr,
	}
}

