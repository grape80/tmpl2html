package tmpl2html_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/grape80/tmpl2html"
)

type OSArgsTest struct {
	programName string
	command     string
	options     []string
	tmpl        string
}

const (
	programName = "tmpl2html"
)

func TestCLI_Run(t *testing.T) {
	tests := map[string]struct {
		OSArgsTest
		exitCodeExpected int
	}{
		"help":              {OSArgsTest{programName, tmpl2html.CMD_HELP, []string{}, ""}, 0},
		"help_useless_args": {OSArgsTest{programName, tmpl2html.CMD_HELP, []string{"--option=useless"}, "useless.tmpl"}, 0},

		"version":              {OSArgsTest{programName, tmpl2html.CMD_VERSION, []string{}, ""}, 0},
		"version_useless_args": {OSArgsTest{programName, tmpl2html.CMD_VERSION, []string{"--option=useless"}, "useless.tmpl"}, 0},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(tt *testing.T) {
			args := []string{
				test.programName,
				test.command,
				strings.Join(test.options, " "),
				test.tmpl,
			}

			var stdout = new(bytes.Buffer)
			var stderr = new(bytes.Buffer)
			cli := &tmpl2html.CLI{
				Args:   args,
				Stdout: stdout,
				Stderr: stderr,
			}

			exitCode := cli.Run()
			if exitCode != test.exitCodeExpected {
				t.Errorf("Run() = %v; want %v", exitCode, test.exitCodeExpected)
			}
		})
	}
}
