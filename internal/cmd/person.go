package cmd

import (
	"sync"

	"github.com/abhisekp/conf42-golang-2023-building-cli-app/pkg/person"
)

// GenPersons generates a `n` number of Person(s)
func GenPersons(n int) []*person.Person {
	persons := make([]*person.Person, n)
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			p := person.NewPerson()
			persons[i] = p
		}(i)
	}
	wg.Wait()
	return persons
}
