package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"rsc.io/getopt"

	"github.com/abhisekp/conf42-golang-2023-building-cli-app/internal/cmd"
)

//go:embed person.schema.json
var personSchema []byte

func GetCurrentDir() string {
	currDir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return currDir
}

func main() {
	var genPersonSchema bool
	flag.BoolVar(&genPersonSchema, "generate", false, "Generate a person schema file")

	getopt.Aliases(
		"g", "generate",
	)

	getopt.Parse()

	if genPersonSchema {
		fullGeneratePath := filepath.Join(GetCurrentDir(), "person.schema.json")

		err := cmd.GeneratePersonSchemaToFile(personSchema, fullGeneratePath)
		if err != nil {
			_ = fmt.Errorf("failed to generate person schema: %w", err)
			return
		}
		fmt.Println("Generated file:", fullGeneratePath)
	}

}
