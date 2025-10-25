package types

// WorkItemType represents the type of Azure DevOps work item
type WorkItemType string

const (
	Initiative  WorkItemType = "Initiative"
	Requirement WorkItemType = "Requirement"
	UserStory   WorkItemType = "User Story"
	Task        WorkItemType = "Task"
	Bug         WorkItemType = "Bug"
)

// Comment represents a discussion comment on a work item
type Comment struct {
	Author  string
	Date    string
	Content string
}

// BacklogItem represents a work item from Azure DevOps
type BacklogItem struct {
	ID          int
	Type        WorkItemType
	Title       string
	AssignedTo  string
	State       string
	Priority    int
	Description string
	CreatedBy   string
	CreatedDate string
	Tags        []string
	AreaPath    string
	Iteration   string
	Comments    []Comment
}
