package kvalid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	req := require.New(t)
	data := &intType{}
	rules := kvalid.New(data).Field(&data.Field, kvalid.MinNum(0))
	req.NoError(rules.Validate(&intType{Field: 1}), "Big enough")
	req.NoError(rules.Validate(&intType{Field: 0}), "Exactly hit min")
	ass.Len(rules.Validate(&intType{Field: -1}).(kvalid.Errors), 1, "Too low")
	// message
	ass.Equal(_msg, kvalid.New(data).Field(&data.Field, kvalid.MinNum(0).SetMessage(_msg)).
		Validate(&intType{Field: -1}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(data).Field(&data.Field, kvalid.MinNum(0)).
		Validate(&intType{Field: -1}).(kvalid.Errors)[0].Error(), "Default error message")
	// optional
	rules = kvalid.New(data).Field(&data.Field, kvalid.MinNum(5).Optional())
	req.NoError(rules.Validate(&intType{Field: 0}), "Invalid but zero")
	ass.Len(rules.Validate(&intType{Field: 1}).(kvalid.Errors), 1, "Invalid and not zero")
	req.NoError(rules.Validate(&intType{Field: 5}), "Valid and not zero")
	// null
	rules = kvalid.New(data).Field(&data.Null, kvalid.MinNullInt(0))
	req.NoError(rules.Validate(&intType{Null: null.IntFrom(1)}), "Big enough")
	req.NoError(rules.Validate(&intType{Null: null.IntFrom(0)}), "Exactly hit min")
	ass.Len(rules.Validate(&intType{Null: null.IntFrom(-1)}).(kvalid.Errors), 1, "Too low")
	ass.Equal(_msg, kvalid.New(data).Field(&data.Null, kvalid.MinNullInt(0).SetMessage(_msg)).
		Validate(&intType{Null: null.IntFrom(-1)}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(data).Field(&data.Null, kvalid.MinNullInt(0)).
		Validate(&intType{Null: null.IntFrom(-1)}).(kvalid.Errors)[0].Error(), "Default error message")
	// optional
	rules = kvalid.New(data).Field(&data.Null, kvalid.MinNullInt(5).Optional())
	req.NoError(rules.Validate(&intType{Null: null.IntFrom(0)}), "Invalid but zero")
	ass.Len(rules.Validate(&intType{Null: null.IntFrom(1)}).(kvalid.Errors), 1, "Invalid and not zero")
	req.NoError(rules.Validate(&intType{Null: null.IntFrom(5)}), "Valid and not zero")
}

func TestMinNum_MarshalJSON(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	num := &intType{}
	rules := kvalid.New(num).
		Field(&num.Null, kvalid.MinNullInt(0)).
		Field(&num.Null, kvalid.MaxNullInt(100))
	data, _ := json.Marshal(rules)

	ass.Equal(`{"Null":[{"rule":"minNum"},{"rule":"maxNum","max":100}]}`, string(data), "MarshalJSON")
}
