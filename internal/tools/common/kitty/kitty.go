package kitty

import "parsertry/internal/renderer"

type Kitty struct {
 	Cursor   string
 	Fg       string
 	Bg       string
 	SelectBg string
 	SelectFg string
 	Opacity  string

 	Color0   string
 	Color1   string
 	Color2   string
 	Color3   string
 	Color4   string
 	Color5   string
 	Color6   string
 	Color7   string
 	Color8   string
 	Color9   string
 	Color10  string
 	Color11  string
 	Color12  string
 	Color13  string
 	Color14  string
 	Color15  string
}

func (tool Kitty) Render(templatePath, outputPath string) error {
 	return renderer.Render(templatePath, outputPath, tool);
}
