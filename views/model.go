package views

import (
	"fazure/azure"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	view           View
	user           string
	azure          azure.MockAzureClient
	terminalWidth  int
	terminalHeight int
}

func NewModel() Model {
	return Model{
		user:  "john",
		view:  &BacklogView{},
		azure: azure.MockAzureClient{},
	}
}

func (m Model) Init() tea.Cmd {
	return m.view.Init(m)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		return m, nil
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
