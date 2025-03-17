package nocodbgo

import (
	"fmt"
	"net/http"
)

// deleteLinkBuilder provides a fluent interface for configuring an operation to unlink a single target record
// from a local record via a specified link field. It encapsulates the local record and target record identifiers
// along with necessary context without revealing implementation details.
type deleteLinkBuilder struct {
	table            *Table
	localLinkFieldID string
	localRecordID    int
	targetRecordID   int

	contextProvider[*deleteLinkBuilder]
}

// DeleteLink begins constructing an unlink operation for a single record.
//
// It takes the local link field ID, the identifier of the local record from which to remove the link,
// and the target record ID that is to be unlinked.
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

// Execute finalizes the unlink operation configured in the builder.
//
// It internally delegates to the multi-unlink operation by converting the single target record ID into a slice,
// and returns an error if the unlink process fails.
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

// deleteLinksBuilder provides a fluent interface for configuring an operation to unlink multiple target records
// from a local record via a specified link field. It bundles together all required identifiers and context.
type deleteLinksBuilder struct {
	table            *Table
	localLinkFieldID string
	localRecordID    int
	targetRecordIDs  []int

	contextProvider[*deleteLinksBuilder]
}

// DeleteLinks begins constructing an unlink operation for multiple target records.
//
// It accepts the local link field ID, the local record ID from which links need to be removed,
// and a slice of target record IDs to be unlinked.
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

// Execute finalizes and executes the unlink operation using the builder's configuration.
//
// It performs the removal of links for all specified target records and returns an error if the process fails.
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
