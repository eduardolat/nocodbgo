package nocodbgo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// createRecordBuilder is used to build a create query with a fluent API
type createRecordBuilder struct {
	table    *Table
	data     map[string]any
	chainErr error // Stores any error in the chain of methods

	contextProvider[*createRecordBuilder]
}

// CreateRecord initiates the construction of a create operation for a single record.
//
// It accepts data which can be either a map[string]any or a struct with JSON tags.
func (t *Table) CreateRecord(data any) *createRecordBuilder {
	var dataMap map[string]any
	var err error

	switch v := data.(type) {
	case map[string]any:
		dataMap = v
	default:
		dataMap, err = structToMap(data)
	}

	b := &createRecordBuilder{
		table:    t,
		data:     dataMap,
		chainErr: err,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute performs the create operation with the configured parameters.
// Returns the ID of the created record or an error if the operation fails.
func (b *createRecordBuilder) Execute() (int, error) {
	if b.chainErr != nil {
		return 0, fmt.Errorf("error in the chain of methods: %w", b.chainErr)
	}

	records, err := b.table.
		CreateRecords([]map[string]any{b.data}).
		WithContext(b.contextProvider.ctx).
		Execute()
	if err != nil {
		return 0, fmt.Errorf("failed to create record: %w", err)
	}

	if len(records) == 0 {
		return 0, fmt.Errorf("no record created")
	}

	return records[0], nil
}

// createRecordsBuilder is used to build a bulk create query with a fluent API
type createRecordsBuilder struct {
	table    *Table
	data     []map[string]any
	chainErr error // Stores any error in the chain of methods

	contextProvider[*createRecordsBuilder]
}

// CreateRecords initiates the construction of a bulk create operation for multiple records.
//
// It accepts data which can be either a []map[string]any or a slice of structs with JSON tags.
func (t *Table) CreateRecords(data any) *createRecordsBuilder {
	var dataMaps []map[string]any
	var err error

	switch v := data.(type) {
	case []map[string]any:
		dataMaps = v
	default:
		dataMaps, err = structsToMaps(data)
	}

	b := &createRecordsBuilder{
		table:    t,
		data:     dataMaps,
		chainErr: err,
	}

	b.contextProvider = newContextProvider(b)

	return b
}

// Execute performs the bulk create operation with the configured parameters.
// Returns a slice of IDs for the created records or an error if the operation fails.
func (b *createRecordsBuilder) Execute() ([]int, error) {
	if b.chainErr != nil {
		return nil, fmt.Errorf("error in the chain of methods: %w", b.chainErr)
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	respBody, err := b.table.client.request(b.contextProvider.ctx, http.MethodPost, path, b.data, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create records: %w", err)
	}

	var response []map[string]any
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal create response: %w", err)
	}

	var ids []int
	for _, record := range response {
		if id, ok := record["Id"].(float64); ok {
			ids = append(ids, int(id))
		}
	}

	return ids, nil
}
