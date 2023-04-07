package main

import (
	"errors"
	"flag"
	"fmt"
	"regexp"
	"strconv"
)

var reNum = regexp.MustCompile(`\d+`)

func main() {
	var num int
	flag.Func("num", "number of records to generate", func(s string) error {
		var err error
		num, err = strconv.Atoi(reNum.FindString(s))
		if err != nil {
			return errors.New("string must have number")
		}
		return nil
	})

	flag.Parse()

	fmt.Println("num:", num)
}
