package person

import (
	"fmt"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Address struct {
	Street  string `fake:"{street}"`
	City    string `fake:"{city}"`
	Pincode string `fake:"{zip}"`
	State   string `fake:"{state}"`
	Country string `fake:"{country}"`
}

type Person struct {
	FirstName string `json:"firstName" fake:"{firstname}"`
	LastName  string `json:"lastName" fake:"{lastname}"`
	Address   `json:"address"`
	Meta      `json:"meta"`
}

type Meta struct {
	Age int       `json:"age"`
	Dob time.Time `json:"dob"`
}

func NewPerson() *Person {
	seed := time.Now().Unix()
	faker := gofakeit.New(seed)

	var person Person
	var addr Address
	err := faker.Struct(&addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	person.Address = addr

	return &person
}
