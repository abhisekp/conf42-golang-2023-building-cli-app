package person

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Address struct {
	Street  string `json:"street" xml:"Street" fake:"{street}"`
	City    string `json:"city" xml:"City" fake:"{city}"`
	Pincode string `json:"pincode" xml:"Pincode" fake:"{zip}"`
	State   string `json:"state" xml:"State" fake:"{state}"`
	Country string `json:"country" xml:"Country" fake:"{country}"`
}

type Person struct {
	FirstName string `json:"firstName" xml:"FirstName" fake:"{firstname}"`
	LastName  string `json:"lastName" xml:"LastName" fake:"{lastname}"`
	Address   `json:"address" xml:"Address"`
	Meta      `json:"meta" xml:"Meta"`
}

type Meta struct {
	Age int        `json:"age" xml:"Age" fake:"{number:0,100}"`
	Dob *time.Time `json:"dob" xml:"-"`
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

type GenPersonOptions struct {
	_           struct{}
	Concurrency int
}

var defaultOption = GenPersonOptions{
	Concurrency: 10,
}

// GenPersons generates a `n` number of Person(s)
func GenPersons(n int, options ...GenPersonOptions) []*Person {
	option := defaultOption

	if len(options) > 0 {
		option = options[0]
	}

	persons := make([]*Person, n)
	// Semaphore to limit the number of concurrent goroutines
	sem := make(chan int, option.Concurrency)

	for i := 0; i < n; i++ {
		sem <- i
		go func(i int) {
			p := NewPerson()
			persons[i] = p
			<-sem
		}(i)
	}

	// Wait for all goroutines to finish
	for i := 0; i < cap(sem); i++ {
		sem <- i
	}

	return persons
}
