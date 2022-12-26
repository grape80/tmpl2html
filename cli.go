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
	EXIT_SUCCESS = 0
	EXIT_FAILURE = 1
)

func (cli *CLI) Run() (exitCode int) {
	if err := cli.runContext(); err != nil {
		fmt.Fprintf(cli.Stderr, "%v\n", err)
		return EXIT_FAILURE
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
	CMD_HELP    = "help"
	CMD_VERSION = "version"
)

//go:embed embed/usage.txt
var usage []byte

//go:embed embed/version.txt
var version string

func (cli *CLI) parseArgs() (exit bool, err error) {
	if len(cli.Args) < 2 {
		return true, &cliError{code: ERR_NO_ARGS}
	}

	switch cli.Args[1] {
	case CMD_HELP:
		_, err = fmt.Fprintf(cli.Stdout, "%s\n", usage)
		return true, err
	case CMD_VERSION:
		_, err = fmt.Fprintf(cli.Stdout, "%s\n", version)
		return true, err
	}

	return exit, err
}
