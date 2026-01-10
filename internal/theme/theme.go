package theme

type Theme struct {
	Kitty Kitty `json:"kitty"`
	Cava Cava `json:"cava"`
	Hyprland Hyprland `json:"hyprland"`
	Foot Foot `json:"foot"`
}

type Renderer interface {
	Render(templatePath, outputPath string) error
}
