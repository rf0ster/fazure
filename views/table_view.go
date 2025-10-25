package views

import (
	"fazure/types"
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

// CreateTable creates and configures a table from backlog items
func CreateTable(items []types.BacklogItem) table.Model {
	columns := []table.Column{
		{Title: "ID", Width: 8},
		{Title: "Type", Width: 12},
		{Title: "Title", Width: 45},
		{Title: "Assigned To", Width: 15},
		{Title: "State", Width: 12},
		{Title: "Priority", Width: 8},
	}

	rows := []table.Row{}
	for _, item := range items {
		rows = append(rows, table.Row{
			strconv.Itoa(item.ID),
			string(item.Type),
			item.Title,
			item.AssignedTo,
			item.State,
			strconv.Itoa(item.Priority),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(15),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)
	return t
}

// RenderTableView renders the work items table screen
func RenderTableView(searchName string, tableView string, hasResults bool) string {
	var s string
	s += types.TitleStyle.Render(fmt.Sprintf("Work Items for: %s", searchName))
	s += "\n\n"

	if !hasResults {
		s += "No work items found for this user.\n\n"
	} else {
		s += tableView
		s += "\n\n"
	}

	s += types.HelpStyle.Render("Press 'enter' to view details • 'esc' to search again • 'q' to quit")
	return s
}
