package shared

import (
	"fmt"
	"runtime"
)

type MyError struct {
	status  int
	message string
	caller  string
	err     error
}

func (e *MyError) Status() int {
	return e.status
}

func (e *MyError) Error() string {
	return e.message
}

func (e *MyError) Unwrap() error {
	return e.err
}

func (e *MyError) Print() {
	wrappedErr := e.Unwrap()
	if wrappedErr != nil {
		fmt.Print(e.status, e.message, wrappedErr.Error(), e.caller)
	} else {
		fmt.Print(e.status, e.message, e.caller)
	}
}

func NewError(status int, msg string, err error) error {
	return &MyError{
		status:  status,
		message: msg,
		caller:  caller(),
		err:     err,
	}
}

func caller() string {
	i := 2
	caller := ""
	for {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		caller += fmt.Sprintf("%v,line: %v\n", file, line)
		i++
	}

	return caller
}
