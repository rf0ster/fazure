package forms

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
var focusedTabStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("15")).
	Bold(true).
	Padding(0, 1)

var tabStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("240")).
	Padding(0, 1)

// Tabs implements a tabbed interface for switching between multiple FormFields.
// It is a FormField itself, just forwarding the appropriate calls to the currently focused field.
type Tabs struct {
	label        string
	labels       []string
	fields       []FormField
	focusedIndex int
	isFocused	 bool
	isEditing    bool
}

func NewTabs(label string, labels []string, fields []FormField) *Tabs {
	if len(labels) != len(fields) {
		panic("number of labels must match number of fields")
	}
	fields[0].Focus()
	return &Tabs{
		label:  label,
		labels: labels,
		fields: fields,
	}
}

func (t *Tabs) Label() string {
	return t.label
}

func (t *Tabs) Update(form *Form, msg tea.Msg) tea.Cmd {
	if t.isEditing {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				t.isEditing = false
				t.fields[t.focusedIndex].Save()
				return nil
			case t.fields[t.focusedIndex].Terminator():
				t.isEditing = false
				t.fields[t.focusedIndex].Save()
				return nil
			default:
				return t.fields[t.focusedIndex].Update(form, msg)
			}
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h", "left":
			return t.focusPrev()
		case "l", "right":
			return t.focusNext()
		case "enter":
			t.isEditing = true
			return t.fields[t.focusedIndex].Edit()
		}
	}

	return nil
}

func (t *Tabs) View(form *Form) string {
	output := ""
	for _, label := range t.labels {
		if t.labels[t.focusedIndex] == label && t.isFocused {
			output += focusedTabStyle.Render(" " + label + " ")
		} else {
			output += tabStyle.Render(" " + label + " ")
		}
	}

	output += t.fields[t.focusedIndex].View(form)
	return output
}

func (t *Tabs) Focus() tea.Cmd {
	t.isFocused = true
	return t.fields[t.focusedIndex].Focus()
}

func (t *Tabs) Blur() {
	t.isFocused = false
	t.fields[t.focusedIndex].Blur()
}

func (t *Tabs) Edit() tea.Cmd {
	t.isEditing = true
	return t.fields[t.focusedIndex].Edit()
}

func (t *Tabs) Save() {
	t.isEditing = false
	t.fields[t.focusedIndex].Save()
}

func (t *Tabs) Terminator() string {
	return t.fields[t.focusedIndex].Terminator()
}

func (t *Tabs) focusNext() tea.Cmd {
	t.fields[t.focusedIndex].Blur()
	t.focusedIndex = (t.focusedIndex + 1) % len(t.fields)
	return t.fields[t.focusedIndex].Focus()
}

func (t *Tabs) focusPrev() tea.Cmd {
	t.fields[t.focusedIndex].Blur()
	t.focusedIndex = (t.focusedIndex - 1 + len(t.fields)) % len(t.fields)
	return t.fields[t.focusedIndex].Focus()
}
