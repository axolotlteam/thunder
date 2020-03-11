package status

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc/codes"
	st "google.golang.org/grpc/status"
)

func Test_GRPCStatus(t *testing.T) {
	//

	x := struct {
		Code int32  `json:"code" yaml:"code"`
		Msg  string `json:"msg" yaml:"msg"`
	}{
		Code: 1234,
		Msg:  "hello",
	}

	msg, err := json.Marshal(&x)
	if err != nil {
		panic(err)
	}

	err = st.New(codes.NotFound, string(msg)).Err()

	fmt.Println(err)

	m := GRPCConvert(err)

	spew.Dump(m)
}

//
func Test_NewGRPCStatus(t *testing.T) {
	x := NewGRPCStatus(Unknown, 0, "123").Err()
	spew.Dump(x)
}
