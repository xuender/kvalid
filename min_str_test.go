package kvalid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kvalid"
)

const _msg = "custom message"

type strType struct {
	Field string
}

// nolint: errorlint, forcetypeassert
func TestMinStr(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	req := require.New(t)
	str := &strType{}
	rules := kvalid.New(str).Field(&str.Field, kvalid.MinStr(2))

	req.NoError(rules.Validate(&strType{Field: "123"}), "Long enough")
	req.NoError(rules.Validate(&strType{Field: "12"}), "Exactly hit min")
	ass.Len(rules.Validate(&strType{Field: "1"}).(kvalid.Errors), 1, "Too short")
	ass.Len(rules.Validate(&strType{Field: "£"}).(kvalid.Errors), 1, "Multi-byte characters too short")
	// message
	ass.Equal(_msg, kvalid.New(str).Field(&str.Field, kvalid.MinStr(2).SetMessage(_msg)).
		Validate(&strType{Field: "1"}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(str).Field(&str.Field, kvalid.MinStr(2)).
		Validate(&strType{Field: "1"}).(kvalid.Errors)[0].Error(), "Default error message")
	// optional
	rules = kvalid.New(str).Field(&str.Field, kvalid.MinStr(3).Optional())
	req.NoError(rules.Validate(&strType{Field: ""}), "Invalid but zero")
	ass.Len(rules.Validate(&strType{Field: " "}).(kvalid.Errors), 1, "Invalid and not zero")
	req.NoError(rules.Validate(&strType{Field: "123"}), "Valid and not zero")
}
