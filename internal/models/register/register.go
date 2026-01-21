package register

import (
	"parsertry/internal/models/output"
	"parsertry/internal/renderer"
)

func RegisterTool(t output.ThemeOutput) map[string]renderer.Renderer {
	return map[string]renderer.Renderer{
		"foot"	 		: t.Foot,
		"kitty"	 		: t.Kitty,
	}
}

