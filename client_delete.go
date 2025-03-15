package nocodbgo

import (
	"context"
	"fmt"
	"net/http"
)

// deleteBuilder is used to build a delete query with a fluent API
type deleteBuilder struct {
	table    *Table
	ctx      context.Context
	recordID int
}

// DeleteRecord initiates the construction of a delete query
func (t *Table) DeleteRecord(recordID int) *deleteBuilder {
	return &deleteBuilder{
		table:    t,
		ctx:      nil,
		recordID: recordID,
	}
}

// WithContext sets the context for the query
func (b *deleteBuilder) WithContext(ctx context.Context) *deleteBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the delete query
func (b *deleteBuilder) Execute() error {
	if b.recordID == 0 {
		return ErrRowIDRequired
	}

	err := b.table.
		BulkDeleteRecords([]int{b.recordID}).
		WithContext(b.ctx).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to delete record: %w", err)
	}

	return nil
}

// bulkDeleteBuilder is used to build a bulk delete query with a fluent API
type bulkDeleteBuilder struct {
	table     *Table
	ctx       context.Context
	recordIDs []int
}

// BulkDeleteRecords initiates the construction of a bulk delete query
func (t *Table) BulkDeleteRecords(recordIDs []int) *bulkDeleteBuilder {
	return &bulkDeleteBuilder{
		table:     t,
		ctx:       nil,
		recordIDs: recordIDs,
	}
}

// WithContext sets the context for the query
func (b *bulkDeleteBuilder) WithContext(ctx context.Context) *bulkDeleteBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the bulk delete query
func (b *bulkDeleteBuilder) Execute() error {
	if len(b.recordIDs) == 0 {
		return nil
	}

	// Convert IDs to the format expected by the API
	ids := make([]map[string]any, len(b.recordIDs))
	for i, id := range b.recordIDs {
		ids[i] = map[string]any{"Id": id}
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	_, err := b.table.client.request(b.ctx, http.MethodDelete, path, ids, nil)
	if err != nil {
		return fmt.Errorf("failed to delete records: %w", err)
	}

	return nil
}
