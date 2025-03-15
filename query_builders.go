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

// List initiates the construction of a list query
func (t *Table) List(ctx context.Context) *listBuilder {
	return &listBuilder{
		table: t,
		ctx:   ctx,
	}
}

// EqualTo adds an equality filter to the query
func (b *listBuilder) EqualTo(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, equal, value)
	b.filters = append(b.filters, filter)
	return b
}

// NotEqualTo adds an inequality filter to the query
func (b *listBuilder) NotEqualTo(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// GreaterThan adds a greater than filter to the query
func (b *listBuilder) GreaterThan(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// GreaterThanOrEqual adds a greater than or equal filter to the query
func (b *listBuilder) GreaterThanOrEqual(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, greaterThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// LessThan adds a less than filter to the query
func (b *listBuilder) LessThan(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// LessThanOrEqual adds a less than or equal filter to the query
func (b *listBuilder) LessThanOrEqual(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, lessThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// IsNull adds a is null filter to the query
func (b *listBuilder) IsNull(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "null")
	b.filters = append(b.filters, filter)
	return b
}

// IsNotNull adds a is not null filter to the query
func (b *listBuilder) IsNotNull(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isNot, "null")
	b.filters = append(b.filters, filter)
	return b
}

// IsTrue adds a is true filter to the query
func (b *listBuilder) IsTrue(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "true")
	b.filters = append(b.filters, filter)
	return b
}

// IsFalse adds a is false filter to the query
func (b *listBuilder) IsFalse(column string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, is, "false")
	b.filters = append(b.filters, filter)
	return b
}

// In adds an in filter to the query
func (b *listBuilder) In(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, in, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// Between adds a between filter to the query
func (b *listBuilder) Between(column string, min, max string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, between, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// NotBetween adds a not between filter to the query
func (b *listBuilder) NotBetween(column string, min, max string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s,%s)", column, notBetween, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// Like adds a like filter to the query
func (b *listBuilder) Like(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, like, value)
	b.filters = append(b.filters, filter)
	return b
}

// NotLike adds a not like filter to the query
func (b *listBuilder) NotLike(column string, value string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, notLike, value)
	b.filters = append(b.filters, filter)
	return b
}

// IsWithin adds an isWithin filter to the query (Available in Date and DateTime only)
func (b *listBuilder) IsWithin(column string, subOperation string) *listBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, isWithin, subOperation)
	b.filters = append(b.filters, filter)
	return b
}

// AllOf adds an allOf filter to the query (includes all of the values)
func (b *listBuilder) AllOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, allOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// AnyOf adds an anyOf filter to the query (includes any of the values)
func (b *listBuilder) AnyOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, anyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// NotAllOf adds a notAllOf filter to the query (does not include all of the values)
func (b *listBuilder) NotAllOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAllOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// NotAnyOf adds a notAnyOf filter to the query (does not include any of the values)
func (b *listBuilder) NotAnyOf(column string, values ...string) *listBuilder {
	if len(values) == 0 {
		return b
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, notAnyOf, strings.Join(values, ","))
	b.filters = append(b.filters, filter)
	return b
}

// SortAsc adds an ascending sort to the query
func (b *listBuilder) SortAsc(column string) *listBuilder {
	b.sorts = append(b.sorts, column)
	return b
}

// SortDesc adds a descending sort to the query
func (b *listBuilder) SortDesc(column string) *listBuilder {
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

// Fields adds specific fields to the query
func (b *listBuilder) Fields(fields ...string) *listBuilder {
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

// Decode converts the list response into a slice of structs
// It takes a pointer to a slice of structs as destination and populates it with the data
func (r ListResponse) Decode(dest any) error {
	return decode(r.List, dest)
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

// countBuilder is used to build a count query with a fluent API
type countBuilder struct {
	table   *Table
	ctx     context.Context
	filters []string
}

// Count initiates the construction of a count query
func (t *Table) Count(ctx context.Context) *countBuilder {
	return &countBuilder{
		table: t,
		ctx:   ctx,
	}
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

// readBuilder is used to build a read query with a fluent API
type readBuilder struct {
	table    *Table
	ctx      context.Context
	recordID int
	fields   []string
}

// Read initiates the construction of a read query
func (t *Table) Read(ctx context.Context, recordID int) *readBuilder {
	return &readBuilder{
		table:    t,
		ctx:      ctx,
		recordID: recordID,
	}
}

// Fields adds specific fields to the query
func (b *readBuilder) Fields(fields ...string) *readBuilder {
	b.fields = fields
	return b
}

// ReadResponse is the response from a read query
type ReadResponse struct {
	// Data contains the record data
	Data map[string]any
}

// Decode converts the read response into a struct
// It takes a pointer to a struct as destination and populates it with the data
func (r ReadResponse) Decode(dest any) error {
	return decode(r.Data, dest)
}

// Execute executes the read query
func (b *readBuilder) Execute() (ReadResponse, error) {
	if b.recordID == 0 {
		return ReadResponse{}, ErrRowIDRequired
	}

	query := url.Values{}

	// Add fields
	if len(b.fields) > 0 {
		query.Set("fields", strings.Join(b.fields, ","))
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records/%d", b.table.tableID, b.recordID)
	respBody, err := b.table.client.request(b.ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return ReadResponse{}, fmt.Errorf("failed to read record: %w", err)
	}

	var response map[string]any
	if err := json.Unmarshal(respBody, &response); err != nil {
		return ReadResponse{}, fmt.Errorf("failed to unmarshal read response: %w", err)
	}

	return ReadResponse{Data: response}, nil
}
