package main

import (
	"flag"
	"os"

	"rsc.io/getopt"
)

func main() {
	totalArgs := len(os.Args[1:])

	fs := getopt.NewFlagSet("Abhisek's CLI", flag.ExitOnError)

	var help bool
	fs.BoolVar(&help, "help", totalArgs == 0, "Show help")

	fs.Aliases(
		"h", "help",
	)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return
	}

	if help {
		fs.Usage()
		os.Exit(0)
	}
}
