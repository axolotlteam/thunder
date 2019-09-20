package status

// GrpcStatus -
type GrpcStatus interface {
}

type status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Unix int64  `json:"unix"`
}
