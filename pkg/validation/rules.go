package validation

const (
	// Numeric rules
	RuleNumericMinimum = iota
	RuleNumericMaximum
	RuleNumericExclusiveMinimum
	RuleNumericExclusiveMaximum

	// String rules
	RuleStringMinLength
	RuleStringMaxLength
	RuleStringPattern
	RuleStringFormat

	// Array rules
	RuleArrayMinItems
	RuleArrayMaxItems
	RuleArrayUniqueItems

	// Map rules
	RuleMapMinProperties
	RuleMapMaxProperties
	RuleMapRequired
)

type Rule struct {
	Type     uint
	Function int
}

type ValidationFunction func(value interface{}) error
