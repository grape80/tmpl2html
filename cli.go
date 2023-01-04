package tmpl2html

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"
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
		fmt.Fprintf(cli.Stderr, "%v\n%s\n", err, usage)
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

	if err := cli.execute(); err != nil {
		return err
	}

	return err
}

const (
	CMD_RUN     = "run"
	CMD_HELP    = "help"
	CMD_VERSION = "version"
)

//go:embed embed/usage.txt
var usage []byte

//go:embed embed/version.txt
var version string

// Options
var (
	basedir  string
	makedeps bool
)

// Gotemplate
var gotmpl string

func (cli *CLI) parseArgs() (exit bool, err error) {
	if len(cli.Args) < 2 {
		return true, errors.New(ERR_NO_ARGS)
	}

	switch cli.Args[1] {
	case CMD_RUN:
		runCmd := flag.NewFlagSet(CMD_RUN, flag.ContinueOnError)
		runCmd.StringVar(&basedir, OPT_BASEDIR, OPT_BASEDIR_DEFAULT, "")
		runCmd.BoolVar(&makedeps, OPT_MAKEDEPS, OPT_MAKEDEPS_DEFAULT, "")
		runCmd.SetOutput(io.Discard)

		if err = runCmd.Parse(cli.Args[2:]); err != nil {
			m := strings.Replace(err.Error(), " -", " ", 1) // -option -> option
			return true, errors.New(m)
		}

		l := len(runCmd.Args())
		switch {
		case l == 1:
			gotmpl = runCmd.Args()[0]
		case l == 0:
			return true, errors.New(ERR_NO_GOTMPL)
		default:
			return true, errors.New(ERR_MULTI_GOTMPLS)
		}
	case CMD_HELP:
		_, err = fmt.Fprintf(cli.Stdout, "%s\n", usage)
		return true, err
	case CMD_VERSION:
		_, err = fmt.Fprintf(cli.Stdout, "%s\n", version)
		return true, err
	default:
		return true, errors.New(ERR_INVALID_CMD)
	}

	return exit, err
}

func (cli *CLI) execute() (err error) {
	c := &converter{
		stdout:   cli.Stdout,
		stderr:   cli.Stderr,
		basedir:  basedir,
		gotmpl:   gotmpl,
		makedeps: makedeps,
	}

	if err := c.run(); err != nil {
		return err
	}

	return err
}
