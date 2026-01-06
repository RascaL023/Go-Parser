package theme

import (
	"os"
	"text/template"
)

type Kitty struct {
	Cursor   string `json:"cursor"`
	Fg       string `json:"fg"`
	Bg       string `json:"bg"`
	SelectBg string `json:"selectBg"`
	SelectFg string `json:"selectFg"`
	Opacity  string `json:"opacity"`

	Color0  string `json:"color0"`
	Color1  string `json:"color1"`
	Color2  string `json:"color2"`
	Color3  string `json:"color3"`
	Color4  string `json:"color4"`
	Color5  string `json:"color5"`
	Color6  string `json:"color6"`
	Color7  string `json:"color7"`
	Color8  string `json:"color8"`
	Color9  string `json:"color9"`
	Color10 string `json:"color10"`
	Color11 string `json:"color11"`
	Color12 string `json:"color12"`
	Color13 string `json:"color13"`
	Color14 string `json:"color14"`
	Color15 string `json:"color15"`
}

func (tool Kitty) Render(templatePath, outputPath string) error {
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

