package main

import (
	"fmt"

	"github.com/gorilla/schema"
)

func main() {
	form := map[string][]string{
		"FirstName": []string{"Dirk"},
		"LastName":  []string{"Gently"},
	}

	var p Person
	d := schema.NewDecoder()

	d.Decode(&p, form)

	fmt.Println(p)
}

type Person struct {
	FirstName string
	LastName  string
}
