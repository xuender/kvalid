package kvalid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kvalid"
	"gopkg.in/guregu/null.v3"
)

type intType struct {
	Field int
	Null  null.Int
}

// nolint: errorlint, forcetypeassert
func TestMinNum(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	data := &intType{}
	rules := kvalid.New(data).Field(&data.Field, kvalid.MinNum(0))
	ass.Nil(rules.Validate(&intType{Field: 1}), "Big enough")
	ass.Nil(rules.Validate(&intType{Field: 0}), "Exactly hit min")
	ass.Len(rules.Validate(&intType{Field: -1}).(kvalid.Errors), 1, "Too low")
	// message
	ass.Equal(_msg, kvalid.New(data).Field(&data.Field, kvalid.MinNum(0).SetMessage(_msg)).
		Validate(&intType{Field: -1}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(data).Field(&data.Field, kvalid.MinNum(0)).
		Validate(&intType{Field: -1}).(kvalid.Errors)[0].Error(), "Default error message")
	// optional
	rules = kvalid.New(data).Field(&data.Field, kvalid.MinNum(5).Optional())
	ass.Nil(rules.Validate(&intType{Field: 0}), "Invalid but zero")
	ass.Len(rules.Validate(&intType{Field: 1}).(kvalid.Errors), 1, "Invalid and not zero")
	ass.Nil(rules.Validate(&intType{Field: 5}), "Valid and not zero")
	// null
	rules = kvalid.New(data).Field(&data.Null, kvalid.MinNullInt(0))
	ass.Nil(rules.Validate(&intType{Null: null.IntFrom(1)}), "Big enough")
	ass.Nil(rules.Validate(&intType{Null: null.IntFrom(0)}), "Exactly hit min")
	ass.Len(rules.Validate(&intType{Null: null.IntFrom(-1)}).(kvalid.Errors), 1, "Too low")
	ass.Equal(_msg, kvalid.New(data).Field(&data.Null, kvalid.MinNullInt(0).SetMessage(_msg)).
		Validate(&intType{Null: null.IntFrom(-1)}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(data).Field(&data.Null, kvalid.MinNullInt(0)).
		Validate(&intType{Null: null.IntFrom(-1)}).(kvalid.Errors)[0].Error(), "Default error message")
	// optional
	rules = kvalid.New(data).Field(&data.Null, kvalid.MinNullInt(5).Optional())
	ass.Nil(rules.Validate(&intType{Null: null.IntFrom(0)}), "Invalid but zero")
	ass.Len(rules.Validate(&intType{Null: null.IntFrom(1)}).(kvalid.Errors), 1, "Invalid and not zero")
	ass.Nil(rules.Validate(&intType{Null: null.IntFrom(5)}), "Valid and not zero")
}
