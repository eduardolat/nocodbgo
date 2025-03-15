package nocodbgo

import "errors"

var (
	// ErrHTTPClientRequired is returned when an HTTP client is required but not provided
	ErrHTTPClientRequired = errors.New("HTTP client is required")

	// ErrBaseURLRequired is returned when a base URL is required but not provided
	ErrBaseURLRequired = errors.New("base URL is required")

	// ErrAPITokenRequired is returned when an API token is required but not provided
	ErrAPITokenRequired = errors.New("API token is required")

	// ErrRowIDRequired is returned when a row ID is required but not provided
	ErrRowIDRequired = errors.New("row ID is required")
)
