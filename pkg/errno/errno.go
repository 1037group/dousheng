package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(int32(0), "Success")
	ServiceErr             = NewErrNo(int32(10001), "Service is unable to start successfully")
	ParamErr               = NewErrNo(int32(10002), "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(int32(10003), "User already exists")
	UserNotExistErr        = NewErrNo(int32(10004), "User not exists")
	AuthorizationFailedErr = NewErrNo(int32(10005), "Authorization failed")
	RedisLockFailed        = NewErrNo(int32(10006), "redis lock failed")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
