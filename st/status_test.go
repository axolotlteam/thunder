package st

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func Test_NewError(t *testing.T) {

	err := NewError(0, "success", OK)
	spew.Dump(err)
	assert.NoError(t, err)

	err = NewError(1000, "NotFound", NotFound)
	spew.Dump(err)
	assert.Error(t, err, "NotFound")
}

func Test_ConverError(t *testing.T) {
	err := NewError(0, "success", OK)

	sts := ConvertError(err)
	assert.Equal(t, sts.GetCode(), int32(0))
	err = NewError(1000, "NotFound", NotFound)
	sts = ConvertError(err)
	assert.Equal(t, sts.GetCode(), int32(1000))

}
