package nocodbgo

import (
	"net/url"
	"strings"
)

// SortDirection represents the direction of sorting
type SortDirection string

const (
	// SortAsc represents ascending sort direction
	SortAsc SortDirection = "asc"
	// SortDesc represents descending sort direction
	SortDesc SortDirection = "desc"
)

// Sort represents a sort condition
type Sort struct {
	Column    string
	Direction SortDirection
}

// NewSort creates a new sort condition
func NewSort(column string, direction SortDirection) Sort {
	return Sort{
		Column:    column,
		Direction: direction,
	}
}

// String returns the string representation of the sort condition
func (s Sort) String() string {
	if s.Direction == SortDesc {
		return "-" + s.Column
	}
	return s.Column
}

// WithSort adds a sort condition to the query
func WithSort(sort Sort) QueryOption {
	return func(values url.Values) {
		values.Set("sort", sort.String())
	}
}

// WithSorts adds multiple sort conditions to the query
func WithSorts(sorts ...Sort) QueryOption {
	return func(values url.Values) {
		if len(sorts) == 0 {
			return
		}

		var sortStrings []string
		for _, sort := range sorts {
			sortStrings = append(sortStrings, sort.String())
		}

		values.Set("sort", strings.Join(sortStrings, ","))
	}
}

// Asc creates a sort condition with ascending direction
func Asc(column string) Sort {
	return NewSort(column, SortAsc)
}

// Desc creates a sort condition with descending direction
func Desc(column string) Sort {
	return NewSort(column, SortDesc)
}
