package views

import (
	"fazure/azure"

	"github.com/charmbracelet/lipgloss"
)

// Azure DevOps work item type colors (8-bit ANSI color codes)
const (
	InitiativeColor  = "208" // Orange
	RequirementColor = "97"  // Purple
	UserStoryColor   = "38"  // Blue/Cyan
	TaskColor        = "220" // Yellow
	BugColor         = "161" // Red
)

// UI color constants
const (
	ColorPurpleViolet = "#7D56F4" // Primary purple/violet
	ColorGray         = "#626262" // Gray for help text
	ColorWhite        = "15"      // White
	ColorPurpleBg     = "57"      // Purple background
	ColorCyanGreen    = "86"      // Cyan/green for labels
	ColorLightYellow  = "229"     // Light yellow
	ColorDarkGray     = "237"     // Dark gray background
	ColorBorderGray   = "240"     // Border gray
	ColorDateGray     = "241"     // Date text gray
)

// Common styles used throughout the application
var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(ColorPurpleViolet)).
			MarginTop(1).
			MarginBottom(1)

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorGray)).
			MarginTop(1)

	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(ColorWhite)).
			Background(lipgloss.Color(ColorPurpleBg)).
			Padding(0, 2)

	TitleDetailStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(ColorWhite))

	FieldLabelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(ColorCyanGreen)).
			Width(15)

	FieldValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorWhite))

	TagStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorLightYellow)).
			Background(lipgloss.Color(ColorDarkGray)).
			Padding(0, 1)

	TextAreaHeader = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(ColorCyanGreen)).
			MarginBottom(1)

	DescriptionStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(ColorBorderGray)).
				Padding(1).
				MarginBottom(1)

	DiscussionHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(ColorCyanGreen)).
				MarginBottom(1)

	CommentStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(ColorBorderGray)).
			Padding(1).
			MarginBottom(1)

	CommentAuthorStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(ColorCyanGreen))

	CommentDateStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ColorDateGray))

	ActiveOptionStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(ColorCyanGreen))

	InactiveOptionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ColorGray))

	ActiveTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(ColorWhite)).
			Background(lipgloss.Color(ColorPurpleBg)).
			Padding(0, 1)

	InactiveTabStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ColorGray)).
				Padding(0, 1)
)

// GetWorkItemTypeColor returns the ANSI color for a work item type
func GetWorkItemTypeColor(itemType azure.WorkItemType) string {
	switch itemType {
	case azure.Initiative:
		return InitiativeColor
	case azure.Requirement:
		return RequirementColor
	case azure.UserStory:
		return UserStoryColor
	case azure.Task:
		return TaskColor
	case azure.Bug:
		return BugColor
	default:
		return ""
	}
}

// GetWorkItemTypeStyle returns the lipgloss style for a work item type
func GetWorkItemTypeStyle(itemType azure.WorkItemType) lipgloss.Style {
	color := GetWorkItemTypeColor(itemType)
	if color == "" {
		return lipgloss.NewStyle()
	}

	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(color))
}
