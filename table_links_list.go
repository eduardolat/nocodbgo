package nocodbgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// listLinksBuilder provides a fluent interface for constructing a query to retrieve linked records.
// It encapsulates the necessary query options—such as filters, sorting, pagination, and field selection—for
// fetching records associated with a specified link field and local record.
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

// ListLinks lists the target table records linked to a local table record via a specified link field.
//
// Parameters:
//   - localLinkFieldID: the identifier of the link field used to associate records.
//   - localRecordID: the identifier of the local table record whose linked records are being retrieved.
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

// Execute finalizes and executes the operation.
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
