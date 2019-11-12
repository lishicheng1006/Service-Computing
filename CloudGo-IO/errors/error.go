package errors

import . "errors"

var (
	codes = make(map[error]ErrCode)

	ErrInvalidStuId     = New("invalid student id")
	ErrInvalidUsername  = New("invalid username")
	ErrInvalidEmail     = New("invalid email")
	ErrInvalidPhone     = New("invalid phone")
	ErrUserExists       = New("user exists")
	ErrUserDoesNotExist = New("user does not exist")
)

type ErrCode struct {
	Code           int    `json:"code,omitempty"`
	Message        string `json:"msg,omitempty"`
	HTTPStatusCode int    `json:"-"`
}

func init() {
	newErrCode(ErrInvalidStuId, 401, 1)
	newErrCode(ErrInvalidUsername, 401, 2)
	newErrCode(ErrInvalidEmail, 401, 3)
	newErrCode(ErrInvalidPhone, 401, 4)
	newErrCode(ErrUserExists, 401, 5)
	newErrCode(ErrUserDoesNotExist, 401, 6)
}

type OK struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func OKCode(data interface{}) OK {
	return OK{200, data}
}

func FromErrCode(err error) (ErrCode, bool) {
	v, ok := codes[err]
	return v, ok
}

func newErrCode(err error, status int, code int) {
	errCode := ErrCode{
		Code:           code,
		Message:        err.Error(),
		HTTPStatusCode: status,
	}
	codes[err] = errCode
}
