package main

import (
  "fmt"
)

func main() {
  coursesInProg := []string{"Docker Deep Dive", "Docker Clustering", "Docker and Kubernetes"}

  for _, i := range coursesInProg{
    fmt.Println(i)
  }
}
