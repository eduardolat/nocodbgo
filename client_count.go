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

// EqualTo adds an equality filter to the query
func (b *countBuilder) EqualTo(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, equal, value)
	b.filters = append(b.filters, filter)
	return b
}

// NotEqualTo adds an inequality filter to the query
func (b *countBuilder) NotEqualTo(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// GreaterThan adds a greater than filter to the query
func (b *countBuilder) GreaterThan(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// GreaterThanOrEqual adds a greater than or equal filter to the query
func (b *countBuilder) GreaterThanOrEqual(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// LessThan adds a less than filter to the query
func (b *countBuilder) LessThan(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// LessThanOrEqual adds a less than or equal filter to the query
func (b *countBuilder) LessThanOrEqual(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// IsNull adds a is null filter to the query
func (b *countBuilder) IsNull(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "null")
	b.filters = append(b.filters, filter)
	return b
}

// IsNotNull adds a is not null filter to the query
func (b *countBuilder) IsNotNull(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isNot, "null")
	b.filters = append(b.filters, filter)
	return b
}

// IsTrue adds a is true filter to the query
func (b *countBuilder) IsTrue(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "true")
	b.filters = append(b.filters, filter)
	return b
}

// IsFalse adds a is false filter to the query
func (b *countBuilder) IsFalse(column string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "false")
	b.filters = append(b.filters, filter)
	return b
}

// In adds an in filter to the query
func (b *countBuilder) In(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, in, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// Between adds a between filter to the query
func (b *countBuilder) Between(column string, min, max string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, between, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// NotBetween adds a not between filter to the query
func (b *countBuilder) NotBetween(column string, min, max string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, notBetween, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// Like adds a like filter to the query
func (b *countBuilder) Like(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, like, value)
	b.filters = append(b.filters, filter)
	return b
}

// NotLike adds a not like filter to the query
func (b *countBuilder) NotLike(column string, value string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notLike, value)
	b.filters = append(b.filters, filter)
	return b
}

// IsWithin adds an isWithin filter to the query (Available in Date and DateTime only)
func (b *countBuilder) IsWithin(column string, subOperation string) *countBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isWithin, subOperation)
	b.filters = append(b.filters, filter)
	return b
}

// AllOf adds an allOf filter to the query (includes all of the values)
func (b *countBuilder) AllOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, allOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// AnyOf adds an anyOf filter to the query (includes any of the values)
func (b *countBuilder) AnyOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, anyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// NotAllOf adds a notAllOf filter to the query (does not include all of the values)
func (b *countBuilder) NotAllOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAllOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// NotAnyOf adds a notAnyOf filter to the query (does not include any of the values)
func (b *countBuilder) NotAnyOf(column string, values ...string) *countBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAnyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// Where adds a custom filter expression to the query
func (b *countBuilder) Where(filter string) *countBuilder {
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
