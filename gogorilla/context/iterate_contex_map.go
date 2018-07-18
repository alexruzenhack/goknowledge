package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

const key = "MyKey"

func main() {
	request1 := &http.Request{}

	context.Set(request1, key, "foo")
	context.Set(request1, "SecondKey", "bar")

	vals := context.GetAll(request1)

	fmt.Println(vals[key])
	fmt.Println(vals["SecondKey"])
}
