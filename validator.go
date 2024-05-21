// Package kvalid is lightweight validation library, based on Go 1.18+ Generics.
package kvalid

import "encoding/json"

// Validator to implement a rule.
type Validator interface {
	SetName(name string)
	Name() string
	HTMLCompatible() bool
	SetMessage(message string) Validator
}

// RuleHolder needs to be Rules.
type RuleHolder[T any] interface {
	Validation(rule string) *Rules[T]
	Validate(val string) error
}

// ValidJSONer to implement a Validator JSON map.
type ValidJSONer interface {
	ValidJSON() map[string]json.Marshaler
}

// jsonStruce to JSON.
type jsonStruct[N Number] struct {
	Rule    string `json:"rule"`
	Min     N      `json:"min,omitempty"`
	Max     N      `json:"max,omitempty"`
	Pattern string `json:"pattern,omitempty"`
	Msg     string `json:"msg,omitempty"`
}
