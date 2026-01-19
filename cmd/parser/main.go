package main

import (
	"fmt"
	"os"
	"parsertry/internal/loader"
	"parsertry/internal/tools/common"
	"path/filepath"
)

func main() {
	assetsMap := os.ExpandEnv("$MYENV/map");
	state, err := loader.LoadState(assetsMap + "/.state.json");
	if err != nil {
		panic("Error on parsing state!");
	}
	fmt.Println("[+] Using state", state.Version);

	tools, err := loader.LoadToolMap(assetsMap + "/path.txt", state);
	if err != nil {
		panic("Error on parsing map path!");
	}


	themeName 	 := state.Theme.Name;
	theme, err := loader.LoadToolFromTheme(filepath.Join("themes", themeName, "theme.json"));
	if err != nil {
		panic(fmt.Sprintf("Theme %s not found!", themeName));
	}

	waybar, err := loader.LoadWaybar(filepath.Join("themes", themeName, "waybar.json"));
	if err != nil {
		panic(fmt.Sprintf("Waybar Preset %s error!", themeName));
	}


	registerTools := common.RegisterTool(*theme, *state, waybar);
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
