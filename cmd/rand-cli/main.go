package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/davecgh/go-spew/spew"
	"github.com/hexops/valast"
)

type Address struct {
	Street  string
	City    string
	Pincode string
	State   string
	Country string
}

func main() {
	seed := time.Now().Unix()
	faker := gofakeit.New(seed)

	fakeAddr := faker.Address()

	addr := Address{
		Street:  fakeAddr.Street,
		City:    fakeAddr.City,
		Pincode: fakeAddr.Zip,
		State:   fakeAddr.State,
		Country: fakeAddr.Country,
	}

	spew.Dump(addr)
	fmt.Println(valast.String(addr))
}
