package mobile

import (
	"fmt"

	"github.com/gopub/errors"
)

type Error errors.Error

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) String() string {
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func ToError(err error) *Error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*Error); ok {
		return e
	}

	if e, ok := err.(*errors.Error); ok {
		return (*Error)(e)
	}

	return NewError(0, err.Error())
}
