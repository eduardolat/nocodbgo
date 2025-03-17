package nocodbgo

import (
	"fmt"
	"net/http"
)

// createLinkBuilder provides a fluent interface for building an operation that links a single target record
// to a local record through a designated link field. It holds the necessary identifiers and execution context.
type createLinkBuilder struct {
	table            *Table
	localLinkFieldID string
	localRecordID    int
	targetRecordID   int

	contextProvider[*createLinkBuilder]
}

// CreateLink initializes a builder for creating a link between a single target record and a local record.
//
// Parameters:
//   - localLinkFieldID: The identifier for the link field on the local table.
//   - localRecordID:    The identifier for the local record to which the target will be linked.
//   - targetRecordID:   The identifier for the target record that will be linked.
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

// Execute carries out the linking operation defined in the builder. It ensures that the required parameters are set
// and delegates to the multi-link operation if necessary.
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

// createLinksBuilder provides a fluent interface for constructing an operation to link multiple target records
// to a single local record via a specified link field. It bundles the required identifiers and execution context.
type createLinksBuilder struct {
	table            *Table
	localLinkFieldID string
	localRecordID    int
	targetRecordIDs  []int

	contextProvider[*createLinksBuilder]
}

// CreateLinks initializes a builder for creating links between a local record and multiple target records.
//
// Parameters:
//   - localLinkFieldID: The identifier for the link field on the local table.
//   - localRecordID:    The identifier for the local record to which the targets will be linked.
//   - targetRecordIDs:  A slice of identifiers corresponding to the target records to be linked.
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

// Execute finalizes and executes the bulk linking operation as configured in the builder.
// It validates the essential parameters, transforms the target record IDs into the expected payload format,
// and performs the API call to establish the links.
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

	// Convert IDs to the payload format expected by the API
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
