package kvalid

// IgnoreValidator only for bind.
type IgnoreValidator struct {
	name string
}

// Name of the field.
func (p *IgnoreValidator) Name() string {
	return p.name
}

// SetName of the field.
func (p *IgnoreValidator) SetName(name string) {
	p.name = name
}

// SetMessage set error message.
func (p *IgnoreValidator) SetMessage(_ string) Validator {
	return p
}

// Validate the value.
func (p *IgnoreValidator) Validate(_ any) {}

// HTMLCompatible for this validator.
func (p *IgnoreValidator) HTMLCompatible() bool {
	return false
}

// Ignore only for bind.
func Ignore() *IgnoreValidator {
	return &IgnoreValidator{}
}
