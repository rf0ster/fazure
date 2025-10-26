package views

import (
	"fazure/azure"
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type BacklogView struct {
	workItems []azure.WorkItem
	table     table.Model
}

func (v *BacklogView) Init(m Model) tea.Cmd {
	return func() tea.Msg {
		return m.azure.SearchWorkItems(m.user)
	}
}

func (v *BacklogView) View(m Model) string {
	var s string
	s += TitleStyle.Render(fmt.Sprintf("Backlog: %s", m.user))
	s += "\n\n"

	if len(v.workItems) == 0 {
		s += "No work items found for this user.\n\n"
	} else {
		s += v.table.View()
		s += "\n\n"
	}

	s += HelpStyle.Render("Press 'enter' to view details • 'esc' to search again • 'q' to quit")
	return s
}

func (v *BacklogView) Update(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.view = &DetailsView{
				item: v.GetSelectedWorkItem(),
			}
			return m, m.view.Init(m)
		case "esc":
			m.view = &LoginView{}
			return m, m.view.Init(m)
		}
	case []azure.WorkItem:
		v.workItems = msg
		v.table = createTable(v.workItems)
	}

	v.table, cmd = v.table.Update(msg)
	return m, cmd
}

// CreateTable creates and configures a table from backlog items
func createTable(items []azure.WorkItem) table.Model {
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

func (v *BacklogView) GetSelectedWorkItem() *azure.WorkItem {
	if len(v.workItems) == 0 {
		return nil
	}
	selectedIndex := v.table.Cursor()
	return &v.workItems[selectedIndex]
}
