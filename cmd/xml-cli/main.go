package main

import "flag"

func main() {
	name := ""
	flag.StringVar(&name, "", "XML CLI", "The name to say hello to.")
	flag.Parse()
	println("Hello,", name)
}
