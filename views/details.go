package views

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type DetailsView struct{}

func (v *DetailsView) Init(m *Model) tea.Cmd {
	return nil
}

func (v *DetailsView) View(m *Model) string {
	var sb strings.Builder
	sb.WriteString("Viewing Backlog Item\n")
	return sb.String()
}

func (v *DetailsView) Update(m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.view = &BacklogView{}
			return m, m.view.Init(m)
		}
	}

	return m, nil
}

