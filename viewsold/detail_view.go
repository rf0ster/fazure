// Package viewsold contains modules for rendering different parts of the application.
package viewsold

import (
	"fazure/azure"

	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

// CreateCommentsViewport creates a viewport for scrollable comments
func CreateCommentsViewport(item *azure.WorkItem, width int) viewport.Model {
	vp := viewport.New(width, 20)
	vp.SetContent(RenderComments(item, width))
	return vp
}

// RenderComments renders the comments section as a string
func RenderComments(item *azure.WorkItem, width int) string {
	if item == nil || len(item.Comments) == 0 {
		return "No comments yet."
	}

	var content strings.Builder

	for _, comment := range item.Comments {
		// Wrap comment content to width
		wrappedContent := wrapText(comment.Content, width-4) // -4 for padding

		commentBox := fmt.Sprintf("%s  %s\n\n%s",
			azure.CommentAuthorStyle.Render(comment.Author),
			azure.CommentDateStyle.Render(comment.Date),
			wrappedContent,
		)
		content.WriteString(azure.CommentStyle.Render(commentBox))
		content.WriteString("\n")
	}

	return content.String()
}

// wrapText wraps text to the specified width
func wrapText(text string, width int) string {
	if width <= 0 {
		return text
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	var lines []string
	var currentLine strings.Builder

	for _, word := range words {
		// If adding this word would exceed width, start a new line
		if currentLine.Len() > 0 && currentLine.Len()+1+len(word) > width {
			lines = append(lines, currentLine.String())
			currentLine.Reset()
		}

		if currentLine.Len() > 0 {
			currentLine.WriteString(" ")
		}
		currentLine.WriteString(word)
	}

	if currentLine.Len() > 0 {
		lines = append(lines, currentLine.String())
	}

	return strings.Join(lines, "\n")
}

// Field index constants - must match main.go
const (
	FieldDescription = 0
	FieldAddComment  = 1
)

// RenderDetailView renders the full detail view for a work item
func RenderDetailView(item *azure.WorkItem, selectedFieldIndex int, isDirty bool, width int, viewportView string) string {
	if item == nil {
		return "No item selected"
	}

	var s strings.Builder

	// Header with ID and Type
	header := fmt.Sprintf("%s #%d", item.Type, item.ID)
	if isDirty {
		// Add dirty indicator
		dirtyStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("214")).
			Render(" ●")
		header += dirtyStyle
	}
	s.WriteString(azure.HeaderStyle.Render(header))
	s.WriteString("\n\n")

	// Title - wrap to width
	wrappedTitle := wrapText(item.Title, width-2)
	s.WriteString(azure.TitleDetailStyle.Render(wrappedTitle))
	s.WriteString("\n\n")

	// Fields section
	s.WriteString(azure.FieldLabelStyle.Render("State:") + " " + azure.FieldValueStyle.Render(item.State) + "\n")
	s.WriteString(azure.FieldLabelStyle.Render("Assigned To:") + " " + azure.FieldValueStyle.Render(item.AssignedTo) + "\n")
	s.WriteString(azure.FieldLabelStyle.Render("Priority:") + " " + azure.FieldValueStyle.Render(strconv.Itoa(item.Priority)) + "\n")
	s.WriteString(azure.FieldLabelStyle.Render("Created By:") + " " + azure.FieldValueStyle.Render(item.CreatedBy) + "\n")
	s.WriteString(azure.FieldLabelStyle.Render("Created Date:") + " " + azure.FieldValueStyle.Render(item.CreatedDate) + "\n")
	s.WriteString(azure.FieldLabelStyle.Render("Area Path:") + " " + azure.FieldValueStyle.Render(item.AreaPath) + "\n")
	s.WriteString(azure.FieldLabelStyle.Render("Iteration:") + " " + azure.FieldValueStyle.Render(item.Iteration) + "\n")

	// Tags
	if len(item.Tags) > 0 {
		var tags []string
		for _, tag := range item.Tags {
			tags = append(tags, azure.TagStyle.Render(tag))
		}
		s.WriteString(azure.FieldLabelStyle.Render("Tags:") + " " + strings.Join(tags, " ") + "\n")
	}

	s.WriteString("\n")

	// Description section - show as editable field
	descriptionLabel := azure.FieldLabelStyle.Render("Description:")
	if selectedFieldIndex == FieldDescription {
		// Highlight selected field
		selectedStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("226")).
			Background(lipgloss.Color("237"))
		descriptionLabel = selectedStyle.Render("▶ Description:") + " " +
			lipgloss.NewStyle().Foreground(lipgloss.Color("243")).Render("[press Enter to edit]")
	}

	s.WriteString(descriptionLabel + "\n")
	// Wrap description text to width
	wrappedDescription := wrapText(item.Description, width-2) // -2 for padding
	s.WriteString(azure.DescriptionStyle.Render(wrappedDescription))
	s.WriteString("\n\n")

	// Discussion section
	discussionHeader := azure.DiscussionHeaderStyle.Render(fmt.Sprintf("Discussion (%d)", len(item.Comments)))

	// Add new comment field - show inline when not selected
	if selectedFieldIndex == FieldAddComment {
		// Show on separate line when selected
		s.WriteString(discussionHeader + "\n\n")
		selectedStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("226")).
			Background(lipgloss.Color("237"))
		addCommentLabel := selectedStyle.Render("▶ Add new comment:") + " " +
			lipgloss.NewStyle().Foreground(lipgloss.Color("243")).Render("[press Enter to add]")
		s.WriteString(addCommentLabel + "\n\n")
	} else {
		// Show header on separate line
		s.WriteString(discussionHeader + "\n\n")
		addCommentLabel := azure.DiscussionHeaderStyle.Render("➤ Add new comment:")
		s.WriteString(addCommentLabel + "\n\n")
	}

	// Viewport for scrollable comments
	s.WriteString(viewportView)
	s.WriteString("\n\n")

	// Help text with new key bindings
	helpText := "j/k or ↑/↓: navigate fields • PgUp/PgDn or Ctrl+U/D: scroll comments • Enter: edit field"
	if isDirty {
		helpText += " • Ctrl+S: save • Ctrl+R: revert"
	}
	helpText += " • Esc: back • q: quit"
	s.WriteString(azure.HelpStyle.Render(helpText))

	return s.String()
}

// RenderEditView renders the edit mode view for modifying fields
func RenderEditView(item *azure.WorkItem, selectedFieldIndex int, width int, textareaView string) string {
	if item == nil {
		return "No item selected"
	}

	var s strings.Builder

	// Header with ID and Type
	s.WriteString(.HeaderStyle.Render(fmt.Sprintf("EDITING: %s #%d", item.Type, item.ID)))
	s.WriteString("\n\n")

	// Title - wrap to width
	wrappedTitle := wrapText(item.Title, width-2)
	s.WriteString(azure.TitleDetailStyle.Render(wrappedTitle))
	s.WriteString("\n\n")

	// Edit instruction
	editLabelStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("226"))

	// Show which field is being edited
	fieldName := "Field"
	switch selectedFieldIndex {
	case FieldDescription:
		fieldName = "Description"
	case FieldAddComment:
		fieldName = "New Comment"
	}

	s.WriteString(editLabelStyle.Render(fmt.Sprintf("Edit %s:", fieldName)))
	s.WriteString("\n")

	// Textarea for editing
	s.WriteString(textareaView)
	s.WriteString("\n\n")

	s.WriteString(azure.HelpStyle.Render("Ctrl+S: save to modified item • Esc: cancel"))

	return s.String()
}
