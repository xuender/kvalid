package kvalid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kvalid"
)

type patternType struct {
	Field string
}

// nolint: errorlint, forcetypeassert
func TestPattern(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	data := &patternType{}
	rules := kvalid.New(data).Field(&data.Field, kvalid.Pattern(`\d{2}`))
	ass.Nil(rules.Validate(&patternType{Field: "00"}), "Exact match")
	ass.Nil(rules.Validate(&patternType{Field: "1234"}), "Submatch also works")
	ass.Len(rules.Validate(&patternType{Field: "wrong"}).(kvalid.Errors), 1, "Pattern is wrong")
	// message
	ass.Equal(_msg, kvalid.New(data).Field(&data.Field, kvalid.Pattern(`\d{2}`).SetMessage(_msg)).
		Validate(&patternType{Field: "message"}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(data).Field(&data.Field, kvalid.Pattern(`\d{2}`)).
		Validate(&patternType{Field: "message"}).(kvalid.Errors)[0].Error(), "Default error message")
	// optional
	rules = kvalid.New(data).Field(&data.Field, kvalid.Pattern(`\w{3,}`).Optional())
	ass.Nil(rules.Validate(&patternType{Field: ""}), "Invalid but zero")
	ass.Len(rules.Validate(&patternType{Field: " "}).(kvalid.Errors), 1, "Invalid and not zero")
	ass.Nil(rules.Validate(&patternType{Field: "123"}), "Valid and not zero")
}
