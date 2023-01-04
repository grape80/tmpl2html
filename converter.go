package tmpl2html

import (
	"fmt"
	"io"
)

const (
	OPT_BASEDIR  = "basedir"
	OPT_MAKEDEPS = "makedeps"
	OPT_VERBOSE  = "verbose"
)

const (
	OPT_BASEDIR_DEFAULT  = "./"
	OPT_MAKEDEPS_DEFAULT = false
	OPT_VERBOSE_DEFAULT  = false
)

type converter struct {
	stdout   io.Writer
	stderr   io.Writer
	makedeps bool
	verbose  bool
	basedir  string
	gotmpl   string
	partials []string
}

func (c *converter) run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(c.stderr, "%s %v\n", ERR_RUNTIME, r)
		}
	}()

	return err
}
