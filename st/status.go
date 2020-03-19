package st

import (
	"github.com/axolotlteam/thunder/logger"
	"google.golang.org/grpc/codes"
	gs "google.golang.org/grpc/status"
)

// NewError -
func NewError(code int32, msg string, gc codes.Code) error {
	if _, ok := list[code]; ok {
		logger.Panicf("code : %d is duplicate", code)
	}

	if gc == codes.OK {
		return nil
	}

	x := &body{
		Code: code,
		Msg:  msg,
	}

	str, _ := json.Marshal(x)
	s := &err{
		gst: gs.New(gc, string(str)),
		st:  x,
	}

	err := s.Err()

	list[code] = err

	return err
}

// List -
func List() map[int32]error {
	return list
}

// ConvertError -
func ConvertError(errs error) Errors {

	if errs == nil {
		return &err{
			gst: gs.New(OK, "success"),
			st:  &body{Code: 0, Msg: "success"},
		}
	}

	gt, ok := gs.FromError(errs)

	temp := &err{
		gst: gt,
	}

	if ok {
		temp.parseMsg()
	} else {
		logger.Panicf("convert error failed with errr : %s", errs.Error())
	}

	return temp
}
