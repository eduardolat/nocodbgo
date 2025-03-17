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

// DeleteLink unlinks a single target table record from a local table record via a specified link field.
//
// Parameters:
//   - localLinkFieldID: The identifier for the link field on the local table.
//   - localRecordID:    The identifier for the local table record from which the link needs to be removed.
//   - targetRecordID:   The identifier for the target table record that needs to be unlinked.
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

// Execute finalizes and executes the operation.
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

// DeleteLinks unlinks multiple target table records from a local table record via a specified link field.
//
// Parameters:
//   - localLinkFieldID: The identifier for the link field on the local table.
//   - localRecordID:    The identifier for the local table record from which the links need to be removed.
//   - targetRecordIDs:  A slice of identifiers for the target table records that need to be unlinked.
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

// Execute finalizes and executes the operation.
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
