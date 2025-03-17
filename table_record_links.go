package nocodbgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// linkedRecordsListBuilder is used to build a query to list linked records with a fluent API
type linkedRecordsListBuilder struct {
	table       *Table
	linkFieldID string
	recordID    int

	contextable[*linkedRecordsListBuilder]
	filterable[*linkedRecordsListBuilder]
	sortable[*linkedRecordsListBuilder]
	pageable[*linkedRecordsListBuilder]
	fieldable[*linkedRecordsListBuilder]
}

// ListLinkedRecords initiates the construction of a query to list linked records for a specific link field and record ID.
// It accepts a link field ID and a record ID to identify which linked records to retrieve.
// Returns a linkedRecordsListBuilder for further configuration and execution.
func (t *Table) ListLinkedRecords(linkFieldID string, recordID int) *linkedRecordsListBuilder {
	b := &linkedRecordsListBuilder{
		table:       t,
		linkFieldID: linkFieldID,
		recordID:    recordID,
	}

	b.contextable = newContextable(b)
	b.filterable = newFilterable(b)
	b.sortable = newSortable(b)
	b.pageable = newPageable(b)
	b.fieldable = newFieldable(b)

	return b
}

// Execute performs the list linked records operation with the configured parameters.
// Returns a ListResponse containing the linked records and pagination information, or an error if the operation fails.
func (b *linkedRecordsListBuilder) Execute() (ListResponse, error) {
	if b.linkFieldID == "" {
		return ListResponse{}, ErrLinkFieldIDRequired
	}

	if b.recordID == 0 {
		return ListResponse{}, ErrRowIDRequired
	}

	query := url.Values{}
	query = b.filterable.apply(query)
	query = b.sortable.apply(query)
	query = b.pageable.apply(query)
	query = b.fieldable.apply(query)

	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%d", b.table.tableID, b.linkFieldID, b.recordID)
	respBody, err := b.table.client.request(b.contextable.ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return ListResponse{}, fmt.Errorf("failed to list linked records: %w", err)
	}

	var response ListResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return ListResponse{}, fmt.Errorf("failed to unmarshal linked records response: %w", err)
	}

	return response, nil
}

// linkRecordsBuilder is used to build a link records operation with a fluent API
type linkRecordsBuilder struct {
	table       *Table
	linkFieldID string
	recordID    int
	recordIDs   []int

	contextable[*linkRecordsBuilder]
}

// LinkRecords initiates the construction of an operation to link records to a specific link field and record ID.
// It accepts a link field ID, a record ID, and a slice of record IDs to link.
// Returns a linkRecordsBuilder for further configuration and execution.
func (t *Table) LinkRecords(linkFieldID string, recordID int, recordIDs []int) *linkRecordsBuilder {
	b := &linkRecordsBuilder{
		table:       t,
		linkFieldID: linkFieldID,
		recordID:    recordID,
		recordIDs:   recordIDs,
	}

	b.contextable = newContextable(b)

	return b
}

// Execute performs the link records operation with the configured parameters.
// Returns an error if the operation fails.
func (b *linkRecordsBuilder) Execute() error {
	if b.linkFieldID == "" {
		return ErrLinkFieldIDRequired
	}

	if b.recordID == 0 {
		return ErrRowIDRequired
	}

	if len(b.recordIDs) == 0 {
		return nil
	}

	// Convert IDs to the format expected by the API
	ids := make([]map[string]any, len(b.recordIDs))
	for i, id := range b.recordIDs {
		ids[i] = map[string]any{"Id": id}
	}

	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%d", b.table.tableID, b.linkFieldID, b.recordID)
	_, err := b.table.client.request(b.contextable.ctx, http.MethodPost, path, ids, nil)
	if err != nil {
		return fmt.Errorf("failed to link records: %w", err)
	}

	return nil
}

// unlinkRecordsBuilder is used to build an unlink records operation with a fluent API
type unlinkRecordsBuilder struct {
	table       *Table
	linkFieldID string
	recordID    int
	recordIDs   []int

	contextable[*unlinkRecordsBuilder]
}

// UnlinkRecords initiates the construction of an operation to unlink records from a specific link field and record ID.
// It accepts a link field ID, a record ID, and a slice of record IDs to unlink.
// Returns an unlinkRecordsBuilder for further configuration and execution.
func (t *Table) UnlinkRecords(linkFieldID string, recordID int, recordIDs []int) *unlinkRecordsBuilder {
	b := &unlinkRecordsBuilder{
		table:       t,
		linkFieldID: linkFieldID,
		recordID:    recordID,
		recordIDs:   recordIDs,
	}

	b.contextable = newContextable(b)

	return b
}

// Execute performs the unlink records operation with the configured parameters.
// Returns an error if the operation fails.
func (b *unlinkRecordsBuilder) Execute() error {
	if b.linkFieldID == "" {
		return ErrLinkFieldIDRequired
	}

	if b.recordID == 0 {
		return ErrRowIDRequired
	}

	if len(b.recordIDs) == 0 {
		return nil
	}

	// Convert IDs to the format expected by the API
	ids := make([]map[string]any, len(b.recordIDs))
	for i, id := range b.recordIDs {
		ids[i] = map[string]any{"Id": id}
	}

	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%d", b.table.tableID, b.linkFieldID, b.recordID)
	_, err := b.table.client.request(b.contextable.ctx, http.MethodDelete, path, ids, nil)
	if err != nil {
		return fmt.Errorf("failed to unlink records: %w", err)
	}

	return nil
}
