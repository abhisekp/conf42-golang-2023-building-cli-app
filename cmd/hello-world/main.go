package main

import "flag"

func main() {
	name := ""
	flag.StringVar(&name, "name", "World", "The name to say hello to.")
	flag.StringVar(&name, "n", "World", "The name to say hello to.")
	flag.Parse()
	println("Hello,", name)
}
