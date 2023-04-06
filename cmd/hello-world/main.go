package main

import (
	"flag"
	"fmt"
	"os"

	"rsc.io/getopt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "The name to say hello to.")
	flag.StringVar(&name, "n", "", "The name to say hello to.")
	var repeat int
	flag.IntVar(&repeat, "repeat", 1, "The number of times to repeat.")

	totalArgs := len(os.Args[1:])

	var help bool
	flag.BoolVar(&help, "help", totalArgs == 0, "Show help")

	getopt.Aliases(
		"n", "name",
		"r", "repeat",
		"h", "help",
	)

	getopt.Parse()

	if help {
		getopt.PrintDefaults()
		os.Exit(0)
	}

	if name != "" {
		for i := 0; i < repeat; i++ {
			fmt.Println("Hello,", name)
		}
	}
}
