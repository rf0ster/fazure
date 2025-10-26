package views

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type LoginView struct {
	userInput textinput.Model
}

func (v *LoginView) Init(m *Model) tea.Cmd {
	v.userInput = textinput.New()
	v.userInput.Placeholder = "Username"
	v.userInput.CharLimit = 32
	v.userInput.Width = 20
	v.userInput.Focus()

	return nil
}

func (v *LoginView) View(m *Model) string {
	var sb strings.Builder
	sb.WriteString("Login View\n")
	sb.WriteString(v.userInput.View())
	return sb.String()
}

func (v *LoginView) Update(m *Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, tea.Quit
		case "enter":
			m.view = &BacklogView{}
			return m, m.view.Init(m)
		}
	}

	var cmd tea.Cmd
	v.userInput, cmd = v.userInput.Update(msg)
	return m, cmd
}

