package waybar

import (
	"fmt"
	"os"
	"text/template"
)

type Preset map[string]string

type Waybar struct {
	WaybarPreset map[string]Preset `json:"waybar"`
}

func (w Waybar) RenderPreset(name, templatePath, outputPath string) error {
	preset, ok := w.WaybarPreset[name];
	if !ok {
		return fmt.Errorf("Waybar preset [%s] not found\n", name);
	}
	
	tmpl, err := template.ParseFiles(templatePath);
	if err != nil {
		return fmt.Errorf("Waybar path not found!");
	}

	out, err := os.Create(outputPath);
	if err != nil {
		return fmt.Errorf("Waybar output path not found!");
	}

	return tmpl.Execute(out, preset);
}
