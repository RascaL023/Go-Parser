package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"parsertry/internal/renderer"
	"parsertry/internal/theme"
)

func main() {
	t, err := renderer.LoadTheme("themes/amethyst/theme.json")
	if err != nil {
		log.Fatal(err)
	}

	wb, err := renderer.LoadWaybar("themes/amethyst/waybar.json")
	if err != nil {
		log.Fatal(err)
	}


	renderers := map[string]theme.Renderer{
		"kitty": &t.Kitty,
		"hypr":  &t.Hyprland,
		"cava":  &t.Cava,
		"waybar": &renderer.WaybarRenderer{
			Waybar: wb,
			Preset: "mechabox",
		},
	}


	file, err := os.Open("assets/templates/paths.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "$")

		if len(parts) != 3 {
			log.Fatalf("invalid line: %s", line)
		}
		
		tool := parts[0]
		templatePath := parts[1]
		outputPath := parts[2]
		fmt.Printf("[+] Rendering %s...\n", tool)

		r, ok := renderers[tool]
		if !ok {
			log.Fatalf("unknown renderer: %s", tool)
		}

		if err := r.Render(templatePath, outputPath); err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("[+] Render Completed!")
}

