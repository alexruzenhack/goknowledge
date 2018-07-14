package main

import (
  "fmt"
  "reflect"
)

func main() {

  name := "Nigel"
  module := 3.2
  ptr := &module

  fmt.Println("Name is set to", reflect.TypeOf(name))
  fmt.Println("Module is set to", reflect.TypeOf(module))
  fmt.Println("Memory address of *module* variable is",
    ptr, "and the value of *module* is", *ptr)
}
