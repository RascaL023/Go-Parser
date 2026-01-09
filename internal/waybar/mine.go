package waybar

type Mine struct {
	WaybarBg string `json:"waybarBg"`
	WaybarFg string `json:"waybarFg"`

	ModulesBg string `json:"modulesBg"`
	TooltipBg string `json:"tooltipBg"`

	ModuleHoverFg string `json:"moduleHoverFg"`

	WorkspacesBtn       string `json:"workspacesBtn"`
	WorkspacesActiveBg  string `json:"workspacesActiveBg"`
	WorkspacesActiveFg  string `json:"workspacesActiveFg"`

	Accent1     string `json:"accent1"`
	Accent2     string `json:"accent2"`
	BorderColor string `json:"borderColor"`
}

