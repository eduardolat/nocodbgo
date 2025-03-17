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

// DeleteRecord deletes a single record in the table.
//
// Parameters:
//   - recordID: The identifier of the record to delete.
func (t *Table) DeleteRecord(recordID int) *deleteRecordBuilder {
	b := &deleteRecordBuilder{
		table:    t,
		recordID: recordID,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute finalizes and executes the operation.
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

// DeleteRecords deletes multiple records in the table.
//
// Parameters:
//   - recordIDs: A slice of record IDs to identify which records to delete.
func (t *Table) DeleteRecords(recordIDs []int) *deleteRecordsBuilder {
	b := &deleteRecordsBuilder{
		table:     t,
		recordIDs: recordIDs,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute finalizes and executes the operation.
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
