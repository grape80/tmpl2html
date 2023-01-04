package main

import (
	"os"

	"github.com/grape80/tmpl2html"
)

func main() {
	cli := &tmpl2html.CLI{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Args:   os.Args,
	}
	os.Exit(cli.Run())
}
