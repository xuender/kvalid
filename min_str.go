package kvalid

import (
	"encoding/json"
	"fmt"
)

// MinStrValidator field must have minimum length.
type MinStrValidator struct {
	name     string
	message  string
	min      int
	optional bool
}

// Name of the field.
func (p *MinStrValidator) Name() string {
	return p.name
}

// SetName of the field.
func (p *MinStrValidator) SetName(name string) {
	p.name = name
}

// SetMessage set error message.
func (p *MinStrValidator) SetMessage(msg string) Validator {
	p.message = msg

	return p
}

// Optional don't validate if the value is zero.
func (p *MinStrValidator) Optional() Validator {
	p.optional = true

	return p
}

// Validate the value.
func (p *MinStrValidator) Validate(value string) Error {
	if p.optional && value == "" {
		return nil
	}

	if len([]rune(value)) < p.min {
		return createError(p.name, p.message, fmt.Sprintf("Please lengthen %s to %d characters or more", p.name, p.min))
	}

	return nil
}

// MarshalJSON for this validator.
func (p *MinStrValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct[int]{
		Rule: "minStr",
		Min:  p.min,
		Msg:  p.message,
	})
}

// HTMLCompatible for this validator.
func (p *MinStrValidator) HTMLCompatible() bool {
	return true
}

// MinStr field must have minimum length.
func MinStr(min int) *MinStrValidator {
	return &MinStrValidator{
		min: min,
	}
}
