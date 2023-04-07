package cmd

import (
	"os"
)

func GeneratePersonSchemaToFile(personSchema []byte, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = f.Write(personSchema)
	return err
}
