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

// 204
func NewNoContentFoundError(op string, err error, msg string) error {
	return &CustomError{
		Op:   op,
		Code: http.StatusNoContent,
		Err:  err,
		Msg:  msg,
	}
}

// 400
func NewBadRequestError(op string, err error, msg string) error {
	return &CustomError{
		Op:   op,
		Code: http.StatusBadRequest,
		Err:  err,
		Msg:  msg,
	}
}

// 401
func NewAuthenticationError(op string, err error, msg string) error {
	return &CustomError{
		Op:   op,
		Code: http.StatusUnauthorized,
		Err:  err,
		Msg:  msg,
	}
}

// 403
func NewAuthorizationError(op string, err error, msg string) error {
	return &CustomError{
		Op:   op,
		Code: http.StatusForbidden,
		Err:  err,
		Msg:  msg,
	}
}

// 500
func NewUnexpectedError(op string, err error, msg string) error {
	return &CustomError{
		Op:   op,
		Code: http.StatusInternalServerError,
		Err:  err,
		Msg:  msg,
	}
}

// 500
func NewDatabaseError(op string, err error, msg string) error {
	return &CustomError{
		Op:   op,
		Code: http.StatusInternalServerError,
		Err:  err,
		Msg:  msg,
	}
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("Operation:%s :Log Message:%s :Message:%s", c.Op, c.Err, c.Msg)
}
