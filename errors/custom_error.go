package custom_error

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Op   string
	Code int
	Err  error
	Msg  string
}

func NewAuthenticationError(op string, err error, msg string) error {
	return &CustomError{
		Op:   op,
		Code: http.StatusUnauthorized,
		Err:  err,
		Msg:  msg,
	}
}

func NewAuthorizationError(op string, err error, msg string) error {
	return &CustomError{
		Op:   op,
		Code: http.StatusForbidden,
		Err:  err,
		Msg:  msg,
	}
}

func NewUnexpectedError(op string, err error, msg string) error {
	return &CustomError{
		Op:   op,
		Code: http.StatusInternalServerError,
		Err:  err,
		Msg:  msg,
	}
}

func NewDatabaseError(op string, err error, msg string) error {
	return &CustomError{
		Op:   op,
		Code: http.StatusInternalServerError,
		Err:  err,
		Msg:  msg,
	}
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("Operation:%s :LogMsg:%s :Msg:%s", c.Op, c.Err, c.Msg)
}
