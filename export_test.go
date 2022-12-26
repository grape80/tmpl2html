package tmpl2html

import (
	"fmt"
	"io/fs"
)

var (
	ErrWrapped = &t2hError{code: ERR_UNEXPECTED, err: fs.ErrExist}

	MsgUnexpected = fmt.Sprintf("%s\n", errMsgs[ERR_UNEXPECTED])
	ErrUnexpected = &t2hError{code: ERR_UNEXPECTED}

	MsgNoArgs = fmt.Sprintf("%s\n%s", errMsgs[ERR_NO_ARGS], usage)
	ErrNoArgs = &t2hError{code: ERR_NO_ARGS}
)
