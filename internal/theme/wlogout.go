package theme

import (
	"os"
	"text/template"
)

type Wlogout struct {
	WindowBg string `json:"window-bg"`
	ButtonBg string `json:"button-bg"`
	ButtonFocusBg string `json:"button-focus-bg"`
	ButtonFocusBorder string `json:"button-focus-border"`
}

func (tool Wlogout) Render(templatePath, outputPath string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	return tmpl.Execute(out, tool)
}

