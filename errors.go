package nocodbgo

import "errors"

var (
	// ErrHTTPClientRequired is returned when attempting to create a client without providing an HTTP client
	ErrHTTPClientRequired = errors.New("HTTP client is required")

	// ErrBaseURLRequired is returned when attempting to create a client without providing a base URL
	ErrBaseURLRequired = errors.New("base URL is required")

	// ErrAPITokenRequired is returned when attempting to create a client without providing an API token
	ErrAPITokenRequired = errors.New("API token is required")

	// ErrRowIDRequired is returned when attempting to perform an operation that requires a row ID without providing one
	ErrRowIDRequired = errors.New("row ID is required")

	// ErrLinkFieldIDRequired is returned when attempting to perform an operation that requires a link field ID without providing one
	ErrLinkFieldIDRequired = errors.New("link field ID is required")
)
