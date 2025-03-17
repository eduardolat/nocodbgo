package nocodbgo

import "net/url"

// shuffleable provides a reusable set of methods for building query with support for shuffling using
// the "shuffle" query parameter.
//
// It is designed to be embedded in builder types to provide consistent shuffling capabilities.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
type shuffleable[T any] struct {
	builder    T
	rawShuffle bool
}

// newShuffleable creates a new shuffleable instance with the given builder and apply function.
// The apply function is used to add a shuffle to the builder and return the builder for chaining.
func newShuffleable[T any](builder T) shuffleable[T] {
	return shuffleable[T]{
		builder:    builder,
		rawShuffle: false,
	}
}

// apply takes the url.Values and adds the "shuffle" query parameter to it with the value
// that has been added to the shuffleable instance.
//
// It returns a new copy of the provided url.Values with the "shuffle" query parameter added.
func (s *shuffleable[T]) apply(query url.Values) url.Values {
	if query == nil || s.rawShuffle == false {
		return query
	}

	query.Set("shuffle", "1")
	return query
}

// Shuffle enables random ordering of results.
//
// When enabled, results will be returned in a random order.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
func (s *shuffleable[T]) Shuffle() T {
	s.rawShuffle = true
	return s.builder
}
