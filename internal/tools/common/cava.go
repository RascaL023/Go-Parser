package common

import "parsertry/internal/renderer"

type Cava struct {
	Gradient1 string `json:"gradient1"`
	Gradient2 string `json:"gradient2"`
	Gradient3 string `json:"gradient3"`
	Gradient4 string `json:"gradient4"`
	Gradient5 string `json:"gradient5"`
}

func (c Cava) Render(templatePath, OutputPath string) error {
	return renderer.Render(templatePath, OutputPath, c);
}
