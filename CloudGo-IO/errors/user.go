package errors

import . "errors"

var (
	codes = make(map[error]ErrorCode)

	ErrInvalidStuId     = New("invalid student id")
	ErrInvalidUsername  = New("invalid username")
	ErrInvalidEmail     = New("invalid email")
	ErrInvalidPhone     = New("invalid phone")
	ErrUserExists       = New("user exists")
	ErrUserDoesNotExist = New("user does not exist")
)

type ErrorCode struct {
	Code           int    `json:"code,omitempty"`
	Message        string `json:"msg,omitempty"`
	HTTPStatusCode int    `json:"-"`
}

func init() {
	newErrorCode(ErrInvalidStuId, 401, 1)
	newErrorCode(ErrInvalidUsername, 401, 2)
	newErrorCode(ErrInvalidEmail, 401, 3)
	newErrorCode(ErrInvalidPhone, 401, 4)
	newErrorCode(ErrUserExists, 401, 5)
	newErrorCode(ErrUserDoesNotExist, 401, 6)
}

type OK struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func OKCode(data interface{}) OK {
	return OK{200, data}
}

func FromErrorCode(err error) (ErrorCode, bool) {
	v, ok := codes[err]
	return v, ok
}

func newErrorCode(err error, status int, code int, ) {
	errCode := ErrorCode{
		Code:           code,
		Message:        err.Error(),
		HTTPStatusCode: status,
	}
	codes[err] = errCode
}
