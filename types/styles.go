package types

import "github.com/charmbracelet/lipgloss"

// Azure DevOps work item type colors (8-bit ANSI color codes)
const (
	InitiativeColor  = "208" // Orange
	RequirementColor = "97"  // Purple
	UserStoryColor   = "38"  // Blue/Cyan
	TaskColor        = "220" // Yellow
	BugColor         = "161" // Red
)

// Common styles used throughout the application
var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			MarginTop(1).
			MarginBottom(1)

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(1)

	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("15")).
			Background(lipgloss.Color("57")).
			Padding(0, 2)

	TitleDetailStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("15"))

	FieldLabelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("86")).
			Width(15)

	FieldValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15"))

	TagStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("237")).
			Padding(0, 1)

	DescriptionStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("240")).
				Padding(1).
				MarginBottom(1)

	DiscussionHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("86")).
				MarginBottom(1)

	CommentStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(1).
			MarginBottom(1)

	CommentAuthorStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("86"))

	CommentDateStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("241"))
)

// GetWorkItemTypeColor returns the ANSI color for a work item type
func GetWorkItemTypeColor(itemType WorkItemType) string {
	switch itemType {
	case Initiative:
		return InitiativeColor
	case Requirement:
		return RequirementColor
	case UserStory:
		return UserStoryColor
	case Task:
		return TaskColor
	case Bug:
		return BugColor
	default:
		return ""
	}
}

// GetWorkItemTypeStyle returns the lipgloss style for a work item type
func GetWorkItemTypeStyle(itemType WorkItemType) lipgloss.Style {
	color := GetWorkItemTypeColor(itemType)
	if color == "" {
		return lipgloss.NewStyle()
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(color))
}
