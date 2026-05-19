package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	urls    []string
	results []string
	loading bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if len(m.results) == 0 {
		return "Kontrol ediliyor...\n"
	}
	output := "URL Check Results:\n\n"
	for _, r := range m.results {
		output += r
	}
	return output
}

func NewModel(urls []string, results []string) model {
	return model{
		urls:    urls,
		results: results,
		loading: false,
	}
}
