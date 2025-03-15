package nocodbgo

import (
	"fmt"
	"net/url"
)

// QueryOption is a function that configures a query
type QueryOption func(url.Values)

// PageInfo contains pagination information
type PageInfo struct {
	TotalRows   int  `json:"totalRows"`
	Page        int  `json:"page"`
	PageSize    int  `json:"pageSize"`
	IsFirstPage bool `json:"isFirstPage"`
	IsLastPage  bool `json:"isLastPage"`
}

// ListResponse is the response from a list request
type ListResponse struct {
	List     []map[string]any `json:"list"`
	PageInfo PageInfo         `json:"pageInfo"`
}

// CountResponse is the response from a count request
type CountResponse struct {
	Count int `json:"count"`
}

// APIError represents an error returned by the NocoDB API
type APIError struct {
	Message string `json:"msg"`
	Code    string `json:"code"`
}

// Error implements the error interface
func (e APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("API error: %s (code: %s)", e.Message, e.Code)
	}
	return fmt.Sprintf("API error: %s", e.Message)
}

// ComparisonOperator represents a comparison operator
type ComparisonOperator string

const (
	// Equal operator
	Equal ComparisonOperator = "eq"
	// NotEqual operator
	NotEqual ComparisonOperator = "neq"
	// GreaterThan operator
	GreaterThan ComparisonOperator = "gt"
	// GreaterThanOrEqual operator
	GreaterThanOrEqual ComparisonOperator = "ge"
	// LessThan operator
	LessThan ComparisonOperator = "lt"
	// LessThanOrEqual operator
	LessThanOrEqual ComparisonOperator = "le"
	// Is operator
	Is ComparisonOperator = "is"
	// IsNot operator
	IsNot ComparisonOperator = "isnot"
	// In operator
	In ComparisonOperator = "in"
	// Between operator
	Between ComparisonOperator = "btw"
	// NotBetween operator
	NotBetween ComparisonOperator = "nbtw"
	// Like operator
	Like ComparisonOperator = "like"
	// NotLike operator
	NotLike ComparisonOperator = "nlike"
	// IsWithin operator
	IsWithin ComparisonOperator = "isWithin"
	// AllOf operator
	AllOf ComparisonOperator = "allof"
	// AnyOf operator
	AnyOf ComparisonOperator = "anyof"
	// NotAllOf operator
	NotAllOf ComparisonOperator = "nallof"
	// NotAnyOf operator
	NotAnyOf ComparisonOperator = "nanyof"
)

// LogicalOperator represents a logical operator
type LogicalOperator string

const (
	// And operator
	And LogicalOperator = "~and"
	// Or operator
	Or LogicalOperator = "~or"
	// Not operator
	Not LogicalOperator = "~not"
)

// DateSubOperator represents a date sub-operator
type DateSubOperator string

const (
	// Today sub-operator
	Today DateSubOperator = "today"
	// Tomorrow sub-operator
	Tomorrow DateSubOperator = "tomorrow"
	// Yesterday sub-operator
	Yesterday DateSubOperator = "yesterday"
	// OneWeekAgo sub-operator
	OneWeekAgo DateSubOperator = "oneWeekAgo"
	// OneWeekFromNow sub-operator
	OneWeekFromNow DateSubOperator = "oneWeekFromNow"
	// OneMonthAgo sub-operator
	OneMonthAgo DateSubOperator = "oneMonthAgo"
	// OneMonthFromNow sub-operator
	OneMonthFromNow DateSubOperator = "oneMonthFromNow"
	// DaysAgo sub-operator
	DaysAgo DateSubOperator = "daysAgo"
	// DaysFromNow sub-operator
	DaysFromNow DateSubOperator = "daysFromNow"
	// ExactDate sub-operator
	ExactDate DateSubOperator = "exactDate"
)

// DateWithinSubOperator represents a date within sub-operator
type DateWithinSubOperator string

const (
	// PastWeek sub-operator
	PastWeek DateWithinSubOperator = "pastWeek"
	// PastMonth sub-operator
	PastMonth DateWithinSubOperator = "pastMonth"
	// PastYear sub-operator
	PastYear DateWithinSubOperator = "pastYear"
	// NextWeek sub-operator
	NextWeek DateWithinSubOperator = "nextWeek"
	// NextMonth sub-operator
	NextMonth DateWithinSubOperator = "nextMonth"
	// NextYear sub-operator
	NextYear DateWithinSubOperator = "nextYear"
	// NextNumberOfDays sub-operator
	NextNumberOfDays DateWithinSubOperator = "nextNumberOfDays"
	// PastNumberOfDays sub-operator
	PastNumberOfDays DateWithinSubOperator = "pastNumberOfDays"
)
