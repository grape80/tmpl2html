package tmpl2html

import "fmt"

type errorCode uint8

type cliError struct {
	code errorCode
	err  error
}

func (e *cliError) Error() string {
	if e.code == ERR_UNEXPECTED {
		return fmt.Sprintf("%s\n", errMsgs[e.code])
	}
	return fmt.Sprintf("%s\n%s", errMsgs[e.code], usage)
}

func (e *cliError) Unwrap() error {
	return e.err
}

const (
	ERR_UNEXPECTED errorCode = iota
)

var errMsgs = map[errorCode]string{
	ERR_UNEXPECTED: "An unexpected error has occurred.",
}
