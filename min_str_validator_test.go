package kvalid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kvalid"
)

type strType struct {
	Field string
}

// nolint: errorlint, forcetypeassert
func TestMinStr(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	str := &strType{}
	rules := kvalid.New(str).Field(&str.Field, kvalid.MinStr(2))

	ass.Nil(rules.Validate(&strType{Field: "123"}), "Long enough")
	ass.Nil(rules.Validate(&strType{Field: "12"}), "Exactly hit min")
	ass.Len(rules.Validate(&strType{Field: "1"}).(kvalid.Errors), 1, "Too short")
	ass.Len(rules.Validate(&strType{Field: "£"}).(kvalid.Errors), 1, "Multi-byte characters too short")

	msg := "custom message"
	ass.Equal(msg, kvalid.New(str).Field(&str.Field, kvalid.MinStr(2).SetMessage(msg)).
		Validate(&strType{Field: "1"}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(msg, kvalid.New(str).Field(&str.Field, kvalid.MinStr(2)).
		Validate(&strType{Field: "1"}).(kvalid.Errors)[0].Error(), "Default error message")
	// optional
	rules = kvalid.New(str).Field(&str.Field, kvalid.MinStr(3).Optional())
	ass.Nil(rules.Validate(&strType{Field: ""}), "Invalid but zero")
	ass.Len(rules.Validate(&strType{Field: " "}).(kvalid.Errors), 1, "Invalid and not zero")
	ass.Nil(rules.Validate(&strType{Field: "123"}), "Valid and not zero")
}
