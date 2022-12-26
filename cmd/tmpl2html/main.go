package main

import (
	"os"

	"github.com/grape80/tmpl2html"
)

func main() {
	cli := &tmpl2html.CLI{
		Args:   os.Args,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	os.Exit(cli.Run())
}
