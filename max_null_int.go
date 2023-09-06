package kvalid

import (
	"encoding/json"
	"fmt"

	"gopkg.in/guregu/null.v3"
)

// MaxNullIntValidator field have minimum value.
type MaxNullIntValidator struct {
	name    string
	message string
	max     int64
}

// Name of the field.
func (p *MaxNullIntValidator) Name() string {
	return p.name
}

// SetName of the field.
func (p *MaxNullIntValidator) SetName(name string) {
	p.name = name
}

// SetMessage set error message.
func (p *MaxNullIntValidator) SetMessage(msg string) Validator {
	p.message = msg

	return p
}

// Validate the value.
func (p *MaxNullIntValidator) Validate(value null.Int) Error {
	if value.Int64 > p.max {
		return createError(p.name, p.message, fmt.Sprintf("Please decrease %s to be %d or less", p.name, p.max))
	}

	return nil
}

// MarshalJSON for this validator.
func (p *MaxNullIntValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct[int64]{
		Rule: "maxNum",
		Max:  p.max,
		Msg:  p.message,
	})
}

// HTMLCompatible for this validator.
func (p *MaxNullIntValidator) HTMLCompatible() bool {
	return true
}

// MaxNullInt field have minimum value.
func MaxNullInt(max int64) *MaxNullIntValidator {
	return &MaxNullIntValidator{
		max: max,
	}
}
