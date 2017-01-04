package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mitchellh/colorstring"
)

// CLI is the command line interface object.
type CLI struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

const (
	// EnvDebug is environmental var to handle debug mode
	EnvDebug = "GO_DEBUG"
)

// Exit codes are values representing an exit code for a error type.
const (
	ExitCodeOK int = 0

	// Errors start at 10
	ExitCodeError = 10 + iota
	ExitCodeParseFlagsError
)

// PrintGreenF prints green success message on console
func PrintGreenF(writer io.Writer, format string, args ...interface{}) {
	PrintColorF(writer, "green", format, args...)
}

// PrintRedF prints red error message on console
func PrintRedF(writer io.Writer, format string, args ...interface{}) {
	PrintColorF(writer, "red", format, args...)
}

// PrintColorF prints colored message on console
func PrintColorF(writer io.Writer, color, format string, args ...interface{}) {
	format = fmt.Sprintf("[%s]%s[reset]", color, format)
	fmt.Fprint(writer, colorstring.Color(fmt.Sprintf(format, args...)))
}

// DebugF prints debug message when EnvDebug is true
func DebugF(format string, args ...interface{}) {
	if env := os.Getenv(EnvDebug); len(env) != 0 {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

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

	DebugF("Example of utils for printing messages.")
	PrintRedF(cli.errStream, "You should change following lines.\n")
	fmt.Fprintf(cli.errStream, "Makefile:1:NAME := ")
	PrintRedF(cli.errStream, "%s\n", Name)

	return ExitCodeOK
}
