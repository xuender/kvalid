package kvalid_test

import (
	"os"
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

func (c *compareValidator) Validate(value *structSubject) kvalid.Error {
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
	ass.Nil(rules.Validate(&structSubject{Less: 1, More: 2}), "Valid")
	ass.Len(rules.Validate(&structSubject{Less: 2, More: 1}).(kvalid.Errors), 1, "Invalid")
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	str := &strType{}
	rules := kvalid.New(str).
		Field(&str.Field, kvalid.MinStr(2)).
		Field(&str.Field, kvalid.FieldFunc(func(field, value string) kvalid.Error {
			return nil
		})).
		Field(&str.Field, kvalid.StructFunc(func(value string) kvalid.Error {
			return nil
		}))
	data, _ := json.Marshal(rules)

	ass.Equal(string(data), `{"Field":[{"rule":"minStr","min":2}]}`)

	rules = kvalid.New(str).Field(&str.Field, kvalid.MinStr(2).SetMessage("length minimum 2"))
	data, _ = json.Marshal(rules)

	ass.Equal(string(data), `{"Field":[{"rule":"minStr","min":2,"msg":"length minimum 2"}]}`)
}

func TestRules_OnlyFor(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	req := &requiredType{}
	rules := kvalid.New(req).
		Field(&req.String, kvalid.MinStr(2)).
		Field(&req.Time, kvalid.Required())

	ass.Len(rules.Validators(), 2)
	ass.Len(rules.OnlyFor("string").Validators(), 1)
}

func TestNew(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)

	ass.Panics(func() {
		kvalid.New(Book{})
	})

	ass.Panics(func() {
		var book *Book
		kvalid.New(book)
	})
}

func TestRules_Field(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	book := &Book{}

	ass.Panics(func() {
		kvalid.New(book).Field(book, kvalid.Required())
	})

	ass.Panics(func() {
		kvalid.New(book).Field(book.Title, kvalid.Required())
	})
}

type (
	bedValidator  struct{ name string }
	bed2Validator struct{ bedValidator }
	bed3Validator struct{ bedValidator }
)

func (p *bedValidator) SetName(name string)                { p.name = name }
func (p *bedValidator) Name() string                       { return p.name }
func (p *bedValidator) HTMLCompatible() bool               { return false }
func (p *bedValidator) SetMessage(string) kvalid.Validator { return p }

func (p *bed2Validator) Validate(_ string) error { return os.ErrClosed }
func (p *bed3Validator) Validate(_ string)       {}

func TestRules_Validate(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	book := &Book{}

	ass.Panics(func() {
		_ = kvalid.New(book).Field(&book.Title, &bedValidator{}).Validate(book)
	})
	ass.Panics(func() {
		_ = kvalid.New(book).Field(&book.Title, &bed2Validator{}).Validate(book)
	})

	_ = kvalid.New(book).Field(&book.Title, &bed3Validator{}).Validate(book)
}
