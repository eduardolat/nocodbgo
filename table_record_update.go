package nocodbgo

import (
	"fmt"
	"net/http"
)

// updateRecordBuilder is used to build an update query with a fluent API
type updateRecordBuilder struct {
	table    *Table
	data     map[string]any
	chainErr error // Stores any error in the chain of methods

	contextProvider[*updateRecordBuilder]
}

// UpdateRecord updates a single record in the table.
//
// Parameters:
//   - data: The data to update the record with, can be a map[string]any or a struct with JSON tags that match the table columns.
//
// Notes:
//   - The "data" parameter must contain an "Id" field to identify which record to update.
//   - It will update the fields that are present in the "data" parameter even if they are zero values, so if you want to update a single field, you can use a map.
func (t *Table) UpdateRecord(data any) *updateRecordBuilder {
	var dataMap map[string]any
	var err error

	switch v := data.(type) {
	case map[string]any:
		dataMap = v
	default:
		dataMap, err = structToMap(data)
	}

	b := &updateRecordBuilder{
		table:    t,
		data:     dataMap,
		chainErr: err,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute finalizes and executes the operation.
func (b *updateRecordBuilder) Execute() error {
	if b.chainErr != nil {
		return fmt.Errorf("error in the chain of methods: %w", b.chainErr)
	}

	err := b.table.
		UpdateRecords([]map[string]any{b.data}).
		WithContext(b.contextProvider.ctx).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil
}

// updateRecordsBuilder is used to build a bulk update query with a fluent API
type updateRecordsBuilder struct {
	table    *Table
	data     []map[string]any
	chainErr error // Stores any error in the chain of methods

	contextProvider[*updateRecordsBuilder]
}

// UpdateRecords updates multiple records in the table.
//
// Parameters:
//   - data: The data to update the records with, can be a []map[string]any or a slice of structs with JSON tags that match the table columns.
//
// Notes:
//   - Each record in the "data" parameter must have an "Id" field to identify which record to update.
//   - It will update the fields that are present in the "data" parameter even if they are zero values, so if you want to update a single field, you can use a map.
func (t *Table) UpdateRecords(data any) *updateRecordsBuilder {
	var dataMaps []map[string]any
	var err error

	switch v := data.(type) {
	case []map[string]any:
		dataMaps = v
	default:
		dataMaps, err = structsToMaps(data)
	}

	b := &updateRecordsBuilder{
		table:    t,
		data:     dataMaps,
		chainErr: err,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute finalizes and executes the operation.
func (b *updateRecordsBuilder) Execute() error {
	if b.chainErr != nil {
		return fmt.Errorf("error in the chain of methods: %w", b.chainErr)
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	_, err := b.table.client.request(b.contextProvider.ctx, http.MethodPatch, path, b.data, nil)
	if err != nil {
		return fmt.Errorf("failed to update records: %w", err)
	}

	return nil
}
