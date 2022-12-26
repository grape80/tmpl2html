package tmpl2html

import "fmt"

type errorCode uint8

type t2hError struct {
	code errorCode
	err  error
}

func (e *t2hError) Error() string {
	if e.code == ERR_UNEXPECTED {
		return fmt.Sprintf("%s\n", errMsgs[e.code])
	}
	return fmt.Sprintf("%s\n%s", errMsgs[e.code], usage)
}

func (e *t2hError) Unwrap() error {
	return e.err
}

const (
	ERR_UNEXPECTED errorCode = iota
	ERR_NO_ARGS
)

var errMsgs = map[errorCode]string{
	ERR_UNEXPECTED: "An unexpected error has occurred.",
	ERR_NO_ARGS:    "No argument specified.",
}
