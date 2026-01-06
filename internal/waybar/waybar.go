package waybar

import (
	"os"
	"text/template"
)

type Preset map[string]string

type Waybar struct {
	Presets map[string]Preset `json:"waybar"`
}

func RenderWaybar(
	preset Preset,
	templatePath string,
	outputPath string,
) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	return tmpl.Execute(out, preset)
}
