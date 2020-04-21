package st

import (
	"google.golang.org/grpc/codes"
	gs "google.golang.org/grpc/status"
)

//
type err struct {
	gst *gs.Status
	st  *body
}

//
type body struct {
	Code int32  `json:"code" yaml:"code" `
	Msg  string `json:"msg" yaml:"msg"`
}

func (e *err) parseMsg() error {

	e.st = &body{}

	err := json.Unmarshal([]byte(e.gst.Message()), e.st)

	return err
}

// String -
func (e *err) String() string {
	return e.gst.Message()
}

func (e *err) Error() string {
	return e.gst.Err().Error()
}

func (e *err) Err() error {
	return e.gst.Err()
}

func (e *err) GetCode() int32 {
	return e.st.Code
}

func (e *err) GetMsg() string {
	return e.st.Msg
}
func (e *err) GetGRPCCode() codes.Code {
	return e.gst.Code()
}

func (e *err) Equal(er error) bool {
	if x, ok := er.(*err); ok {
		if x.GetCode() == e.GetCode() {
			return true
		}
	}
	return false
}
