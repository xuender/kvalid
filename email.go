package kvalid

import (
	"encoding/json"
	"regexp"
)

// EmailValidator field must be a valid email address.
type EmailValidator struct {
	PatternValidator
}

var _emailRegex = regexp.MustCompile(`^[\w.!#$%&'*+/=?^_{|}~-]+@\w(?:[\w-]{0,61}\w)?(?:\.\w(?:[\w-]{0,61}\w)?)*$`)

// Email field must be a valid email address.
func Email() *EmailValidator {
	return &EmailValidator{
		PatternValidator{
			re:      _emailRegex,
			message: "Please use a valid email address",
		},
	}
}

// MarshalJSON for this validator.
func (p *EmailValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct[int]{
		Rule: "email",
		Msg:  p.message,
	})
}

// IsEmail returns true if the string is an email.
func IsEmail(email string) bool {
	return _emailRegex.MatchString(email)
}
