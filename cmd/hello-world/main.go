package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "The name to say hello to.")
	flag.StringVar(&name, "n", "", "The name to say hello to.")

	flag.Parse()

	if name == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
	fmt.Println("Hello,", name)
}
