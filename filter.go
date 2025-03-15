package nocodbgo

import (
	"fmt"
	"net/url"
	"strings"
)

// Filter represents a filter condition
type Filter struct {
	Column    string
	Operator  ComparisonOperator
	Value     any
	SubValues []any
}

// FilterGroup represents a group of filters with a logical operator
type FilterGroup struct {
	Filters  []Filter
	Groups   []FilterGroup
	Operator LogicalOperator
}

// NewFilter creates a new filter
func NewFilter(column string, operator ComparisonOperator, value any, subValues ...any) Filter {
	return Filter{
		Column:    column,
		Operator:  operator,
		Value:     value,
		SubValues: subValues,
	}
}

// NewFilterGroup creates a new filter group with the given operator
func NewFilterGroup(operator LogicalOperator, filters ...Filter) FilterGroup {
	return FilterGroup{
		Filters:  filters,
		Operator: operator,
	}
}

// AddFilter adds a filter to the filter group
func (fg *FilterGroup) AddFilter(filter Filter) *FilterGroup {
	fg.Filters = append(fg.Filters, filter)
	return fg
}

// AddGroup adds a filter group to the filter group
func (fg *FilterGroup) AddGroup(group FilterGroup) *FilterGroup {
	fg.Groups = append(fg.Groups, group)
	return fg
}

// String returns the string representation of the filter
func (f Filter) String() string {
	if len(f.SubValues) == 0 {
		return fmt.Sprintf("(%s,%s,%v)", f.Column, f.Operator, f.Value)
	}

	values := []string{fmt.Sprintf("%v", f.Value)}
	for _, v := range f.SubValues {
		values = append(values, fmt.Sprintf("%v", v))
	}

	return fmt.Sprintf("(%s,%s,%s)", f.Column, f.Operator, strings.Join(values, ","))
}

// String returns the string representation of the filter group
func (fg FilterGroup) String() string {
	var parts []string

	for _, f := range fg.Filters {
		parts = append(parts, f.String())
	}

	for _, g := range fg.Groups {
		parts = append(parts, g.String())
	}

	if len(parts) == 1 {
		return parts[0]
	}

	return fmt.Sprintf("(%s)", strings.Join(parts, string(fg.Operator)))
}

// WithFilter adds a filter to the query
func WithFilter(filter Filter) QueryOption {
	return func(values url.Values) {
		values.Set("where", filter.String())
	}
}

// WithFilterGroup adds a filter group to the query
func WithFilterGroup(filterGroup FilterGroup) QueryOption {
	return func(values url.Values) {
		values.Set("where", filterGroup.String())
	}
}

// EqualFilter creates a filter with the Equal operator
func EqualFilter(column string, value any) Filter {
	return NewFilter(column, Equal, value)
}

// NotEqualFilter creates a filter with the NotEqual operator
func NotEqualFilter(column string, value any) Filter {
	return NewFilter(column, NotEqual, value)
}

// GreaterThanFilter creates a filter with the GreaterThan operator
func GreaterThanFilter(column string, value any) Filter {
	return NewFilter(column, GreaterThan, value)
}

// GreaterThanOrEqualFilter creates a filter with the GreaterThanOrEqual operator
func GreaterThanOrEqualFilter(column string, value any) Filter {
	return NewFilter(column, GreaterThanOrEqual, value)
}

// LessThanFilter creates a filter with the LessThan operator
func LessThanFilter(column string, value any) Filter {
	return NewFilter(column, LessThan, value)
}

// LessThanOrEqualFilter creates a filter with the LessThanOrEqual operator
func LessThanOrEqualFilter(column string, value any) Filter {
	return NewFilter(column, LessThanOrEqual, value)
}

// IsNullFilter creates a filter with the Is operator and null value
func IsNullFilter(column string) Filter {
	return NewFilter(column, Is, "null")
}

// IsNotNullFilter creates a filter with the IsNot operator and null value
func IsNotNullFilter(column string) Filter {
	return NewFilter(column, IsNot, "null")
}

// IsTrueFilter creates a filter with the Is operator and true value
func IsTrueFilter(column string) Filter {
	return NewFilter(column, Is, "true")
}

// IsFalseFilter creates a filter with the Is operator and false value
func IsFalseFilter(column string) Filter {
	return NewFilter(column, Is, "false")
}

// InFilter creates a filter with the In operator
func InFilter(column string, values ...any) Filter {
	if len(values) == 0 {
		return NewFilter(column, In, "")
	}
	return NewFilter(column, In, values[0], values[1:]...)
}

// BetweenFilter creates a filter with the Between operator
func BetweenFilter(column string, min, max any) Filter {
	return NewFilter(column, Between, min, max)
}

// NotBetweenFilter creates a filter with the NotBetween operator
func NotBetweenFilter(column string, min, max any) Filter {
	return NewFilter(column, NotBetween, min, max)
}

// LikeFilter creates a filter with the Like operator
func LikeFilter(column string, value any) Filter {
	return NewFilter(column, Like, value)
}

// NotLikeFilter creates a filter with the NotLike operator
func NotLikeFilter(column string, value any) Filter {
	return NewFilter(column, NotLike, value)
}

// IsWithinDateFilter creates a filter with the IsWithin operator for dates
func IsWithinDateFilter(column string, subOperator DateWithinSubOperator, value ...any) Filter {
	if len(value) == 0 {
		return NewFilter(column, IsWithin, subOperator)
	}
	return NewFilter(column, IsWithin, subOperator, value...)
}

// AllOfFilter creates a filter with the AllOf operator
func AllOfFilter(column string, values ...any) Filter {
	if len(values) == 0 {
		return NewFilter(column, AllOf, "")
	}
	return NewFilter(column, AllOf, values[0], values[1:]...)
}

// AnyOfFilter creates a filter with the AnyOf operator
func AnyOfFilter(column string, values ...any) Filter {
	if len(values) == 0 {
		return NewFilter(column, AnyOf, "")
	}
	return NewFilter(column, AnyOf, values[0], values[1:]...)
}

// NotAllOfFilter creates a filter with the NotAllOf operator
func NotAllOfFilter(column string, values ...any) Filter {
	if len(values) == 0 {
		return NewFilter(column, NotAllOf, "")
	}
	return NewFilter(column, NotAllOf, values[0], values[1:]...)
}

// NotAnyOfFilter creates a filter with the NotAnyOf operator
func NotAnyOfFilter(column string, values ...any) Filter {
	if len(values) == 0 {
		return NewFilter(column, NotAnyOf, "")
	}
	return NewFilter(column, NotAnyOf, values[0], values[1:]...)
}

// AndGroup creates a filter group with the And operator
func AndGroup(filters ...Filter) FilterGroup {
	return NewFilterGroup(And, filters...)
}

// OrGroup creates a filter group with the Or operator
func OrGroup(filters ...Filter) FilterGroup {
	return NewFilterGroup(Or, filters...)
}

// NotGroup creates a filter group with the Not operator
func NotGroup(filter Filter) FilterGroup {
	return NewFilterGroup(Not, filter)
}
