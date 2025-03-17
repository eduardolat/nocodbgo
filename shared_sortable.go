package nocodbgo

import (
	"net/url"
	"strings"
)

// sortable provides a reusable set of sort methods for building query with support for sorting using
// the "sort" query parameter.
//
// It is designed to be embedded in builder types to provide consistent sorting capabilities.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
type sortable[T any] struct {
	builder  T
	rawSorts []string
}

// newSortable creates a new sortable instance with the given builder and apply function.
// The apply function is used to add a sort to the builder and return the builder for chaining.
func newSortable[T any](builder T) sortable[T] {
	return sortable[T]{
		builder:  builder,
		rawSorts: []string{},
	}
}

// apply takes the url.Values and adds the "sort" query parameter to it with all the sorts
// that have been added to the sortable instance.
//
// It returns a new copy of the provided url.Values with the "sort" query parameter added.
func (s *sortable[T]) apply(query url.Values) url.Values {
	if query == nil || len(s.rawSorts) < 1 {
		return query
	}

	query.Set("sort", strings.Join(s.rawSorts, ","))
	return query
}

// SortAscBy adds an ascending sort on the specified column.
//
// You can call this method multiple times to add multiple sort criteria, they will
// be applied in the order they were added.
//
// Example:
//
//	// Sort by Name ascending and Age descending
//	query = query.SortAscBy("Name").SortDescBy("Age")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
func (s *sortable[T]) SortAscBy(column string) T {
	s.rawSorts = append(s.rawSorts, column)
	return s.builder
}

// SortDescBy adds a descending sort on the specified column.
//
// You can call this method multiple times to add multiple sort criteria, they will
// be applied in the order they were added.
//
// Example:
//
//	// Sort by Name descending and Age ascending
//	query = query.SortDescBy("Name").SortAscBy("Age")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
func (s *sortable[T]) SortDescBy(column string) T {
	s.rawSorts = append(s.rawSorts, "-"+column)
	return s.builder
}
