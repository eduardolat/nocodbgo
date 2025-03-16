package nocodbgo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// readBuilder is used to build a read query with a fluent API
type readBuilder struct {
	table    *Table
	ctx      context.Context
	recordID int
	fields   []string
}

// ReadRecord initiates the construction of a read query for a single record.
// It accepts a record ID to identify which record to retrieve.
// Returns a readBuilder for further configuration and execution.
func (t *Table) ReadRecord(recordID int) *readBuilder {
	return &readBuilder{
		table:    t,
		ctx:      nil,
		recordID: recordID,
	}
}

// WithContext sets the context for the read operation.
// This allows for request cancellation and timeout control.
// Returns the readBuilder for method chaining.
func (b *readBuilder) WithContext(ctx context.Context) *readBuilder {
	b.ctx = ctx
	return b
}

// ReturnFields specifies which fields to include in the response.
// If not called, all fields will be returned.
// Returns the readBuilder for method chaining.
func (b *readBuilder) ReturnFields(fields ...string) *readBuilder {
	b.fields = fields
	return b
}

// ReadResponse is the response from a read query
type ReadResponse struct {
	// Data contains the record data
	Data map[string]any
}

// DecodeInto converts the read response data into the provided struct.
// It takes a pointer to a struct as destination and populates it with the data.
// Returns an error if the conversion fails.
func (r ReadResponse) DecodeInto(dest any) error {
	return decodeInto(r.Data, dest)
}

// Execute performs the read operation with the configured parameters.
// Returns a ReadResponse containing the record data or an error if the operation fails.
func (b *readBuilder) Execute() (ReadResponse, error) {
	if b.recordID == 0 {
		return ReadResponse{}, ErrRowIDRequired
	}

	query := url.Values{}

	// Add fields
	if len(b.fields) > 0 {
		query.Set("fields", strings.Join(b.fields, ","))
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records/%d", b.table.tableID, b.recordID)
	respBody, err := b.table.client.request(b.ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return ReadResponse{}, fmt.Errorf("failed to read record: %w", err)
	}

	var response map[string]any
	if err := json.Unmarshal(respBody, &response); err != nil {
		return ReadResponse{}, fmt.Errorf("failed to unmarshal read response: %w", err)
	}

	return ReadResponse{Data: response}, nil
}
