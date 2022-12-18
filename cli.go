package tmpl2html

import (
	_ "embed"
	"fmt"
	"io"
)

type CLI struct {
	Stdout io.Writer
	Stderr io.Writer
	Args   []string
}

const (
	errorExitCode = 1
)

func (cli *CLI) Run() (exitCode int) {
	if err := cli.runContext(); err != nil {
		fmt.Fprintf(cli.Stderr, "%v\n", err)
		return errorExitCode
	}
	return exitCode
}

func (cli *CLI) runContext() (err error) {
	if exit, err := cli.parseArgs(); err != nil {
		return err
	} else if exit == true {
		return nil
	}

	return err
}

const (
	CMD_HELP = "help"
)

//go:embed embed/usage.txt
var usage []byte

func (cli *CLI) parseArgs() (exit bool, err error) {
	switch cli.Args[1] {
	case CMD_HELP:
		_, err = fmt.Fprintf(cli.Stdout, "%s\n", usage)
		return true, err
	}

	return exit, err
}
