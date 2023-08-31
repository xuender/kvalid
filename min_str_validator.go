package kvalid

import (
	"encoding/json"
	"fmt"
)

// MinStrValidator field must have minimum length.
type MinStrValidator struct {
	name     string
	message  string
	min      int64
	optional bool
}

// Name of the field.
func (c *MinStrValidator) Name() string {
	return c.name
}

// SetName of the field.
func (c *MinStrValidator) SetName(name string) {
	c.name = name
}

// SetMessage set error message.
func (c *MinStrValidator) SetMessage(msg string) Validator {
	c.message = msg

	return c
}

// Optional don't validate if the value is zero.
func (c *MinStrValidator) Optional() Validator {
	c.optional = true

	return c
}

// Validate the value.
func (c *MinStrValidator) Validate(value any) Error {
	str, _ := value.(string)
	if c.optional && str == "" {
		return nil
	}

	if len([]rune(str)) < int(c.min) {
		return createError(c.name, c.message, fmt.Sprintf("Please lengthen %s to %d characters or more", c.name, c.min))
	}

	return nil
}

// MarshalJSON for this validator.
func (c *MinStrValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct{
		Rule: "minStr",
		Min:  c.min,
		Msg:  c.message,
	})
}

// HTMLCompatible for this validator.
func (c *MinStrValidator) HTMLCompatible() bool {
	return true
}

// MinStr field must have minimum length.
func MinStr(min int64) *MinStrValidator {
	return &MinStrValidator{
		min: min,
	}
}
