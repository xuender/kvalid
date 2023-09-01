package kvalid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kvalid"
	"github.com/xuender/kvalid/json"
)

// nolint: errorlint, forcetypeassert
func TestMaxStr(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	str := &strType{}
	rules := kvalid.New(str).Field(&str.Field, kvalid.MaxStr(2))
	ass.Len(rules.Validate(&strType{Field: "123"}).(kvalid.Errors), 1, "Short enough")
	ass.Nil(rules.Validate(&strType{Field: "12"}), "Exactly hit max")
	ass.Nil(rules.Validate(&strType{Field: "1"}), "Short enough")
	ass.Nil(rules.Validate(&strType{Field: "世界"}), "Multi-byte characters are short enough")
	// message
	ass.Equal(_msg, kvalid.New(str).Field(&str.Field, kvalid.MaxStr(0).SetMessage(_msg)).
		Validate(&strType{Field: "1"}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(str).Field(&str.Field, kvalid.MaxStr(0)).
		Validate(&strType{Field: "1"}).(kvalid.Errors)[0].Error(), "Default error message")
}

func TestMaxStrValidator_MarshalJSON(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	str := &strType{}
	rules := kvalid.New(str).Field(&str.Field, kvalid.MaxStr(2).SetMessage(_msg))
	data, _ := json.Marshal(rules)

	ass.Equal(`{"Field":[{"rule":"maxStr","max":2,"msg":"custom message"}]}`, string(data), "MarshalJSON")
}
