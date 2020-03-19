package st

// 0	    success
// 11000	warning
// 12000	other warning(by project)
// 21000	error
// 22000	other error(by project)
//
var (
	list = make(map[int32]error, 0)
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
	ErrorConnectFailed   = NewError(21000, "connect failed", Unavailable)
	ErrorConnectTimeOut  = NewError(21001, "connect time out", DeadlineExceeded)
	ErrorInvalidProtocol = NewError(21002, "invalid protocol", Aborted)
	// Daatabase
	ErrorDatabaseConnectFailed = NewError(21102, "database connect failed", Unavailable)
	ErrorDatabaseCreateFailed  = NewError(21103, "database create failed", Aborted)
	ErrorDatabaseUpdateFailed  = NewError(21104, "database update failed", Aborted)
	ErrorDatabaseDeleteFailed  = NewError(21105, "database delete failed", Aborted)
	ErrorDataNotFound          = NewError(21106, "data not found", NotFound)
	ErrorDataIsExists          = NewError(21107, "data is exists", AlreadyExists)

	// auth
	ErrorInvalidOAuthToken = NewError(21201, "invalid oauth access token", Unauthenticated)
	ErrorInvalidParameter  = NewError(21202, "invalid argument", InvalidArgument)
	ErrorPermissionsFailed = NewError(21203, "perrmission failed", PermissionDenied)
)
