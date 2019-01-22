package main

import (
	"github.com/jessevdk/go-flags"
	"os"
)
type Options struct {
	Logging string `short:"l" long:"logging" description:"Enables logging." default:"warn" choice:"debug" choice:"warn"`
}
var options Options
var parser = flags.NewParser(&options, flags.Default)

func main() {
	_, err := parser.Parse() // reads program arguments, executes the program
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0) // if --help was called
		} else {
			os.Exit(1) // if incorrect parameters have been passed
		}
	}
	//os.Exit(0)
}