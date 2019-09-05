package errno

import (
	"fmt"
)

/*
This function defines all error functions.
*/
type Errno struct {
	Code    int
	Message string
}

// define a function for go struct by this format.
func (err Errno) Error() string {
	return err.Message
}

type Err struct {
	Code    int
	Message string
	Err     error
}

// construct new error code
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

func (err *Err) Add(message string) *Err {
	err.Message += " " + message
	return err
}

func (err *Err) addf(format string, args ...interface{}) *Err {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error %s", err.Code, err.Message, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}
	return InternalServerError.Code, err.Error()
}

// define constant code for error
var (
	// common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrValidation = &Errno{Code:20001, Message:"Validation failed."}
	ErrDatabase = &Errno{Code:20002, Message:"Database Error."}
	// user errors
	ErrUserNotFound = &Errno{Code:20101, Message:"The user was not found."}
	ErrPasswordIncorrect = &Errno{Code:20102, Message:"The pass word was incorrect."}
)
