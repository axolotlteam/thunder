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

	// error list

	// Connect
	ErrorConnectFailed   = NewError(21000, "connect failed", Unavailable)
	ErrorConnectTimeOut  = NewError(21001, "connect time out", DeadlineExceeded)
	ErrorInvalidProtocol = NewError(21002, "invalid protocol", Aborted)
	ErrorHostNotFound    = NewError(21003, "host not found", Unavailable)
	ErrorServiceNotFound = NewError(21004, "service not found", Unavailable)

	// Database
	ErrorDatabaseConnectFailed = NewError(21102, "database connect failed", Unavailable)
	ErrorDatabaseCreateFailed  = NewError(21103, "database create failed", Aborted)
	ErrorDatabaseUpdateFailed  = NewError(21104, "database update failed", Aborted)
	ErrorDatabaseDeleteFailed  = NewError(21105, "database delete failed", Aborted)
	ErrorDataNotFound          = NewError(21106, "data not found", NotFound)
	ErrorDataIsExists          = NewError(21107, "data is exists", AlreadyExists)

	// Auth
	ErrorInvalidOAuthToken = NewError(21201, "invalid oauth access token", Unauthenticated)
	ErrorPermissionsFailed = NewError(21205, "permission failed", PermissionDenied)
	ErrorUserDisabled      = NewError(21207, "user was disaabled", PermissionDenied)
	ErrorUserFreeze        = NewError(21208, "user was freeze", PermissionDenied)
	ErrorInvalidCheckCode  = NewError(21209, "invalid check code", PermissionDenied)

	// File
	ErrorUploadFileTooLarge = NewError(21304, "upload file too large", Aborted)
	ErrorUploadFileNotFound = NewError(21305, "upload file not found", Aborted)
	ErrorOpenFileFailed     = NewError(21310, "open file failed", Aborted)
	ErrorUploadPhotoFailed  = NewError(21311, "upload photo failed", Aborted)

	// Parameters
	ErrorInvalidParameter      = NewError(21402, "invalid argument", InvalidArgument)           // 錯誤的參數
	ErrorMissRequiredParameter = NewError(21403, "missing required parameter", InvalidArgument) // 缺少必填參數
	ErrroOverParameter         = NewError(21404, "over max parameter", InvalidArgument)         // 超過最大可輸入參數
	ErrorParameterOutOfRange   = NewError(21405, "parameter out of range", OutOfRange)          // 參數超過範圍
	ErrorParameterNotAllowed   = NewError(21406, "parameter not allowed", InvalidArgument)      // 參數不被允許

	// DataCheck
	ErrorInvalidUserID      = NewError(21601, "invalid user id", InvalidArgument)
	ErrorInvalidEmail       = NewError(21602, "invalid Email", InvalidArgument)
	ErrorInvalidPhoneNumber = NewError(21603, "invalid phone number", InvalidArgument)
	ErrorPhotoEditFailed    = NewError(21630, "photo edit failed", Aborted)
)
