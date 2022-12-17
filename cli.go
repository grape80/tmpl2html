package tmpl2html

import "io"

type CLI struct {
	Stdout io.Writer
	Stderr io.Writer
	Args   []string
}

func (cli *CLI) Run() (exitcode int) {
	return exitcode
}
