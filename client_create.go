package nocodbgo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// createBuilder is used to build a create query with a fluent API
type createBuilder struct {
	table    *Table
	ctx      context.Context
	data     map[string]any
	chainErr error // Stores any error in the chain of methods
}

// CreateRecord initiates the construction of a create query
// The data parameter can be either a map[string]any or a struct with JSON tags
func (t *Table) CreateRecord(data any) *createBuilder {
	var dataMap map[string]any
	var err error

	switch v := data.(type) {
	case map[string]any:
		dataMap = v
	default:
		dataMap, err = structToMap(data)
	}

	return &createBuilder{
		table:    t,
		ctx:      nil,
		data:     dataMap,
		chainErr: err,
	}
}

// WithContext sets the context for the query
func (b *createBuilder) WithContext(ctx context.Context) *createBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the create query
func (b *createBuilder) Execute() (int, error) {
	if b.chainErr != nil {
		return 0, fmt.Errorf("error in the chain of methods: %w", b.chainErr)
	}

	records, err := b.table.
		BulkCreateRecords([]map[string]any{b.data}).
		WithContext(b.ctx).
		Execute()
	if err != nil {
		return 0, fmt.Errorf("failed to create record: %w", err)
	}

	if len(records) == 0 {
		return 0, fmt.Errorf("no record created")
	}

	return records[0], nil
}

// bulkCreateBuilder is used to build a bulk create query with a fluent API
type bulkCreateBuilder struct {
	table    *Table
	ctx      context.Context
	data     []map[string]any
	chainErr error // Stores any error in the chain of methods
}

// BulkCreateRecords initiates the construction of a bulk create query
// The data parameter can be either a []map[string]any or a slice of structs with JSON tags
func (t *Table) BulkCreateRecords(data any) *bulkCreateBuilder {
	var dataMaps []map[string]any
	var err error

	switch v := data.(type) {
	case []map[string]any:
		dataMaps = v
	default:
		dataMaps, err = structsToMaps(data)
	}

	return &bulkCreateBuilder{
		table:    t,
		ctx:      nil,
		data:     dataMaps,
		chainErr: err,
	}
}

// WithContext sets the context for the query
func (b *bulkCreateBuilder) WithContext(ctx context.Context) *bulkCreateBuilder {
	b.ctx = ctx
	return b
}

// Execute executes the bulk create query
func (b *bulkCreateBuilder) Execute() ([]int, error) {
	if b.chainErr != nil {
		return nil, fmt.Errorf("error in the chain of methods: %w", b.chainErr)
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	respBody, err := b.table.client.request(b.ctx, http.MethodPost, path, b.data, nil)
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
