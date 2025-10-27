package views

import (
	"fazure/azure"
	"fazure/forms"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type DetailsView struct {
	item *azure.WorkItem
	form *forms.Form
}

func (v *DetailsView) Init(m Model) tea.Cmd {
	v.form = forms.NewForm(
		forms.NewRadioField("Assigned To", []string{"Unassigned", "John Doe", "Jane Smith", "Alice Johnson", "Bob Brown"}, true),
		forms.NewRadioField("State", []string{"New", "Active", "Resolved", "Closed"}, true),
		forms.NewRadioField("Priority", []string{"1", "2", "3", "4", "5"}, true),
		forms.NewReadonly("Iteration Path", v.item.Iteration),
		forms.NewReadonly("Area Path", v.item.AreaPath),
		forms.NewReadonly("Created By", v.item.CreatedBy),
		forms.NewReadonly("Created Date", v.item.CreatedDate),
		forms.NewTabs("",
			[]string{"Description", "Acceptance Criteria", "Discussion"},
			[]forms.FormField{
				forms.NewTextAreaField("", v.item.Description, true),
				forms.NewTextAreaField("", v.item.AcceptanceCriteria, true),
				forms.NewTextAreaField("", "Discussion notes", true),
			},
		),
	)

	return nil
}

func (v *DetailsView) View(m Model) string {
	if v.item == nil {
		return "No item selected"
	}
	item := v.item

	// Display header
	var s strings.Builder
	header := fmt.Sprintf("%s #%d", item.Type, item.ID)
	s.WriteString(GetWorkItemTypeStyle(item.Type).Render(header))
	s.WriteString("\n")

	// Displat work item title
	wrappedTitle := wrapText(v.item.Title, getContentWidth(m.terminalWidth))
	s.WriteString(GetWorkItemTypeStyle(item.Type).Render(wrappedTitle))
	s.WriteString("\n\n")

	s.WriteString(v.form.View())
	return s.String()
}

func (v *DetailsView) Update(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := v.form.Update(m, msg)
	return m, cmd
}
