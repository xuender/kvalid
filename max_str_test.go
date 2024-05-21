package kvalid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kvalid"
)

// nolint: errorlint, forcetypeassert
func TestMaxStr(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	req := require.New(t)
	str := &strType{}
	rules := kvalid.New(str).Field(&str.Field, kvalid.MaxStr(2))
	ass.Len(rules.Validate(&strType{Field: "123"}).(kvalid.Errors), 1, "Short enough")
	req.NoError(rules.Validate(&strType{Field: "12"}), "Exactly hit max")
	req.NoError(rules.Validate(&strType{Field: "1"}), "Short enough")
	req.NoError(rules.Validate(&strType{Field: "世界"}), "Multi-byte characters are short enough")
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
