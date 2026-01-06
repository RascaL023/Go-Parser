package main

import (
	"fmt"
	"log"

	"parsertry/internal/renderer"
)

func main() {
	t, err := renderer.LoadTheme("themes/amethyst/theme.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Kitty.Render(
		"assets/templates/kitty/kitty.conf.tmpl",
		"output/kitty/kitty.conf",
	); err != nil {
		log.Fatal(err)
	}

	if err := t.Hyprland.Render(
		"assets/templates/hypr/hyprland.conf.tmpl",
		"output/hypr/hyprland.conf",
	); err != nil {
		log.Fatal(err)
	}

	if err := t.Cava.Render(
		"assets/templates/cava/cava.tmpl",
		"output/cava/cava_extra",
	); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Render Done!")
}

