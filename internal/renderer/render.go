package renderer

import (
	"os"
	"text/template"
)
//
// type Renderer interface {
// 	Render(inputPath, outputPath string) error
// }
//
func Render(templatePath, outputPath string, tool any) error {
	tmpl, err := template.ParseFiles(templatePath);
	if err != nil {
		return err;
	}

	out, err := os.Create(outputPath);
	if err != nil {
		return err;
	}
	defer out.Close();

	return tmpl.Execute(out, tool);
}
