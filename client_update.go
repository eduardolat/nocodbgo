package nocodbgo

import (
	"context"
	"fmt"
	"net/http"
)

// updateBuilder is used to build an update query with a fluent API
type updateBuilder struct {
	table    *Table
	ctx      context.Context
	recordID int
	data     map[string]any
}

// UpdateRecord initiates the construction of an update query
func (t *Table) UpdateRecord(recordID int, data map[string]any) *updateBuilder {
	return &updateBuilder{
		table:    t,
		ctx:      nil,
		recordID: recordID,
		data:     data,
	}
}

// WithContext sets the context for the query
func (b *updateBuilder) WithContext(ctx context.Context) *updateBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the update query
func (b *updateBuilder) Execute() error {
	if b.recordID == 0 {
		return ErrRowIDRequired
	}

	// Add ID to the data
	updateData := make(map[string]any)
	for k, v := range b.data {
		updateData[k] = v
	}
	updateData["Id"] = b.recordID

	err := b.table.
		BulkUpdateRecords([]map[string]any{updateData}).
		WithContext(b.ctx).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil
}

// bulkUpdateBuilder is used to build a bulk update query with a fluent API
type bulkUpdateBuilder struct {
	table *Table
	ctx   context.Context
	data  []map[string]any
}

// BulkUpdateRecords initiates the construction of a bulk update query
func (t *Table) BulkUpdateRecords(data []map[string]any) *bulkUpdateBuilder {
	return &bulkUpdateBuilder{
		table: t,
		ctx:   nil,
		data:  data,
	}
}

// WithContext sets the context for the query
func (b *bulkUpdateBuilder) WithContext(ctx context.Context) *bulkUpdateBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the bulk update query
func (b *bulkUpdateBuilder) Execute() error {
	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	_, err := b.table.client.request(b.ctx, http.MethodPatch, path, b.data, nil)
	if err != nil {
		return fmt.Errorf("failed to update records: %w", err)
	}

	return nil
}
