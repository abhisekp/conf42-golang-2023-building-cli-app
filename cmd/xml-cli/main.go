package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"rsc.io/getopt"

	"github.com/abhisekp/conf42-golang-2023-building-cli-app/internal/cmd"
	"github.com/abhisekp/conf42-golang-2023-building-cli-app/pkg/person"
)

var numOfCPUs = runtime.NumCPU()

var currDir, _ = os.Getwd()

func main() {
	var numPersons int
	var concurrency int
	var isPrint bool
	var name string
	var outputFilepath string
	var forceCreate bool
	var inputFilepath string
	var generateInputJSON bool
	flag.IntVar(&numPersons, "num", 1, "Number of persons to generate")
	flag.IntVar(&concurrency, "concurrency", numOfCPUs, "Number of threads to run concurrently")
	flag.BoolVar(&isPrint, "print", false, "Print the generated persons")
	flag.StringVar(&name, "name", "", "Name of the xml file")
	flag.StringVar(&outputFilepath, "output", currDir, "Output filepath")
	flag.StringVar(&inputFilepath, "input", "", "Input filepath")
	flag.BoolVar(&forceCreate, "force", false, "Overwrite existing")
	flag.BoolVar(&generateInputJSON, "generate", false, "Generate Input JSON and schema")

	getopt.Aliases(
		"n", "num",
		"c", "concurrency",
		"p", "print",
		"o", "output",
		"i", "input",
		"N", "name",
		"f", "force",
		"g", "generate",
	)

	getopt.Parse()

	if generateInputJSON {
		dir := outputFilepath
		if inputFilepath != "" {
			dir = inputFilepath
		}
		dir, err := filepath.Abs(dir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = cmd.CreateInputJSON(dir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if name == "" {
		panic("name is required")
	}

	outputFilepath = filepath.Join(currDir, outputFilepath, fmt.Sprintf("%s.xml", name))

	var personInfo *person.Person
	if inputFilepath != "" {
		inputFilepath = filepath.Join(currDir, inputFilepath)
		personInfo = cmd.GetInputPersonData(inputFilepath)
	}

	err := cmd.CreateXMLFile(outputFilepath, name, numPersons, cmd.CreateXMLFileOptions{
		Concurrency:      concurrency,
		ForceCreate:      forceCreate,
		PredefinedPerson: personInfo,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data created using concurrency of", concurrency)
}
