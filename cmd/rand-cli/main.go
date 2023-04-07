package main

import (
	"flag"
	"fmt"

	"rsc.io/getopt"

	"github.com/abhisekp/conf42-golang-2023-building-cli-app/internal/cmd"
)

func main() {
	var numPersons int
	flag.IntVar(&numPersons, "num", 1, "Number of persons to generate")
	getopt.Aliases(
		"n", "num",
	)

	getopt.Parse()

	persons := cmd.GenPersons(numPersons)
	for _, p := range persons {
		fmt.Println(p)
	}
}
