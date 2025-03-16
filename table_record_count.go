package nocodbgo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// countBuilder is used to build a count query with a fluent API
type countBuilder struct {
	table   *Table
	ctx     context.Context
	filters []string
}

// CountRecords initiates the construction of a query to count records in a table.
// Returns a countBuilder for further configuration and execution.
func (t *Table) CountRecords() *countBuilder {
	return &countBuilder{
		table: t,
		ctx:   nil,
	}
}

// WithContext sets the context for the count operation.
// This allows for request cancellation and timeout control.
// Returns the countBuilder for method chaining.
func (b *countBuilder) WithContext(ctx context.Context) *countBuilder {
	b.ctx = ctx
	return b
}

// FilterEqualTo adds a filter to count records where the specified column equals the given value.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterEqualTo(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, equal, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotEqualTo adds a filter to count records where the specified column does not equal the given value.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterNotEqualTo(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterGreaterThan adds a filter to count records where the specified column is greater than the given value.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterGreaterThan(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterGreaterThanOrEqual adds a filter to count records where the specified column is greater than or equal to the given value.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterGreaterThanOrEqual(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLessThan adds a filter to count records where the specified column is less than the given value.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterLessThan(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLessThanOrEqual adds a filter to count records where the specified column is less than or equal to the given value.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterLessThanOrEqual(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsNull adds a filter to count records where the specified column is null.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterIsNull(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "null")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsNotNull adds a filter to count records where the specified column is not null.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterIsNotNull(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isNot, "null")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsTrue adds a filter to count records where the specified column is true.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterIsTrue(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "true")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsFalse adds a filter to count records where the specified column is false.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterIsFalse(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "false")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIn adds a filter to count records where the specified column's value is in the provided list of values.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterIn(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, in, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterBetween adds a filter to count records where the specified column's value is between the min and max values (inclusive).
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterBetween(column string, min, max string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, between, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotBetween adds a filter to count records where the specified column's value is not between the min and max values.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterNotBetween(column string, min, max string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, notBetween, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLike adds a filter to count records where the specified column's value matches the given pattern.
// The pattern can include wildcards (% for any sequence of characters, _ for a single character).
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterLike(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, like, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotLike adds a filter to count records where the specified column's value does not match the given pattern.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterNotLike(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notLike, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsWithin adds a filter for date/datetime columns to count records within a specific time range.
// The subOperation parameter specifies the time range (e.g., "today", "yesterday", "thisWeek").
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterIsWithin(column string, subOperation string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isWithin, subOperation)
	b.filters = append(b.filters, filter)
	return b
}

// FilterAllOf adds a filter to count records where the specified column contains all of the provided values.
// Typically used with multi-select or array columns.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterAllOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, allOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterAnyOf adds a filter to count records where the specified column contains any of the provided values.
// Typically used with multi-select or array columns.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterAnyOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, anyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotAllOf adds a filter to count records where the specified column does not contain all of the provided values.
// Typically used with multi-select or array columns.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterNotAllOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAllOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotAnyOf adds a filter to count records where the specified column does not contain any of the provided values.
// Typically used with multi-select or array columns.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterNotAnyOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAnyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterWhere adds a custom filter expression to the count query.
// This allows for more complex filtering logic than the predefined filter methods.
// Returns the countBuilder for method chaining.
func (b *countBuilder) FilterWhere(filter string) *countBuilder {
	if filter != "" {
		b.filters = append(b.filters, filter)
	}
	return b
}

// Execute performs the count operation with the configured parameters.
// Returns the number of records that match the filters or an error if the operation fails.
func (b *countBuilder) Execute() (int, error) {
	query := url.Values{}

	// Add filters
	for _, filter := range b.filters {
		query.Add("where", filter)
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records/count", b.table.tableID)
	respBody, err := b.table.client.request(b.ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return 0, fmt.Errorf("failed to count records: %w", err)
	}

	var response struct {
		Count int `json:"count"`
	}
	if err := json.Unmarshal(respBody, &response); err != nil {
		return 0, fmt.Errorf("failed to unmarshal count response: %w", err)
	}

	return response.Count, nil
}
