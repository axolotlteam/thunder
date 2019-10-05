package status

// GRPCConvert -
func GRPCConvert(err error) Status {
	return fromGRPCStatus(err)
}

// NewGRPCStatus -
func NewGRPCStatus(gcode GRPCCode, errcode int32, msg string) Status {
	return newGrpcStatus(gcode, errcode, msg)
}
