package kvalid

import (
	"encoding/json"
	"fmt"
)

// MaxNumValidator field have maximum value.
type MaxNumValidator[N Number] struct {
	name    string
	message string
	max     N
}

// Name of the field.
func (p *MaxNumValidator[N]) Name() string {
	return p.name
}

// SetName of the field.
func (p *MaxNumValidator[N]) SetName(name string) {
	p.name = name
}

// SetMessage set error message.
func (p *MaxNumValidator[N]) SetMessage(msg string) Validator {
	p.message = msg

	return p
}

// Validate the value.
func (p *MaxNumValidator[N]) Validate(value N) Error {
	if value > p.max {
		return createError(p.name, p.message, fmt.Sprintf("Please decrease %s to be %v or less", p.name, p.max))
	}

	return nil
}

// MarshalJSON for this validator.
func (p *MaxNumValidator[N]) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct[N]{
		Rule: "maxNum",
		Max:  p.max,
		Msg:  p.message,
	})
}

// HTMLCompatible for this validator.
func (p *MaxNumValidator[N]) HTMLCompatible() bool {
	return true
}

// MaxNum field have maximum value.
func MaxNum[N Number](max N) *MaxNumValidator[N] {
	return &MaxNumValidator[N]{
		max: max,
	}
}
