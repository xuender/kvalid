package kvalid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kvalid"
)

type funcTest struct {
	Field string
}

// nolint: errorlint, forcetypeassert
func TestFieldFunc(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	req := require.New(t)
	checker := func(fieldName string, value string) kvalid.Error {
		if value == "invalid" {
			return kvalid.NewError("Invalid field", fieldName)
		}

		return nil
	}
	data := &funcTest{}
	rules := kvalid.New(data).Field(&data.Field, kvalid.FieldFunc(checker).SetMessage(_msg))
	req.NoError(rules.Validate(&funcTest{Field: "valid"}), "Valid")

	errs := rules.Validate(&funcTest{Field: "invalid"}).(kvalid.Errors)
	ass.Len(errs, 1, "Invalid")
	ass.Equal("Field", errs[0].Field())
}
