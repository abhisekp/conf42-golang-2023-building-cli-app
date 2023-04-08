package main

import (
	"flag"
	"fmt"
	"runtime"

	"rsc.io/getopt"

	"github.com/abhisekp/conf42-golang-2023-building-cli-app/internal/cmd"
)

var numOfCPUs = runtime.NumCPU()

func main() {
	var numPersons int
	var concurrency int
	var isPrint bool
	flag.IntVar(&numPersons, "num", 1, "Number of persons to generate")
	flag.IntVar(&concurrency, "concurrency", numOfCPUs, "Number of threads to run concurrently")
	flag.BoolVar(&isPrint, "print", false, "Print the generated persons")

	getopt.Aliases(
		"n", "num",
		"c", "concurrency",
		"p", "print",
	)

	getopt.Parse()

	persons := cmd.GenPersons(numPersons, cmd.GenPersonOptions{
		Concurrency: concurrency,
	})

	if isPrint {
		for _, p := range persons {
			fmt.Println(p)
		}
		fmt.Println("Using concurrency of", concurrency)
	}
}
