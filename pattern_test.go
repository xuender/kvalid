package kvalid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kvalid"
)

type patternType struct {
	Field string
}

// nolint: errorlint, forcetypeassert
func TestPattern(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	req := require.New(t)
	data := &patternType{}
	rules := kvalid.New(data).Field(&data.Field, kvalid.Pattern(`\d{2}`))
	req.NoError(rules.Validate(&patternType{Field: "00"}), "Exact match")
	req.NoError(rules.Validate(&patternType{Field: "1234"}), "Submatch also works")
	ass.Len(rules.Validate(&patternType{Field: "wrong"}).(kvalid.Errors), 1, "Pattern is wrong")
	// message
	ass.Equal(_msg, kvalid.New(data).Field(&data.Field, kvalid.Pattern(`\d{2}`).SetMessage(_msg)).
		Validate(&patternType{Field: "message"}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(data).Field(&data.Field, kvalid.Pattern(`\d{2}`)).
		Validate(&patternType{Field: "message"}).(kvalid.Errors)[0].Error(), "Default error message")
	// optional
	rules = kvalid.New(data).Field(&data.Field, kvalid.Pattern(`\w{3,}`).Optional())
	req.NoError(rules.Validate(&patternType{Field: ""}), "Invalid but zero")
	ass.Len(rules.Validate(&patternType{Field: " "}).(kvalid.Errors), 1, "Invalid and not zero")
	req.NoError(rules.Validate(&patternType{Field: "123"}), "Valid and not zero")
}

func TestPattern_int(t *testing.T) {
	t.Parallel()

	req := require.New(t)
	data := &intType{}
	rules := kvalid.New(data).Field(&data.Field, kvalid.Pattern(`\d{2}`))
	req.NoError(rules.Validate(&intType{Field: 12}), "Exact match")
	req.Error(rules.Validate(&intType{Field: 3}), &kvalid.Errors{}, "Pattern is wrong")
}

func TestPatternValidator_MarshalJSON(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	pattern := &patternType{}
	rules := kvalid.New(pattern).Field(&pattern.Field, kvalid.Pattern(`\d{2}`).SetMessage(_msg))
	data, _ := json.Marshal(rules)

	ass.Equal(`{"Field":[{"rule":"pattern","pattern":"\\d{2}","msg":"custom message"}]}`, string(data), "MarshalJSON")
}
