package kvalid_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kvalid"
	"gopkg.in/guregu/null.v3"
)

type requiredType struct {
	Number int
	String string `json:"string,omitempty"`
	Time   time.Time
	Null   null.Int
	Ptr    *string
}

// nolint: errorlint, forcetypeassert
func TestRequired(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	requ := require.New(t)
	req := &requiredType{}
	errs := kvalid.New(req).
		Field(&req.Number, kvalid.Required()).
		Field(&req.String, kvalid.Required()).
		Field(&req.Time, kvalid.Required()).
		Field(&req.Null, kvalid.Required()).
		Field(&req.Ptr, kvalid.Required()).
		Validate(req)

	requ.ErrorAs(errs, &kvalid.Errors{})
	ass.Len(errs.(kvalid.Errors), 5, "All not set")

	str := ""
	req = &requiredType{
		Null: null.NewInt(0, false),
		Time: time.Time{},
		Ptr:  &str,
	}
	errs = kvalid.New(req).
		Field(&req.Time, kvalid.Required()).
		Field(&req.Null, kvalid.Required()).
		Field(&req.Ptr, kvalid.Required()).
		Validate(req)

	ass.Len(errs.(kvalid.Errors), 3, "Complex type set but empty")

	str = "ok"
	req = &requiredType{
		Number: 1,
		String: "ok",
		Null:   null.NewInt(0, true),
		Time:   time.Now(),
		Ptr:    &str,
	}
	errs = kvalid.New(req).
		Field(&req.Number, kvalid.Required()).
		Field(&req.String, kvalid.Required()).
		Field(&req.Time, kvalid.Required()).
		Field(&req.Null, kvalid.Required()).
		Field(&req.Ptr, kvalid.Required()).
		Validate(req)
	requ.NoError(errs, "All value are set")
	// message
	ass.Equal(_msg, kvalid.New(req).Field(&req.Number, kvalid.Required().SetMessage(_msg)).
		Validate(&requiredType{}).(kvalid.Errors)[0].Error(), "Custom error message")
	ass.NotEqual(_msg, kvalid.New(req).Field(&req.Number, kvalid.Required()).
		Validate(&requiredType{}).(kvalid.Errors)[0].Error(), "Default error message")
}
