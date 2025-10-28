package azure

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AzureClient represents an Azure DevOps API client
type AzureClient struct {
	Organization string
	Project      string
	PAT          string
	HTTPClient   *http.Client
}

// NewClient creates a new Azure DevOps client
func NewClient(organization, project, pat string) *AzureClient {
	return &AzureClient{
		Organization: organization,
		Project:      project,
		PAT:          pat,
		HTTPClient:   &http.Client{},
	}
}

// QueryParams represents parameters for querying work items
type QueryParams struct {
	AssignedTo    string
	State         string
	IterationPath string
	AreaPath      string
}

// wiqlQuery represents the request body for WIQL queries
type wiqlQuery struct {
	Query string `json:"query"`
}

// wiqlResponse represents the response from a WIQL query
type wiqlResponse struct {
	WorkItems []struct {
		ID  int    `json:"id"`
		URL string `json:"url"`
	} `json:"workItems"`
}

// workItemsResponse represents the response when fetching work item details
type workItemsResponse struct {
	Count int `json:"count"`
	Value []struct {
		ID     int               `json:"id"`
		Fields map[string]any `json:"fields"`
	} `json:"value"`
}

// QueryWorkItems queries work items from Azure DevOps based on the provided parameters
func (c *AzureClient) QueryWorkItems(params QueryParams) ([]WorkItem, error) {
	// Build WIQL query
	wiql := c.buildWIQL(params)

	// Execute WIQL query to get work item IDs
	ids, err := c.executeWIQL(wiql)
	if err != nil {
		return nil, fmt.Errorf("failed to execute WIQL query: %w", err)
	}

	if len(ids) == 0 {
		return []WorkItem{}, nil
	}

	// Fetch full work item details
	workItems, err := c.getWorkItemDetails(ids)
	if err != nil {
		return nil, fmt.Errorf("failed to get work item details: %w", err)
	}

	return workItems, nil
}

// buildWIQL constructs a WIQL query based on the provided parameters
func (c *AzureClient) buildWIQL(params QueryParams) string {
	conditions := []string{
		fmt.Sprintf("[System.TeamProject] = '%s'", c.Project),
	}

	if params.AssignedTo != "" {
		conditions = append(conditions, fmt.Sprintf("[System.AssignedTo] = '%s'", params.AssignedTo))
	}

	if params.State != "" {
		conditions = append(conditions, fmt.Sprintf("[System.State] = '%s'", params.State))
	}

	if params.IterationPath != "" {
		conditions = append(conditions, fmt.Sprintf("[System.IterationPath] = '%s'", params.IterationPath))
	}

	if params.AreaPath != "" {
		conditions = append(conditions, fmt.Sprintf("[System.AreaPath] = '%s'", params.AreaPath))
	}

	whereClause := strings.Join(conditions, " AND ")
	return fmt.Sprintf("SELECT [System.Id] FROM WorkItems WHERE %s", whereClause)
}

// executeWIQL executes a WIQL query and returns work item IDs
func (c *AzureClient) executeWIQL(wiql string) ([]int, error) {
	apiURL := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/wit/wiql?api-version=7.0",
		url.PathEscape(c.Organization),
		url.PathEscape(c.Project))

	query := wiqlQuery{Query: wiql}
	body, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var wiqlResp wiqlResponse
	if err := json.NewDecoder(resp.Body).Decode(&wiqlResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	ids := make([]int, len(wiqlResp.WorkItems))
	for i, item := range wiqlResp.WorkItems {
		ids[i] = item.ID
	}

	return ids, nil
}

// getWorkItemDetails fetches full details for the given work item IDs
func (c *AzureClient) getWorkItemDetails(ids []int) ([]WorkItem, error) {
	if len(ids) == 0 {
		return []WorkItem{}, nil
	}

	// Convert IDs to comma-separated string
	idStrs := make([]string, len(ids))
	for i, id := range ids {
		idStrs[i] = fmt.Sprintf("%d", id)
	}
	idsParam := strings.Join(idStrs, ",")

	// Build URL with all required fields
	fields := []string{
		"System.Id",
		"System.WorkItemType",
		"System.Title",
		"System.AssignedTo",
		"System.State",
		"Microsoft.VSTS.Common.Priority",
		"System.Description",
		"Microsoft.VSTS.Common.AcceptanceCriteria",
		"System.CreatedBy",
		"System.CreatedDate",
		"System.Tags",
		"System.AreaPath",
		"System.IterationPath",
	}

	apiURL := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/wit/workitems?ids=%s&fields=%s&api-version=7.0",
		url.PathEscape(c.Organization),
		url.PathEscape(c.Project),
		idsParam,
		strings.Join(fields, ","))

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var wiResp workItemsResponse
	if err := json.NewDecoder(resp.Body).Decode(&wiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert to WorkItem structs
	workItems := make([]WorkItem, 0, len(wiResp.Value))
	for _, item := range wiResp.Value {
		wi := c.convertToWorkItem(item.ID, item.Fields)
		workItems = append(workItems, wi)
	}

	return workItems, nil
}

// convertToWorkItem converts Azure DevOps API response fields to a WorkItem struct
func (c *AzureClient) convertToWorkItem(id int, fields map[string]any) WorkItem {
	wi := WorkItem{
		ID: id,
	}

	if v, ok := fields["System.WorkItemType"].(string); ok {
		wi.Type = WorkItemType(v)
	}

	if v, ok := fields["System.Title"].(string); ok {
		wi.Title = v
	}

	if v, ok := fields["System.AssignedTo"].(map[string]any); ok {
		if name, ok := v["displayName"].(string); ok {
			wi.AssignedTo = name
		}
	}

	if v, ok := fields["System.State"].(string); ok {
		wi.State = v
	}

	if v, ok := fields["Microsoft.VSTS.Common.Priority"].(float64); ok {
		wi.Priority = int(v)
	}

	if v, ok := fields["System.Description"].(string); ok {
		wi.Description = v
	}

	if v, ok := fields["Microsoft.VSTS.Common.AcceptanceCriteria"].(string); ok {
		wi.AcceptanceCriteria = v
	}

	if v, ok := fields["System.CreatedBy"].(map[string]any); ok {
		if name, ok := v["displayName"].(string); ok {
			wi.CreatedBy = name
		}
	}

	if v, ok := fields["System.CreatedDate"].(string); ok {
		wi.CreatedDate = v
	}

	if v, ok := fields["System.Tags"].(string); ok {
		if v != "" {
			tags := strings.Split(v, "; ")
			wi.Tags = tags
		}
	}

	if v, ok := fields["System.AreaPath"].(string); ok {
		wi.AreaPath = v
	}

	if v, ok := fields["System.IterationPath"].(string); ok {
		wi.Iteration = v
	}

	return wi
}

// setHeaders sets common headers for Azure DevOps API requests
func (c *AzureClient) setHeaders(req *http.Request) {
	auth := base64.StdEncoding.EncodeToString([]byte(":" + c.PAT))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
}
