package nocodbgo

import (
	"fmt"
	"net/http"
)

// createLinkBuilder is a fluent builder for constructing an operation that links a single target record
// to a local record via a specified link field. It encapsulates the necessary configuration of IDs
// and context for performing this singular linking operation.
type createLinkBuilder struct {
	table            *Table
	localLinkFieldID string
	localRecordID    int
	targetRecordID   int

	contextProvider[*createLinkBuilder]
}

// CreateLink begins the construction of an operation to create a link between a single target record
// and a local record using a designated link field.
//
// Parameters:
//   - localLinkFieldID: The identifier of the link field in the local table.
//   - localRecordID:    The identifier of the local record that will be linked.
//   - targetRecordID:   The identifier of the target record (from the related table) to be linked.
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

// Execute carries out the operation to create a link using the parameters configured in the builder.
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

// createLinksBuilder is a fluent builder that facilitates constructing an operation to link multiple target records
// to a local record through a designated link field. It bundles together the necessary identifiers and execution context.
type createLinksBuilder struct {
	table            *Table
	localLinkFieldID string
	localRecordID    int
	targetRecordIDs  []int

	contextProvider[*createLinksBuilder]
}

// CreateLinks initiates the construction of an operation to create links between a set of target records
// and a local record using a specified link field.
//
// Parameters:
//   - localLinkFieldID: The identifier of the link field in the local table.
//   - localRecordID:    The identifier of the local record that will be linked.
//   - targetRecordIDs:  A slice of identifiers for the target records to be linked.
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

// Execute finalizes and carries out the bulk link operation based on the builder's configuration.
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
