package views

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	view View
}

func NewModel() Model {
	return Model{
		view: &LoginView{},
	}
}

func (m Model) Init() tea.Cmd {
	return m.view.Init(m)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			return m.view.Update(m, msg)
		}
	default:
		return m.view.Update(m, msg)
	}
}

func (m Model) View() string {
	return m.view.View(m)
}

