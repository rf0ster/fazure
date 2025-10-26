package views

import (
	"fazure/azure"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type DetailsView struct {
	item *azure.WorkItem
}

func (v *DetailsView) Init(m Model) tea.Cmd {
	return nil
}

func (v *DetailsView) View(m Model) string {
	if v.item == nil {
		return "No item selected"
	}

	var s strings.Builder
	header := fmt.Sprintf("%s #%d", v.item.Type, v.item.ID)
	s.WriteString(GetWorkItemTypeStyle(v.item.Type).Render(header))
	s.WriteString("\n\n")

	return s.String()
}

func (v *DetailsView) Update(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
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
