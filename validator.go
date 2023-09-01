package kvalid

// Validator to implement a rule.
type Validator interface {
	SetName(string)
	Name() string
	HTMLCompatible() bool
	SetMessage(string) Validator
}

type RuleHolder[T any] interface {
	Validation(string) *Rules[T]
	Validate(string) error
}
