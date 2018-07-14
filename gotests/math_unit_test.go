package pack

import (
  "testing"
  "time"
  "os"
)

func TestMain(m *testing.M) {
  println("Test are about to run")
  result := m.Run()
  println("Test done executing")

  os.Exit(result)
}

var addTable = []struct {
  in []int
  out int
} {
  {[]int{1,2}, 3},
  {[]int{1,2,3,4}, 10},
}

func TestCanAddNumbers(t *testing.T) {
  t.Parallel()
  for _, entry := range addTable {
    result := Add(entry.in...)
    if result !=  entry.out {
      t.Error("Failed to add numbers as expected")
    }
  }
}

func TestCanSubtractNumbers(t *testing.T) {
  t.Parallel()
  result := Subtract(10, 5)
  time.Sleep(1 * time.Second)

  if result != 5 {
    t.Error("Failed to subtract two numbers properly")
  }
}
