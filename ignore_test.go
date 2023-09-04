package kvalid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kvalid"
	"github.com/xuender/kvalid/json"
)

func TestIgnore(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	str := &strType{}
	rules := kvalid.New(str).Field(&str.Field, kvalid.Ignore().SetMessage(_msg))
	ass.Nil(rules.Validate(str), "Valid")

	data, _ := json.Marshal(rules)
	ass.Equal(`{}`, string(data))
}
