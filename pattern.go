package kvalid

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// PatternValidator field must match regexp.
type PatternValidator struct {
	name     string
	message  string
	re       *regexp.Regexp
	optional bool
}

// Name of the field.
func (p *PatternValidator) Name() string {
	return p.name
}

// SetName of the field.
func (p *PatternValidator) SetName(name string) {
	p.name = name
}

// SetMessage set error message.
func (p *PatternValidator) SetMessage(msg string) Validator {
	p.message = msg

	return p
}

// Optional don't validate if the value is zero.
func (p *PatternValidator) Optional() Validator {
	p.optional = true

	return p
}

// Validate the value.
func (p *PatternValidator) Validate(value any) Error {
	val := toString(value)
	if p.optional && val == "" {
		return nil
	}

	if p.re.MatchString(val) {
		return nil
	}

	return createError(p.name, p.message, fmt.Sprintf("Please correct %s into a valid format", p.name))
}

// MarshalJSON for this validator.
func (p *PatternValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct[int]{
		Rule:    "pattern",
		Pattern: p.re.String(),
		Msg:     p.message,
	})
}

// HTMLCompatible for this validator.
func (p *PatternValidator) HTMLCompatible() bool {
	return true
}

func toString(pattern any) string {
	switch val := pattern.(type) {
	case string:
		return val
	case fmt.Stringer:
		return val.String()
	default:
		return fmt.Sprintf("%v", val)
	}
}

// Pattern field must match regexp.
func Pattern(pattern any) *PatternValidator {
	return &PatternValidator{
		re: regexp.MustCompile(toString(pattern)),
	}
}
