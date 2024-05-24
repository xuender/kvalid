package kvalid

import "encoding/json"

// IgnoreValidator only for bind.
type IgnoreValidator struct {
	name    string
	message string
}

// Name of the field.
func (p *IgnoreValidator) Name() string {
	return p.name
}

// SetName of the field.
func (p *IgnoreValidator) SetName(name string) {
	p.name = name
}

// SetMessage set error message.
func (p *IgnoreValidator) SetMessage(message string) Validator {
	p.message = message

	return p
}

// Validate the value.
func (p *IgnoreValidator) Validate(_ any) {}

func (p *IgnoreValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct[int]{
		Rule: "ignore",
		Msg:  p.message,
	})
}

// HTMLCompatible for this validator.
func (p *IgnoreValidator) HTMLCompatible() bool {
	return true
}

// Ignore only for bind.
func Ignore() *IgnoreValidator {
	return &IgnoreValidator{}
}
