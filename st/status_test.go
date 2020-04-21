package st

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

}

func Test_ConverErrorPanics(t *testing.T) {

	err := status.Error(codes.Unavailable, "server not unavailable")
	sts := ConvertError(err)
	spew.Dump(err, sts)

}
