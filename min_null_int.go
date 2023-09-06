package kvalid

import (
	"encoding/json"
	"fmt"

	"gopkg.in/guregu/null.v3"
)

// MinNullIntValidator field have minimum value.
type MinNullIntValidator struct {
	name     string
	message  string
	min      int64
	optional bool
}

// Name of the field.
func (p *MinNullIntValidator) Name() string {
	return p.name
}

// SetName of the field.
func (p *MinNullIntValidator) SetName(name string) {
	p.name = name
}

// SetMessage set error message.
func (p *MinNullIntValidator) SetMessage(msg string) Validator {
	p.message = msg

	return p
}

// Optional don't validate if the value is zero.
func (p *MinNullIntValidator) Optional() Validator {
	p.optional = true

	return p
}

// Validate the value.
func (p *MinNullIntValidator) Validate(value null.Int) Error {
	if p.optional && value.Int64 == 0 {
		return nil
	}

	if value.Int64 < p.min {
		return createError(p.name, p.message, fmt.Sprintf("Please increase %s to be %v or more", p.name, p.min))
	}

	return nil
}

// MarshalJSON for this validator.
func (p *MinNullIntValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct[int64]{
		Rule: "minNum",
		Min:  p.min,
		Msg:  p.message,
	})
}

// HTMLCompatible for this validator.
func (p *MinNullIntValidator) HTMLCompatible() bool {
	return true
}

// MinNullInt field have minimum value.
func MinNullInt(min int64) *MinNullIntValidator {
	return &MinNullIntValidator{
		min: min,
	}
}
