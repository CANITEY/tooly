package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices []string
	cursor int
	selected map[int]bool
}

func initialModel(toolList []string) model {
	return model{
		choices: toolList,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg :
		switch msg.String() {
		case "k":
			if !(m.cursor <= 0) {
				m.cursor--
			}
		case "j":
			if m.cursor < len(m.choices) - 1 {
				m.cursor++
			}
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "tool list\n\n"

	for i, toolName := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n",cursor, toolName)
	}

	return s
}



func main() {
	toolist := []string{
		"subfinder",
		"amass",
		"assetfinder",
		"nmap",
		"naabu",
	}
	model := initialModel(toolist)
    p := tea.NewProgram(model)
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
