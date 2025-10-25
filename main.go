package main

import (
	"fazure/types"
	"fazure/views"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput    textinput.Model
	table        table.Model
	viewport     viewport.Model
	textarea     textarea.Model
	azureClient  *MockAzureClient
	searchResult []types.BacklogItem
	showResults  bool
	showDetail   bool
	editMode     bool
	selectedItem *types.BacklogItem
	err          error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter assignee name (e.g., john, sarah, mike, emma)"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 60

	return model{
		textInput:   ti,
		azureClient: NewMockAzureClient(),
		showResults: false,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			// Don't quit if in edit mode, let user explicitly cancel first
			if m.editMode {
				return m, nil
			}
			return m, tea.Quit

		case "esc":
			// Exit edit mode if editing
			if m.editMode {
				m.editMode = false
				return m, nil
			}
			// Navigate back through views
			if m.showDetail {
				// Go back from detail to table
				m.showDetail = false
				m.selectedItem = nil
				return m, nil
			} else if m.showResults {
				// Go back from table to search
				m.showResults = false
				m.textInput.SetValue("")
				m.textInput.Focus()
				return m, nil
			}

		case "e":
			// Enter edit mode when in detail view
			if m.showDetail && !m.editMode && m.selectedItem != nil {
				m.editMode = true
				// Initialize textarea with current description
				ta := textarea.New()
				ta.SetValue(m.selectedItem.Description)
				ta.Focus()
				ta.SetWidth(80)
				ta.SetHeight(10)
				m.textarea = ta
				return m, nil
			}

		case "ctrl+s":
			// Save changes when in edit mode
			if m.editMode && m.selectedItem != nil {
				// Update the description
				m.selectedItem.Description = m.textarea.Value()
				// Also update in the search results slice
				for i := range m.searchResult {
					if m.searchResult[i].ID == m.selectedItem.ID {
						m.searchResult[i].Description = m.textarea.Value()
						break
					}
				}
				m.editMode = false
				return m, nil
			}

		case "enter":
			if m.showDetail && !m.editMode {
				// Do nothing on enter in detail view
				return m, nil
			} else if m.showResults {
				// Open detail view for selected item
				selectedRow := m.table.Cursor()
				if selectedRow >= 0 && selectedRow < len(m.searchResult) {
					m.selectedItem = &m.searchResult[selectedRow]
					m.showDetail = true
					m.viewport = views.CreateCommentsViewport(m.selectedItem)
					return m, nil
				}
			} else if !m.editMode {
				// Perform search
				searchName := m.textInput.Value()
				if searchName != "" {
					m.searchResult = m.azureClient.SearchWorkItems(searchName)
					m.showResults = true
					m.table = views.CreateTable(m.searchResult)
				}
				return m, nil
			}
		}
	}

	// Update the appropriate component based on current state
	if m.editMode {
		m.textarea, cmd = m.textarea.Update(msg)
	} else if m.showDetail {
		m.viewport, cmd = m.viewport.Update(msg)
	} else if m.showResults {
		m.table, cmd = m.table.Update(msg)
	} else {
		m.textInput, cmd = m.textInput.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	if m.editMode && m.selectedItem != nil {
		return views.RenderEditView(m.selectedItem, m.textarea.View())
	}

	if m.showDetail && m.selectedItem != nil {
		return views.RenderDetailView(m.selectedItem, m.viewport.View())
	}

	if m.showResults {
		return views.RenderTableView(m.textInput.Value(), m.table.View(), len(m.searchResult) > 0)
	}

	// Show search input
	return views.RenderSearchView(m.textInput.View())
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
