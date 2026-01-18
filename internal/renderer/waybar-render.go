package renderer

import "parsertry/internal/tools/waybar"

type WaybarRender struct {
	Waybar waybar.Waybar 
	Preset string
}

func (w WaybarRender) Render(template, output string) error {
	return w.Waybar.RenderPreset(w.Preset, template, output);
}
