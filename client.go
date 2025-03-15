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
	"strconv"
	"strings"
	"time"
)

const (
	// defaultTimeout is the default timeout for HTTP requests
	defaultTimeout = 30 * time.Second
)

/******************
 * Client Options *
 *****************/

// Client is the NocoDB API client
type Client struct {
	// baseURL is the base URL for the NocoDB API
	baseURL string

	// apiToken is the API token for authentication
	apiToken string

	// httpClient is the HTTP client used to make requests
	httpClient *http.Client
}

// ClientOption is a function that configures a Client
type ClientOption func(*Client)

// WithBaseURL sets the base URL for the client
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		// Ensure the base URL doesn't end with a slash
		c.baseURL = strings.TrimSuffix(baseURL, "/")
	}
}

// WithAPIToken sets the API token for the client
func WithAPIToken(apiToken string) ClientOption {
	return func(c *Client) {
		c.apiToken = apiToken
	}
}

// WithHTTPClient sets the HTTP client for the client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithHTTPTimeout sets the timeout for the HTTP client
func WithHTTPTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		if c.httpClient == nil {
			c.httpClient = &http.Client{}
		}
		c.httpClient.Timeout = timeout
	}
}

// NewClient creates a new NocoDB client with the given options
func NewClient(clientOptions ...ClientOption) (*Client, error) {
	client := &Client{
		baseURL:    "",
		apiToken:   "",
		httpClient: &http.Client{Timeout: defaultTimeout},
	}

	for _, clientOptionFn := range clientOptions {
		clientOptionFn(client)
	}

	if client.baseURL == "" {
		return nil, ErrBaseURLRequired
	}

	if client.apiToken == "" {
		return nil, ErrAPITokenRequired
	}

	if client.httpClient == nil {
		return nil, ErrHTTPClientRequired
	}

	return client, nil
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
		var apiErr APIError
		if err := json.Unmarshal(respBody, &apiErr); err != nil {
			return nil, fmt.Errorf("status code %d: failed to unmarshal API error: %w", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("status code %d: API error: %s", resp.StatusCode, apiErr.Message)
	}

	return respBody, nil
}

// buildTablePath builds the path for table API endpoints
func (c *Client) buildTablePath(tableID string, recordID string) string {
	if recordID != "" {
		return fmt.Sprintf("/api/v2/tables/%s/records/%s", tableID, recordID)
	}
	return fmt.Sprintf("/api/v2/tables/%s/records", tableID)
}

/********************
 * Table Operations *
 *******************/

// Table represents a table in NocoDB
type Table struct {
	client  *Client
	tableID string
}

// Table returns a new table with the given ID
func (c *Client) Table(tableID string) *Table {
	return &Table{
		client:  c,
		tableID: tableID,
	}
}

// List retrieves a list of records from a table
func (t *Table) List(ctx context.Context, options ...QueryOption) (ListResponse, error) {
	query := url.Values{}
	for _, option := range options {
		option(query)
	}

	path := t.client.buildTablePath(t.tableID, "")
	respBody, err := t.client.request(ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return ListResponse{}, fmt.Errorf("failed to list records: %w", err)
	}

	var response ListResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return ListResponse{}, fmt.Errorf("failed to unmarshal list response: %w", err)
	}

	return response, nil
}

// Count returns the number of records in a table
func (t *Table) Count(ctx context.Context, options ...QueryOption) (int, error) {
	query := url.Values{}
	for _, option := range options {
		option(query)
	}

	path := t.client.buildTablePath(t.tableID, "") + "/count"
	respBody, err := t.client.request(ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return 0, fmt.Errorf("failed to count records: %w", err)
	}

	var response CountResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return 0, fmt.Errorf("failed to unmarshal count response: %w", err)
	}

	return response.Count, nil
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

// Read retrieves a record by ID from a table
func (t *Table) Read(ctx context.Context, recordID int, options ...QueryOption) (map[string]any, error) {
	if recordID == 0 {
		return nil, ErrRowIDRequired
	}

	query := url.Values{}
	for _, option := range options {
		option(query)
	}

	path := t.client.buildTablePath(t.tableID, strconv.Itoa(recordID))
	respBody, err := t.client.request(ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return nil, fmt.Errorf("failed to read record: %w", err)
	}

	var response map[string]any
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal read response: %w", err)
	}

	return response, nil
}

// Update updates a record by ID in a table
func (t *Table) Update(ctx context.Context, recordID int, data map[string]any) error {
	if recordID == 0 {
		return ErrRowIDRequired
	}

	// Add the ID to the data for bulk update
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

// Delete deletes a record by ID from a table
func (t *Table) Delete(ctx context.Context, recordID int) error {
	if recordID == 0 {
		return ErrRowIDRequired
	}

	return t.BulkDelete(ctx, []map[string]any{{"Id": strconv.Itoa(recordID)}})
}

// BulkCreate creates multiple records in a table
func (t *Table) BulkCreate(ctx context.Context, data []map[string]any) ([]int, error) {
	path := t.client.buildTablePath(t.tableID, "")
	respBody, err := t.client.request(ctx, http.MethodPost, path, data, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create records: %w", err)
	}

	var response []map[string]any
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal create response: %w", err)
	}

	ids := make([]int, len(response))
	for i, record := range response {
		id, ok := record["Id"].(int)
		if !ok {
			return nil, fmt.Errorf("created record ID is not an integer for record %d", i+1)
		}
		ids[i] = id
	}

	return ids, nil
}

// BulkUpdate updates multiple records in a table
func (t *Table) BulkUpdate(ctx context.Context, data []map[string]any) error {
	path := t.client.buildTablePath(t.tableID, "")
	respBody, err := t.client.request(ctx, http.MethodPatch, path, data, nil)
	if err != nil {
		return fmt.Errorf("failed to update records: %w", err)
	}

	var response []map[string]any
	if err := json.Unmarshal(respBody, &response); err != nil {
		return fmt.Errorf("failed to unmarshal update response: %w", err)
	}

	if len(response) != len(data) {
		return fmt.Errorf("number of updated records does not match number of records in request")
	}

	return nil
}

// BulkDelete deletes multiple records in a table
func (t *Table) BulkDelete(ctx context.Context, ids []map[string]any) error {
	path := t.client.buildTablePath(t.tableID, "")

	_, err := t.client.request(ctx, http.MethodDelete, path, ids, nil)
	if err != nil {
		return fmt.Errorf("failed to delete records: %w", err)
	}

	return nil
}
