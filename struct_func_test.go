package kvalid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kvalid"
)

// nolint: errorlint, forcetypeassert
func TestStructFunc(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	checker := func(value string) kvalid.Error {
		if value == "invalid" {
			return kvalid.NewError("Invalid field", "")
		}

		return nil
	}
	data := &funcTest{}
	rules := kvalid.New(data).Field(&data.Field, kvalid.StructFunc(checker))
	ass.Nil(rules.Validate(&funcTest{Field: "valid"}), "Valid")
	ass.Len(rules.Validate(&funcTest{Field: "invalid"}).(kvalid.Errors), 1, "Invalid")
}
