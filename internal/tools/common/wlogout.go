package common

import "parsertry/internal/renderer"

type Wlogout struct {
	WindowBg string `json:"window-bg"`
	ButtonBg string `json:"button-bg"`
	ButtonFocusBg string `json:"button-focus-bg"`
	ButtonFocusBorder string `json:"button-focus-border"`
}

func (t Wlogout) Render(templatePath, outputPath string) error {
	return renderer.Render(templatePath, outputPath, t);
}
