package kvalid_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kvalid"
)

func TestIgnore(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	req := require.New(t)
	str := &strType{}
	rules := kvalid.New(str).Field(&str.Field, kvalid.Ignore().SetMessage(_msg))
	req.NoError(rules.Validate(str), "Valid")

	data, _ := json.Marshal(rules)
	ass.Equal(`{"Field":[{"rule":"ignore","msg":"custom message"}]}`, string(data))
}
