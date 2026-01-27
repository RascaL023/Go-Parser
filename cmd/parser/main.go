package main

import (
	"fmt"
	"os"
	"parsertry/internal/context"
	"parsertry/internal/domain/terminal"
	"parsertry/internal/loader"
	"parsertry/internal/register"

	"parsertry/internal/state"
	"parsertry/internal/theme"
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

	theme, err := loader.LoadJSON[theme.Theme](
		filepath.Join("themes", state.Theme.Name, "theme.json"),
	);
	if err != nil {
		fmt.Println(err);
	}

	// ================================ Domain ================================
	rawTerminal := theme.Theme["terminal"];
	var terminalProc terminal.TerminalProcessor;

	parsedTerminal, _ := terminalProc.Parse(rawTerminal);
		resolvedTerminal, _ := terminalProc.Resolve(parsedTerminal);

	ctx := &context.Context{
		Terminal: resolvedTerminal,
	}
	// ================================ Domain ================================

	for toolName, raw := range theme.Theme {
		processor, ok := register.Get(toolName);
		if !ok {
			fmt.Printf("Unknown %s!\n", toolName);
			continue;
		}

		toolPath, ok := tools[toolName];
		if !ok {
			fmt.Printf("Path for %s not found!", toolName);
			continue;
		}

		parsed, err := processor.Parse(raw);
		if err != nil {
			panic(err);
		}

		resolved, err := processor.Resolve(parsed, ctx);

		// processor.Render(toolPath.TemplatePath, toolPath.OutputPath, resolved);
		if r, ok := processor.(register.Renderer); ok {
			r.Render(toolPath.TemplatePath, toolPath.OutputPath, resolved);
		}
	}
}
