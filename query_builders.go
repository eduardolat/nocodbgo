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

// ListBuilder is used to build a list query with a fluent API
type ListBuilder struct {
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
func (t *Table) List(ctx context.Context) *ListBuilder {
	return &ListBuilder{
		table: t,
		ctx:   ctx,
	}
}

// EqualTo adds an equality filter to the query
func (b *ListBuilder) EqualTo(column string, value any) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, Equal, value)
	b.filters = append(b.filters, filter)
	return b
}

// NotEqualTo adds an inequality filter to the query
func (b *ListBuilder) NotEqualTo(column string, value any) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, NotEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// GreaterThan adds a greater than filter to the query
func (b *ListBuilder) GreaterThan(column string, value any) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, GreaterThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// GreaterThanOrEqual adds a greater than or equal filter to the query
func (b *ListBuilder) GreaterThanOrEqual(column string, value any) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, GreaterThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// LessThan adds a less than filter to the query
func (b *ListBuilder) LessThan(column string, value any) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, LessThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// LessThanOrEqual adds a less than or equal filter to the query
func (b *ListBuilder) LessThanOrEqual(column string, value any) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, LessThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// IsNull adds a is null filter to the query
func (b *ListBuilder) IsNull(column string) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, Is, "null")
	b.filters = append(b.filters, filter)
	return b
}

// IsNotNull adds a is not null filter to the query
func (b *ListBuilder) IsNotNull(column string) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, IsNot, "null")
	b.filters = append(b.filters, filter)
	return b
}

// IsTrue adds a is true filter to the query
func (b *ListBuilder) IsTrue(column string) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, Is, "true")
	b.filters = append(b.filters, filter)
	return b
}

// IsFalse adds a is false filter to the query
func (b *ListBuilder) IsFalse(column string) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, Is, "false")
	b.filters = append(b.filters, filter)
	return b
}

// In adds an in filter to the query
func (b *ListBuilder) In(column string, values ...any) *ListBuilder {
	if len(values) == 0 {
		return b
	}

	var valuesStr []string
	for _, v := range values {
		valuesStr = append(valuesStr, fmt.Sprintf("%v", v))
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, In, strings.Join(valuesStr, ","))
	b.filters = append(b.filters, filter)
	return b
}

// Between adds a between filter to the query
func (b *ListBuilder) Between(column string, min, max any) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%v,%v)", column, Between, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// NotBetween adds a not between filter to the query
func (b *ListBuilder) NotBetween(column string, min, max any) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%v,%v)", column, NotBetween, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// Like adds a like filter to the query
func (b *ListBuilder) Like(column string, value any) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, Like, value)
	b.filters = append(b.filters, filter)
	return b
}

// NotLike adds a not like filter to the query
func (b *ListBuilder) NotLike(column string, value any) *ListBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, NotLike, value)
	b.filters = append(b.filters, filter)
	return b
}

// SortAsc adds an ascending sort to the query
func (b *ListBuilder) SortAsc(column string) *ListBuilder {
	b.sorts = append(b.sorts, column)
	return b
}

// SortDesc adds a descending sort to the query
func (b *ListBuilder) SortDesc(column string) *ListBuilder {
	b.sorts = append(b.sorts, "-"+column)
	return b
}

// Limit adds a limit to the query
func (b *ListBuilder) Limit(limit int) *ListBuilder {
	b.limit = limit
	return b
}

// Offset adds an offset to the query
func (b *ListBuilder) Offset(offset int) *ListBuilder {
	b.offset = offset
	return b
}

// Page adds pagination to the query
func (b *ListBuilder) Page(page, pageSize int) *ListBuilder {
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
func (b *ListBuilder) Fields(fields ...string) *ListBuilder {
	b.fields = fields
	return b
}

// Shuffle adds a shuffle parameter to the query
func (b *ListBuilder) Shuffle(shuffle bool) *ListBuilder {
	b.shuffle = shuffle
	return b
}

// Execute executes the list query
func (b *ListBuilder) Execute() (*ListResponse, error) {
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
		return nil, fmt.Errorf("failed to list records: %w", err)
	}

	var response listResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal list response: %w", err)
	}

	return &ListResponse{
		List:     response.List,
		PageInfo: PageInfo(response.PageInfo),
	}, nil
}

// ListResponse is the response from a list query
type ListResponse struct {
	List     []map[string]any
	PageInfo PageInfo
}

// PageInfo contains pagination information
type PageInfo struct {
	TotalRows   int
	Page        int
	PageSize    int
	IsFirstPage bool
	IsLastPage  bool
}

// Decode decodes the list response into a slice of structs
func (r *ListResponse) Decode(dest any) error {
	return decode(r.List, dest)
}

// CountBuilder is used to build a count query with a fluent API
type CountBuilder struct {
	table   *Table
	ctx     context.Context
	filters []string
}

// Count initiates the construction of a count query
func (t *Table) Count(ctx context.Context) *CountBuilder {
	return &CountBuilder{
		table: t,
		ctx:   ctx,
	}
}

// EqualTo adds an equality filter to the query
func (b *CountBuilder) EqualTo(column string, value any) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, Equal, value)
	b.filters = append(b.filters, filter)
	return b
}

// NotEqualTo adds an inequality filter to the query
func (b *CountBuilder) NotEqualTo(column string, value any) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, NotEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// GreaterThan adds a greater than filter to the query
func (b *CountBuilder) GreaterThan(column string, value any) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, GreaterThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// GreaterThanOrEqual adds a greater than or equal filter to the query
func (b *CountBuilder) GreaterThanOrEqual(column string, value any) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, GreaterThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// LessThan adds a less than filter to the query
func (b *CountBuilder) LessThan(column string, value any) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, LessThan, value)
	b.filters = append(b.filters, filter)
	return b
}

// LessThanOrEqual adds a less than or equal filter to the query
func (b *CountBuilder) LessThanOrEqual(column string, value any) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, LessThanOrEqual, value)
	b.filters = append(b.filters, filter)
	return b
}

// IsNull adds a is null filter to the query
func (b *CountBuilder) IsNull(column string) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, Is, "null")
	b.filters = append(b.filters, filter)
	return b
}

// IsNotNull adds a is not null filter to the query
func (b *CountBuilder) IsNotNull(column string) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, IsNot, "null")
	b.filters = append(b.filters, filter)
	return b
}

// IsTrue adds a is true filter to the query
func (b *CountBuilder) IsTrue(column string) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, Is, "true")
	b.filters = append(b.filters, filter)
	return b
}

// IsFalse adds a is false filter to the query
func (b *CountBuilder) IsFalse(column string) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%s)", column, Is, "false")
	b.filters = append(b.filters, filter)
	return b
}

// In adds an in filter to the query
func (b *CountBuilder) In(column string, values ...any) *CountBuilder {
	if len(values) == 0 {
		return b
	}

	var valuesStr []string
	for _, v := range values {
		valuesStr = append(valuesStr, fmt.Sprintf("%v", v))
	}

	filter := fmt.Sprintf("(%s,%s,%s)", column, In, strings.Join(valuesStr, ","))
	b.filters = append(b.filters, filter)
	return b
}

// Between adds a between filter to the query
func (b *CountBuilder) Between(column string, min, max any) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%v,%v)", column, Between, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// NotBetween adds a not between filter to the query
func (b *CountBuilder) NotBetween(column string, min, max any) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%v,%v)", column, NotBetween, min, max)
	b.filters = append(b.filters, filter)
	return b
}

// Like adds a like filter to the query
func (b *CountBuilder) Like(column string, value any) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, Like, value)
	b.filters = append(b.filters, filter)
	return b
}

// NotLike adds a not like filter to the query
func (b *CountBuilder) NotLike(column string, value any) *CountBuilder {
	filter := fmt.Sprintf("(%s,%s,%v)", column, NotLike, value)
	b.filters = append(b.filters, filter)
	return b
}

// Execute executes the count query
func (b *CountBuilder) Execute() (int, error) {
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

	var response countResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return 0, fmt.Errorf("failed to unmarshal count response: %w", err)
	}

	return response.Count, nil
}

// ReadBuilder is used to build a read query with a fluent API
type ReadBuilder struct {
	table    *Table
	ctx      context.Context
	recordID int
	fields   []string
}

// Read initiates the construction of a read query
func (t *Table) Read(ctx context.Context, recordID int) *ReadBuilder {
	return &ReadBuilder{
		table:    t,
		ctx:      ctx,
		recordID: recordID,
	}
}

// Fields adds specific fields to the query
func (b *ReadBuilder) Fields(fields ...string) *ReadBuilder {
	b.fields = fields
	return b
}

// Execute executes the read query
func (b *ReadBuilder) Execute() (*ReadResponse, error) {
	if b.recordID == 0 {
		return nil, ErrRowIDRequired
	}

	query := url.Values{}

	// Add fields
	if len(b.fields) > 0 {
		query.Set("fields", strings.Join(b.fields, ","))
	}

	path := fmt.Sprintf("/api/v2/tables/%s/records/%d", b.table.tableID, b.recordID)
	respBody, err := b.table.client.request(b.ctx, http.MethodGet, path, nil, query)
	if err != nil {
		return nil, fmt.Errorf("failed to read record: %w", err)
	}

	var response map[string]any
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal read response: %w", err)
	}

	return &ReadResponse{
		Data: response,
	}, nil
}

// ReadResponse is the response from a read query
type ReadResponse struct {
	Data map[string]any
}

// Decode decodes the read response into a struct
func (r *ReadResponse) Decode(dest any) error {
	return decode(r.Data, dest)
}
