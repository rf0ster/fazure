package views

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type BacklogView struct{}

func (v *BacklogView) Init(m *Model) tea.Cmd {
	return nil
}

func (v *BacklogView) View(m *Model) string {
	var sb strings.Builder
	sb.WriteString("Viewing Backlog\n")
	return sb.String()
}

func (v *BacklogView) Update(m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.view = &DetailsView{}
			return m, m.view.Init(m)
		case "esc":
			m.view = &LoginView{}
			return m, m.view.Init(m)
		}
	}

	return m, nil
}


