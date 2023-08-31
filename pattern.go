package kvalid

import (
	"fmt"
	"regexp"

	"github.com/xuender/kvalid/json"
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
func (p *PatternValidator) Validate(value string) Error {
	if p.optional && value == "" {
		return nil
	}

	if p.re.MatchString(value) {
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

// Pattern field must match regexp.
func Pattern(pattern string) *PatternValidator {
	return &PatternValidator{
		re: regexp.MustCompile(pattern),
	}
}
