package vo

type ErrorCode int64

const (
	errorCodeBind     ErrorCode = 420 // 参数绑定错误
	errorCodeValidate           = 421 // 参数校验错误
)

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Detail  string    `json:"detail"`
}

func (e Error) Error() string {
	return e.Message
}

func ErrorNew(code ErrorCode, msg string, detail string) Error {
	return Error{code, msg, detail}
}

func ErrorWrap(code ErrorCode, msg string, err error) Error {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	return Error{code, msg, errMsg}
}

func ErrorBind(err error) Error {
	return Error{errorCodeBind, "参数错误", err.Error()}
}
func ErrorValidate(detail string) Error {
	return Error{errorCodeValidate, "参数错误", detail}
}
