package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)


type model struct {
	choices []string
	cursor int
	selected map[int]bool
	next bool
}

func InitialModel(toolList []string) model {
	return model{
		choices: toolList,
		selected: make(map[int]bool),
	}
}

func Run(m model) error {
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg :
		switch msg.String() {
		case "enter":
			m.next = true
			return m, nil
		case "k", "up":
			if !(m.cursor <= 0) {
				m.cursor--
			}
		case "j", "down":
			if m.cursor < len(m.choices) - 1 {
				m.cursor++
			}
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				m.selected[m.cursor] = false
			} else {
				m.selected[m.cursor] = true
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "tool list\n\n"
	if m.next == true {
		output := sliceToMap(m.choices, m.selected)
		return fmt.Sprintf("output: %+v", output)
	}

	for i, toolName := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n",cursor, toolName)
	}

	return s
}
func sliceToMap(lside []string, rside map[int]bool) map[string]bool {
	output := map[string]bool{}
	for i, item := range lside {
		output[item] = rside[i]
	}

	return output
}
