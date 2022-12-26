package tmpl2html_test

import (
	"errors"
	"io/fs"
	"testing"

	. "github.com/grape80/tmpl2html"
)

func TestCliError_Unwrap(t *testing.T) {
	var expected = fs.ErrExist
	if !errors.Is(ErrWrapped, expected) {
		t.Errorf("Error() = %T; want %T", ErrWrapped, expected)
	}
}

func TestCliError_Error(t *testing.T) {

	tests := map[string]struct {
		errMsgExpected string
		err            error
	}{
		"err_unexpected": {ErrMsgUnexpected, ErrUnexpected},
		"err_no_args":    {ErrMsgNoArgs, ErrNoArgs},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(tt *testing.T) {
			if test.err.Error() != test.errMsgExpected {
				t.Errorf("Error() = %v; want %v", test.err.Error(), test.errMsgExpected)
			}
		})
	}
}
