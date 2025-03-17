package nocodbgo

import (
	"fmt"
	"net/http"
)

// deleteLinkBuilder is used to build an unlink records operation with a fluent API
type deleteLinkBuilder struct {
	table            *Table
	localLinkFieldID string
	localRecordID    int
	targetRecordID   int

	contextProvider[*deleteLinkBuilder]
}

// DeleteLink initiates the construction of an operation to unlink a record from a specific link field and record ID.
//
// It accepts a link field ID, a record ID, and a target record ID to unlink.
func (t *Table) DeleteLink(localLinkFieldID string, localRecordID int, targetRecordID int) *deleteLinkBuilder {
	b := &deleteLinkBuilder{
		table:            t,
		localLinkFieldID: localLinkFieldID,
		localRecordID:    localRecordID,
		targetRecordID:   targetRecordID,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute performs the unlink record operation with the configured parameters.
func (b *deleteLinkBuilder) Execute() error {
	if b.localLinkFieldID == "" {
		return ErrLinkFieldIDRequired
	}

	if b.localRecordID == 0 {
		return ErrRowIDRequired
	}

	if b.targetRecordID == 0 {
		return nil
	}

	return b.table.
		DeleteLinks(b.localLinkFieldID, b.localRecordID, []int{b.targetRecordID}).
		WithContext(b.contextProvider.ctx).
		Execute()
}

// deleteLinksBuilder is used to build an unlink records operation with a fluent API
type deleteLinksBuilder struct {
	table            *Table
	localLinkFieldID string
	localRecordID    int
	targetRecordIDs  []int

	contextProvider[*deleteLinksBuilder]
}

// DeleteLinks initiates the construction of an operation to unlink records from a specific link field and record ID.
//
// It accepts a link field ID, a record ID, and a slice of record IDs to unlink.
func (t *Table) DeleteLinks(localLinkFieldID string, localRecordID int, targetRecordIDs []int) *deleteLinksBuilder {
	b := &deleteLinksBuilder{
		table:            t,
		localLinkFieldID: localLinkFieldID,
		localRecordID:    localRecordID,
		targetRecordIDs:  targetRecordIDs,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute performs the unlink records operation with the configured parameters.
func (b *deleteLinksBuilder) Execute() error {
	if b.localLinkFieldID == "" {
		return ErrLinkFieldIDRequired
	}

	if b.localRecordID == 0 {
		return ErrRowIDRequired
	}

	if len(b.targetRecordIDs) == 0 {
		return nil
	}

	// Convert IDs to the format expected by the API
	ids := make([]map[string]any, len(b.targetRecordIDs))
	for i, id := range b.targetRecordIDs {
		ids[i] = map[string]any{"Id": id}
	}

	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%d", b.table.tableID, b.localLinkFieldID, b.localRecordID)
	_, err := b.table.client.request(b.contextProvider.ctx, http.MethodDelete, path, ids, nil)
	if err != nil {
		return fmt.Errorf("failed to unlink records: %w", err)
	}

	return nil
}
