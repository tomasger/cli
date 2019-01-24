package main

import (
	"github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	Logging string `short:"l" long:"logging" description:"Enables logging." choice:"debug" choice:"warn"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func main() {
	// parser.Parse() reads program arguments, executes the program.
	// Check servers.go and login.go contain Execute methods where the program starts.
	_, err := parser.Parse()
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0) // if --help was called
		} else {
			os.Exit(1) // if incorrect parameters have been passed
		}
	}
}
