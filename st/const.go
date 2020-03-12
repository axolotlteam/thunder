package st

import (
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc/codes"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// -
const (
	OK                 = codes.OK
	Canceled           = codes.Canceled
	Unknown            = codes.Unknown
	InvalidArgument    = codes.InvalidArgument
	DeadlineExceeded   = codes.DeadlineExceeded
	NotFound           = codes.NotFound
	AlreadyExists      = codes.AlreadyExists
	PermissionDenied   = codes.PermissionDenied
	ResourceExhausted  = codes.ResourceExhausted
	FailedPrecondition = codes.FailedPrecondition
	Aborted            = codes.Aborted
	OutOfRange         = codes.OutOfRange
	Unimplemented      = codes.Unimplemented
	Internal           = codes.Internal
	Unavailable        = codes.Unavailable
	DataLoss           = codes.DataLoss
	Unauthenticated    = codes.Unauthenticated
)
