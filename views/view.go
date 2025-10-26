// Package views defines the interface for different views in the application.
package views

import tea "github.com/charmbracelet/bubbletea"

type View interface {
	Init(Model) tea.Cmd
	View(Model) string
	Update(Model, tea.Msg) (tea.Model, tea.Cmd)
}

func getContentWidth(width int) int {
	contentWidth := width / 2
	if contentWidth > 100 {
		contentWidth = 100 // Cap at 100 for very wide terminals
	}

	if contentWidth < 40 {
		contentWidth = 40 // Minimum width for readability
	}
	return contentWidth
}
