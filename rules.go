package kvalid

import (
	"reflect"
	"strings"
)

// Rules for creating a chain of rules for validating a struct.
type Rules[T any] struct {
	validators []Validator
	structPtr  T
}

// New rule chain.
func New[T any](structPtr T) *Rules[T] {
	return &Rules[T]{
		structPtr:  structPtr,
		validators: make([]Validator, 0),
	}
}

// Field adds validators for a field.
func (r *Rules[T]) Field(fieldPtr any, validators ...Validator) *Rules[T] {
	for _, validator := range validators {
		validator.SetName(r.getFieldName(fieldPtr))
		r.validators = append(r.validators, validator)
	}

	return r
}

// Struct adds validators for the struct.
func (r *Rules[T]) Struct(validators ...Validator) *Rules[T] {
	r.validators = append(r.validators, validators...)

	return r
}

// Validate a struct and return Errors.
func (r *Rules[T]) Validate(subject T) error {
	errs := make(Errors, 0)
	vmap := r.structToMap(subject)

	for _, validator := range r.validators {
		var err Error
		if validator.Name() == "" {
			err = validator.Validate(subject)
		} else {
			err = validator.Validate(vmap[validator.Name()])
		}

		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

// OnlyFor filters the validators to match only the fields.
func (r *Rules[T]) OnlyFor(name string) *Rules[T] {
	validators := r.validators
	r.validators = make([]Validator, 0)

	for _, val := range validators {
		if val.Name() == name {
			r.validators = append(r.validators, val)
		}
	}

	return r
}

// Validators for this chain.
func (r *Rules[T]) Validators() []Validator {
	return r.validators
}

// structToMap converts struct to map and uses the json name if available.
func (r *Rules[T]) structToMap(structPtr T) map[string]any {
	vmap := make(map[string]any)
	structValue := reflect.ValueOf(structPtr)

	if structValue.Kind() == reflect.Ptr {
		structValue = structValue.Elem()
	}

	for index := structValue.NumField() - 1; index >= 0; index-- {
		sf := structValue.Type().Field(index)
		name := sf.Tag.Get("json")

		if name == "" {
			name = sf.Name
		}

		if index := strings.Index(name, ","); index > 0 {
			name = name[:index]
		}

		f := structValue.Field(index)
		if f.CanInterface() {
			vmap[name] = f.Interface()
		}
	}

	return vmap
}

func (r *Rules[T]) getFieldName(fieldPtr any) string {
	value := reflect.ValueOf(r.structPtr)
	if value.Kind() != reflect.Ptr || !value.IsNil() && value.Elem().Kind() != reflect.Struct {
		panic(ErrStructNotPointer)
	}

	if value.IsNil() {
		panic(ErrIsNil)
	}

	value = value.Elem()

	fval := reflect.ValueOf(fieldPtr)
	if fval.Kind() != reflect.Ptr {
		panic(ErrFieldNotPointer)
	}

	fsf := findStructField(value, fval)
	if fsf == nil {
		panic(ErrFindField)
	}

	tag := fsf.Tag.Get("json")
	if tag == "" {
		tag = fsf.Name
	}

	if index := strings.Index(tag, ","); index > 0 {
		tag = tag[:index]
	}

	return tag
}

// findStructField looks for a field in the given struct.
// The field being looked for should be a pointer to the actual struct field.
// If found, the field info will be returned. Otherwise, nil will be returned.
func findStructField(structValue reflect.Value, fieldValue reflect.Value) *reflect.StructField {
	ptr := fieldValue.Pointer()

	for i := structValue.NumField() - 1; i >= 0; i-- {
		sf := structValue.Type().Field(i)
		if ptr == structValue.Field(i).UnsafeAddr() {
			// do additional type comparison because it's possible that the address of
			// an embedded struct is the same as the first field of the embedded struct
			if sf.Type == fieldValue.Elem().Type() {
				return &sf
			}
		}
	}

	return nil
}
