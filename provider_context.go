package nocodbgo

import "context"

// contextProvider is a helper struct that contains a context.Context and it's intended
// to be embedded in builder types to provide a context for the request.
type contextProvider[T any] struct {
	builder T
	ctx     context.Context
}

// newContextProvider creates a new contextProvider instance with the given builder.
func newContextProvider[T any](builder T) contextProvider[T] {
	return contextProvider[T]{
		builder: builder,
		ctx:     context.Background(),
	}
}

// WithContext sets the context for the operation.
//
// This allows for request cancellation and timeout control.
//
// If not set, the context will default to context.Background().
func (c *contextProvider[T]) WithContext(ctx context.Context) T {
	c.ctx = ctx
	return c.builder
}
