package kvalid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kvalid"
)

type emailType struct {
	Field string
}

// nolint: errorlint, forcetypeassert
func TestEmail(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	data := &emailType{}
	rules := kvalid.New(data).Field(&data.Field, kvalid.Email())
	ass.Len(rules.Validate(&emailType{Field: "fake"}).(kvalid.Errors), 1, "Invalid email address")
	ass.Nil(rules.Validate(&emailType{Field: "test@mail.com"}), "Valid email address")
	ass.Equal(_msg, kvalid.New(data).Field(&data.Field, kvalid.Email().SetMessage(_msg)).
		Validate(&emailType{Field: "invalid"}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(data).Field(&data.Field, kvalid.Email()).
		Validate(&emailType{Field: "invalid"}).(kvalid.Errors)[0].Error(), "Default error message")
	// optional
	rules = kvalid.New(data).Field(&data.Field, kvalid.Email().Optional())
	ass.Nil(rules.Validate(&emailType{Field: ""}), "Invalid but zero")
	ass.Len(rules.Validate(&emailType{Field: "fake"}).(kvalid.Errors), 1, "Invalid and not zero")
	ass.Nil(rules.Validate(&emailType{Field: "test@mail.com"}), "Valid and not zero")
}

func TestEmailValidator_MarshalJSON(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	email := kvalid.Email()
	data, _ := json.Marshal(email)

	ass.Equal(`{"rule":"email","msg":"Please use a valid email address"}`, string(data), "email Marshal")
}
