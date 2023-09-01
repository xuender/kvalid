package kvalid

// StructFuncValidator validate struct with custom function.
type StructFuncValidator[T any] struct {
	name    string
	checker func(T) Error
}

// Name of the field.
func (c *StructFuncValidator[T]) Name() string {
	return c.name
}

// SetName of the field.
func (c *StructFuncValidator[T]) SetName(name string) {
	c.name = name
}

// SetMessage set error message.
func (c *StructFuncValidator[T]) SetMessage(_ string) Validator {
	return c
}

// Validate the value.
func (c *StructFuncValidator[T]) Validate(value T) Error {
	return c.checker(value)
}

// HTMLCompatible for this validator.
func (c *StructFuncValidator[T]) HTMLCompatible() bool {
	return false
}

// StructFunc validate struct with custom function.
func StructFunc[T any](checker func(T) Error) Validator {
	return &StructFuncValidator[T]{
		checker: checker,
	}
}
