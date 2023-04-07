package main

import (
	"fmt"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/davecgh/go-spew/spew"
	"github.com/hexops/valast"
)

type Address struct {
	Street  string `fake:"{street}"`
	City    string `fake:"{city}"`
	Pincode string `fake:"{zip}"`
	State   string `fake:"{state}"`
	Country string `fake:"{country}"`
}

func main() {
	seed := time.Now().Unix()
	faker := gofakeit.New(seed)

	var addr Address
	err := faker.Struct(&addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(addr)
	fmt.Println(valast.String(addr))
}
