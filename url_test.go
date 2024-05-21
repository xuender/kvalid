package kvalid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kvalid"
)

// nolint: errorlint, forcetypeassert
func TestURL(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	req := require.New(t)
	urlData := &strType{}
	rules := kvalid.New(urlData).Field(&urlData.Field, kvalid.URL())

	req.NoError(rules.Validate(&strType{Field: "http://web.com"}), "Valid url")
	ass.Len(rules.Validate(&strType{Field: "fake url"}).(kvalid.Errors), 1, "Invalid url")
	ass.Equal(_msg, kvalid.New(urlData).Field(&urlData.Field, kvalid.URL().SetMessage(_msg)).
		Validate(&strType{Field: "invalid"}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(urlData).Field(&urlData.Field, kvalid.URL()).
		Validate(&strType{Field: "invalid"}).(kvalid.Errors)[0].Error(), "Default error message")
	// optional
	rules = kvalid.New(urlData).Field(&urlData.Field, kvalid.URL().Optional())

	req.NoError(rules.Validate(&strType{Field: ""}), "Invalid but zero")
	ass.Len(rules.Validate(&strType{Field: "fake url"}).(kvalid.Errors), 1, "Invalid and not zero")
	req.NoError(rules.Validate(&strType{Field: "http://web.com"}), "Valid and not zero")
}

func TestURLValidator_MarshalJSON(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	url := kvalid.URL()
	data, _ := json.Marshal(url)

	ass.Equal(`{"rule":"url","msg":"Please use a valid URL"}`, string(data), "url Marshal")
}
