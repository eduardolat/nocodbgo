package nocodbgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// listLinksBuilder is used to build a query to list linked records with a fluent API
type listLinksBuilder struct {
	table            *Table
	localLinkFieldID string
	localRecordID    int

	contextProvider[*listLinksBuilder]
	filterProvider[*listLinksBuilder]
	sortProvider[*listLinksBuilder]
	paginationProvider[*listLinksBuilder]
	fieldProvider[*listLinksBuilder]
}

// ListLinks initiates the construction of a query to list linked records for a specific link field and record ID.
//
// It accepts a link field ID and a record ID to identify which linked records to retrieve.
func (t *Table) ListLinks(localLinkFieldID string, localRecordID int) *listLinksBuilder {
	b := &listLinksBuilder{
		table:            t,
		localLinkFieldID: localLinkFieldID,
		localRecordID:    localRecordID,
	}

	b.contextProvider = newContextProvider(b)
	b.filterProvider = newFilterProvider(b)
	b.sortProvider = newSortProvider(b)
	b.paginationProvider = newPaginationProvider(b)
	b.fieldProvider = newFieldProvider(b)

	return b
}

// Execute performs the list linked records operation with the configured parameters.
// Returns a ListResponse containing the linked records and pagination information, or an error if the operation fails.
func (b *listLinksBuilder) Execute() (ListResponse, error) {
	if b.localLinkFieldID == "" {
		return ListResponse{}, ErrLinkFieldIDRequired
	}

	if b.localRecordID == 0 {
		return ListResponse{}, ErrRowIDRequired
	}

	query := url.Values{}
	query = b.filterProvider.apply(query)
	query = b.sortProvider.apply(query)
	query = b.paginationProvider.apply(query)
	query = b.fieldProvider.apply(query)

	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%d", b.table.tableID, b.localLinkFieldID, b.localRecordID)
	respBody, err := b.table.client.request(b.contextProvider.ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return ListResponse{}, fmt.Errorf("failed to list linked records: %w", err)
	}

	var response ListResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return ListResponse{}, fmt.Errorf("failed to unmarshal linked records response: %w", err)
	}

	return response, nil
}
