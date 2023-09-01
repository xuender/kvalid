package kvalid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kvalid"
	"gopkg.in/guregu/null.v3"
)

// nolint: errorlint, forcetypeassert
func TestMaxNum(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	data := &intType{}
	rules := kvalid.New(data).Field(&data.Field, kvalid.MaxNum(0))
	ass.Len(rules.Validate(&intType{Field: 1}).(kvalid.Errors), 1, "Too big")
	ass.Nil(rules.Validate(&intType{Field: 0}), "Exactly hit max")
	ass.Nil(rules.Validate(&intType{Field: -1}), "Low engouh")
	// message
	ass.Equal(_msg, kvalid.New(data).Field(&data.Field, kvalid.MaxNum(0).SetMessage(_msg)).
		Validate(&intType{Field: 1}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(data).Field(&data.Field, kvalid.MaxNum(0)).
		Validate(&intType{Field: 1}).(kvalid.Errors)[0].Error(), "Default error message")
	// null
	rules = kvalid.New(data).Field(&data.Null, kvalid.MaxNullInt(0))
	ass.Len(rules.Validate(&intType{Null: null.IntFrom(1)}).(kvalid.Errors), 1, "Too big")
	ass.Nil(rules.Validate(&intType{Null: null.IntFrom(0)}), "Exactly hit max")
	ass.Nil(rules.Validate(&intType{Null: null.IntFrom(-1)}), "Low engouh")
	ass.Equal(_msg, kvalid.New(data).Field(&data.Null, kvalid.MaxNullInt(0).SetMessage(_msg)).
		Validate(&intType{Null: null.IntFrom(1)}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(data).Field(&data.Null, kvalid.MaxNullInt(0)).
		Validate(&intType{Null: null.IntFrom(1)}).(kvalid.Errors)[0].Error(), "Default error message")
}
