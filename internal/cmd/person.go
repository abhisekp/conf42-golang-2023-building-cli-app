package cmd

import (
	"github.com/abhisekp/conf42-golang-2023-building-cli-app/pkg/person"
)

// GenPersons generates a `n` number of Person(s)
func GenPersons(n int) []*person.Person {
	persons := make([]*person.Person, n)
	for i := 0; i < n; i++ {
		p := person.NewPerson()
		persons[i] = p
	}
	return persons
}
