# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Fazure is a terminal-based Azure DevOps work item manager built with Go and Bubble Tea. It provides a TUI (Terminal User Interface) for querying and viewing Azure DevOps backlog items including initiatives, requirements, user stories, tasks, and bugs. Currently uses mock data for testing, with plans to integrate with actual Azure DevOps APIs.

## Build and Run Commands

```bash
# Build the application
go build -o fazure

# Run the application
./fazure

# Or build and run in one step
go run .

# Install dependencies (if needed)
go mod tidy
```

## Architecture

### The Elm Architecture (TEA) Pattern

This application follows the Elm Architecture via Bubble Tea:
- **Model**: Single source of truth (main.go `model` struct)
- **Update**: State transitions based on messages (keyboard input)
- **View**: Pure rendering functions that take state and return strings

### State Machine Flow

The application has three distinct view states managed via boolean flags:

1. **Search View** (`showResults=false, showDetail=false`): Text input for searching users
2. **Table View** (`showResults=true, showDetail=false`): Display search results in a table
3. **Detail View** (`showResults=true, showDetail=true`): Full work item details with scrollable comments

Navigation:
- `Enter` advances to the next view (search → table → detail)
- `Esc` navigates backward through views
- `Q` quits from any view

### Package Structure

**`types/`** - Core data structures and styling
- `work_items.go`: Azure DevOps domain models (BacklogItem, Comment, WorkItemType)
- `styles.go`: All lipgloss styles and Azure DevOps color constants (8-bit ANSI codes)

**`views/`** - View rendering logic (stateless)
- `search_view.go`: Renders the search input screen
- `table_view.go`: Creates and renders the work items table
- `detail_view.go`: Renders full work item details with scrollable comments viewport

**`main.go`** - Application orchestration
- Defines the Bubble Tea model and state
- Handles all user input and state transitions in `Update()`
- Delegates rendering to appropriate view functions in `View()`

**`mock_azure_client.go`** - Mock data provider
- Returns hard-coded work items for test users: john, sarah, mike, emma
- Will be replaced with real Azure DevOps API client

### Key Design Decisions

**View Components are Stateless**: All view functions in `views/` are pure functions that take data and return strings. They don't hold state. The Bubble Tea components (textinput, table, viewport) are stored in the main model.

**Styles Centralized in types/**: All lipgloss styles live in `types/styles.go` and are imported as `types.TitleStyle`, `types.HelpStyle`, etc. This allows views to reference styles without circular imports.

**Azure DevOps Color Matching**: Work item type colors use 8-bit ANSI codes (not hex) to ensure proper terminal rendering. These match Azure DevOps official colors as closely as possible.

**Viewport for Scrolling**: The detail view uses Bubble Tea's `viewport` component for scrollable comments. The viewport is created fresh when entering detail view and its rendered output is passed to `RenderDetailView()`.

## Adding New Features

**Adding a new work item field:**
1. Add field to `types.BacklogItem` struct
2. Update mock data in `mock_azure_client.go`
3. Update detail view rendering in `views/detail_view.go`

**Adding a new view:**
1. Create new file in `views/` package with render function(s)
2. Add state flags to main `model` struct if needed
3. Add navigation logic in `Update()` function
4. Call new view's render function from main `View()` function

**Integrating real Azure DevOps API:**
1. Replace `MockAzureClient` with real client implementation
2. Keep the same `SearchWorkItems(assignedTo string) []types.BacklogItem` interface
3. Update `NewMockAzureClient()` call in `initialModel()`
