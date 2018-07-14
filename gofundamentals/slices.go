package main

import (
  "fmt"
)

func main() {
  //Declares a slice called myCourses
  myCourses := []string{"Docker", "Puppet", "Python"}

  fmt.Printf("Length is: %d.\nCapacity is: %d", len(myCourses), cap(myCourses))
}
