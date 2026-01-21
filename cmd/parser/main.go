package main

import (
	"fmt"
	"os"
	"parsertry/internal/loader"
	"parsertry/internal/state"
	"parsertry/internal/tools/common"
	"parsertry/internal/tools/common/kitty"
	"parsertry/internal/tools/waybar"
	"path/filepath"
)

func main() {
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


	// Get all the config of each tools from theme.json(not the rendered one!)
	themeName 	 := state.Theme.Name;
	theme, err := loader.LoadJSON[common.Tool](filepath.Join("themes", themeName, "theme.json"));
	if err != nil {
		panic(fmt.Sprintf("[Go] Theme %s not found!", themeName));
	}


	// Get the waybar json by the theme
	waybar, err := loader.LoadJSON[waybar.Waybar](filepath.Join("themes", themeName, "waybar.json"));
	if err != nil {
		panic(fmt.Sprintf("Waybar Preset %s error!", themeName));
	}

	// Build here!
	var renderTool common.RenderTool;
	renderTool.Kitty = kitty.BuildKitty(theme.Kitty);
	

	// Render the tools
	registerTools := common.RegisterTool(renderTool , *state, waybar);
	for _, tool := range tools {
		reg, ok := registerTools[tool.Name];
		if !ok {
			fmt.Println("Unknown tool:", tool.Name);
			continue;
		}
		fmt.Println("Rendering", tool.Name);

		if err := reg.Render(tool.TemplatePath, tool.OutputPath); err != nil {
			panic(err);
		}
	}
}
