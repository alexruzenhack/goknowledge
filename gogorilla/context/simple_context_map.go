package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

const key = "MyKey"

func main() {
	request1 := &http.Request{}
	request2 := &http.Request{}

	context.Set(request1, key, "foo")
	context.Set(request2, key, "bar")

	fmt.Println(context.Get(request1, key)) // print foo

	context.Delete(request1, key)

	fmt.Println(context.Get(request1, key)) // print nil

	// separeted context
	fmt.Println(context.Get(request2, key)) // print bar

	context.Clear(request2) // remove all the key set

	fmt.Println(context.Get(request2, key)) // print nil
}
