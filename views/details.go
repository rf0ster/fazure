package views

import (
	"fazure/azure"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type DetailsView struct {
	item *azure.WorkItem
}

func (v *DetailsView) Init(m Model) tea.Cmd {
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

	width := getContentWidth(m.terminalWidth)

	// Displat work item title
	wrappedTitle := wrapText(v.item.Title, width)
	s.WriteString(GetWorkItemTypeStyle(item.Type).Render(wrappedTitle))
	s.WriteString("\n\n")

	// Display work item fields
	s.WriteString(FieldLabelStyle.Render("State:") + " " + FieldValueStyle.Render(item.State) + "\n")
	s.WriteString(FieldLabelStyle.Render("Assigned To:") + " " + FieldValueStyle.Render(item.AssignedTo) + "\n")
	s.WriteString(FieldLabelStyle.Render("Priority:") + " " + FieldValueStyle.Render(strconv.Itoa(item.Priority)) + "\n")
	s.WriteString(FieldLabelStyle.Render("Created By:") + " " + FieldValueStyle.Render(item.CreatedBy) + "\n")
	s.WriteString(FieldLabelStyle.Render("Created Date:") + " " + FieldValueStyle.Render(item.CreatedDate) + "\n")
	s.WriteString(FieldLabelStyle.Render("Area Path:") + " " + FieldValueStyle.Render(item.AreaPath) + "\n")
	s.WriteString(FieldLabelStyle.Render("Iteration:") + " " + FieldValueStyle.Render(item.Iteration) + "\n")

	// Display tags
	if len(item.Tags) > 0 {
		var tags []string
		for _, tag := range item.Tags {
			tags = append(tags, TagStyle.Render(tag))
		}
		s.WriteString(FieldLabelStyle.Render("Tags:") + " " + strings.Join(tags, " ") + "\n")
	}
	s.WriteString("\n")

	// Display description
	s.WriteString(FieldLabelStyle.Render("Description:") + "\n")
	wrappedDescription := wrapText(item.Description, width-2) // -2 for padding
	s.WriteString(DescriptionStyle.Render(wrappedDescription))
	s.WriteString("\n\n")


	// Display comments section
	discussionHeader := DiscussionHeaderStyle.Render(fmt.Sprintf("Discussion (%d)", len(item.Comments)))
	s.WriteString(discussionHeader + "\n")
	viewportView := CreateCommentsViewport(item, width).View()
	s.WriteString(viewportView)
	s.WriteString("\n\n")

	return s.String()
}

func (v *DetailsView) Update(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.view = &BacklogView{}
			return m, m.view.Init(m)
		}
	}

	return m, nil
}

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
			CommentAuthorStyle.Render(comment.Author),
			CommentDateStyle.Render(comment.Date),
			wrappedContent,
		)
		content.WriteString(CommentStyle.Render(commentBox))
		content.WriteString("\n")
	}

	return content.String()
}

