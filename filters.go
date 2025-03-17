package nocodbgo

import (
	"fmt"
	"strings"
)

// Filters provides a reusable set of filter methods for building query filters.
// It is designed to be embedded in builder types to provide consistent filtering capabilities.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
type filters[T any] struct {
	builder    T
	rawFilters []string
}

// newFilters creates a new Filters instance with the given builder and apply function.
// The apply function is used to add a filter to the builder and return the builder for chaining.
func newFilters[T any](builder T) filters[T] {
	return filters[T]{
		builder:    builder,
		rawFilters: []string{},
	}
}

// Where adds a custom filter expression to the "where" query parameter of the request.
// This allows for more complex filtering logic than the predefined filter methods.
//
// Example:
//
//	query = query.Where("(Check,eq,55)~or((Amount,gt,10)~and(Amount,lt,20))")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#logical-operators
func (f *filters[T]) Where(filter string) T {
	if filter != "" {
		f.rawFilters = append(f.rawFilters, filter)
	}
	return f.builder
}

// WhereEqualsTo adds a filter to the "where" query parameter of the request that matches
// records where the specified column equals the given value.
//
// Example:
//
//	// Where MyField equals foo
//	query = query.WhereEqualsTo("MyField", "foo")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereEqualsTo(column string, value string) T {
	filter := fmt.Sprintf("(%s,eq,%s)", column, value)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereNotEqualsTo adds a filter to the "where" query parameter of the request that matches
// records where the specified column does not equal the given value.
//
// Example:
//
//	// Where MyField does not equal foo
//	query = query.WhereNotEqualsTo("MyField", "foo")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereNotEqualsTo(column string, value string) T {
	filter := fmt.Sprintf("(%s,neq,%s)", column, value)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereGreaterThan adds a filter to the "where" query parameter of the request that matches
// records where the specified column is greater than the given value.
//
// Example:
//
//	// Where MyField is greater than 55
//	query = query.WhereGreaterThan("MyField", "55")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereGreaterThan(column string, value string) T {
	filter := fmt.Sprintf("(%s,gt,%s)", column, value)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereGreaterThanOrEqual adds a filter to the "where" query parameter of the request that matches
// records where the specified column is greater than or equal to the given value.
//
// Example:
//
//	// Where MyField is greater than or equal to 55
//	query = query.WhereGreaterThanOrEqual("MyField", "55")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereGreaterThanOrEqual(column string, value string) T {
	filter := fmt.Sprintf("(%s,ge,%s)", column, value)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereLessThan adds a filter to the "where" query parameter of the request that matches
// records where the specified column is less than the given value.
//
// Example:
//
//	// Where MyField is less than 55
//	query = query.WhereLessThan("MyField", "55")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereLessThan(column string, value string) T {
	filter := fmt.Sprintf("(%s,lt,%s)", column, value)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereLessThanOrEqual adds a filter to the "where" query parameter of the request that matches
// records where the specified column is less than or equal to the given value.
//
// Example:
//
//	// Where MyField is less than or equal to 55
//	query = query.WhereLessThanOrEqual("MyField", "55")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereLessThanOrEqual(column string, value string) T {
	filter := fmt.Sprintf("(%s,le,%s)", column, value)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereIsNull adds a filter to the "where" query parameter of the request that matches
// records where the specified column is null.
//
// Example:
//
//	// Where MyField is null
//	query = query.WhereIsNull("MyField")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereIsNull(column string) T {
	filter := fmt.Sprintf("(%s,is,null)", column)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereIsNotNull adds a filter to the "where" query parameter of the request that matches
// records where the specified column is not null.
//
// Example:
//
//	// Where MyField is not null
//	query = query.WhereIsNotNull("MyField")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereIsNotNull(column string) T {
	filter := fmt.Sprintf("(%s,isnot,null)", column)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereIsTrue adds a filter to the "where" query parameter of the request that matches
// records where the specified column is true.
//
// Example:
//
//	// Where MyField is true
//	query = query.WhereIsTrue("MyField")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereIsTrue(column string) T {
	filter := fmt.Sprintf("(%s,is,true)", column)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereIsFalse adds a filter to the "where" query parameter of the request that matches
// records where the specified column is false.
//
// Example:
//
//	// Where MyField is false
//	query = query.WhereIsFalse("MyField")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereIsFalse(column string) T {
	filter := fmt.Sprintf("(%s,is,false)", column)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereIsIn adds a filter to the "where" query parameter of the request that matches
// records where the specified column's value is in the provided list of values.
//
// Example:
//
//	// Where MyField is in the list of values
//	query = query.WhereIsIn("MyField", "55", "66", "77")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereIsIn(column string, values ...string) T {
	if len(values) == 0 {
		return f.builder
	}

	filter := fmt.Sprintf("(%s,in,%s)", column, strings.Join(values, ","))
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereIsBetween adds a filter to the "where" query parameter of the request that matches
// records where the specified column's value is between the min and max values (inclusive).
//
// Example:
//
//	// Where MyField is between 55 and 66
//	query = query.WhereIsBetween("MyField", "55", "66")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereIsBetween(column string, min, max string) T {
	filter := fmt.Sprintf("(%s,btw,%s,%s)", column, min, max)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereNotIsBetween adds a filter to the "where" query parameter of the request that matches
// records where the specified column's value is not between the min and max values.
//
// Example:
//
//	// Where MyField is not between 55 and 66
//	query = query.WhereNotIsBetween("MyField", "55", "66")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereNotIsBetween(column string, min, max string) T {
	filter := fmt.Sprintf("(%s,nbtw,%s,%s)", column, min, max)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereIsLike adds a filter to match records where the specified column's value matches the given pattern.
// The pattern can include "%" as a wildcard for any sequence of characters.
//
// Example:
//
//	// Where MyField is like "Foo%"
//	// This will include "Foo Foo", "FooBar", "FooBaz", etc.
//	query = query.WhereIsLike("MyField", "Foo%")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereIsLike(column string, value string) T {
	filter := fmt.Sprintf("(%s,like,%s)", column, value)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereIsNotLike adds a filter to match records where the specified column's value does not match the given pattern.
// The pattern can include "%" as a wildcard for any sequence of characters.
//
// Example:
//
//	// Where MyField is not like "Foo%"
//	// This will not include "Foo Foo", "FooBar", "FooBaz", etc.
//	query = query.WhereIsNotLike("MyField", "Foo%")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereIsNotLike(column string, value string) T {
	filter := fmt.Sprintf("(%s,nlike,%s)", column, value)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereIsWithin adds a filter for date/datetime columns to match records within a specific time range.
// The subOperation parameter specifies the time range (e.g., "today", "yesterday", "thisWeek").
//
// This is only available for Date/DateTime columns and you can use the following subOperations:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-sub-operators
//
// Example:
//
//	// Where MyField is within the last week
//	query = query.WhereIsWithin("MyField", "oneWeekAgo")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereIsWithin(column string, subOperation string) T {
	filter := fmt.Sprintf("(%s,within,%s)", column, subOperation)
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereAllOf adds a filter to the "where" query parameter of the request that matches
// records where the specified column contains all of the provided values.
// Typically used with multi-select or array columns.
//
// Example:
//
//	// Where MyField contains all of the values
//	query = query.WhereAllOf("MyField", "55", "66", "77")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereAllOf(column string, values ...string) T {
	if len(values) == 0 {
		return f.builder
	}

	filter := fmt.Sprintf("(%s,allof,%s)", column, strings.Join(values, ","))
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereAnyOf adds a filter to the "where" query parameter of the request that matches
// records where the specified column contains any of the provided values.
// Typically used with multi-select or array columns.
//
// Example:
//
//	// Where MyField contains any of the values
//	query = query.WhereAnyOf("MyField", "55", "66", "77")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereAnyOf(column string, values ...string) T {
	if len(values) == 0 {
		return f.builder
	}

	filter := fmt.Sprintf("(%s,anyof,%s)", column, strings.Join(values, ","))
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereNotAllOf adds a filter to the "where" query parameter of the request that matches
// records where the specified column does not contain all of the provided values.
// Typically used with multi-select or array columns.
//
// Example:
//
//	// Where MyField does not contain all of the values
//	query = query.WhereNotAllOf("MyField", "55", "66", "77")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereNotAllOf(column string, values ...string) T {
	if len(values) == 0 {
		return f.builder
	}

	filter := fmt.Sprintf("(%s,nallof,%s)", column, strings.Join(values, ","))
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}

// WhereNotAnyOf adds a filter to the "where" query parameter of the request that matches
// records where the specified column does not contain any of the provided values.
// Typically used with multi-select or array columns.
//
// Example:
//
//	// Where MyField does not contain any of the values
//	query = query.WhereNotAnyOf("MyField", "55", "66", "77")
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#comparison-operators
func (f *filters[T]) WhereNotAnyOf(column string, values ...string) T {
	if len(values) == 0 {
		return f.builder
	}

	filter := fmt.Sprintf("(%s,nanyof,%s)", column, strings.Join(values, ","))
	f.rawFilters = append(f.rawFilters, filter)
	return f.builder
}
