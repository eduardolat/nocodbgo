package nocodbgo

import (
	"net/url"
	"strconv"
	"strings"
)

// WithLimit adds a limit to the query
func WithLimit(limit int) QueryOption {
	return func(values url.Values) {
		values.Set("limit", strconv.Itoa(limit))
	}
}

// WithOffset adds an offset to the query
func WithOffset(offset int) QueryOption {
	return func(values url.Values) {
		values.Set("offset", strconv.Itoa(offset))
	}
}

// WithPage adds a page to the query
func WithPage(page, pageSize int) QueryOption {
	return func(values url.Values) {
		if page < 1 {
			page = 1
		}
		if pageSize < 1 {
			pageSize = 10
		}
		values.Set("limit", strconv.Itoa(pageSize))
		values.Set("offset", strconv.Itoa((page-1)*pageSize))
	}
}

// WithFields adds fields to the query
func WithFields(fields ...string) QueryOption {
	return func(values url.Values) {
		if len(fields) == 0 {
			return
		}
		values.Set("fields", strings.Join(fields, ","))
	}
}

// WithShuffle adds a shuffle parameter to the query
func WithShuffle(shuffle bool) QueryOption {
	return func(values url.Values) {
		if shuffle {
			values.Set("shuffle", "1")
		} else {
			values.Set("shuffle", "0")
		}
	}
}
