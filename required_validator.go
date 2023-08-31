package kvalid

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// RequiredValidator field must not be zero.
type RequiredValidator struct {
	name    string
	message string
}

// Name of the field.
func (c *RequiredValidator) Name() string {
	return c.name
}

// SetName of the field.
func (c *RequiredValidator) SetName(name string) {
	c.name = name
}

// SetMessage set error message.
func (c *RequiredValidator) SetMessage(msg string) Validator {
	c.message = msg

	return c
}

// Validate the value.
func (c *RequiredValidator) Validate(value any) Error {
	v := reflect.ValueOf(value)
	kind := v.Kind()

	if !v.IsValid() ||
		v.IsZero() ||
		((kind == reflect.Ptr || kind == reflect.Interface) && v.Elem().IsZero()) {
		return createError(c.name, c.message, fmt.Sprintf("Please enter the %v", c.name))
	}

	return nil
}

// MarshalJSON for this validator.
func (c *RequiredValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct{
		Rule: "required",
		Msg:  c.message,
	})
}

// HTMLCompatible for this validator.
func (c *RequiredValidator) HTMLCompatible() bool {
	return true
}

// Required fields must not be zero.
func Required() *RequiredValidator {
	return &RequiredValidator{}
}
