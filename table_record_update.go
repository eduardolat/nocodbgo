package nocodbgo

import (
	"fmt"
	"net/http"
)

// updateRecordBuilder is used to build an update query with a fluent API
type updateRecordBuilder struct {
	table    *Table
	recordID int
	data     map[string]any
	chainErr error // Stores any error in the chain of methods

	contextProvider[*updateRecordBuilder]
}

// UpdateRecord initiates the construction of an update query for a single record.
//
// It accepts a record ID and the data to update, which can be either a map[string]any or a struct with JSON tags.
func (t *Table) UpdateRecord(recordID int, data any) *updateRecordBuilder {
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
		recordID: recordID,
		data:     dataMap,
		chainErr: err,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute performs the update operation with the configured parameters.
// Returns an error if the operation fails.
func (b *updateRecordBuilder) Execute() error {
	if b.recordID == 0 {
		return ErrRowIDRequired
	}

	if b.chainErr != nil {
		return fmt.Errorf("error in the chain of methods: %w", b.chainErr)
	}

	// Add ID to the data
	updateData := make(map[string]any)
	for k, v := range b.data {
		updateData[k] = v
	}
	updateData["Id"] = b.recordID

	err := b.table.
		UpdateRecords([]map[string]any{updateData}).
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

// UpdateRecords initiates the construction of a bulk update query for multiple records.
//
// It accepts data which can be either a []map[string]any or a slice of structs with JSON tags.
//
// Each record must have an "Id" field to identify which record to update.
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

// Execute performs the bulk update operation with the configured parameters.
// Returns an error if the operation fails.
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
