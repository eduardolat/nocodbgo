// Package nocodbgo provides a client for the NocoDB v2 API
//
// More info and docs:
//
//   - Overview: https://docs.nocodb.com/developer-resources/rest-APIs/overview
//   - API Explorer: https://data-apis-v2.nocodb.com
//   - OpenAPI spec: https://data-apis-v2.nocodb.com/swagger-v2.json
package nocodbgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// defaultTimeout is the default timeout for HTTP requests
	defaultTimeout = 30 * time.Second
)

// Client is the NocoDB API client
type Client struct {
	// baseURL is the base URL for the NocoDB API
	baseURL string

	// apiToken is the API token for authentication
	apiToken string

	// httpClient is the HTTP client used to make requests
	httpClient *http.Client
}

// NewClient creates a new client builder
func NewClient() *clientBuilder {
	return &clientBuilder{
		httpClient: &http.Client{Timeout: defaultTimeout},
	}
}

// apiError represents an error returned by the NocoDB API
type apiError struct {
	Message string `json:"msg"`
	Code    string `json:"code"`
}

// Error implements the error interface
func (e apiError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("API error: %s (code: %s)", e.Message, e.Code)
	}
	return fmt.Sprintf("API error: %s", e.Message)
}

// request makes an HTTP request to the NocoDB API, it includes the api token in the request header
func (c *Client) request(ctx context.Context, method string, path string, body any, query url.Values) ([]byte, error) {
	parsedUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.baseURL, strings.TrimPrefix(path, "/")))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	if query != nil {
		parsedUrl.RawQuery = query.Encode()
	}

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, parsedUrl.String(), reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("xc-token", c.apiToken)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiErr apiError
		if err := json.Unmarshal(respBody, &apiErr); err != nil {
			return nil, fmt.Errorf("status code %d: failed to unmarshal API error: %w", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("status code %d: API error: %s", resp.StatusCode, apiErr.Message)
	}

	return respBody, nil
}

// Table returns a new table with the given ID
func (c *Client) Table(tableID string) *Table {
	return &Table{
		client:  c,
		tableID: tableID,
	}
}

// Table represents a table in NocoDB
type Table struct {
	client  *Client
	tableID string
}

// Create creates a new record in a table and returns the ID
func (t *Table) Create(ctx context.Context, data map[string]any) (int, error) {
	records, err := t.BulkCreate(ctx, []map[string]any{data})
	if err != nil {
		return 0, fmt.Errorf("failed to create record: %w", err)
	}

	if len(records) == 0 {
		return 0, fmt.Errorf("no record created")
	}

	return records[0], nil
}

// Update updates a record in a table
func (t *Table) Update(ctx context.Context, recordID int, data map[string]any) error {
	if recordID == 0 {
		return ErrRowIDRequired
	}

	// Add ID to the data
	updateData := make(map[string]any)
	for k, v := range data {
		updateData[k] = v
	}
	updateData["Id"] = recordID

	err := t.BulkUpdate(ctx, []map[string]any{updateData})
	if err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil
}

// Delete deletes a record from a table
func (t *Table) Delete(ctx context.Context, recordID int) error {
	if recordID == 0 {
		return ErrRowIDRequired
	}

	err := t.BulkDelete(ctx, []int{recordID})
	if err != nil {
		return fmt.Errorf("failed to delete record: %w", err)
	}

	return nil
}

// BulkCreate creates multiple records in a table and returns the IDs
func (t *Table) BulkCreate(ctx context.Context, data []map[string]any) ([]int, error) {
	path := fmt.Sprintf("/api/v2/tables/%s/records", t.tableID)
	respBody, err := t.client.request(ctx, http.MethodPost, path, data, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create records: %w", err)
	}

	var response []map[string]any
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal create response: %w", err)
	}

	var ids []int
	for _, record := range response {
		if id, ok := record["Id"].(float64); ok {
			ids = append(ids, int(id))
		}
	}

	return ids, nil
}

// BulkUpdate updates multiple records in a table
func (t *Table) BulkUpdate(ctx context.Context, data []map[string]any) error {
	path := fmt.Sprintf("/api/v2/tables/%s/records", t.tableID)
	_, err := t.client.request(ctx, http.MethodPatch, path, data, nil)
	if err != nil {
		return fmt.Errorf("failed to update records: %w", err)
	}

	return nil
}

// BulkDelete deletes multiple records from a table
func (t *Table) BulkDelete(ctx context.Context, recordIDs []int) error {
	if len(recordIDs) == 0 {
		return nil
	}

	// Convert IDs to the format expected by the API
	ids := make([]map[string]any, len(recordIDs))
	for i, id := range recordIDs {
		ids[i] = map[string]any{"Id": id}
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records", t.tableID)
	_, err := t.client.request(ctx, http.MethodDelete, path, ids, nil)
	if err != nil {
		return fmt.Errorf("failed to delete records: %w", err)
	}

	return nil
}
