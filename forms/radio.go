package forms

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var radioLabelStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("15"))

// RadioField implements FormField for a radio button selection.
type RadioField struct {
	label         string
	labelPad      int
	focused       bool
	editing       bool
	options       []string
	selectedIndex int
	horizontal    bool
}

func NewRadioField(label string, options []string, horizontal bool) *RadioField {
	return &RadioField{
		label:      label,
		options:    options,
		horizontal: horizontal,
	}
}

func (r *RadioField) Label() string {
	return r.label
}

func (r *RadioField) Update(form *Form, msg tea.Msg) tea.Cmd {
	if r.horizontal || !r.editing {
		return r.updateHorizontal(msg)
	}
	return r.updateVertical(msg)
}

func (r *RadioField) updateHorizontal(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h", "left":
			if r.selectedIndex > 0 {
				r.selectedIndex--
			}
		case "l", "right":
			if r.selectedIndex < len(r.options)-1 {
				r.selectedIndex++
			}
		}
	}
	return nil
}

func (r *RadioField) updateVertical(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if r.selectedIndex < len(r.options)-1 {
				r.selectedIndex++
			}
		case "k", "up":
			if r.selectedIndex > 0 {
				r.selectedIndex--
			}
		}
	}
	return nil
}

func (r *RadioField) View(form *Form) string {
	if r.horizontal {
		return r.viewHorizontal(form)
	}
	return r.viewVertical(form)
}

func (r *RadioField) viewHorizontal(form *Form) string {
	label := form.Pad(r.label + ":")
	option := r.options[r.selectedIndex]

	if !r.focused && !r.editing {
		output := label + option
		return output
	}

	if r.editing {
		all := ""
		for i, opt := range r.options {
			if i == r.selectedIndex {
				all += " (*) " + opt
			} else {
				all += " ( ) " + opt
			}
		}
		return radioLabelStyle.Render(label + all[1:])
	}

	output := label
	if r.selectedIndex > 0 {
		// Append left arrow nerd font icon to output
		output += "◀ "
	}
	output += option
	if r.selectedIndex < len(r.options)-1 {
		output += " ▶"
	} 
	return radioLabelStyle.Render(output)
}

func (r *RadioField) viewVertical(form *Form) string {
	label := form.Pad(r.label + ":")
	option := r.options[r.selectedIndex]

	if !r.focused && !r.editing {
		output := label + option
		return output
	}

	if r.editing {
		all := "\n"
		for i, opt := range r.options {
			if i == r.selectedIndex {
				all += " (*) " + opt + "\n"
			} else {
				all += " ( ) " + opt + "\n"
			}
		}
		return radioLabelStyle.Render(label + all)
	}

	output := label
	if r.selectedIndex > 0 {
		output += "◀ "
	}
	output += option
	if r.selectedIndex < len(r.options)-1 {
		output += " ▶"
	} 
	return radioLabelStyle.Render(output)
}

func (r *RadioField) Focus() tea.Cmd {
	r.focused = true
	return nil
}

func (r *RadioField) Blur() {
	r.focused = false
}

func (r *RadioField) Edit() tea.Cmd {
	r.editing = true
	return nil
}

func (r *RadioField) Save() {
	r.editing = false
}

func (r *RadioField) Terminator() string {
	return "enter"
}
