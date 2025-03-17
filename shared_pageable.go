package nocodbgo

import (
	"net/url"
	"strconv"
)

// pageable provides a reusable set of methods for building query with support for pagination using
// the "limit" and "offset" query parameters.
//
// It is designed to be embedded in builder types to provide consistent pagination capabilities.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
type pageable[T any] struct {
	builder T
	limit   int
	offset  int
}

// newPageable creates a new pageable instance with the given builder and apply function.
// The apply function is used to add a page to the builder and return the builder for chaining.
func newPageable[T any](builder T) pageable[T] {
	return pageable[T]{
		builder: builder,
	}
}

// apply takes the url.Values and adds the "limit" and "offset" query parameters to it with the values
// that have been added to the pageable instance.
func (p *pageable[T]) apply(query url.Values) {
	if p.limit > 0 {
		query.Set("limit", strconv.Itoa(p.limit))
	}
	if p.offset > 0 {
		query.Set("offset", strconv.Itoa(p.offset))
	}
}

// Limit sets the limit for the number of records to return from the query.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
func (p *pageable[T]) Limit(limit int) T {
	if limit < 1 {
		return p.builder
	}

	p.limit = limit
	return p.builder
}

// Offset sets the offset for the number of records to skip from the query.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
func (p *pageable[T]) Offset(offset int) T {
	if offset < 0 {
		return p.builder
	}

	p.offset = offset
	return p.builder
}

// Page sets the page number and page size for the query to return a specific amount of records.
//
// It converts the page number and page size to a limit and offset and sets them on the pageable instance.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
func (p *pageable[T]) Page(page int, pageSize int) T {
	if page < 1 || pageSize < 1 {
		return p.builder
	}

	p.limit = pageSize
	p.offset = (page - 1) * pageSize
	return p.builder
}
