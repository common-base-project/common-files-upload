package code

import "fmt"

type Errno struct {
	ErrNo   int
	ErrMsg  string
	ErrData interface{}
}

func (err Errno) Error() string {
	return err.ErrMsg
}

// Err represents an error
type Err struct {
	Errno  int
	Errmsg string
	Err    error
}

func New(errno *Errno, err error) *Err {
	return &Err{Errno: errno.ErrNo, Errmsg: errno.ErrMsg, Err: err}
}

func (err *Err) Add(message string) error {
	err.Errmsg += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Errmsg += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - errno: %d, errmsg: %s, error: %s", err.Errno, err.Errmsg, err.Err)
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return Success.ErrNo, Success.ErrMsg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Errno, typed.Errmsg
	case *Errno:
		return typed.ErrNo, typed.ErrMsg
	default:
	}

	return InternalServerError.ErrNo, err.Error()
}
