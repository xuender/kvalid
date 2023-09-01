package kvalid

import (
	"errors"
	"strings"
)

var (
	ErrStructNotPointer = errors.New("struct is not pointer")
	ErrIsNil            = errors.New("struct is nil")
	ErrFieldNotPointer  = errors.New("field is not pointer")
	ErrFindField        = errors.New("can't find field")
	ErrMissValidate     = errors.New("miss Validate func")
)

// Error when a rule is broken.
type Error interface {
	Error() string
	Field() string
}

// validationError implements Error interface.
type validationError struct {
	Message   string `json:"message"`
	FieldName string `json:"field"`
}

// Error message.
func (p validationError) Error() string {
	return p.Message
}

// Field name.
func (p validationError) Field() string {
	return p.FieldName
}

// NewError creates new validation error.
func NewError(message, fieldName string) Error {
	return &validationError{
		FieldName: fieldName,
		Message:   message,
	}
}

// Errors is a list of Error.
type Errors []Error

// Error will combine all errors into a list of sentences.
func (p Errors) Error() string {
	list := make([]string, len(p))
	for i := range p {
		list[i] = p[i].Error()
	}

	return joinSentences(list)
}

// joinSentences converts a list of strings to a paragraph.
func joinSentences(list []string) string {
	return strings.Join(list, ". ") + "."
}

func createError(name, custom, fallback string) Error {
	if custom != "" {
		return NewError(custom, name)
	}

	return NewError(fallback, name)
}
