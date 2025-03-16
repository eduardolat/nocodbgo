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

// listBuilder is used to build a list query with a fluent API
type listBuilder struct {
	table   *Table
	ctx     context.Context
	filters []string
	sorts   []string
	limit   int
	offset  int
	fields  []string
	shuffle bool
}

// ListRecords initiates the construction of a query to list records from a table.
// Returns a listBuilder for further configuration and execution.
func (t *Table) ListRecords() *listBuilder {
	return &listBuilder{
		table: t,
		ctx:   nil,
	}
}

// WithContext sets the context for the list operation.
// This allows for request cancellation and timeout control.
// Returns the listBuilder for method chaining.
func (b *listBuilder) WithContext(ctx context.Context) *listBuilder {
	b.ctx = ctx
	return b
}

// FilterEqualTo adds a filter to match records where the specified column equals the given value.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterEqualTo(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, equal, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotEqualTo adds a filter to match records where the specified column does not equal the given value.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterNotEqualTo(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterGreaterThan adds a filter to match records where the specified column is greater than the given value.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterGreaterThan(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterGreaterThanOrEqual adds a filter to match records where the specified column is greater than or equal to the given value.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterGreaterThanOrEqual(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLessThan adds a filter to match records where the specified column is less than the given value.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterLessThan(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLessThanOrEqual adds a filter to match records where the specified column is less than or equal to the given value.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterLessThanOrEqual(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsNull adds a filter to match records where the specified column is null.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterIsNull(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "null")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsNotNull adds a filter to match records where the specified column is not null.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterIsNotNull(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isNot, "null")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsTrue adds a filter to match records where the specified column is true.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterIsTrue(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "true")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsFalse adds a filter to match records where the specified column is false.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterIsFalse(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "false")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIn adds a filter to match records where the specified column's value is in the provided list of values.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterIn(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, in, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterBetween adds a filter to match records where the specified column's value is between the min and max values (inclusive).
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterBetween(column string, min, max string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, between, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotBetween adds a filter to match records where the specified column's value is not between the min and max values.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterNotBetween(column string, min, max string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, notBetween, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLike adds a filter to match records where the specified column's value matches the given pattern.
// The pattern can include wildcards (% for any sequence of characters, _ for a single character).
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterLike(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, like, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotLike adds a filter to match records where the specified column's value does not match the given pattern.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterNotLike(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notLike, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsWithin adds a filter for date/datetime columns to match records within a specific time range.
// The subOperation parameter specifies the time range (e.g., "today", "yesterday", "thisWeek").
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterIsWithin(column string, subOperation string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isWithin, subOperation)
	b.filters = append(b.filters, filter)
	return b
}

// FilterAllOf adds a filter to match records where the specified column contains all of the provided values.
// Typically used with multi-select or array columns.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterAllOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, allOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterAnyOf adds a filter to match records where the specified column contains any of the provided values.
// Typically used with multi-select or array columns.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterAnyOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, anyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotAllOf adds a filter to match records where the specified column does not contain all of the provided values.
// Typically used with multi-select or array columns.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterNotAllOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAllOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotAnyOf adds a filter to match records where the specified column does not contain any of the provided values.
// Typically used with multi-select or array columns.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterNotAnyOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAnyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterWhere adds a custom filter expression to the query.
// This allows for more complex filtering logic than the predefined filter methods.
// Returns the listBuilder for method chaining.
func (b *listBuilder) FilterWhere(filter string) *listBuilder {
	if filter != "" {
		b.filters = append(b.filters, filter)
	}
	return b
}

// SortAscBy adds an ascending sort on the specified column.
// Multiple sort criteria can be added and will be applied in the order they were added.
// Returns the listBuilder for method chaining.
func (b *listBuilder) SortAscBy(column string) *listBuilder {
	b.sorts = append(b.sorts, column)
	return b
}

// SortDescBy adds a descending sort on the specified column.
// Multiple sort criteria can be added and will be applied in the order they were added.
// Returns the listBuilder for method chaining.
func (b *listBuilder) SortDescBy(column string) *listBuilder {
	b.sorts = append(b.sorts, "-"+column)
	return b
}

// Limit sets the maximum number of records to return.
// Returns the listBuilder for method chaining.
func (b *listBuilder) Limit(limit int) *listBuilder {
	b.limit = limit
	return b
}

// Offset sets the number of records to skip before returning results.
// Used for pagination in conjunction with Limit.
// Returns the listBuilder for method chaining.
func (b *listBuilder) Offset(offset int) *listBuilder {
	b.offset = offset
	return b
}

// Page configures pagination by setting both limit and offset based on page number and page size.
// Page numbers start at 1. If page is less than 1, it defaults to 1.
// If pageSize is less than 1, it defaults to 10.
// Returns the listBuilder for method chaining.
func (b *listBuilder) Page(page, pageSize int) *listBuilder {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	b.limit = pageSize
	b.offset = (page - 1) * pageSize
	return b
}

// ReturnFields specifies which fields to include in the response.
// If not called, all fields will be returned.
// Returns the listBuilder for method chaining.
func (b *listBuilder) ReturnFields(fields ...string) *listBuilder {
	b.fields = fields
	return b
}

// Shuffle enables or disables random ordering of results.
// When enabled, results will be returned in a random order.
// Returns the listBuilder for method chaining.
func (b *listBuilder) Shuffle(shuffle bool) *listBuilder {
	b.shuffle = shuffle
	return b
}

// ListResponse is the response from a list query with pagination information
type ListResponse struct {
	// List contains the records returned by the query
	List []map[string]any `json:"list"`
	// PageInfo contains pagination information
	PageInfo PageInfo `json:"pageInfo"`
}

// PageInfo contains pagination information for list queries
type PageInfo struct {
	// TotalRows is the total number of rows in the table
	TotalRows int `json:"totalRows"`
	// Page is the current page number
	Page int `json:"page"`
	// PageSize is the number of records per page
	PageSize int `json:"pageSize"`
	// IsFirstPage indicates if this is the first page
	IsFirstPage bool `json:"isFirstPage"`
	// IsLastPage indicates if this is the last page
	IsLastPage bool `json:"isLastPage"`
}

// DecodeInto converts the list response data into a slice of the provided struct type.
// It takes a pointer to a slice of structs as destination and populates it with the data.
// Returns an error if the conversion fails.
func (r ListResponse) DecodeInto(dest any) error {
	return decodeInto(r.List, dest)
}

// Execute performs the list operation with the configured parameters.
// Returns a ListResponse containing the records and pagination information, or an error if the operation fails.
func (b *listBuilder) Execute() (ListResponse, error) {
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

	// Add shuffle
	if b.shuffle {
		query.Set("shuffle", "1")
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records", b.table.tableID)
	respBody, err := b.table.client.request(b.ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return ListResponse{}, fmt.Errorf("failed to list records: %w", err)
	}

	var response ListResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return ListResponse{}, fmt.Errorf("failed to unmarshal list response: %w", err)
	}

	return response, nil
}
