package kvalid

import (
	"reflect"
	"strings"

	"github.com/xuender/kvalid/json"
)

// Rules for creating a chain of rules for validating a struct.
type Rules struct {
	validators []Validator
	value      reflect.Value
}

// New rule chain.
func New(structPtr any) *Rules {
	value := reflect.ValueOf(structPtr)
	if value.Kind() != reflect.Ptr || !value.IsNil() && value.Elem().Kind() != reflect.Struct {
		panic(ErrStructNotPointer)
	}

	if value.IsNil() {
		panic(ErrIsNil)
	}

	return &Rules{
		value:      value.Elem(),
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
	var (
		errs = make(Errors, 0)
		vmap = p.structToMap(subject)
	)

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

func (p *Rules) Bind(source, target any) error {
	if err := p.Validate(source); err != nil {
		return err
	}

	var (
		sVal = reflect.ValueOf(source)
		tVal = reflect.ValueOf(target)
	)

	if sVal.Kind() == reflect.Ptr {
		sVal = sVal.Elem()
	}

	if tVal.Kind() == reflect.Ptr {
		tVal = tVal.Elem()
	}

	for _, index := range p.getFieldIndexes(source) {
		tVal.Field(index).Set(sVal.Field(index))
	}

	return nil
}

func (p *Rules) getFieldIndexes(elem any) []int {
	var (
		typ     = reflect.TypeOf(elem)
		names   = map[string]struct{}{}
		indexes = []int{}
	)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for _, vali := range p.validators {
		names[vali.Name()] = struct{}{}
	}

	for index := 0; index < typ.NumField(); index++ {
		field := typ.Field(index)
		if _, has := names[fieldName(&field)]; has {
			indexes = append(indexes, index)
		}
	}

	return indexes
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
	var (
		vmap  = make(map[string]any)
		value = reflect.ValueOf(structPtr)
	)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	for index := value.NumField() - 1; index >= 0; index-- {
		typeField := value.Type().Field(index)
		name := fieldName(&typeField)

		field := value.Field(index)
		if field.CanInterface() {
			vmap[name] = field.Interface()
		}
	}

	return vmap
}

func (p *Rules) getFieldName(fieldPtr any) string {
	value := reflect.ValueOf(fieldPtr)
	if value.Kind() != reflect.Ptr {
		panic(ErrFieldNotPointer)
	}

	field := p.findStructField(value)
	if field == nil {
		panic(ErrFindField)
	}

	return fieldName(field)
}

func fieldName(fsf *reflect.StructField) string {
	name := fsf.Tag.Get("json")
	if name == "" {
		name = fsf.Name
	}

	if index := strings.Index(name, ","); index > 0 {
		name = name[:index]
	}

	return name
}

// findStructField looks for a field in the given struct.
// The field being looked for should be a pointer to the actual struct field.
// If found, the field info will be returned. Otherwise, nil will be returned.
func (p *Rules) findStructField(fieldValue reflect.Value) *reflect.StructField {
	ptr := fieldValue.Pointer()

	for index := p.value.NumField() - 1; index >= 0; index-- {
		field := p.value.Type().Field(index)
		if ptr == p.value.Field(index).UnsafeAddr() {
			// do additional type comparison because it's possible that the address of
			// an embedded struct is the same as the first field of the embedded struct
			if field.Type == fieldValue.Elem().Type() {
				return &field
			}
		}
	}

	return nil
}
