package kvalid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kvalid"
	"github.com/xuender/kvalid/json"
)

type structSubject struct {
	Less int
	More int
}

type compareValidator struct {
	name    string
	message string
}

func (c *compareValidator) Name() string {
	return c.name
}

func (c *compareValidator) SetName(name string) {
	c.name = name
}

func (c *compareValidator) SetMessage(msg string) kvalid.Validator {
	c.message = msg

	return c
}

func (c *compareValidator) HTMLCompatible() bool {
	return true
}

func (c *compareValidator) Validate(value structSubject) kvalid.Error {
	if value.Less > value.More {
		return kvalid.NewError("comparison failed", "")
	}

	return nil
}

// nolint: errorlint, forcetypeassert
func TestRules_Struct(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	sub := &structSubject{}
	rules := kvalid.New(sub).Struct(&compareValidator{})
	ass.Nil(rules.Validate(structSubject{Less: 1, More: 2}), "Valid")
	ass.Len(rules.Validate(structSubject{Less: 2, More: 1}).(kvalid.Errors), 1, "Invalid")
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	str := &strType{}
	rules := kvalid.New(str).Field(&str.Field, kvalid.MinStr(2))
	data, _ := json.Marshal(rules)

	ass.Equal(string(data), `{"Field":[{"rule":"minStr","min":2}]}`)

	rules = kvalid.New(str).Field(&str.Field, kvalid.MinStr(2).SetMessage("length minimum 2"))
	data, _ = json.Marshal(rules)

	ass.Equal(string(data), `{"Field":[{"rule":"minStr","min":2,"msg":"length minimum 2"}]}`)
}
