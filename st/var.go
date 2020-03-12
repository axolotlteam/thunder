package st

// 0	    success
// 11000	warning
// 12000	other warning(by project)
// 21000	error
// 22000	other error(by project)
//
var (
	List = make(map[int32]error, 0)
)

// Error List
var (
	NoError = NewError(0, "success", OK)
	// waring
	WarningLoginFail        = NewError(11000, "loging failed", NotFound)
	WarningRegisterFail     = NewError(11001, "register failed", NotFound)
	WarningRegisterExists   = NewError(11007, "register data was exists", NotFound)
	WarningInvalidParameter = NewError(11008, "invalid rrgument", InvalidArgument)

	// error
	ErrorConnectFailed         = NewError(21000, "connect failed", Unavailable)
	ErrorConnectTimeOut        = NewError(21001, "connect time out", DeadlineExceeded)
	ErrorDatabaseConnectFailed = NewError(21002, "database connect failed", Unavailable)
	ErrorDatabaseCreateFailed  = NewError(21003, "database create failed", Aborted)
	ErrorDataNotFound          = NewError(21004, "data not found", NotFound)
	ErrorDatabaseUpdateFailed  = NewError(21005, "database update failed", Aborted)
	ErrorDatabaseDeleteFailed  = NewError(21006, "database delete failed", Aborted)
	ErrorDataIsExists          = NewError(21007, "data is exists", AlreadyExists)
	ErrorInvalidParameter      = NewError(21008, "invalid argument", InvalidArgument)
)
