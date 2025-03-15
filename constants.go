package nocodbgo

// comparisonOperator represents a comparison operator
type comparisonOperator string

const (
	// equal operator for equality
	equal comparisonOperator = "eq"
	// notEqual operator for inequality
	notEqual comparisonOperator = "neq"
	// greaterThan operator for greater than
	greaterThan comparisonOperator = "gt"
	// greaterThanOrEqual operator for greater than or equal
	greaterThanOrEqual comparisonOperator = "ge"
	// lessThan operator for less than
	lessThan comparisonOperator = "lt"
	// lessThanOrEqual operator for less than or equal
	lessThanOrEqual comparisonOperator = "le"
	// is operator for is
	is comparisonOperator = "is"
	// isNot operator for is not
	isNot comparisonOperator = "isnot"
	// in operator for in
	in comparisonOperator = "in"
	// between operator for between
	between comparisonOperator = "btw"
	// notBetween operator for not between
	notBetween comparisonOperator = "nbtw"
	// like operator for like
	like comparisonOperator = "like"
	// notLike operator for not like
	notLike comparisonOperator = "nlike"
	// isWithin operator for is within
	isWithin comparisonOperator = "isWithin"
	// allOf operator for all of
	allOf comparisonOperator = "allof"
	// anyOf operator for any of
	anyOf comparisonOperator = "anyof"
	// notAllOf operator for not all of
	notAllOf comparisonOperator = "nallof"
	// notAnyOf operator for not any of
	notAnyOf comparisonOperator = "nanyof"
)
