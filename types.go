package nocodbgo

import (
	"fmt"
)

// pageInfo contains pagination information
type pageInfo struct {
	TotalRows   int  `json:"totalRows"`
	Page        int  `json:"page"`
	PageSize    int  `json:"pageSize"`
	IsFirstPage bool `json:"isFirstPage"`
	IsLastPage  bool `json:"isLastPage"`
}

// listResponse is the response from a list query
type listResponse struct {
	List     []map[string]any `json:"list"`
	PageInfo pageInfo         `json:"pageInfo"`
}

// countResponse is the response from a count query
type countResponse struct {
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
type comparisonOperator string

const (
	// Equal operator for equality
	Equal comparisonOperator = "eq"
	// NotEqual operator for inequality
	NotEqual comparisonOperator = "neq"
	// GreaterThan operator for greater than
	GreaterThan comparisonOperator = "gt"
	// GreaterThanOrEqual operator for greater than or equal
	GreaterThanOrEqual comparisonOperator = "ge"
	// LessThan operator for less than
	LessThan comparisonOperator = "lt"
	// LessThanOrEqual operator for less than or equal
	LessThanOrEqual comparisonOperator = "le"
	// Is operator for is
	Is comparisonOperator = "is"
	// IsNot operator for is not
	IsNot comparisonOperator = "isnot"
	// In operator for in
	In comparisonOperator = "in"
	// Between operator for between
	Between comparisonOperator = "btw"
	// NotBetween operator for not between
	NotBetween comparisonOperator = "nbtw"
	// Like operator for like
	Like comparisonOperator = "like"
	// NotLike operator for not like
	NotLike comparisonOperator = "nlike"
	// IsWithin operator for is within
	IsWithin comparisonOperator = "isWithin"
	// AllOf operator for all of
	AllOf comparisonOperator = "allof"
	// AnyOf operator for any of
	AnyOf comparisonOperator = "anyof"
	// NotAllOf operator for not all of
	NotAllOf comparisonOperator = "nallof"
	// NotAnyOf operator for not any of
	NotAnyOf comparisonOperator = "nanyof"
)

// dateSubOperator represents a date sub-operator
type dateSubOperator string

const (
	// Today sub-operator for today
	Today dateSubOperator = "today"
	// Tomorrow sub-operator for tomorrow
	Tomorrow dateSubOperator = "tomorrow"
	// Yesterday sub-operator for yesterday
	Yesterday dateSubOperator = "yesterday"
	// OneWeekAgo sub-operator for one week ago
	OneWeekAgo dateSubOperator = "oneWeekAgo"
	// OneWeekFromNow sub-operator for one week from now
	OneWeekFromNow dateSubOperator = "oneWeekFromNow"
	// OneMonthAgo sub-operator for one month ago
	OneMonthAgo dateSubOperator = "oneMonthAgo"
	// OneMonthFromNow sub-operator for one month from now
	OneMonthFromNow dateSubOperator = "oneMonthFromNow"
	// DaysAgo sub-operator for days ago
	DaysAgo dateSubOperator = "daysAgo"
	// DaysFromNow sub-operator for days from now
	DaysFromNow dateSubOperator = "daysFromNow"
	// ExactDate sub-operator for exact date
	ExactDate dateSubOperator = "exactDate"
)

// dateWithinSubOperator represents a date within sub-operator
type dateWithinSubOperator string

const (
	// PastWeek sub-operator for past week
	PastWeek dateWithinSubOperator = "pastWeek"
	// PastMonth sub-operator for past month
	PastMonth dateWithinSubOperator = "pastMonth"
	// PastYear sub-operator for past year
	PastYear dateWithinSubOperator = "pastYear"
	// NextWeek sub-operator for next week
	NextWeek dateWithinSubOperator = "nextWeek"
	// NextMonth sub-operator for next month
	NextMonth dateWithinSubOperator = "nextMonth"
	// NextYear sub-operator for next year
	NextYear dateWithinSubOperator = "nextYear"
	// NextNumberOfDays sub-operator for next number of days
	NextNumberOfDays dateWithinSubOperator = "nextNumberOfDays"
	// PastNumberOfDays sub-operator for past number of days
	PastNumberOfDays dateWithinSubOperator = "pastNumberOfDays"
)
