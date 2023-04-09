package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/abhisekp/conf42-golang-2023-building-cli-app/pkg/person"
)

//go:embed input.schema.json
var jsonSchema []byte

func GetJSONSchema() string {
	return string(jsonSchema)
}

func ReadInputJSON(inputJSONFilepath string) ([]byte, error) {
	return os.ReadFile(inputJSONFilepath)
}

func GetInputPersonData(inputJSONFilepath string) *person.Person {
	personInfoBytes, err := ReadInputJSON(inputJSONFilepath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var personInfo person.Person

	err = json.Unmarshal(personInfoBytes, &personInfo)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &personInfo
}

//go:embed input.json
var inputJSON []byte

func CreateInputJSON(baseDir string) error {
	inputFileCreated := false
	// Check for existing input.json file
	if _, err := os.Stat(filepath.Join(baseDir, "input.json")); err == nil {
		fmt.Println("input.json already exists")
		inputFileCreated = true
	} else {
		file, err := os.Create(filepath.Join(baseDir, "input.json"))
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println(err)
			}

			if !inputFileCreated {
				err := os.Remove(file.Name())
				if err != nil {
					fmt.Println(err)
				}
			}
		}(file)

		_, err = file.Write(inputJSON)
		if err != nil {
			return err
		}
		inputFileCreated = true
	}

	schemaFileCreated := false
	// Create a schema file
	file, err := os.Create(filepath.Join(baseDir, "input.schema.json"))
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}

		if !schemaFileCreated {
			err := os.Remove(file.Name())
			if err != nil {
				fmt.Println(err)
			}
		}
	}(file)

	_, err = file.WriteString(GetJSONSchema())
	if err != nil {
		return err
	}
	schemaFileCreated = true

	return nil
}
