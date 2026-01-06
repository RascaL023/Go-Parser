package theme

import (
	"os"
	"text/template"
)

type Cava struct {
	Gradient1 string `json:"gradient1"`
	Gradient2 string `json:"gradient2"`
	Gradient3 string `json:"gradient3"`
	Gradient4 string `json:"gradient4"`
	Gradient5 string `json:"gradient5"`
}

func (tool Cava) Render(templatePath, outputPath string) error {
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
