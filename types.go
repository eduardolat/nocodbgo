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

// dateSubOperator represents a date sub-operator
type dateSubOperator string

const (
	// today sub-operator for today
	today dateSubOperator = "today"
	// tomorrow sub-operator for tomorrow
	tomorrow dateSubOperator = "tomorrow"
	// yesterday sub-operator for yesterday
	yesterday dateSubOperator = "yesterday"
	// oneWeekAgo sub-operator for one week ago
	oneWeekAgo dateSubOperator = "oneWeekAgo"
	// oneWeekFromNow sub-operator for one week from now
	oneWeekFromNow dateSubOperator = "oneWeekFromNow"
	// oneMonthAgo sub-operator for one month ago
	oneMonthAgo dateSubOperator = "oneMonthAgo"
	// oneMonthFromNow sub-operator for one month from now
	oneMonthFromNow dateSubOperator = "oneMonthFromNow"
	// daysAgo sub-operator for days ago
	daysAgo dateSubOperator = "daysAgo"
	// daysFromNow sub-operator for days from now
	daysFromNow dateSubOperator = "daysFromNow"
	// exactDate sub-operator for exact date
	exactDate dateSubOperator = "exactDate"
)

// dateWithinSubOperator represents a date within sub-operator
type dateWithinSubOperator string

const (
	// pastWeek sub-operator for past week
	pastWeek dateWithinSubOperator = "pastWeek"
	// pastMonth sub-operator for past month
	pastMonth dateWithinSubOperator = "pastMonth"
	// pastYear sub-operator for past year
	pastYear dateWithinSubOperator = "pastYear"
	// nextWeek sub-operator for next week
	nextWeek dateWithinSubOperator = "nextWeek"
	// nextMonth sub-operator for next month
	nextMonth dateWithinSubOperator = "nextMonth"
	// nextYear sub-operator for next year
	nextYear dateWithinSubOperator = "nextYear"
	// nextNumberOfDays sub-operator for next number of days
	nextNumberOfDays dateWithinSubOperator = "nextNumberOfDays"
	// pastNumberOfDays sub-operator for past number of days
	pastNumberOfDays dateWithinSubOperator = "pastNumberOfDays"
)
