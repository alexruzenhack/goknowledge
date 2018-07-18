package main

import (
	"fmt"

	"github.com/gorilla/schema"
)

func main() {
	form := map[string][]string{
		"FirstName":         []string{"Dirk"},
		"LastName":          []string{"Gently"},
		"Addresses.0.City":  []string{"Rio de Janeiro"},
		"Addresses.0.State": []string{"RJ"},
		"Addresses.1.City":  []string{"Sao Paulo"},
		"Addresses.1.State": []string{"SP"},
	}

	var p Person
	d := schema.NewDecoder()

	d.Decode(&p, form)

	fmt.Println(p)
}

type Person struct {
	FirstName string
	LastName  string
	Addresses []Address
}

type Address struct {
	City  string
	State string
}
