package nocodbgo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// linkedRecordsListBuilder is used to build a query to list linked records with a fluent API
type linkedRecordsListBuilder struct {
	table       *Table
	ctx         context.Context
	linkFieldID string
	recordID    int
	filters     []string
	sorts       []string
	limit       int
	offset      int
	fields      []string
}

// ListLinkedRecords initiates the construction of a query to list linked records for a specific link field and record ID.
// It accepts a link field ID and a record ID to identify which linked records to retrieve.
// Returns a linkedRecordsListBuilder for further configuration and execution.
func (t *Table) ListLinkedRecords(linkFieldID string, recordID int) *linkedRecordsListBuilder {
	return &linkedRecordsListBuilder{
		table:       t,
		ctx:         nil,
		linkFieldID: linkFieldID,
		recordID:    recordID,
	}
}

// WithContext sets the context for the list linked records operation.
// This allows for request cancellation and timeout control.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) WithContext(ctx context.Context) *linkedRecordsListBuilder {
	b.ctx = ctx
	return b
}

// FilterEqualTo adds a filter to match linked records where the specified column equals the given value.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterEqualTo(column string, value string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, equal, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotEqualTo adds a filter to match linked records where the specified column does not equal the given value.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterNotEqualTo(column string, value string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterGreaterThan adds a filter to match linked records where the specified column is greater than the given value.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterGreaterThan(column string, value string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterGreaterThanOrEqual adds a filter to match linked records where the specified column is greater than or equal to the given value.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterGreaterThanOrEqual(column string, value string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLessThan adds a filter to match linked records where the specified column is less than the given value.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterLessThan(column string, value string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLessThanOrEqual adds a filter to match linked records where the specified column is less than or equal to the given value.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterLessThanOrEqual(column string, value string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsNull adds a filter to match linked records where the specified column is null.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterIsNull(column string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "null")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsNotNull adds a filter to match linked records where the specified column is not null.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterIsNotNull(column string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isNot, "null")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsTrue adds a filter to match linked records where the specified column is true.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterIsTrue(column string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "true")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsFalse adds a filter to match linked records where the specified column is false.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterIsFalse(column string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "false")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIn adds a filter to match linked records where the specified column's value is in the provided list of values.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterIn(column string, values ...string) *linkedRecordsListBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, in, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterBetween adds a filter to match linked records where the specified column's value is between the min and max values (inclusive).
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterBetween(column string, min, max string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, between, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotBetween adds a filter to match linked records where the specified column's value is not between the min and max values.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterNotBetween(column string, min, max string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, notBetween, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLike adds a filter to match linked records where the specified column's value matches the given pattern.
// The pattern can include wildcards (% for any sequence of characters, _ for a single character).
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterLike(column string, value string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, like, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotLike adds a filter to match linked records where the specified column's value does not match the given pattern.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterNotLike(column string, value string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notLike, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsWithin adds a filter for date/datetime columns to match linked records within a specific time range.
// The subOperation parameter specifies the time range (e.g., "today", "yesterday", "thisWeek").
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterIsWithin(column string, subOperation string) *linkedRecordsListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isWithin, subOperation)
	b.filters = append(b.filters, filter)
	return b
}

// FilterAllOf adds a filter to match linked records where the specified column contains all of the provided values.
// Typically used with multi-select or array columns.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterAllOf(column string, values ...string) *linkedRecordsListBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, allOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterAnyOf adds a filter to match linked records where the specified column contains any of the provided values.
// Typically used with multi-select or array columns.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterAnyOf(column string, values ...string) *linkedRecordsListBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, anyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotAllOf adds a filter to match linked records where the specified column does not contain all of the provided values.
// Typically used with multi-select or array columns.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterNotAllOf(column string, values ...string) *linkedRecordsListBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAllOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotAnyOf adds a filter to match linked records where the specified column does not contain any of the provided values.
// Typically used with multi-select or array columns.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterNotAnyOf(column string, values ...string) *linkedRecordsListBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAnyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterWhere adds a custom filter expression to the query.
// This allows for more complex filtering logic than the predefined filter methods.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) FilterWhere(filter string) *linkedRecordsListBuilder {
	if filter != "" {
		b.filters = append(b.filters, filter)
	}
	return b
}

// SortAscBy adds an ascending sort on the specified column.
// Multiple sort criteria can be added and will be applied in the order they were added.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) SortAscBy(column string) *linkedRecordsListBuilder {
	b.sorts = append(b.sorts, column)
	return b
}

// SortDescBy adds a descending sort on the specified column.
// Multiple sort criteria can be added and will be applied in the order they were added.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) SortDescBy(column string) *linkedRecordsListBuilder {
	b.sorts = append(b.sorts, "-"+column)
	return b
}

// Limit sets the maximum number of linked records to return.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) Limit(limit int) *linkedRecordsListBuilder {
	b.limit = limit
	return b
}

// Offset sets the number of linked records to skip before returning results.
// Used for pagination in conjunction with Limit.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) Offset(offset int) *linkedRecordsListBuilder {
	b.offset = offset
	return b
}

// ReturnFields specifies which fields to include in the response.
// If not called, only the primary key and display value fields will be returned.
// Returns the linkedRecordsListBuilder for method chaining.
func (b *linkedRecordsListBuilder) ReturnFields(fields ...string) *linkedRecordsListBuilder {
	b.fields = fields
	return b
}

// Execute performs the list linked records operation with the configured parameters.
// Returns a ListResponse containing the linked records and pagination information, or an error if the operation fails.
func (b *linkedRecordsListBuilder) Execute() (ListResponse, error) {
	if b.linkFieldID == "" {
		return ListResponse{}, fmt.Errorf("link field ID is required")
	}

	if b.recordID == 0 {
		return ListResponse{}, ErrRowIDRequired
	}

	query := url.Values{}

	// Add filters
	for _, filter := range b.filters {
		query.Add("where", filter)
	}

	// Add sorts
	if len(b.sorts) > 0 {
		query.Set("sort", strings.Join(b.sorts, ","))
	}

	// Add limit
	if b.limit > 0 {
		query.Set("limit", strconv.Itoa(b.limit))
	}

	// Add offset
	if b.offset > 0 {
		query.Set("offset", strconv.Itoa(b.offset))
	}

	// Add fields
	if len(b.fields) > 0 {
		query.Set("fields", strings.Join(b.fields, ","))
	}

	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%d", b.table.tableID, b.linkFieldID, b.recordID)
	respBody, err := b.table.client.request(b.ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return ListResponse{}, fmt.Errorf("failed to list linked records: %w", err)
	}

	var response ListResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return ListResponse{}, fmt.Errorf("failed to unmarshal linked records response: %w", err)
	}

	return response, nil
}

// linkRecordsBuilder is used to build a link records operation with a fluent API
type linkRecordsBuilder struct {
	table       *Table
	ctx         context.Context
	linkFieldID string
	recordID    int
	recordIDs   []int
}

// LinkRecords initiates the construction of an operation to link records to a specific link field and record ID.
// It accepts a link field ID, a record ID, and a slice of record IDs to link.
// Returns a linkRecordsBuilder for further configuration and execution.
func (t *Table) LinkRecords(linkFieldID string, recordID int, recordIDs []int) *linkRecordsBuilder {
	return &linkRecordsBuilder{
		table:       t,
		ctx:         nil,
		linkFieldID: linkFieldID,
		recordID:    recordID,
		recordIDs:   recordIDs,
	}
}

// WithContext sets the context for the link records operation.
// This allows for request cancellation and timeout control.
// Returns the linkRecordsBuilder for method chaining.
func (b *linkRecordsBuilder) WithContext(ctx context.Context) *linkRecordsBuilder {
	b.ctx = ctx
	return b
}

// Execute performs the link records operation with the configured parameters.
// Returns an error if the operation fails.
func (b *linkRecordsBuilder) Execute() error {
	if b.linkFieldID == "" {
		return fmt.Errorf("link field ID is required")
	}

	if b.recordID == 0 {
		return ErrRowIDRequired
	}

	if len(b.recordIDs) == 0 {
		return nil
	}

	// Convert IDs to the format expected by the API
	ids := make([]map[string]any, len(b.recordIDs))
	for i, id := range b.recordIDs {
		ids[i] = map[string]any{"Id": id}
	}

	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%d", b.table.tableID, b.linkFieldID, b.recordID)
	_, err := b.table.client.request(b.ctx, http.MethodPost, path, ids, nil)
	if err != nil {
		return fmt.Errorf("failed to link records: %w", err)
	}

	return nil
}

// unlinkRecordsBuilder is used to build an unlink records operation with a fluent API
type unlinkRecordsBuilder struct {
	table       *Table
	ctx         context.Context
	linkFieldID string
	recordID    int
	recordIDs   []int
}

// UnlinkRecords initiates the construction of an operation to unlink records from a specific link field and record ID.
// It accepts a link field ID, a record ID, and a slice of record IDs to unlink.
// Returns an unlinkRecordsBuilder for further configuration and execution.
func (t *Table) UnlinkRecords(linkFieldID string, recordID int, recordIDs []int) *unlinkRecordsBuilder {
	return &unlinkRecordsBuilder{
		table:       t,
		ctx:         nil,
		linkFieldID: linkFieldID,
		recordID:    recordID,
		recordIDs:   recordIDs,
	}
}

// WithContext sets the context for the unlink records operation.
// This allows for request cancellation and timeout control.
// Returns the unlinkRecordsBuilder for method chaining.
func (b *unlinkRecordsBuilder) WithContext(ctx context.Context) *unlinkRecordsBuilder {
	b.ctx = ctx
	return b
}

// Execute performs the unlink records operation with the configured parameters.
// Returns an error if the operation fails.
func (b *unlinkRecordsBuilder) Execute() error {
	if b.linkFieldID == "" {
		return fmt.Errorf("link field ID is required")
	}

	if b.recordID == 0 {
		return ErrRowIDRequired
	}

	if len(b.recordIDs) == 0 {
		return nil
	}

	// Convert IDs to the format expected by the API
	ids := make([]map[string]any, len(b.recordIDs))
	for i, id := range b.recordIDs {
		ids[i] = map[string]any{"Id": id}
	}

	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%d", b.table.tableID, b.linkFieldID, b.recordID)
	_, err := b.table.client.request(b.ctx, http.MethodDelete, path, ids, nil)
	if err != nil {
		return fmt.Errorf("failed to unlink records: %w", err)
	}

	return nil
}
