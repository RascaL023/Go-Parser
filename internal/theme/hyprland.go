package theme

import (
	"os"
	"text/template"
)

type Hyprland struct {
	Accent string `json:"accent"`
	AccentAlpha string `json:"accentAlpha"`
	PrimaryFont string `json:"primaryFont"`
	SecondaryFont string `json:"secondaryFont"`

	GapsIn int `json:"gapsIn"`
	GapsOut int `json:"gapsOut"`
	BorderSize int `json:"borderSize"`
	ColActiveBorder1 string `json:"colActiveBorder1"`
	ColActiveBorder2 string `json:"colActiveBorder2"`
	ColActiveBorderDeg string `json:"colActiveBorderDeg"`
	ColInactiveBorder string `json:"colInactiveBorder"`

	Blur bool `json:"blur"`
	BlurSize int `json:"blurSize"`
	BlurPasses int `json:"blurPasses"`
	Rounding int `json:"rounding"`
	Shadow bool `json:"shadow"`
	ShadowRange int `json:"shadowRange"`
	ShadowRenderPower int `json:"shadowRenderPower"`
	ShadowColor string `json:"shadowColor"`

	FootOpacity float64 `json:"footOpacity"`
}

func (tool Hyprland) Render(templatePath, outputPath string) error {
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
