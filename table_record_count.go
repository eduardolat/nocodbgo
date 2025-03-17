package nocodbgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// countBuilder is used to build a count query with a fluent API
type countBuilder struct {
	table *Table

	contextable[*countBuilder]
	filterable[*countBuilder]
	viewSelectable[*countBuilder]
}

// CountRecords initiates the construction of a query to count records in a table.
// Returns a countBuilder for further configuration and execution.
func (t *Table) CountRecords() *countBuilder {
	b := &countBuilder{
		table: t,
	}

	b.contextable = newContextable(b)
	b.filterable = newFilterable(b)
	b.viewSelectable = newViewSelectable(b)

	return b
}

// Execute performs the count operation with the configured parameters.
// Returns the number of records that match the filters or an error if the operation fails.
func (b *countBuilder) Execute() (int, error) {
	query := url.Values{}
	query = b.filterable.apply(query)
	query = b.viewSelectable.apply(query)

	path := fmt.Sprintf("/api/v2/tables/%s/records/count", b.table.tableID)
	respBody, err := b.table.client.request(b.contextable.ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return 0, fmt.Errorf("failed to count records: %w", err)
	}

	var response struct {
		Count int `json:"count"`
	}
	if err := json.Unmarshal(respBody, &response); err != nil {
		return 0, fmt.Errorf("failed to unmarshal count response: %w", err)
	}

	return response.Count, nil
}
