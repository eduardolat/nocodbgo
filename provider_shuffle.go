package nocodbgo

import "net/url"

// shuffleProvider provides a reusable set of methods for building query with support for shuffling using
// the "shuffle" query parameter.
//
// It is designed to be embedded in builder types to provide consistent shuffling capabilities.
//
// Documentation:
//   - https://docs.nocodb.com/developer-resources/rest-apis/overview/#query-params
type shuffleProvider[T any] struct {
	builder    T
	rawShuffle bool
}

// newShuffleProvider creates a new shuffleProvider instance with the given builder and apply function.
// The apply function is used to add a shuffle to the builder and return the builder for chaining.
func newShuffleProvider[T any](builder T) shuffleProvider[T] {
	return shuffleProvider[T]{
		builder:    builder,
		rawShuffle: false,
	}
}

// apply takes the url.Values and adds the "shuffle" query parameter to it with the value
// that has been added to the shuffleProvider instance.
//
// It returns a new copy of the provided url.Values with the "shuffle" query parameter added.
func (s *shuffleProvider[T]) apply(query url.Values) url.Values {
	if query == nil || !s.rawShuffle {
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
func (s *shuffleProvider[T]) Shuffle() T {
	s.rawShuffle = true
	return s.builder
}
