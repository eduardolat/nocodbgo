package nocodbgo

import (
	"fmt"
	"net/http"
)

// createLinkBuilder is used to build a link records operation with a fluent API
type createLinkBuilder struct {
	table            *Table
	localLinkFieldID string
	localRecordID    int
	targetRecordID   int

	contextProvider[*createLinkBuilder]
}

// CreateLink initiates the construction of an operation to link records to a specific link field and record ID.
//
// It accepts a link field ID, a record ID, and a slice of record IDs to link.
func (t *Table) CreateLink(localLinkFieldID string, localRecordID int, targetRecordID int) *createLinkBuilder {
	b := &createLinkBuilder{
		table:            t,
		localLinkFieldID: localLinkFieldID,
		localRecordID:    localRecordID,
		targetRecordID:   targetRecordID,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute performs the link record operation with the configured parameters.
func (b *createLinkBuilder) Execute() error {
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
		CreateLinks(b.localLinkFieldID, b.localRecordID, []int{b.targetRecordID}).
		WithContext(b.contextProvider.ctx).
		Execute()
}

// createLinksBuilder is used to build a link records operation with a fluent API
type createLinksBuilder struct {
	table            *Table
	localLinkFieldID string
	localRecordID    int
	targetRecordIDs  []int

	contextProvider[*createLinksBuilder]
}

// CreateLinks initiates the construction of an operation to link records to a specific link field and record ID.
//
// It accepts a link field ID, a record ID, and a slice of record IDs to link.
func (t *Table) CreateLinks(localLinkFieldID string, localRecordID int, targetRecordIDs []int) *createLinksBuilder {
	b := &createLinksBuilder{
		table:            t,
		localLinkFieldID: localLinkFieldID,
		localRecordID:    localRecordID,
		targetRecordIDs:  targetRecordIDs,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute performs the link records operation with the configured parameters.
func (b *createLinksBuilder) Execute() error {
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
	targetIDS := make([]map[string]any, len(b.targetRecordIDs))
	for i, id := range b.targetRecordIDs {
		targetIDS[i] = map[string]any{"Id": id}
	}

	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%d", b.table.tableID, b.localLinkFieldID, b.localRecordID)
	_, err := b.table.client.request(b.contextProvider.ctx, http.MethodPost, path, targetIDS, nil)
	if err != nil {
		return fmt.Errorf("failed to link records: %w", err)
	}

	return nil
}
