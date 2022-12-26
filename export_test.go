package tmpl2html

import (
	"fmt"
	"io/fs"
)

var (
	ErrWrapped = &cliError{code: ERR_UNEXPECTED, err: fs.ErrExist}

	ErrMsgUnexpected  = fmt.Sprintf("%s\n", errMsgs[ERR_UNEXPECTED])
	ErrUnexpected     = &cliError{code: ERR_UNEXPECTED}
)
