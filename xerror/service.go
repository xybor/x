package xerror

import (
	"fmt"
)

type ServiceError struct {
	err error
	msg string
}

func (err ServiceError) Error() string {
	return err.err.Error()
}

func (err ServiceError) Message() string {
	return err.msg
}

func (err ServiceError) Unwrap() error {
	return err.err
}

func Wrap(err error, msg string, a ...any) ServiceError {
	return ServiceError{err: err, msg: fmt.Sprintf(msg, a...)}
}
