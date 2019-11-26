package status

import (
	"github.com/axolotlteam/thunder/logger"
	"google.golang.org/grpc/codes"
	gs "google.golang.org/grpc/status"
)

// GRPCCode -
type GRPCCode = codes.Code

// Status -
type status struct {
	Code int32  `json:"code" yaml:"code"`
	Msg  string `json:"msg" yaml:"msg"`
}

//
type grpcstatus struct {
	gst *gs.Status
	s   *status
}

func (s *grpcstatus) Error() string {
	return s.gst.Message()
}

func (s *grpcstatus) String() string {
	return s.gst.Message()
}

func (s *grpcstatus) GetCode() int32 {
	return s.s.Code
}

func (s *grpcstatus) GetGRPCCode() GRPCCode {
	return s.gst.Code()
}

func (s *grpcstatus) GetMsg() string {
	return s.s.Msg

}

func (s *grpcstatus) Err() error {
	return s.gst.Err()
}

func newGrpcStatus(gcode codes.Code, code int32, msg string) Status {

	ns := &grpcstatus{
		s: &status{
			Code: code,
			Msg:  msg,
		},
	}

	x, _ := json.Marshal(ns.s)
	ns.gst = gs.New(gcode, string(x))
	return ns
}

//
func fromGRPCStatus(err error) Status {
	st, ok := gs.FromError(err)

	if ok {
		temp := &grpcstatus{
			gst: st,
		}
		temp.parseMsg()
		return temp
	}

	return &grpcstatus{
		s: &status{
			Code: 0,
			Msg:  err.Error(),
		},
	}
}

func (s *grpcstatus) parseMsg() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	// msg = >
	s.s = &status{}
	err := json.Unmarshal([]byte(s.gst.Message()), s.s)
	if err != nil {
		logger.Error("Grpc status Unmarshal Error")
	}
}
