package nocodbgo

import (
	"net/url"
	"strconv"
)

// paginationProvider provides a reusable set of methods for building query with support for pagination using
// the "limit" and "offset" query parameters.
//
// It is designed to be embedded in builder types to provide consistent pagination capabilities.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
type paginationProvider[T any] struct {
	builder   T
	rawLimit  int
	rawOffset int
}

// newPaginationProvider creates a new paginationProvider instance with the given builder and apply function.
// The apply function is used to add a page to the builder and return the builder for chaining.
func newPaginationProvider[T any](builder T) paginationProvider[T] {
	return paginationProvider[T]{
		builder: builder,
	}
}

// apply takes the url.Values and adds the "limit" and "offset" query parameters to it with the values
// that have been added to the paginationProvider instance.
//
// It returns a new copy of the provided url.Values with the "limit" and "offset" query parameters added.
func (p *paginationProvider[T]) apply(query url.Values) url.Values {
	if query == nil {
		return query
	}

	if p.rawLimit > 0 {
		query.Set("limit", strconv.Itoa(p.rawLimit))
	}
	if p.rawOffset > 0 {
		query.Set("offset", strconv.Itoa(p.rawOffset))
	}

	return query
}

// Limit sets the limit for the number of records to return from the query.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
func (p *paginationProvider[T]) Limit(limit int) T {
	if limit < 1 {
		return p.builder
	}

	p.rawLimit = limit
	return p.builder
}

// Offset sets the offset for the number of records to skip from the query.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
func (p *paginationProvider[T]) Offset(offset int) T {
	if offset < 0 {
		return p.builder
	}

	p.rawOffset = offset
	return p.builder
}

// Page sets the page number and page size for the query to return a specific amount of records.
//
// It converts the page number and page size to a limit and offset and sets them on the pageable instance.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
func (p *paginationProvider[T]) Page(page int, pageSize int) T {
	if page < 1 || pageSize < 1 {
		return p.builder
	}

	p.rawLimit = pageSize
	p.rawOffset = (page - 1) * pageSize
	return p.builder
}
