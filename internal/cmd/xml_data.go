package cmd

import (
	_ "embed"
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/abhisekp/conf42-golang-2023-building-cli-app/pkg/person"
	"github.com/abhisekp/conf42-golang-2023-building-cli-app/pkg/xml_data"
)

//go:embed output.xsd
var xmlSchema string

func GenXMLData(name string, persons []*person.Person) xml_data.Root {
	return xml_data.Root{
		XMLName: xml.Name{
			Local: "Persona",
		},
		XMLSchema:    "./output.xsd",
		XMLNamespace: "http://www.w3.org/2001/XMLSchema-instance",
		Name:         name,
		DateCreated:  time.Now().Format(time.RFC3339),
		Items:        persons,
	}
}

func GenXML(xmlData xml_data.Root, writer io.Writer) error {
	_, err := writer.Write([]byte(xml.Header))
	if err != nil {
		return err
	}

	enc := xml.NewEncoder(writer)
	enc.Indent("", "  ")

	if err := enc.Encode(xmlData); err != nil {
		fmt.Println("Error encoding XML:", err)
		return err
	}

	return nil
}

type CreateXMLFileOptions struct {
	Concurrency int
	ForceCreate bool
}

var defaultCreateXMLFileOptions = CreateXMLFileOptions{
	Concurrency: 1,
	ForceCreate: false,
}

func CreateXMLFile(absFilepath, name string, numPersons int, options ...CreateXMLFileOptions) error {
	option := defaultCreateXMLFileOptions
	if len(options) > 0 {
		option = options[0]
	}

	concurrency := option.Concurrency
	forceCreate := option.ForceCreate

	// Check existing file
	baseDir := filepath.Dir(absFilepath)
	err := os.MkdirAll(baseDir, fs.ModePerm)
	if err != nil {
		return err
	}

	_, err = os.Stat(absFilepath)
	if _, ok := err.(*fs.PathError); !forceCreate && !ok {
		fmt.Println("File already exists:", absFilepath)
		return err
	}

	dataWritten := false
	file, err := os.Create(absFilepath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing output file:", err)
			os.Exit(1)
		}

		if !dataWritten {
			_ = os.Remove(absFilepath)
		}
	}(file)

	persons := person.GenPersons(numPersons, person.GenPersonOptions{
		Concurrency: concurrency,
	})
	xmlData := GenXMLData(name, persons)

	// Create XML output file and write records to it
	err = GenXML(xmlData, file)
	if err != nil {
		return err
	}

	err = CreateSchemaFile(absFilepath)
	if err != nil {
		return err
	}

	dataWritten = true
	return nil
}

func CreateSchemaFile(xmlFilepath string) error {
	basePath := filepath.Dir(xmlFilepath)
	schemaFilePath := filepath.Join(basePath, "output.xsd")
	file, err := os.Create(schemaFilePath)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing output file:", err)
			os.Exit(1)
		}
	}(file)

	_, err = file.WriteString(xmlSchema)
	if err != nil {
		return err
	}

	return nil
}
