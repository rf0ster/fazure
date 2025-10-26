// Package forms is a package for creating interactive forms
// in terminal applications using the Bubble Tea framework.
package forms

import (
	tea "github.com/charmbracelet/bubbletea"
)

type FormField interface {
	Update(form *Form, msg tea.Msg) tea.Cmd
	View(form *Form) string

	Label() string
	Terminator() string

	Focus() tea.Cmd
	Blur()

	Edit() tea.Cmd
	Save()
}

type Form struct {
	title        string
	fields       []FormField
	focusedIndex int
	isEditing    bool
	labelPad     int
}

func NewForm(fields ...FormField) *Form {
	fields[0].Focus()

	pad := 0
	for _, field := range fields {
		if len(field.Label()) > pad {
			pad = len(field.Label())
		}
	}
	pad += 2 // extra padding

	return &Form{
		fields:   fields,
		labelPad: pad,
	}
}

func (f *Form) Update(model tea.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	if f.isEditing {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				f.isEditing = false
				f.fields[f.focusedIndex].Save()
				return model, nil
			case f.fields[f.focusedIndex].Terminator():
				f.isEditing = false
				f.fields[f.focusedIndex].Save()
				return model, nil
			default:
				return model, f.fields[f.focusedIndex].Update(f, msg)
			}
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			return model, f.focusNext()
		case "k", "up":
			return model, f.focusPrev()
		case "enter":
			f.isEditing = true
			return model, f.fields[f.focusedIndex].Edit()
		}
	}

	return model, f.fields[f.focusedIndex].Update(f, msg)
}

func (f *Form) View() string {
	output := ""
	if f.title != "" {
		output += f.title + "\n\n"
	}

	for _, field := range f.fields {
		output += field.View(f) + "\n"
	}
	return output
}

func (f *Form) focusPrev() tea.Cmd {
	if f.focusedIndex > 0 {
		f.fields[f.focusedIndex].Blur()
		f.focusedIndex--
		return f.fields[f.focusedIndex].Focus()
	}
	return nil
}

func (f *Form) focusNext() tea.Cmd {
	if f.focusedIndex < len(f.fields)-1 {
		f.fields[f.focusedIndex].Blur()
		f.focusedIndex++
		return f.fields[f.focusedIndex].Focus()
	}
	return nil
}

func (f *Form) Pad(label string) string {
	for range f.labelPad - len(label) {
		label += " "
	}
	return label
}
