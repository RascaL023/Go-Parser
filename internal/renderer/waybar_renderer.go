package renderer

import (
	"fmt"
	"parsertry/internal/waybar"
)

type WaybarRenderer struct {
    Waybar *waybar.Waybar
    Preset string
}

func (w *WaybarRenderer) Render(templatePath, outputPath string) error {
    preset, ok := w.Waybar.Presets[w.Preset]
    if !ok {
        return fmt.Errorf("unknown waybar preset: %s", w.Preset)
    }

    return waybar.RenderWaybar(preset, templatePath, outputPath)
}
