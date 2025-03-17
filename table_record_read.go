package nocodbgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// readRecordBuilder is used to build a read query with a fluent API
type readRecordBuilder struct {
	table    *Table
	recordID int

	contextProvider[*readRecordBuilder]
	fieldProvider[*readRecordBuilder]
}

// ReadRecord initiates the construction of a read query for a single record.
// It accepts a record ID to identify which record to retrieve.
// Returns a readBuilder for further configuration and execution.
func (t *Table) ReadRecord(recordID int) *readRecordBuilder {
	b := &readRecordBuilder{
		table:    t,
		recordID: recordID,
	}

	b.contextProvider = newContextProvider(b)
	b.fieldProvider = newFieldProvider(b)

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

// Execute finalizes and executes the operation.
func (b *readRecordBuilder) Execute() (ReadResponse, error) {
	if b.recordID == 0 {
		return ReadResponse{}, ErrRowIDRequired
	}

	query := url.Values{}
	query = b.fieldProvider.apply(query)

	path := fmt.Sprintf("/api/v2/tables/%s/records/%d", b.table.tableID, b.recordID)
	respBody, err := b.table.client.request(b.contextProvider.ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return ReadResponse{}, fmt.Errorf("failed to read record: %w", err)
	}

	var response map[string]any
	if err := json.Unmarshal(respBody, &response); err != nil {
		return ReadResponse{}, fmt.Errorf("failed to unmarshal read response: %w", err)
	}

	return ReadResponse{Data: response}, nil
}
