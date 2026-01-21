package output

import "parsertry/internal/renderer"

type FootOutput struct {
	TerminalOutput
}

func (tool FootOutput) Render(templatePath, outputPath string) error {
 	return renderer.Render(templatePath, outputPath, tool);
}
