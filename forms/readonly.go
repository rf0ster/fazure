package forms

import (
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)

var readonlyLabelStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("15"))

// Readonly implements FormField but does is a readonly field.
type Readonly struct {
	label string
	value string
	focused bool
}

func NewReadonly(label string, value string) *Readonly {
	return &Readonly{
		label: label,
		value: value,
	}
}

func (r *Readonly) Update(form *Form, msg tea.Msg) tea.Cmd {
	return nil
}

func (r *Readonly) View(form *Form) string {
	output := form.Pad(r.label + ":") + r.value
	if r.focused {
		 return readonlyLabelStyle.Render(output)
	}
	return output
}

func (r *Readonly) Label() string {
	return r.label
}

func (r *Readonly) Terminator() string {
	return ""
}

func (r *Readonly) Focus() tea.Cmd {
	r.focused = true
	return nil
}

func (r *Readonly) Blur() {
	r.focused = false
}

func (r *Readonly) Edit() tea.Cmd {
	return nil
}

func (r *Readonly) Save() {
}

