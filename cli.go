package main

import (
	"flag"
	"fmt"
	"io"
)

// CLI is the command line interface object.
type CLI struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

// Exit codes are values representing an exit code for a error type.
const (
	ExitCodeOK int = 0

	// Errors start at 10
	ExitCodeError = 10 + iota
	ExitCodeParseFlagsError
)

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		version bool
	)

	flags := flag.NewFlagSet(Name, flag.ExitOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&version, "version", false, "")
	flags.BoolVar(&version, "v", false, "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if version {
		fmt.Fprintf(cli.outStream, OutputVersion())
		return ExitCodeOK
	}

	return ExitCodeOK
}
