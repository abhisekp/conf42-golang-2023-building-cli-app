package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"rsc.io/getopt"

	"github.com/abhisekp/conf42-golang-2023-building-cli-app/internal/cmd"
)

var numOfCPUs = runtime.NumCPU()

var dir, _ = os.Getwd()

func main() {
	var numPersons int
	var concurrency int
	var isPrint bool
	var name string
	var outputFilepath string
	var forceCreate bool
	flag.IntVar(&numPersons, "num", 1, "Number of persons to generate")
	flag.IntVar(&concurrency, "concurrency", numOfCPUs, "Number of threads to run concurrently")
	flag.BoolVar(&isPrint, "print", false, "Print the generated persons")
	flag.StringVar(&name, "name", "", "Name of the xml file")
	flag.StringVar(&outputFilepath, "output", dir, "Output filepath")
	flag.BoolVar(&forceCreate, "force", false, "Overwrite existing")

	getopt.Aliases(
		"n", "num",
		"c", "concurrency",
		"p", "print",
		"o", "output",
		"N", "name",
		"f", "force",
	)

	getopt.Parse()

	if name == "" {
		panic("name is required")
	}

	outputFilepath = filepath.Join(dir, outputFilepath, fmt.Sprintf("%s.xml", name))

	err := cmd.CreateXMLFile(outputFilepath, name, numPersons, cmd.CreateXMLFileOptions{
		Concurrency: concurrency,
		ForceCreate: forceCreate,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Using concurrency of", concurrency)
}
