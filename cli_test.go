package tmpl2html_test

import (
	"bytes"
	"testing"

	. "github.com/grape80/tmpl2html"
)

func TestCLI_Run(t *testing.T) {

	tests := map[string]struct {
		exitCodeExpected int
		args             []string
	}{
		"no_args":     {EXIT_FAILURE, []string{}},
		"invalid_cmd": {EXIT_FAILURE, []string{"invalid", "go.tmpl"}},

		// run
		"run_no_gotmpl":     {EXIT_FAILURE, []string{CMD_RUN}},
		"run_invalid_opt":   {EXIT_FAILURE, []string{CMD_RUN, "--option=invalid", "go.tmpl"}},
		"run_multi_gotmpls": {EXIT_FAILURE, []string{CMD_RUN, "go.tmpl", "go2.tmpl"}},

		"run_no_opts": {EXIT_SUCCESS, []string{CMD_RUN, "go.tmpl"}},

		// help
		"help":              {EXIT_SUCCESS, []string{CMD_HELP}},
		"help_useless_args": {EXIT_SUCCESS, []string{CMD_HELP, "--option=useless", "useless.tmpl"}},

		// version
		"version":              {EXIT_SUCCESS, []string{CMD_VERSION}},
		"version_useless_args": {EXIT_SUCCESS, []string{CMD_VERSION, "--option=useless", "useless.tmpl"}},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(tt *testing.T) {
			args := []string{"tmpl2html"}
			args = append(args, test.args...)

			stdout := new(bytes.Buffer)
			stderr := new(bytes.Buffer)
			cli := &CLI{
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
