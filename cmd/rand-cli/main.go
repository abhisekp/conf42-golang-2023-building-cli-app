package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"rsc.io/getopt"
)

func main() {
	min, max := 0, 1

	flag.IntVar(&min, "min", min, "Minimum value")
	flag.IntVar(&max, "max", max, "Maximum value")

	getopt.Aliases(
		"m", "min",
		"M", "max",
	)

	getopt.Parse()

	if max <= min {
		fmt.Println("max must be greater than min")
		os.Exit(1)
	}

	// Seed the random number generator with the current time
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomVal := min + rnd.Intn(max-min)
	fmt.Println(randomVal)
}
