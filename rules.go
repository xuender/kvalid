package kvalid

import (
	"reflect"
	"strings"

	"github.com/xuender/kvalid/json"
)

// Rules for creating a chain of rules for validating a struct.
type Rules struct {
	validators []Validator
	structPtr  any
}

// New rule chain.
func New(structPtr any) *Rules {
	return &Rules{
		structPtr:  structPtr,
		validators: make([]Validator, 0),
	}
}

// Field adds validators for a field.
func (p *Rules) Field(fieldPtr any, validators ...Validator) *Rules {
	for _, validator := range validators {
		validator.SetName(p.getFieldName(fieldPtr))
		p.validators = append(p.validators, validator)
	}

	return p
}

// Struct adds validators for the struct.
func (p *Rules) Struct(validators ...Validator) *Rules {
	p.validators = append(p.validators, validators...)

	return p
}

// Validate a struct and return Errors.
func (p *Rules) Validate(subject any) error {
	errs := make(Errors, 0)
	vmap := p.structToMap(subject)

	for _, validator := range p.validators {
		var err Error
		if validator.Name() == "" {
			err = call(validator, subject)
		} else {
			err = call(validator, vmap[validator.Name()])
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

func call(valid Validator, value any) Error {
	method := reflect.ValueOf(valid).MethodByName("Validate")

	if !method.IsValid() {
		panic(ErrMissValidate)
	}

	var (
		val = reflect.ValueOf(value)
		ret = method.Call([]reflect.Value{val})
	)

	if length := len(ret); length > 0 {
		if ret[length-1].IsNil() {
			return nil
		}

		if err, ok := ret[length-1].Interface().(Error); ok {
			return err
		}

		panic(ErrMissValidate)
	}

	return nil
}

// OnlyFor filters the validators to match only the fields.
func (p *Rules) OnlyFor(name string) *Rules {
	validators := p.validators
	p.validators = make([]Validator, 0)

	for _, val := range validators {
		if val.Name() == name {
			p.validators = append(p.validators, val)
		}
	}

	return p
}

// Validators for this chain.
func (p *Rules) Validators() []Validator {
	return p.validators
}

func (p *Rules) MarshalJSON() ([]byte, error) {
	htmls := map[string][]Validator{}

	for _, val := range p.validators {
		if !val.HTMLCompatible() {
			continue
		}

		rules, has := htmls[val.Name()]
		if !has {
			rules = []Validator{}
		}

		rules = append(rules, val)
		htmls[val.Name()] = rules
	}

	return json.Marshal(htmls)
}

// structToMap converts struct to map and uses the json name if available.
func (p *Rules) structToMap(structPtr any) map[string]any {
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

func (p *Rules) getFieldName(fieldPtr any) string {
	value := reflect.ValueOf(p.structPtr)
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
