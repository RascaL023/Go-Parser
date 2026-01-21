package main

import (
	"fmt"
	"os"
	"parsertry/internal/loader"
	"parsertry/internal/models/input"
	"parsertry/internal/models/output"
	"parsertry/internal/models/register"
	"parsertry/internal/models/resolver"
	"parsertry/internal/state"
	"path/filepath"
)

func main() {
	// ================================ FIX ================================
	// Get the state of theme
	assetsMap := os.ExpandEnv("$MYENV/map");
	state, err := loader.LoadJSON[state.State](assetsMap + "/.state.json");
	if err != nil {
		panic("Error on parsing state!");
	}
	fmt.Println("[Go] Using state", state.Version);


	// Get the input - output path of each tools
	tools, err := loader.LoadToolMap(assetsMap + "/path.txt", state);
	if err != nil {
		panic("Error on parsing map path!");
	}
	// ================================ FIX ================================


	// Get all the config of each tools from theme.json(not the rendered one!)
	themeName 	 := state.Theme.Name;
	theme, err := loader.LoadJSON[input.ThemeInput](filepath.Join("themes", themeName, "theme.json"));
	if err != nil {
		panic(fmt.Sprintf("[Go] Theme %s not found!", themeName));
	}

	
	resolved := resolver.ResolveTerminal(theme.Terminal);
	foot := resolver.ResolveFoot(resolved);
	kitty := resolver.ResolveKitty(resolved, theme.Terminal);
	toRegistered := output.ThemeOutput{
		Foot: foot,
		Kitty: kitty,
	}

	registeredTools := register.RegisterTool(toRegistered);
	for _, tool := range tools {
		reg, ok := registeredTools[tool.Name];
		if !ok {
			fmt.Println("Unknown tool:", tool.Name);
			continue;
		}

		if err := reg.Render(tool.TemplatePath, tool.OutputPath); err != nil {
			panic(err);
		}
	}
}
