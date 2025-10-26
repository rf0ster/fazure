// Package azure provides Azure DevOps client implementations.
package azure

// MockAzureClient provides mock data for testing
type MockAzureClient struct{}

// NewMockAzureClient creates a new mock Azure DevOps client
func NewMockAzureClient() *MockAzureClient {
	return &MockAzureClient{}
}

// SearchWorkItems returns mock backlog items for a given user
func (c *MockAzureClient) SearchWorkItems(assignedTo string) []WorkItem {
	mockData := map[string][]WorkItem{
		"john": {
			{
				ID: 1001, Type: Initiative, Title: "Implement new authentication system",
				AssignedTo: "john", State: "In Progress", Priority: 1,
				Description: "Modernize our authentication system to support OAuth2, SAML, and social logins. This initiative will improve security and user experience.",
				CreatedBy:   "alice", CreatedDate: "2024-01-15",
				Tags:        []string{"security", "authentication", "phase-1"},
				AreaPath:    "FazureApp\\Backend\\Security", Iteration: "Sprint 23",
				Comments: []Comment{
					{Author: "alice", Date: "2024-01-15 09:30", Content: "Created this initiative to track our auth modernization efforts. Let's aim to complete this by end of Q1."},
					{Author: "john", Date: "2024-01-16 14:20", Content: "Started working on the OAuth2 implementation. Will need help with the UI components."},
					{Author: "mike", Date: "2024-01-17 10:15", Content: "I can help with the frontend once you have the API ready!"},
				},
			},
			{
				ID: 1002, Type: UserStory, Title: "Add OAuth2 login support",
				AssignedTo: "john", State: "Active", Priority: 1,
				Description: "As a user, I want to log in using my Google or GitHub account so that I don't have to remember another password.",
				CreatedBy:   "alice", CreatedDate: "2024-01-16",
				Tags:        []string{"oauth2", "user-story"},
				AreaPath:    "FazureApp\\Backend\\Auth", Iteration: "Sprint 23",
				Comments: []Comment{
					{Author: "alice", Date: "2024-01-16 11:00", Content: "This should integrate with Google and GitHub OAuth providers initially."},
					{Author: "john", Date: "2024-01-18 15:45", Content: "Google OAuth is working! GitHub integration next."},
				},
			},
			{
				ID: 1003, Type: Task, Title: "Create login page UI",
				AssignedTo: "john", State: "Done", Priority: 2,
				Description: "Design and implement the new login page with OAuth buttons and traditional login form.",
				CreatedBy:   "john", CreatedDate: "2024-01-17",
				Tags:        []string{"ui", "frontend"},
				AreaPath:    "FazureApp\\Frontend", Iteration: "Sprint 23",
				Comments: []Comment{
					{Author: "john", Date: "2024-01-17 09:00", Content: "Starting the UI work today. Using our design system components."},
					{Author: "sarah", Date: "2024-01-17 16:30", Content: "Looks great! Make sure it's mobile responsive."},
					{Author: "john", Date: "2024-01-18 17:00", Content: "Done! Tested on mobile and desktop."},
				},
			},
			{
				ID: 1004, Type: Bug, Title: "Fix login redirect issue",
				AssignedTo: "john", State: "Active", Priority: 1,
				Description: "After successful OAuth login, users are redirected to /home instead of their originally requested page. Need to preserve the redirect URL through the OAuth flow.",
				CreatedBy:   "emma", CreatedDate: "2024-01-20",
				Tags:        []string{"bug", "oauth", "critical"},
				AreaPath:    "FazureApp\\Backend\\Auth", Iteration: "Sprint 23",
				Comments: []Comment{
					{Author: "emma", Date: "2024-01-20 10:15", Content: "Found this while testing. Steps to reproduce: 1) Navigate to /dashboard, 2) Click login, 3) Complete OAuth, 4) You end up at /home instead of /dashboard"},
					{Author: "john", Date: "2024-01-20 11:30", Content: "Good catch! I'll fix this by storing the original URL in the session state."},
				},
			},
		},
		"sarah": {
			{
				ID: 2001, Type: Requirement, Title: "API rate limiting requirements",
				AssignedTo: "sarah", State: "New", Priority: 1,
				Description: "Define requirements for API rate limiting to protect our services from abuse and ensure fair usage.",
				CreatedBy:   "alice", CreatedDate: "2024-01-10",
				Tags:        []string{"api", "requirements"},
				AreaPath:    "FazureApp\\Backend\\API", Iteration: "Sprint 24",
				Comments: []Comment{
					{Author: "alice", Date: "2024-01-10 14:00", Content: "We need to define rate limits per user tier and implement proper throttling."},
				},
			},
			{
				ID: 2002, Type: UserStory, Title: "Implement rate limiting middleware",
				AssignedTo: "sarah", State: "In Progress", Priority: 1,
				Description: "As an API owner, I want to rate limit requests per user to prevent abuse and ensure system stability.",
				CreatedBy:   "sarah", CreatedDate: "2024-01-12",
				Tags:        []string{"api", "middleware"},
				AreaPath:    "FazureApp\\Backend\\API", Iteration: "Sprint 24",
				Comments: []Comment{
					{Author: "sarah", Date: "2024-01-12 09:00", Content: "Starting implementation using Redis for distributed rate limiting."},
					{Author: "john", Date: "2024-01-13 10:30", Content: "Make sure to add proper headers for rate limit status!"},
				},
			},
		},
		"mike": {
			{
				ID: 3001, Type: Initiative, Title: "Mobile app redesign",
				AssignedTo: "mike", State: "Planning", Priority: 1,
				Description: "Complete redesign of our mobile application with modern UI/UX patterns and improved performance.",
				CreatedBy:   "alice", CreatedDate: "2024-01-05",
				Tags:        []string{"mobile", "redesign", "ux"},
				AreaPath:    "FazureApp\\Mobile", Iteration: "Sprint 25",
				Comments: []Comment{
					{Author: "alice", Date: "2024-01-05 08:00", Content: "Let's make our mobile app shine! Focus on user experience and performance."},
					{Author: "mike", Date: "2024-01-06 11:00", Content: "Working on the design mockups. Will share soon!"},
				},
			},
		},
		"emma": {
			{
				ID: 4001, Type: UserStory, Title: "Add data export functionality",
				AssignedTo: "emma", State: "New", Priority: 2,
				Description: "As a user, I want to export my data in multiple formats (CSV, JSON, Excel) so I can analyze it in external tools.",
				CreatedBy:   "alice", CreatedDate: "2024-01-08",
				Tags:        []string{"export", "data"},
				AreaPath:    "FazureApp\\Features", Iteration: "Sprint 24",
				Comments: []Comment{
					{Author: "alice", Date: "2024-01-08 13:00", Content: "Users have been requesting this feature. Let's start with CSV."},
				},
			},
		},
	}

	// Return items for the specified user, or empty slice if not found
	if items, exists := mockData[assignedTo]; exists {
		return items
	}
	return []WorkItem{}
}
