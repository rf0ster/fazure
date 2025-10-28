package azure

import (
	"fmt"
	"os"
	"testing"
)

// TestQueryActiveWorkItems queries active work items for a given user
func TestQueryActiveWorkItems(t *testing.T) {
	// Get configuration from environment variables
	org := os.Getenv("AZURE_ORG")
	project := os.Getenv("AZURE_PROJECT")
	pat := os.Getenv("AZURE_PAT")
	user := os.Getenv("AZURE_USER")

	if org == "" || project == "" || pat == "" {
		t.Skip("Skipping test: AZURE_ORG, AZURE_PROJECT, or AZURE_PAT environment variables not set")
	}

	if user == "" {
		t.Skip("Skipping test: AZURE_USER environment variable not set")
	}

	// Create client
	client := NewClient(org, project, pat)

	// Query active work items for the user
	params := QueryParams{
		AssignedTo: user,
		State:      "Active",
	}

	fmt.Printf("\n=== Querying Active Work Items ===\n")
	fmt.Printf("User: %s\n", user)
	fmt.Printf("State: Active\n\n")

	workItems, err := client.QueryWorkItems(params)
	if err != nil {
		t.Fatalf("QueryWorkItems failed: %v", err)
	}

	fmt.Printf("Found %d active work items:\n\n", len(workItems))

	for _, wi := range workItems {
		printWorkItem(wi)
	}
}

// printWorkItem prints a work item in a readable format
func printWorkItem(wi WorkItem) {
	fmt.Printf("┌─ Work Item #%d ─────────────────────────────────────\n", wi.ID)
	fmt.Printf("│ Type:        %s\n", wi.Type)
	fmt.Printf("│ Title:       %s\n", wi.Title)
	fmt.Printf("│ Assigned To: %s\n", wi.AssignedTo)
	fmt.Printf("│ State:       %s\n", wi.State)
	fmt.Printf("│ Priority:    %d\n", wi.Priority)
	fmt.Printf("│ Area Path:   %s\n", wi.AreaPath)
	fmt.Printf("│ Iteration:   %s\n", wi.Iteration)
	fmt.Printf("│ Created By:  %s\n", wi.CreatedBy)
	fmt.Printf("│ Created:     %s\n", wi.CreatedDate)
	if len(wi.Tags) > 0 {
		fmt.Printf("│ Tags:        %v\n", wi.Tags)
	}
	if wi.Description != "" {
		desc := wi.Description
		if len(desc) > 100 {
			desc = desc[:97] + "..."
		}
		fmt.Printf("│ Description: %s\n", desc)
	}
	fmt.Printf("└────────────────────────────────────────────────────────\n\n")
}

