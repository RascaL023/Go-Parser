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

	GapsIn int8 `json:"gapsIn"`
	GapsOut int8 `json:"gapsOut"`
	BorderSize int8 `json:"borderSize"`
	ColActiveBorder1 string `json:"colActiveBorder1"`
	ColActiveBorder2 string `json:"colActiveBorder2"`
	ColActiveBorderDeg string `json:"colActiveBorderDeg"`
	ColInactiveBorder string `json:"colInactiveBorder"`

	Blur bool `json:"blur"`
	BlurSize int8 `json:"blurSize"`
	BlurPasses int8 `json:"blurPasses"`
	Rounding int8 `json:"rounding"`
	Shadow bool `json:"shadow"`
	ShadowRange int8 `json:"shadowRange"`
	ShadowRenderPower int8 `json:"shadowRenderPower"`
	ShadowColor string `json:"shadowColor"`

	FootOpacity float32 `json:"footOpacity"`
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
