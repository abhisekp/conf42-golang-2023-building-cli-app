package main

import (
	"flag"
	"os"

	"rsc.io/getopt"
)

func main() {
	totalArgs := len(os.Args[1:])

	f := getopt.NewFlagSet("Abhisek's CLI", flag.ExitOnError)

	var help bool
	f.BoolVar(&help, "help", totalArgs == 0, "Show help")

	f.Aliases(
		"h", "help",
	)

	err := f.Parse(os.Args[1:])
	if err != nil {
		return
	}

	if help {
		getopt.PrintDefaults()
		os.Exit(0)
	}
}
