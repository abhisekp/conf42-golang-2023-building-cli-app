package person

import (
	"fmt"
	"os"
	"strings"
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

func (p *Person) String() string {
	var personStr strings.Builder

	if p.FirstName != "" {
		personStr.WriteString(p.FirstName)
	}
	if p.LastName != "" {
		if p.FirstName != "" {
			personStr.WriteString(" ")
		}
		personStr.WriteString(p.LastName)
	}
	if p.FirstName == "" && p.LastName == "" {
		personStr.WriteString("Anonymous")
	}

	if p.Dob != nil && !p.Dob.IsZero() {
		personStr.WriteString(fmt.Sprintf(" (born %s)", p.Dob.Format("2006-01-02")))
		p.Age = int(time.Now().Sub(*p.Dob).Hours() / 24 / 365)
	}

	if p.Age != 0 {
		personStr.WriteString(fmt.Sprintf(" (aged %d yrs)", p.Age))
	}

	personStr.WriteString(" lives in ")
	addressStarted := false
	if p.Address.Street != "" {
		personStr.WriteString(p.Address.Street)
		addressStarted = true
	}
	if p.Address.City != "" {
		if addressStarted {
			personStr.WriteString(", ")
		}
		personStr.WriteString(p.Address.City)
	}
	if p.Address.State != "" {
		if addressStarted {
			personStr.WriteString(", ")
		}
		personStr.WriteString(p.Address.State)
	}
	if p.Address.Country != "" {
		if addressStarted {
			personStr.WriteString(", ")
		}
		personStr.WriteString(p.Address.Country)
	} else {
		personStr.WriteString(" lives in an unknown location")
	}

	return personStr.String()
}

type Meta struct {
	Age int        `json:"age" fake:"{number:0,100}"`
	Dob *time.Time `json:"dob"`
}

var (
	seed  = time.Now().Unix()
	faker = gofakeit.New(seed)
)

func NewPerson() *Person {
	var person Person
	err := faker.Struct(&person)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &person
}
