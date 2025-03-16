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
)

// comparisonOperator represents a comparison operator used in query filters
type comparisonOperator string

const (
	// equal represents the equality operator (=)
	equal comparisonOperator = "eq"
	// notEqual represents the inequality operator (!=)
	notEqual comparisonOperator = "neq"
	// greaterThan represents the greater than operator (>)
	greaterThan comparisonOperator = "gt"
	// greaterThanOrEqual represents the greater than or equal operator (>=)
	greaterThanOrEqual comparisonOperator = "ge"
	// lessThan represents the less than operator (<)
	lessThan comparisonOperator = "lt"
	// lessThanOrEqual represents the less than or equal operator (<=)
	lessThanOrEqual comparisonOperator = "le"
	// is represents the IS operator for checking specific states
	is comparisonOperator = "is"
	// isNot represents the IS NOT operator for excluding specific states
	isNot comparisonOperator = "isnot"
	// in represents the IN operator for checking if a value is in a set
	in comparisonOperator = "in"
	// between represents the BETWEEN operator for range checks
	between comparisonOperator = "btw"
	// notBetween represents the NOT BETWEEN operator for excluding ranges
	notBetween comparisonOperator = "nbtw"
	// like represents the LIKE operator for pattern matching
	like comparisonOperator = "like"
	// notLike represents the NOT LIKE operator for excluding patterns
	notLike comparisonOperator = "nlike"
	// isWithin represents the IS WITHIN operator for temporal queries
	isWithin comparisonOperator = "isWithin"
	// allOf represents the ALL OF operator for checking if all values in a set match
	allOf comparisonOperator = "allof"
	// anyOf represents the ANY OF operator for checking if any value in a set matches
	anyOf comparisonOperator = "anyof"
	// notAllOf represents the NOT ALL OF operator for checking if not all values in a set match
	notAllOf comparisonOperator = "nallof"
	// notAnyOf represents the NOT ANY OF operator for checking if none of the values in a set match
	notAnyOf comparisonOperator = "nanyof"
)
