package main

import (
	"flag"
	"fmt"
)

func main() {
	name := ""
	flag.StringVar(&name, "name", "World", "The name to say hello to.")

	flag.Parse()

	fmt.Println("Hello,", name)
}
