package main

import (
	"fmt"

	"github.com/gorilla/schema"
)

func main() {
	form := map[string][]string{
		"firstName":         []string{"Dirk"},
		"lastName":          []string{"Gently"},
		"addresses.0.city":  []string{"Rio de Janeiro"},
		"addresses.0.state": []string{"RJ"},
		"addresses.1.city":  []string{"Sao Paulo"},
		"addresses.1.state": []string{"SP"},
	}

	var p Person
	d := schema.NewDecoder()

	d.Decode(&p, form)

	fmt.Println(p)
}

type Person struct {
	FirstName string    `schema:"-"` // use the dash sign to ignore a field to be mapped
	LastName  string    `schema:"lastName"`
	Addresses []Address `schema:addresses`
}

type Address struct {
	City  string `schema:"city"`
	State string `schema:"state"`
}
