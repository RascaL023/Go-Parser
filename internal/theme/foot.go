package theme

import (
	"os"
	"text/template"
)

type Foot struct {
	Fg string `json:"foreground"`
	Bg string `json:"background"`

	Regular0 string `json:"regular0"`
	Regular1 string `json:"regular1"`
	Regular2 string `json:"regular2"`
	Regular3 string `json:"regular3"`
	Regular4 string `json:"regular4"`
	Regular5 string `json:"regular5"`
	Regular6 string `json:"regular6"`
	Regular7 string `json:"regular7"`

	Bright0 string `json:"bright0"`
	Bright1 string `json:"bright1"`
	Bright2 string `json:"bright2"`
	Bright3 string `json:"bright3"`
	Bright4 string `json:"bright4"`
	Bright5 string `json:"bright5"`
	Bright6 string `json:"bright6"`
	Bright7 string `json:"bright7"`
}

func (tool Foot) Render(templatePath, outputPath string) error {
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

