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

	if val1, ok := context.GetOk(request1, key); ok {
		fmt.Println(val1)
	}

	if val2, ok := context.GetOk(request1, "NotAKey"); ok {
		fmt.Println(val2)
	}
}
