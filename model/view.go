package model

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type View interface {
	Init(*model) tea.Cmd
	View(*model) string
	Update(*model, tea.Msg) (tea.Model, tea.Cmd)
}

type LoginView struct {
	userInput textinput.Model
}

func (v *LoginView) Init(m *model) tea.Cmd {
	v.userInput = textinput.New()
	v.userInput.Placeholder = "Username"
	v.userInput.CharLimit = 32
	v.userInput.Width = 20
	v.userInput.Focus()

	return nil
}

func (v *LoginView) View(m *model) string {
	var sb strings.Builder
	sb.WriteString("Login View\n")
	sb.WriteString(v.userInput.View())
	return sb.String()
}

func (v *LoginView) Update(m *model, msg tea.Msg) (tea.Model, tea.Cmd) {
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

type BacklogView struct{}

func (v *BacklogView) Init(m *model) tea.Cmd {
	return nil
}

func (v *BacklogView) View(m *model) string {
	var sb strings.Builder
	sb.WriteString("Viewing Backlog\n")
	return sb.String()
}

func (v *BacklogView) Update(m *model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.view = &BacklogItemView{}
			return m, m.view.Init(m)
		case "esc":
			m.view = &LoginView{}
			return m, m.view.Init(m)
		}
	}

	return m, nil
}

type BacklogItemView struct{}

func (v *BacklogItemView) Init(m *model) tea.Cmd {
	return nil
}

func (v *BacklogItemView) View(m *model) string {
	var sb strings.Builder
	sb.WriteString("Viewing Backlog Item\n")
	return sb.String()
}

func (v *BacklogItemView) Update(m *model, msg tea.Msg) (tea.Model, tea.Cmd) {
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

