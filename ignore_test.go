package kvalid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kvalid"
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
