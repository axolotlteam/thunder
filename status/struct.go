package status

var list map[int]Error

func init() {

}

// Error -
type Error interface {
	Error() string
	String() string
	GetCode() int
}

type status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Unix int64  `json:"unix"`
}
