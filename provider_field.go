package nocodbgo

import (
	"net/url"
	"strings"
)

// fieldProvider provides a reusable set of methods for building query with support for specifying
// fields to return from the query using the "fields" query parameter.
//
// It is designed to be embedded in builder types to provide consistent field selection capabilities.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
type fieldProvider[T any] struct {
	builder   T
	rawFields []string
}

// newFieldProvider creates a new fieldProvider instance with the given builder and apply function.
// The apply function is used to add a field to the builder and return the builder for chaining.
func newFieldProvider[T any](builder T) fieldProvider[T] {
	return fieldProvider[T]{
		builder:   builder,
		rawFields: []string{},
	}
}

// apply takes the url.Values and adds the "fields" query parameter to it with all the fields
// that have been added to the fieldProvider instance.
//
// It returns a new copy of the provided url.Values with the "fields" query parameter added.
func (f *fieldProvider[T]) apply(query url.Values) url.Values {
	if query == nil || len(f.rawFields) < 1 {
		return query
	}

	query.Set("fields", strings.Join(f.rawFields, ","))
	return query
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
func (f *fieldProvider[T]) ReturnFields(fields ...string) T {
	f.rawFields = fields
	return f.builder
}
