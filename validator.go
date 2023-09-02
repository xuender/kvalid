package kvalid

// Validator to implement a rule.
type Validator interface {
	SetName(string)
	Name() string
	HTMLCompatible() bool
	SetMessage(string) Validator
}

type RuleHolder interface {
	Validation(string) *Rules
	Validate(string) error
}
