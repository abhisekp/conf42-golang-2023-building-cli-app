package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type Address struct {
	Street  string
	City    string
	State   string
	Pincode string
	Country string
}

type AddressValue struct {
	*Address
}

func (a AddressValue) Get() any {
	return a.Address
}

func (a AddressValue) String() string {
	if a.Address == nil {
		return ""
	}
	return fmt.Sprintf("%s, %s, %s, %s, %s", a.Address.Street, a.Address.City, a.Address.State, a.Address.Pincode, a.Address.Country)
}

func (a AddressValue) Set(addressStr string) error {
	addressParts := strings.Split(addressStr, ", ")
	if len(addressParts) == 0 || len(addressParts) > 5 {
		return errors.New("Address should be in the format: Street, City, State, Pincode, Country")
	}

	if len(addressParts) >= 1 {
		a.Address.Street = addressParts[0]
	}

	if len(addressParts) >= 2 {
		a.Address.City = addressParts[1]
	}

	if len(addressParts) >= 3 {
		a.Address.State = addressParts[2]
	}

	if len(addressParts) >= 4 {
		a.Address.Pincode = addressParts[3]
	}

	if len(addressParts) >= 5 {
		a.Address.Country = addressParts[4]
	}

	return nil
}

func main() {
	var address Address
	addressVar := AddressValue{&address}
	flag.Var(&addressVar, "address", "The address of the person in the format: Street, City, State, Pincode, Country")

	flag.Parse()

	fmt.Println("Address")
	fmt.Printf("%q\n", address)
}
