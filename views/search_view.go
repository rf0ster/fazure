package views

import (
	"fazure/types"
)

// RenderSearchView renders the search input screen
func RenderSearchView(textInputView string) string {
	var s string
	s += types.TitleStyle.Render("Azure DevOps Work Item Search")
	s += "\n\n"
	s += textInputView
	s += "\n\n"
	s += types.HelpStyle.Render("Press 'enter' to search • 'q' to quit\n")
	s += types.HelpStyle.Render("Test users: john, sarah, mike, emma")
	return s
}
