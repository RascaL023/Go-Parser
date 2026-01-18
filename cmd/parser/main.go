package main

import (
	"fmt"
	"parsertry/internal/loader"
	"parsertry/internal/renderer"
	"path/filepath"
)

func main() {
	assetsMap := "assets/map";

	state, err := loader.LoadState(assetsMap + "/state.json");
	if err != nil {
		panic("Error on parsing state");
	}
	fmt.Println("[+] Using state", state.Version);

	tools, err := loader.LoadToolMap(assetsMap + "/path.txt", state);
	if err != nil {
		panic("Map path not found!");
	}

	themeName 	 := state.Theme;
	waybarPreset := state.Waybar;

	theme, err := loader.LoadToolFromTheme(filepath.Join("themes", themeName, "theme.json"));
	if err != nil {
		panic(fmt.Sprintf("Theme %s not found!", themeName));
	}

	wb, err := loader.LoadWaybar(filepath.Join("themes", themeName, "waybar.json"));
	if err != nil {
		panic(fmt.Sprintf("Waybar Preset %s error!", themeName));
	}



	registerTools := map[string]renderer.Renderer {
		"wlogout": theme.Wlogout,
		"cava"	 : theme.Cava,
		"waybar" : renderer.WaybarRender{
			Waybar: *wb,
			Preset: waybarPreset,
		},
	}


	for _, tool := range tools {
		reg, ok := registerTools[tool.Name];
		if !ok {
			fmt.Println("Unknown tool:", tool.Name);
			continue;
		}

		if err := reg.Render(tool.TemplatePath, tool.OutputPath); err != nil {
			panic(err);
		}
	}
}
