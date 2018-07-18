package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/context"
)

const key = "MyKey"

func main() {
	request1 := &http.Request{}

	context.Set(request1, key, "foo")
	context.Set(request1, "SecondKey", "bar")

	fmt.Println(context.Get(request1, key))

	time.Sleep(2 * time.Second)

	context.Purge(1) // 1 second

	fmt.Println(context.Get(request1, key))

}
