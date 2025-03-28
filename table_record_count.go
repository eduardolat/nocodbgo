package nocodbgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// countRecordsBuilder is used to build a count query with a fluent API
type countRecordsBuilder struct {
	table *Table

	contextProvider[*countRecordsBuilder]
	filterProvider[*countRecordsBuilder]
	viewIDProvider[*countRecordsBuilder]
}

// CountRecords counts the number of records in the table.
func (t *Table) CountRecords() *countRecordsBuilder {
	b := &countRecordsBuilder{
		table: t,
	}

	b.contextProvider = newContextProvider(b)
	b.filterProvider = newFilterProvider(b)
	b.viewIDProvider = newViewIDProvider(b)

	return b
}

// Execute finalizes and executes the operation.
func (b *countRecordsBuilder) Execute() (int, error) {
	query := url.Values{}
	query = b.filterProvider.apply(query)
	query = b.viewIDProvider.apply(query)

	path := fmt.Sprintf("/api/v2/tables/%s/records/count", b.table.tableID)
	respBody, err := b.table.client.request(b.contextProvider.ctx, http.MethodGet, path, nil, query)
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
