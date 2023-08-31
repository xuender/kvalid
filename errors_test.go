package kvalid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/kvalid"
)

func TestErrors_Error(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	errors := kvalid.Errors{
		kvalid.NewError("a", "f1"),
		kvalid.NewError("b", "f1"),
	}

	ass.Equal("a. b.", errors.Error())
}
