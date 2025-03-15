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

// Read initiates the construction of a read query
func (t *Table) Read(recordID int) *readBuilder {
	return &readBuilder{
		table:    t,
		ctx:      nil,
		recordID: recordID,
	}
}

// WithContext sets the context for the query
func (b *readBuilder) WithContext(ctx context.Context) *readBuilder {
	b.ctx = ctx
	return b
}

// Fields adds specific fields to the query
func (b *readBuilder) Fields(fields ...string) *readBuilder {
	b.fields = fields
	return b
}

// ReadResponse is the response from a read query
type ReadResponse struct {
	// Data contains the record data
	Data map[string]any
}

// DecodeInto converts the read response into a struct
// It takes a pointer to a struct as destination and populates it with the data
func (r ReadResponse) DecodeInto(dest any) error {
	return decodeInto(r.Data, dest)
}

// Execute executes the read query
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
