package st

import "google.golang.org/grpc/codes"

// Errors -
type Errors interface {
	Error() string
	String() string
	GetCode() int32
	GetGRPCCode() codes.Code
	GetMsg() string
	Err() error
	Equal(err error) bool
}
