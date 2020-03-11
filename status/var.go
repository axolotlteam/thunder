package status

// 0	    success
// 11000	warning
// 12000	other warning(by project)
// 21000	error
// 22000	other error(by project)
var (
	NoError = NewGRPCStatus(OK, 0, "success")
	// waring
	WarningLoginFail    = NewGRPCStatus(NotFound, 11000, "loging failed")
	WarningRegisterFail = NewGRPCStatus(NotFound, 11001, "register failed")

	WarningRegisterExists   = NewGRPCStatus(NotFound, 11007, "register data exists")
	WarningInvalidParameter = NewGRPCStatus(InvalidArgument, 11008, "Invalid Argument")

	// error
	ErrorConnectFailed         = NewGRPCStatus(Unavailable, 21000, "connect failed")
	ErrorConnectTimeOut        = NewGRPCStatus(DeadlineExceeded, 21001, "connect time out")
	ErrorDatabaseConnectFailed = NewGRPCStatus(Unavailable, 21002, "database connect failed")
	ErrorDatabaseCreateFailed  = NewGRPCStatus(Aborted, 21003, "database create failed")
	ErrorDataNotFound          = NewGRPCStatus(NotFound, 21004, "data not found")
	ErrorDatabaseUpdateFailed  = NewGRPCStatus(Aborted, 21005, "database update failed")
	ErrorDatabaseDeleteFailed  = NewGRPCStatus(Aborted, 21006, "database delete failed")
	ErrorDataIsExists          = NewGRPCStatus(AlreadyExists, 21007, "data is exists")
	ErrorInvalidParameter      = NewGRPCStatus(InvalidArgument, 21008, "Invalid Argument")
)
