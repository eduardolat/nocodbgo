package nocodbgo

import "context"

// contextable is a helper struct that contains a context.Context and it's intended
// to be embedded in builder types to provide a context for the request.
type contextable[T any] struct {
	builder T
	ctx     context.Context
}

// newContextable creates a new contextable instance with the given builder.
func newContextable[T any](builder T) contextable[T] {
	return contextable[T]{
		builder: builder,
		ctx:     context.Background(),
	}
}

// WithContext sets the context for the operation.
//
// This allows for request cancellation and timeout control.
//
// If not set, the context will default to context.Background().
func (c *contextable[T]) WithContext(ctx context.Context) T {
	c.ctx = ctx
	return c.builder
}
