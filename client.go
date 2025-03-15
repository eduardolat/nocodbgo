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
	Msg     string `json:"msg"`
	Message string `json:"message"`
	ErrMsg  string `json:"error"`
	Code    string `json:"code"`
}

// Error implements the error interface
func (e apiError) Error() string {
	if e.Code != "" {
		if e.Msg != "" {
			return fmt.Sprintf("%s: %s", e.Code, e.Msg)
		}
		if e.Message != "" {
			return fmt.Sprintf("%s: %s", e.Code, e.Message)
		}
		if e.ErrMsg != "" {
			return fmt.Sprintf("%s: %s", e.Code, e.ErrMsg)
		}
	}

	if e.Msg != "" {
		return e.Msg
	}
	if e.Message != "" {
		return e.Message
	}
	if e.ErrMsg != "" {
		return e.ErrMsg
	}

	return "Unknown error"
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

	var req *http.Request
	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, parsedUrl.String(), reqBody)
	} else {
		req, err = http.NewRequest(method, parsedUrl.String(), reqBody)
	}

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
		return nil, fmt.Errorf("status code %d: API error: %s", resp.StatusCode, apiErr.Error())
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

// createBuilder is used to build a create query with a fluent API
type createBuilder struct {
	table *Table
	ctx   context.Context
	data  map[string]any
}

// Create initiates the construction of a create query
func (t *Table) Create(data map[string]any) *createBuilder {
	return &createBuilder{
		table: t,
		ctx:   nil,
		data:  data,
	}
}

// WithContext sets the context for the query
func (b *createBuilder) WithContext(ctx context.Context) *createBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the create query
func (b *createBuilder) Execute() (int, error) {
	records, err := b.table.
		BulkCreate([]map[string]any{b.data}).
		WithContext(b.ctx).
		Execute()
	if err != nil {
		return 0, fmt.Errorf("failed to create record: %w", err)
	}

	if len(records) == 0 {
		return 0, fmt.Errorf("no record created")
	}

	return records[0], nil
}

// updateBuilder is used to build an update query with a fluent API
type updateBuilder struct {
	table    *Table
	ctx      context.Context
	recordID int
	data     map[string]any
}

// Update initiates the construction of an update query
func (t *Table) Update(recordID int, data map[string]any) *updateBuilder {
	return &updateBuilder{
		table:    t,
		ctx:      nil,
		recordID: recordID,
		data:     data,
	}
}

// WithContext sets the context for the query
func (b *updateBuilder) WithContext(ctx context.Context) *updateBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the update query
func (b *updateBuilder) Execute() error {
	if b.recordID == 0 {
		return ErrRowIDRequired
	}

	// Add ID to the data
	updateData := make(map[string]any)
	for k, v := range b.data {
		updateData[k] = v
	}
	updateData["Id"] = b.recordID

	err := b.table.
		BulkUpdate([]map[string]any{updateData}).
		WithContext(b.ctx).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil
}

// deleteBuilder is used to build a delete query with a fluent API
type deleteBuilder struct {
	table    *Table
	ctx      context.Context
	recordID int
}

// Delete initiates the construction of a delete query
func (t *Table) Delete(recordID int) *deleteBuilder {
	return &deleteBuilder{
		table:    t,
		ctx:      nil,
		recordID: recordID,
	}
}

// WithContext sets the context for the query
func (b *deleteBuilder) WithContext(ctx context.Context) *deleteBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the delete query
func (b *deleteBuilder) Execute() error {
	if b.recordID == 0 {
		return ErrRowIDRequired
	}

	err := b.table.
		BulkDelete([]int{b.recordID}).
		WithContext(b.ctx).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to delete record: %w", err)
	}

	return nil
}

// bulkCreateBuilder is used to build a bulk create query with a fluent API
type bulkCreateBuilder struct {
	table *Table
	ctx   context.Context
	data  []map[string]any
}

// BulkCreate initiates the construction of a bulk create query
func (t *Table) BulkCreate(data []map[string]any) *bulkCreateBuilder {
	return &bulkCreateBuilder{
		table: t,
		ctx:   nil,
		data:  data,
	}
}

// WithContext sets the context for the query
func (b *bulkCreateBuilder) WithContext(ctx context.Context) *bulkCreateBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the bulk create query
func (b *bulkCreateBuilder) Execute() ([]int, error) {
	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	respBody, err := b.table.client.request(b.ctx, http.MethodPost, path, b.data, nil)
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

// bulkUpdateBuilder is used to build a bulk update query with a fluent API
type bulkUpdateBuilder struct {
	table *Table
	ctx   context.Context
	data  []map[string]any
}

// BulkUpdate initiates the construction of a bulk update query
func (t *Table) BulkUpdate(data []map[string]any) *bulkUpdateBuilder {
	return &bulkUpdateBuilder{
		table: t,
		ctx:   nil,
		data:  data,
	}
}

// WithContext sets the context for the query
func (b *bulkUpdateBuilder) WithContext(ctx context.Context) *bulkUpdateBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the bulk update query
func (b *bulkUpdateBuilder) Execute() error {
	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	_, err := b.table.client.request(b.ctx, http.MethodPatch, path, b.data, nil)
	if err != nil {
		return fmt.Errorf("failed to update records: %w", err)
	}

	return nil
}

// bulkDeleteBuilder is used to build a bulk delete query with a fluent API
type bulkDeleteBuilder struct {
	table     *Table
	ctx       context.Context
	recordIDs []int
}

// BulkDelete initiates the construction of a bulk delete query
func (t *Table) BulkDelete(recordIDs []int) *bulkDeleteBuilder {
	return &bulkDeleteBuilder{
		table:     t,
		ctx:       nil,
		recordIDs: recordIDs,
	}
}

// WithContext sets the context for the query
func (b *bulkDeleteBuilder) WithContext(ctx context.Context) *bulkDeleteBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the bulk delete query
func (b *bulkDeleteBuilder) Execute() error {
	if len(b.recordIDs) == 0 {
		return nil
	}

	// Convert IDs to the format expected by the API
	ids := make([]map[string]any, len(b.recordIDs))
	for i, id := range b.recordIDs {
		ids[i] = map[string]any{"Id": id}
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	_, err := b.table.client.request(b.ctx, http.MethodDelete, path, ids, nil)
	if err != nil {
		return fmt.Errorf("failed to delete records: %w", err)
	}

	return nil
}
