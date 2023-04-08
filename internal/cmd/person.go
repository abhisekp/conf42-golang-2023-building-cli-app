package cmd

import (
	"github.com/abhisekp/conf42-golang-2023-building-cli-app/pkg/person"
)

type GenPersonOptions struct {
	_           struct{}
	Concurrency int
}

var defaultOption = GenPersonOptions{
	Concurrency: 10,
}

// GenPersons generates a `n` number of Person(s)
func GenPersons(n int, options ...GenPersonOptions) []*person.Person {
	option := defaultOption

	if len(options) > 0 {
		option = options[0]
	}

	persons := make([]*person.Person, n)
	// Semaphore to limit the number of concurrent goroutines
	sem := make(chan int, option.Concurrency)

	for i := 0; i < n; i++ {
		sem <- i
		go func(i int) {
			p := person.NewPerson()
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
