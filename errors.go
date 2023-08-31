package kvalid

import "errors"

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
func (v validationError) Error() string {
	return v.Message
}

// Field name.
func (v validationError) Field() string {
	return v.FieldName
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
func (v Errors) Error() string {
	list := make([]string, len(v))
	for i := range v {
		list[i] = v[i].Error()
	}

	return joinSentences(list)
}

// joinSentences converts a list of strings to a paragraph.
func joinSentences(list []string) string {
	length := len(list)
	if length == 0 {
		return ""
	}

	final := list[0]

	for i := 1; i < length; i++ {
		if i == length-1 {
			final = final + list[i] + "."
		} else {
			final = final + list[i] + ". "
		}
	}

	return final
}

func createError(name, custom, fallback string) Error {
	if custom != "" {
		return NewError(custom, name)
	}

	return NewError(fallback, name)
}
