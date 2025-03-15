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
//
// The data parameter can be either a map[string]any or a struct with JSON tags
func (t *Table) UpdateRecord(recordID int, data any) *updateBuilder {
	var dataMap map[string]any
	var err error

	switch v := data.(type) {
	case map[string]any:
		dataMap = v
	default:
		dataMap, err = structToMap(data)
		if err != nil {
			// Return empty builder, error will be handled in Execute
			return &updateBuilder{
				table:    t,
				ctx:      nil,
				data:     nil,
				recordID: recordID,
			}
		}
	}

	return &updateBuilder{
		table:    t,
		ctx:      nil,
		recordID: recordID,
		data:     dataMap,
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

	if b.data == nil {
		return fmt.Errorf("failed to convert data to map")
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
//
// The data parameter can be either a []map[string]any or a slice of structs with JSON tags
//
// Each record must have an "Id" field
func (t *Table) BulkUpdateRecords(data any) *bulkUpdateBuilder {
	var dataMaps []map[string]any
	var err error

	switch v := data.(type) {
	case []map[string]any:
		dataMaps = v
	default:
		dataMaps, err = structsToMaps(data)
		if err != nil {
			// Return empty builder, error will be handled in Execute
			return &bulkUpdateBuilder{
				table: t,
				ctx:   nil,
				data:  nil,
			}
		}
	}

	return &bulkUpdateBuilder{
		table: t,
		ctx:   nil,
		data:  dataMaps,
	}
}

// WithContext sets the context for the query
func (b *bulkUpdateBuilder) WithContext(ctx context.Context) *bulkUpdateBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the bulk update query
func (b *bulkUpdateBuilder) Execute() error {
	if b.data == nil {
		return fmt.Errorf("failed to convert data to maps")
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	_, err := b.table.client.request(b.ctx, http.MethodPatch, path, b.data, nil)
	if err != nil {
		return fmt.Errorf("failed to update records: %w", err)
	}

	return nil
}
