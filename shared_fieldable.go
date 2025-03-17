package nocodbgo

import (
	"net/url"
	"strings"
)

// fieldable provides a reusable set of methods for building query with support for specifying
// fields to return from the query using the "fields" query parameter.
//
// It is designed to be embedded in builder types to provide consistent field selection capabilities.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
type fieldable[T any] struct {
	builder   T
	rawFields []string
}

// newFieldable creates a new fieldable instance with the given builder and apply function.
// The apply function is used to add a field to the builder and return the builder for chaining.
func newFieldable[T any](builder T) fieldable[T] {
	return fieldable[T]{
		builder:   builder,
		rawFields: []string{},
	}
}

// apply takes the url.Values and adds the "fields" query parameter to it with all the fields
// that have been added to the fieldable instance.
func (f *fieldable[T]) apply(query url.Values) {
	if len(f.rawFields) > 0 {
		query.Set("fields", strings.Join(f.rawFields, ","))
	}
}

// ReturnFields specifies which fields to include in the response.
//
// If not called, all fields will be returned.
//
// Example:
//
//	// Return only the "Name" and "Age" fields
//	query = query.ReturnFields("Name", "Age")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
func (f *fieldable[T]) ReturnFields(fields ...string) *fieldable[T] {
	f.rawFields = fields
	return f
}
