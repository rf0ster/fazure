package views

import (
	"fazure/azure"
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)

type BacklogView struct {
	workItems []azure.WorkItem
	table     table.Model
}

func (v *BacklogView) Init(m Model) tea.Cmd {
	return func() tea.Msg {
		items, _ := m.azure.QueryWorkItems(azure.QueryParams{
			AssignedTo: m.user,
			State:      "Active",
		})
		return items
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
		v.table = createTable(v.workItems, m)
	}

	v.table, cmd = v.table.Update(msg)
	return m, cmd
}

// CreateTable creates and configures a table from backlog items
func createTable(items []azure.WorkItem, m Model) table.Model {
	w := m.terminalWidth - 8 // Adjust for padding/margin
	h := m.terminalHeight - 10

	calcW := func(perc float64) int {
		return int(float64(w) * perc)
	}

	columns := []table.Column{
		{Title: "ID", Width: calcW(0.08)},
		{Title: "Type", Width: calcW(0.12)},
		{Title: "Title", Width: calcW(0.45)},
		{Title: "Assigned To", Width: calcW(0.15)},
		{Title: "State", Width: calcW(0.12)},
		{Title: "Priority", Width: calcW(0.08)},
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
		table.WithHeight(h),
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
