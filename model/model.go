// Package model handles data structures and ELM models used in the application.
package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	view View
}

func InitialModel() model {
	v := LoginView{}
	return model{
		view: &v,
	}
}

func (m model) Init() tea.Cmd {
	return m.view.Init(&m)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			return m.view.Update(&m, msg)
		}
	default:
		return m.view.Update(&m, msg)
	}
}

func (m model) View() string {
	return m.view.View(&m)
}

