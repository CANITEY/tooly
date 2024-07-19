package main

import (
	"tooly/internal/ui"
)



// TODO: Create the debian routine first
func main() {
	toolist := []string{
		"subfinder",
		"amass",
		"assetfinder",
		"nmap",
		"naabu",
	}

	model := ui.InitialModel(toolist)
	ui.Run(model)
}
