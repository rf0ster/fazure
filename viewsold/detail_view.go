// Package viewsold contains modules for rendering different parts of the application.
package viewsold

import (
	"fazure/types"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

// CreateCommentsViewport creates a viewport for scrollable comments
func CreateCommentsViewport(item *types.BacklogItem, width int) viewport.Model {
	vp := viewport.New(width, 20)
	vp.SetContent(RenderComments(item, width))
	return vp
}

// RenderComments renders the comments section as a string
func RenderComments(item *types.BacklogItem, width int) string {
	if item == nil || len(item.Comments) == 0 {
		return "No comments yet."
	}

	var content strings.Builder

	for _, comment := range item.Comments {
		// Wrap comment content to width
		wrappedContent := wrapText(comment.Content, width-4) // -4 for padding

		commentBox := fmt.Sprintf("%s  %s\n\n%s",
			types.CommentAuthorStyle.Render(comment.Author),
			types.CommentDateStyle.Render(comment.Date),
			wrappedContent,
		)
		content.WriteString(types.CommentStyle.Render(commentBox))
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
func RenderDetailView(item *types.BacklogItem, selectedFieldIndex int, isDirty bool, width int, viewportView string) string {
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
	s.WriteString(types.HeaderStyle.Render(header))
	s.WriteString("\n\n")

	// Title - wrap to width
	wrappedTitle := wrapText(item.Title, width-2)
	s.WriteString(types.TitleDetailStyle.Render(wrappedTitle))
	s.WriteString("\n\n")

	// Fields section
	s.WriteString(types.FieldLabelStyle.Render("State:") + " " + types.FieldValueStyle.Render(item.State) + "\n")
	s.WriteString(types.FieldLabelStyle.Render("Assigned To:") + " " + types.FieldValueStyle.Render(item.AssignedTo) + "\n")
	s.WriteString(types.FieldLabelStyle.Render("Priority:") + " " + types.FieldValueStyle.Render(strconv.Itoa(item.Priority)) + "\n")
	s.WriteString(types.FieldLabelStyle.Render("Created By:") + " " + types.FieldValueStyle.Render(item.CreatedBy) + "\n")
	s.WriteString(types.FieldLabelStyle.Render("Created Date:") + " " + types.FieldValueStyle.Render(item.CreatedDate) + "\n")
	s.WriteString(types.FieldLabelStyle.Render("Area Path:") + " " + types.FieldValueStyle.Render(item.AreaPath) + "\n")
	s.WriteString(types.FieldLabelStyle.Render("Iteration:") + " " + types.FieldValueStyle.Render(item.Iteration) + "\n")

	// Tags
	if len(item.Tags) > 0 {
		var tags []string
		for _, tag := range item.Tags {
			tags = append(tags, types.TagStyle.Render(tag))
		}
		s.WriteString(types.FieldLabelStyle.Render("Tags:") + " " + strings.Join(tags, " ") + "\n")
	}

	s.WriteString("\n")

	// Description section - show as editable field
	descriptionLabel := types.FieldLabelStyle.Render("Description:")
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
	s.WriteString(types.DescriptionStyle.Render(wrappedDescription))
	s.WriteString("\n\n")

	// Discussion section
	discussionHeader := types.DiscussionHeaderStyle.Render(fmt.Sprintf("Discussion (%d)", len(item.Comments)))

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
		addCommentLabel := types.DiscussionHeaderStyle.Render("➤ Add new comment:")
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
	s.WriteString(types.HelpStyle.Render(helpText))

	return s.String()
}

// RenderEditView renders the edit mode view for modifying fields
func RenderEditView(item *types.BacklogItem, selectedFieldIndex int, width int, textareaView string) string {
	if item == nil {
		return "No item selected"
	}

	var s strings.Builder

	// Header with ID and Type
	s.WriteString(types.HeaderStyle.Render(fmt.Sprintf("EDITING: %s #%d", item.Type, item.ID)))
	s.WriteString("\n\n")

	// Title - wrap to width
	wrappedTitle := wrapText(item.Title, width-2)
	s.WriteString(types.TitleDetailStyle.Render(wrappedTitle))
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

	s.WriteString(types.HelpStyle.Render("Ctrl+S: save to modified item • Esc: cancel"))

	return s.String()
}
