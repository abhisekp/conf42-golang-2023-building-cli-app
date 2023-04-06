package main

import (
	"flag"
	"fmt"
	"os"

	"rsc.io/getopt"
)

func main() {
	name := flag.String("name", "", "The name to say hello to.")
	// flag.StringVar(&name, "n", "", "The name to say hello to.")
	getopt.Aliases(
		"n", "name",
	)

	getopt.Parse()

	if name != nil && *name == "" {
		getopt.PrintDefaults()
		os.Exit(0)
	}

	if name != nil {
		fmt.Println("Hello,", *name)
	}
}
