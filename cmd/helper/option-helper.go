package helper

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"parsertry/internal/renderer"
	"parsertry/internal/theme"
)

func RunTheme(name string) int {
	t, err := renderer.LoadTheme(filepath.Join("themes", name, "theme.json"))
	if err != nil {
		return Fatal(err, 1)
	}

	wb, err := renderer.LoadWaybar(filepath.Join("themes", name, "waybar.json"))
	if err != nil {
		return Fatal(err, 1)
	}

	state, err := renderer.LoadState("assets/state.json")
	if err != nil {
		return Fatal(err, 1)
	}

	renderers := map[string]theme.Renderer{
		"kitty":  &t.Kitty,
		"hypr":   &t.Hyprland,
		"cava":   &t.Cava,
		"waybar": &renderer.WaybarRenderer{
			Waybar: wb,
			Preset: state.Waybar,
		},
		"foot": &t.Foot,
	}

	file, err := os.Open("assets/templates/paths.txt")
	if err != nil {
		return Fatal(err, 1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "$")

		if len(parts) != 3 {
			return Fatal(fmt.Errorf("invalid line: %s", line), 2)
		}

		tool := parts[0]
		templatePath := parts[1]
		outputPath := parts[2]

		fmt.Printf("[+] Rendering %s...\n", tool)

		r, ok := renderers[tool]
		if !ok {
			return Fatal(fmt.Errorf("unknown renderer: %s", tool), 2)
		}

		if err := r.Render(templatePath, outputPath); err != nil {
			return Fatal(err, 1)
		}
	}

	if err := scanner.Err(); err != nil {
		return Fatal(err, 1)
	}

	fmt.Println("[+] Render Completed!")
	return 0
}

func RunWaybar(name string) int {
	state, err := renderer.LoadState(filepath.Join("assets", "state.json"))
	if err != nil {
		return Fatal(err, 4)
	} else if state.Waybar == name {
		fmt.Println("Waybar", name, "already active...")
		return 0
	}

	wb, err := renderer.LoadWaybar(filepath.Join("themes", state.Theme, "waybar.json"))
	if err != nil {
		return Fatal(err, 1)
	}

	render := map[string]theme.Renderer{
		"waybar":  &renderer.WaybarRenderer{
			Waybar: wb,
			Preset: name,
		},
	}

	file, err := os.Open("assets/templates/paths.txt")
	if err != nil {
		return Fatal(err, 1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "$")
		if parts[0] != "waybar" {
			continue
		}

		if err := render[parts[0]].Render(parts[1], parts[2]); err != nil {
			return Fatal(err, 1)
		}

		break
	}

	fmt.Println("[+] Waybar updated to:", name)
	return 0
}
