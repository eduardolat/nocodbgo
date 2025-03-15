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

// Count initiates the construction of a count query
func (t *Table) Count() *countBuilder {
	return &countBuilder{
		table: t,
		ctx:   nil,
	}
}

// WithContext sets the context for the query
func (b *countBuilder) WithContext(ctx context.Context) *countBuilder {
	b.ctx = ctx
	return b
}

// FilterEqualTo adds an equality filter to the query
func (b *countBuilder) FilterEqualTo(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, equal, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotEqualTo adds an inequality filter to the query
func (b *countBuilder) FilterNotEqualTo(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterGreaterThan adds a greater than filter to the query
func (b *countBuilder) FilterGreaterThan(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterGreaterThanOrEqual adds a greater than or equal filter to the query
func (b *countBuilder) FilterGreaterThanOrEqual(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLessThan adds a less than filter to the query
func (b *countBuilder) FilterLessThan(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLessThanOrEqual adds a less than or equal filter to the query
func (b *countBuilder) FilterLessThanOrEqual(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsNull adds a is null filter to the query
func (b *countBuilder) FilterIsNull(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "null")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsNotNull adds a is not null filter to the query
func (b *countBuilder) FilterIsNotNull(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isNot, "null")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsTrue adds a is true filter to the query
func (b *countBuilder) FilterIsTrue(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "true")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsFalse adds a is false filter to the query
func (b *countBuilder) FilterIsFalse(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "false")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIn adds an in filter to the query
func (b *countBuilder) FilterIn(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, in, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterBetween adds a between filter to the query
func (b *countBuilder) FilterBetween(column string, min, max string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, between, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotBetween adds a not between filter to the query
func (b *countBuilder) FilterNotBetween(column string, min, max string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, notBetween, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLike adds a like filter to the query
func (b *countBuilder) FilterLike(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, like, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotLike adds a not like filter to the query
func (b *countBuilder) FilterNotLike(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notLike, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsWithin adds an isWithin filter to the query (Available in Date and DateTime only)
func (b *countBuilder) FilterIsWithin(column string, subOperation string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isWithin, subOperation)
	b.filters = append(b.filters, filter)
	return b
}

// FilterAllOf adds an allOf filter to the query (includes all of the values)
func (b *countBuilder) FilterAllOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, allOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterAnyOf adds an anyOf filter to the query (includes any of the values)
func (b *countBuilder) FilterAnyOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, anyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotAllOf adds a notAllOf filter to the query (does not include all of the values)
func (b *countBuilder) FilterNotAllOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAllOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotAnyOf adds a notAnyOf filter to the query (does not include any of the values)
func (b *countBuilder) FilterNotAnyOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAnyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterWhere adds a custom filter expression to the query
func (b *countBuilder) FilterWhere(filter string) *countBuilder {
	if filter != "" {
		b.filters = append(b.filters, filter)
	}
	return b
}

// Execute executes the count query
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
