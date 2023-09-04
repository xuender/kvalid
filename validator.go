package kvalid

import "encoding/json"

// Validator to implement a rule.
type Validator interface {
	SetName(string)
	Name() string
	HTMLCompatible() bool
	SetMessage(string) Validator
}

// RuleHolder needs to be Rules.
type RuleHolder[T any] interface {
	Validation(string) *Rules[T]
	Validate(string) error
}

// ValidJSONer to implement a Validator JSON map.
type ValidJSONer interface {
	ValidJSON() map[string]json.Marshaler
}
