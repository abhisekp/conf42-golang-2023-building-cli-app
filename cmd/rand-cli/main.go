package main

import (
	"fmt"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/davecgh/go-spew/spew"
	"github.com/hexops/valast"
)

type Student struct {
	Name       string
	Age        int
	Percentage float32
	Height     int
	Active     bool
	Meta
}

type Meta struct {
	Description string
}

func main() {
	seed := time.Now().Unix()
	faker := gofakeit.New(seed)

	var stu Student
	err := faker.Struct(&stu)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(stu)
	fmt.Println(valast.String(stu))
}
