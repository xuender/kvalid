package kvalid

// FieldFuncValidator for validating with custom function.
type FieldFuncValidator[T any] struct {
	name    string
	checker func(string, T) Error
}

// Name of the field.
func (p *FieldFuncValidator[T]) Name() string {
	return p.name
}

// SetName of the field.
func (p *FieldFuncValidator[T]) SetName(name string) {
	p.name = name
}

// SetMessage set error message.
func (p *FieldFuncValidator[T]) SetMessage(_ string) Validator {
	return p
}

// Validate the value.
func (p *FieldFuncValidator[T]) Validate(value T) Error {
	return p.checker(p.name, value)
}

// HTMLCompatible for this validator.
func (p *FieldFuncValidator[T]) HTMLCompatible() bool {
	return false
}

// FieldFunc for validating with custom function.
func FieldFunc[T any](checker func(string, T) Error) Validator {
	return &FieldFuncValidator[T]{
		checker: checker,
	}
}
