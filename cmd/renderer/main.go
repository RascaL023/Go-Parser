package main

import (
	"os"

	"parsertry/cmd/helper"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(4)
	}

	os.Exit(run(os.Args[1], os.Args[2]))
}

func run(option, name string) int {
	ex := 0

	switch option {
	case "theme":
		ex = helper.RunTheme(name)
	case "waybar":
		ex = helper.RunWaybar(name)
	default: ex = 4
	}

	return ex
}

