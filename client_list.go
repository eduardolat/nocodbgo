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

// ListRecords initiates the construction of a list query
func (t *Table) ListRecords() *listBuilder {
	return &listBuilder{
		table: t,
		ctx:   nil,
	}
}

// WithContext sets the context for the query
func (b *listBuilder) WithContext(ctx context.Context) *listBuilder {
	b.ctx = ctx
	return b
}

// FilterEqualTo adds an equality filter to the query
func (b *listBuilder) FilterEqualTo(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, equal, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotEqualTo adds an inequality filter to the query
func (b *listBuilder) FilterNotEqualTo(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterGreaterThan adds a greater than filter to the query
func (b *listBuilder) FilterGreaterThan(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterGreaterThanOrEqual adds a greater than or equal filter to the query
func (b *listBuilder) FilterGreaterThanOrEqual(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLessThan adds a less than filter to the query
func (b *listBuilder) FilterLessThan(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLessThanOrEqual adds a less than or equal filter to the query
func (b *listBuilder) FilterLessThanOrEqual(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsNull adds a is null filter to the query
func (b *listBuilder) FilterIsNull(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "null")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsNotNull adds a is not null filter to the query
func (b *listBuilder) FilterIsNotNull(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isNot, "null")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsTrue adds a is true filter to the query
func (b *listBuilder) FilterIsTrue(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "true")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsFalse adds a is false filter to the query
func (b *listBuilder) FilterIsFalse(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "false")
	b.filters = append(b.filters, filter)
	return b
}

// FilterIn adds an in filter to the query
func (b *listBuilder) FilterIn(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, in, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterBetween adds a between filter to the query
func (b *listBuilder) FilterBetween(column string, min, max string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, between, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotBetween adds a not between filter to the query
func (b *listBuilder) FilterNotBetween(column string, min, max string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, notBetween, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// FilterLike adds a like filter to the query
func (b *listBuilder) FilterLike(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, like, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotLike adds a not like filter to the query
func (b *listBuilder) FilterNotLike(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notLike, value)
	b.filters = append(b.filters, filter)
	return b
}

// FilterIsWithin adds an isWithin filter to the query (Available in Date and DateTime only)
func (b *listBuilder) FilterIsWithin(column string, subOperation string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isWithin, subOperation)
	b.filters = append(b.filters, filter)
	return b
}

// FilterAllOf adds an allOf filter to the query (includes all of the values)
func (b *listBuilder) FilterAllOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, allOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterAnyOf adds an anyOf filter to the query (includes any of the values)
func (b *listBuilder) FilterAnyOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, anyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotAllOf adds a notAllOf filter to the query (does not include all of the values)
func (b *listBuilder) FilterNotAllOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAllOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterNotAnyOf adds a notAnyOf filter to the query (does not include any of the values)
func (b *listBuilder) FilterNotAnyOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAnyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// FilterWhere adds a custom filter expression to the query
func (b *listBuilder) FilterWhere(filter string) *listBuilder {
	if filter != "" {
		b.filters = append(b.filters, filter)
	}
	return b
}

// SortAscBy adds an ascending sort to the query
func (b *listBuilder) SortAscBy(column string) *listBuilder {
	b.sorts = append(b.sorts, column)
	return b
}

// SortDescBy adds a descending sort to the query
func (b *listBuilder) SortDescBy(column string) *listBuilder {
	b.sorts = append(b.sorts, "-"+column)
	return b
}

// Limit adds a limit to the query
func (b *listBuilder) Limit(limit int) *listBuilder {
	b.limit = limit
	return b
}

// Offset adds an offset to the query
func (b *listBuilder) Offset(offset int) *listBuilder {
	b.offset = offset
	return b
}

// Page adds pagination to the query
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

// ReturnFields adds specific fields to the query
func (b *listBuilder) ReturnFields(fields ...string) *listBuilder {
	b.fields = fields
	return b
}

// Shuffle adds a shuffle parameter to the query
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

// DecodeInto converts the list response into a slice of structs
// It takes a pointer to a slice of structs as destination and populates it with the data
func (r ListResponse) DecodeInto(dest any) error {
	return decodeInto(r.List, dest)
}

// Execute executes the list query
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
