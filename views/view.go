// Package views defines the interface for different views in the application.
package views

import tea "github.com/charmbracelet/bubbletea"

type View interface {
	Init(*Model) tea.Cmd
	View(*Model) string
	Update(*Model, tea.Msg) (tea.Model, tea.Cmd)
}
