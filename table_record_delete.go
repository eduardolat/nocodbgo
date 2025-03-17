package nocodbgo

import (
	"fmt"
	"net/http"
)

// deleteRecordBuilder is used to build a delete query with a fluent API
type deleteRecordBuilder struct {
	table    *Table
	recordID int

	contextProvider[*deleteRecordBuilder]
}

// DeleteRecord initiates the construction of a delete operation for a single record.
//
// It accepts a record ID to identify which record to delete.
func (t *Table) DeleteRecord(recordID int) *deleteRecordBuilder {
	b := &deleteRecordBuilder{
		table:    t,
		recordID: recordID,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute performs the delete operation with the configured parameters.
// Returns an error if the operation fails.
func (b *deleteRecordBuilder) Execute() error {
	if b.recordID == 0 {
		return ErrRowIDRequired
	}

	err := b.table.
		DeleteRecords([]int{b.recordID}).
		WithContext(b.contextProvider.ctx).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to delete record: %w", err)
	}

	return nil
}

// deleteRecordsBuilder is used to build a bulk delete query with a fluent API
type deleteRecordsBuilder struct {
	table     *Table
	recordIDs []int

	contextProvider[*deleteRecordsBuilder]
}

// DeleteRecords initiates the construction of a bulk delete operation for multiple records.
//
// It accepts a slice of record IDs to identify which records to delete.
func (t *Table) DeleteRecords(recordIDs []int) *deleteRecordsBuilder {
	b := &deleteRecordsBuilder{
		table:     t,
		recordIDs: recordIDs,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute performs the bulk delete operation with the configured parameters.
// If no record IDs are provided, the method returns without an error.
// Returns an error if the operation fails.
func (b *deleteRecordsBuilder) Execute() error {
	if len(b.recordIDs) == 0 {
		return nil
	}

	// Convert IDs to the format expected by the API
	ids := make([]map[string]any, len(b.recordIDs))
	for i, id := range b.recordIDs {
		ids[i] = map[string]any{"Id": id}
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	_, err := b.table.client.request(b.contextProvider.ctx, http.MethodDelete, path, ids, nil)
	if err != nil {
		return fmt.Errorf("failed to delete records: %w", err)
	}

	return nil
}
