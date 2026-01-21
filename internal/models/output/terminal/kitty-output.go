package output

import "parsertry/internal/renderer"

type KittyOutput struct {
	TerminalOutput

	SelectBackground	string
	SelectForeground	string

	Opacity float64
	Cursor string
}

func (tool KittyOutput) Render(templatePath, outputPath string) error {
 	return renderer.Render(templatePath, outputPath, tool);
}
