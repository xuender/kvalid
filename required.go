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
func (p *RequiredValidator) Name() string {
	return p.name
}

// SetName of the field.
func (p *RequiredValidator) SetName(name string) {
	p.name = name
}

// SetMessage set error message.
func (p *RequiredValidator) SetMessage(msg string) Validator {
	p.message = msg

	return p
}

// Validate the value.
func (p *RequiredValidator) Validate(value any) Error {
	v := reflect.ValueOf(value)
	kind := v.Kind()

	if !v.IsValid() ||
		v.IsZero() ||
		((kind == reflect.Ptr || kind == reflect.Interface) && v.Elem().IsZero()) {
		return createError(p.name, p.message, fmt.Sprintf("Please enter the %v", p.name))
	}

	return nil
}

// MarshalJSON for this validator.
func (p *RequiredValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct[int]{
		Rule: "required",
		Msg:  p.message,
	})
}

// HTMLCompatible for this validator.
func (p *RequiredValidator) HTMLCompatible() bool {
	return true
}

// Required fields must not be zero.
func Required() *RequiredValidator {
	return &RequiredValidator{}
}
