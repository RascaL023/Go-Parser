// package main
//
// import (
// 	"fmt"
// 	"log"
//
// 	"parsertry/internal/renderer"
// )
//
// func main() {
// 	fmt.Println("[+] Load Theme...")
// 	t, err := renderer.LoadTheme("themes/amethyst/theme.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Println("[+] Rendering Kitty...")
// 	if err := t.Kitty.Render(
// 		"assets/templates/kitty/kitty.conf.tmpl",
// 		"output/kitty/kitty.conf",
// 	); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Println("[+] Rendering Hyprland...")
// 	if err := t.Hyprland.Render(
// 		"assets/templates/hypr/hyprland.conf.tmpl",
// 		"output/hypr/hyprland.conf",
// 	); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Println("[+] Rendering Cava...")
// 	if err := t.Cava.Render(
// 		"assets/templates/cava/cava.tmpl",
// 		"output/cava/cava_extra",
// 	); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Println("Render Done!")
// }
//
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
	// 1. Load theme
	t, err := renderer.LoadTheme("themes/amethyst/theme.json")
	if err != nil {
		log.Fatal(err)
	}

	// 2. Registry renderer
	renderers := map[string]theme.Renderer{
		"kitty": &t.Kitty,
		"hypr":  &t.Hyprland,
		"cava":  &t.Cava,
	}

	// 3. Open paths.txt
	file, err := os.Open("assets/templates/paths.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// 4. ONE LOOP 🔥
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "$")

		if len(parts) != 3 {
			log.Fatalf("invalid line: %s", line)
		}

		tool := parts[0]
		templatePath := parts[1]
		outputPath := parts[2]

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

