package views

import (
	"fazure/azure"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	view           View
	user           string
	azure          *azure.AzureClient
	terminalWidth  int
	terminalHeight int
}

func NewModel() Model {
	return Model{
		user:  os.Getenv("AZURE_USER"),
		view:  &BacklogView{},
		azure: azure.NewClient(
			os.Getenv("AZURE_ORG"), 
			os.Getenv("AZURE_PROJECT"), 
			os.Getenv("AZURE_PAT")),
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
		case "ctrl+c":
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
