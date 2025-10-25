// Package views contains modules for rendering different parts of the application.
package views

import (
	"fazure/types"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

// CreateCommentsViewport creates a viewport for scrollable comments
func CreateCommentsViewport(item *types.BacklogItem) viewport.Model {
	vp := viewport.New(100, 20)
	vp.SetContent(RenderComments(item))
	return vp
}

// RenderComments renders the comments section as a string
func RenderComments(item *types.BacklogItem) string {
	if item == nil || len(item.Comments) == 0 {
		return "No comments yet."
	}

	var content strings.Builder

	for _, comment := range item.Comments {
		commentBox := fmt.Sprintf("%s  %s\n\n%s",
			types.CommentAuthorStyle.Render(comment.Author),
			types.CommentDateStyle.Render(comment.Date),
			comment.Content,
		)
		content.WriteString(types.CommentStyle.Render(commentBox))
		content.WriteString("\n")
	}

	return content.String()
}

// RenderDetailView renders the full detail view for a work item
func RenderDetailView(item *types.BacklogItem, viewportView string) string {
	if item == nil {
		return "No item selected"
	}

	var s strings.Builder

	// Header with ID and Type
	s.WriteString(types.HeaderStyle.Render(fmt.Sprintf("%s #%d", item.Type, item.ID)))
	s.WriteString("\n\n")

	// Title
	s.WriteString(types.TitleDetailStyle.Render(item.Title))
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

	// Description section
	s.WriteString(types.FieldLabelStyle.Render("Description:") + "\n")
	s.WriteString(types.DescriptionStyle.Render(item.Description))
	s.WriteString("\n\n")

	// Discussion section
	s.WriteString(types.DiscussionHeaderStyle.Render(fmt.Sprintf("Discussion (%d)", len(item.Comments))))
	s.WriteString("\n")

	// Viewport for scrollable comments
	s.WriteString(viewportView)
	s.WriteString("\n\n")

	s.WriteString(types.HelpStyle.Render("Use arrow keys to scroll comments • 'e' to edit description • 'esc' to go back • 'q' to quit"))

	return s.String()
}

// RenderEditView renders the edit mode view for modifying the description
func RenderEditView(item *types.BacklogItem, textareaView string) string {
	if item == nil {
		return "No item selected"
	}

	var s strings.Builder

	// Header with ID and Type
	s.WriteString(types.HeaderStyle.Render(fmt.Sprintf("EDITING: %s #%d", item.Type, item.ID)))
	s.WriteString("\n\n")

	// Title
	s.WriteString(types.TitleDetailStyle.Render(item.Title))
	s.WriteString("\n\n")

	// Edit instruction
	editLabelStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("226"))

	s.WriteString(editLabelStyle.Render("Edit Description:"))
	s.WriteString("\n")

	// Textarea for editing
	s.WriteString(textareaView)
	s.WriteString("\n\n")

	s.WriteString(types.HelpStyle.Render("'ctrl+s' to save • 'esc' to cancel"))

	return s.String()
}
