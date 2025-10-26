package forms

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var textareadLabelStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("15"))

var textareadHelpStyle = lipgloss.NewStyle().
	Italic(true).
	Foreground(lipgloss.Color("241"))

// TextAreaField implements FormField for a text area input.
type TextAreaField struct {
	label      string
	focused    bool
	editing    bool
	textarea   textarea.Model
	alwaysShow bool
}

func NewTextAreaField(label string, content string, alwaysShow bool) *TextAreaField {
	ta := textarea.New()
	ta.SetWidth(50)
	ta.SetHeight(10)
	ta.SetValue(content)
	ta.ShowLineNumbers = false

	return &TextAreaField{
		label:      label,
		textarea:   ta,
		alwaysShow: alwaysShow,
	}
}

func (t *TextAreaField) Label() string {
	return t.label
}

func (t *TextAreaField) Update(form *Form, msg tea.Msg) tea.Cmd {
	if t.editing {
		var cmd tea.Cmd
		t.textarea, cmd = t.textarea.Update(msg)
		return cmd
	}
	return nil
}

func (t *TextAreaField) View(form *Form) string {
	output := ""
	if t.label != "" {
		output += t.label + ":"
	}

	if t.focused {
		output = textareadLabelStyle.Render(output)
	}

	if t.focused || t.editing || t.alwaysShow {
		output += "\n\n"
		output += t.textarea.View()
	}
	if t.editing {
		helpText := textareadHelpStyle.Render("(Press " + t.Terminator() + " to save, esc to cancel)")
		output += "\n" + helpText + "\n"
	} else if t.focused {
		helpText := textareadHelpStyle.Render("(Press enter to edit)")
		output += "\n" + helpText + "\n"
	}

	return output
}

func (t *TextAreaField) Focus() tea.Cmd {
	t.focused = true
	return nil
}

func (t *TextAreaField) Blur() {
	t.focused = false
	t.textarea.Blur()
}

func (t *TextAreaField) Edit() tea.Cmd {
	t.editing = true
	return t.textarea.Focus()
}

func (t *TextAreaField) Save() {
	t.editing = false
	t.textarea.Blur()
}

func (t *TextAreaField) Terminator() string {
	return "ctrl+s"
}
