package packer

const DEFAULT_ERROR_CODE = 50000
const DEFAULT_SUCCESS_CODE = 0

const (
	OK = 0

	//Client Error
	UNKNOWN_CLIENT_ERROR           = 40000
	INVALID_PARAMS_ERROR           = 40001
	INVALID_TOKEN_ERROR            = 40002
	INVALID_ACCOUNT_PASSWORD_ERROR = 40003
	PASSWORD_NOT_MATCH_ERROR       = 40004

	//Server Error
	UNKNOWN_SERVER_ERROR     = 50000
	AUTH_DELIBER_TOKEN_ERROR = 50001
	USER_REGISTER_ERROR      = 50002
)

var ErrorMessages = map[int]string{
	0: "ok",

	//Client Error
	40000: "Unknow Client Error",
	40001: "Invalid Params",
	40002: "Invalid Token",
	40003: "Invalid Account or Password",
	40004: "Password Not Match",

	//Server Error
	50000: "Unknow Server Error",
	50001: "Auth Deliever Token Error",
	50002: "User Register Error",
}

type MyError struct {
	Code int
	Err  error
}

func (e *MyError) Error() string {
	return e.Err.Error()
}

func NewMyError(code int, err error) *MyError {
	return &MyError{
		Code: code,
		Err:  err,
	}
}
