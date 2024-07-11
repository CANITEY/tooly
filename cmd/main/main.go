package main

import (
	"tooly/internal/ui"
)




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
