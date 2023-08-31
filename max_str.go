package kvalid

import (
	"fmt"

	"github.com/xuender/kvalid/json"
)

// MaxStrValidator field have maximum length.
type MaxStrValidator struct {
	name    string
	message string
	max     int
}

// Name of the field.
func (p *MaxStrValidator) Name() string {
	return p.name
}

// SetName of the field.
func (p *MaxStrValidator) SetName(name string) {
	p.name = name
}

// SetMessage set error message.
func (p *MaxStrValidator) SetMessage(msg string) Validator {
	p.message = msg

	return p
}

// Validate the value.
func (p *MaxStrValidator) Validate(value string) Error {
	if len([]rune(value)) > p.max {
		return createError(p.name, p.message, fmt.Sprintf("Please shorten %s to %d characters or less", p.name, p.max))
	}

	return nil
}

// MarshalJSON for this validator.
func (p *MaxStrValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct[int]{
		Rule: "maxStr",
		Max:  p.max,
		Msg:  p.message,
	})
}

// HTMLCompatible for this validator.
func (p *MaxStrValidator) HTMLCompatible() bool {
	return true
}

// MaxStr field have maximum length.
func MaxStr(max int) *MaxStrValidator {
	return &MaxStrValidator{
		max: max,
	}
}
