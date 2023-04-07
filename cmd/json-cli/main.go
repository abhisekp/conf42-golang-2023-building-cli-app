package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ImpelsysInc/go-ptr"
	"rsc.io/getopt"
)

type Input struct {
	Name  string  `json:"name"`
	Items []*Item `json:"items"`
}

type Item struct {
	Name       *string  `json:"name"`
	Quantity   *int     `json:"quantity"`
	Percentage *float32 `json:"percentage"`
	Active     *bool    `json:"active"`
	Aliases    []string `json:"aliases"`
}

func (i *Item) UnmarshalJSON(data []byte) error {
	type Alias Item
	aux := &struct {
		Aliases any `json:"aliases"`
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch aliases := aux.Aliases.(type) {
	case string:
		reComma := regexp.MustCompile(`\s*,\s*`)
		i.Aliases = reComma.Split(aliases, -1)
	case []any:
		i.Aliases = make([]string, len(aliases))
		for j, alias := range aliases {
			i.Aliases[j] = fmt.Sprintf("%v", alias)
		}
	}
	return nil
}

func main() {
	var input Input

	fs := getopt.NewFlagSet("JSON Input CLI", flag.ExitOnError)

	fs.Var(&jsonInputFlag{&input}, "input", "Input JSON")
	fs.Aliases("i", "input")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		os.Exit(2)
	}

	fmt.Println()
	for _, item := range input.Items {
		fmt.Println("name:", ptr.StringValue(item.Name))
		fmt.Println("quantity:", ptr.IntValue(item.Quantity))
		fmt.Println("percentage:", ptr.Float32Value(item.Percentage))
		fmt.Println("active:", ptr.BoolValue(item.Active))
		fmt.Println("aliases:", strings.Join(item.Aliases, ", "))
		fmt.Println()
	}
}

type jsonInputFlag struct {
	inputFile *Input
}

func (f *jsonInputFlag) String() string {
	return ""
}

func (f *jsonInputFlag) Set(value string) error {
	abs, err := filepath.Abs(value)
	if err != nil {
		return err
	}

	// Read the JSON file
	file, err := os.ReadFile(abs)
	if err != nil {
		return err
	}

	// Parse the JSON
	err = json.Unmarshal(file, f.inputFile)
	if err != nil {
		return err
	}

	return nil
}
