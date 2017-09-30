package main

import "fmt"

func main() {

  slice := make([]string, 5)
  slice[0] = "apple"
  slice[1] = "banana"
  slice[2] = "pulm"
  slice[3] = "orange"
  slice[4] = "grape"

  inspectSlice(slice)
}

func inspectSlice(slice []string) {
  fmt.Printf("length[%d] capacity[%d]\n", len(slice), cap(slice))
  for i, v := range slice {
    fmt.Printf("[%d] %p %s \n", i, &slice[i], v)
  }
}
