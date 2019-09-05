package errno

import "fmt"

type Errno struct {
	Code int
	Message string
}
// define a function for go struct by this format.
func (err Errno) Error() string {
	return err.Message
}

type Err struct {
	Code int
	Message string
	Err error
}
// construct new error code
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

func (err *Err) Add(message string) *Err {
	err.Message += " " + message
	return err
}

func(err *Err) addf(format string, args... interface{}) *Err {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {

}


