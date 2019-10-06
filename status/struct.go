package status

//
var (
	NoError                    = NewGRPCStatus(OK, 0, "success")
	ErrorConnectFailed         = NewGRPCStatus(Unavailable, 98, "connect failed")
	ErrorConnectTimeOut        = NewGRPCStatus(DeadlineExceeded, 99, "connect time out")
	ErrorDatabaseConnectFailed = NewGRPCStatus(Internal, 102, "database connect failed")
	ErrorDatabaseCreateFailed  = NewGRPCStatus(Aborted, 103, "database create failed")
	ErrorDatabaseDeleteFailed  = NewGRPCStatus(Aborted, 104, "database delete failed")
	ErrorDatabaseUpdateFailed  = NewGRPCStatus(Aborted, 105, "database update failed")
	ErrorDataNotFound          = NewGRPCStatus(NotFound, 106, "data not found")
	ErrorDataIsExists          = NewGRPCStatus(AlreadyExists, 200, "data is exists")
	ErrorInvalidParameter      = NewGRPCStatus(InvalidArgument, 300, "Invalid Argument")
)
