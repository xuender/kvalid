package kvalid

import (
	"encoding/json"
	"fmt"
)

// MinNumValidator field have minimum value.
type MinNumValidator[N Number] struct {
	name     string
	message  string
	min      N
	optional bool
}

// Name of the field.
func (p *MinNumValidator[N]) Name() string {
	return p.name
}

// SetName of the field.
func (p *MinNumValidator[N]) SetName(name string) {
	p.name = name
}

// SetMessage set error message.
func (p *MinNumValidator[N]) SetMessage(msg string) Validator {
	p.message = msg

	return p
}

// Optional don't validate if the value is zero.
func (p *MinNumValidator[N]) Optional() Validator {
	p.optional = true

	return p
}

// Validate the value.
func (p *MinNumValidator[N]) Validate(value N) Error {
	if p.optional && value == 0 {
		return nil
	}

	if value < p.min {
		return createError(p.name, p.message, fmt.Sprintf("Please increase %s to be %v or more", p.name, p.min))
	}

	return nil
}

// MarshalJSON for this validator.
func (p *MinNumValidator[N]) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct[N]{
		Rule: "minNum",
		Min:  p.min,
		Msg:  p.message,
	})
}

// HTMLCompatible for this validator.
func (p *MinNumValidator[N]) HTMLCompatible() bool {
	return true
}

// MinNum field have minimum value.
func MinNum[N Number](min N) *MinNumValidator[N] {
	return &MinNumValidator[N]{
		min: min,
	}
}
